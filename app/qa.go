// 質問対話機能。
// 回答はローカルの Claude Code CLI (`claude -p`) をヘッドレス実行して得る。
// Q&A は app/questions.json に保存し、疑問点まとめ (app/notes.md) と
// 復習問題 (lesson90/chNN.json) の自動生成に使う。
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"
)

var claudeModel string // -claude-model フラグ（既定 "sonnet"）

const reviewLesson = 90 // 自動生成した復習問題を置くレッスン番号

// ---- Q&A ストア ----

type QA struct {
	ID           string `json:"id"`
	Time         string `json:"time"`
	Lesson       int    `json:"lesson"`
	Chapter      int    `json:"chapter"`
	ChapterTitle string `json:"chapterTitle"`
	Question     string `json:"question"`
	Answer       string `json:"answer"`
}

type qaStore struct {
	mu   sync.Mutex
	path string
	list []QA
}

var questions *qaStore

func newQAStore(path string) *qaStore {
	s := &qaStore{path: path}
	if b, err := os.ReadFile(path); err == nil {
		json.Unmarshal(b, &s.list)
	}
	return s
}

func (s *qaStore) save() {
	b, _ := json.MarshalIndent(s.list, "", "  ")
	os.WriteFile(s.path, b, 0o644)
}

func (s *qaStore) add(qa QA) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.list = append(s.list, qa)
	s.save()
}

func (s *qaStore) remove(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, qa := range s.list {
		if qa.ID == id {
			s.list = append(s.list[:i], s.list[i+1:]...)
			s.save()
			return true
		}
	}
	return false
}

func (s *qaStore) all() []QA {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]QA, len(s.list))
	copy(out, s.list)
	return out
}

// recentInChapter は同じチャプターの直近の Q&A を返す（文脈用）。
func (s *qaStore) recentInChapter(lesson, chapter, n int) []QA {
	s.mu.Lock()
	defer s.mu.Unlock()
	var hits []QA
	for _, qa := range s.list {
		if qa.Lesson == lesson && qa.Chapter == chapter {
			hits = append(hits, qa)
		}
	}
	if len(hits) > n {
		hits = hits[len(hits)-n:]
	}
	return hits
}

// ---- Claude CLI 実行 ----

func runClaude(prompt string, timeout time.Duration) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, "claude", "-p", "--model", claudeModel, "--output-format", "text")
	cmd.Dir = courseRoot
	cmd.Stdin = strings.NewReader(prompt)
	var out, errBuf strings.Builder
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	err := cmd.Run()
	if ctx.Err() == context.DeadlineExceeded {
		return "", errors.New("Claude の応答がタイムアウトしました")
	}
	if err != nil {
		msg := strings.TrimSpace(errBuf.String())
		if msg == "" {
			msg = err.Error()
		}
		return "", fmt.Errorf("claude の実行に失敗しました: %s", msg)
	}
	return strings.TrimSpace(out.String()), nil
}

// ---- 質問に回答 ----

type askReq struct {
	Lesson   int    `json:"lesson"`
	Chapter  int    `json:"chapter"`
	Question string `json:"question"`
	Code     string `json:"code"`
}

func handleAsk(w http.ResponseWriter, r *http.Request) {
	var req askReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if strings.TrimSpace(req.Question) == "" {
		http.Error(w, "質問が空です", 400)
		return
	}

	var b strings.Builder
	b.WriteString("あなたは「Go入門編」（Goの基礎を学ぶ入門講座） の講師です。")
	b.WriteString("受講者は Python 入門を終えて Go を学び始めた初心者です。\n\n")

	chapterTitle := ""
	if c, err := loadChapter(req.Lesson, req.Chapter); err == nil {
		chapterTitle = c.Title
		fmt.Fprintf(&b, "現在学習中のチャプター: lesson%02d ch%02d「%s」\n\n教材の説明:\n%s\n\n演習の問題文:\n%s\n\n",
			req.Lesson, req.Chapter, c.Title, c.Text, c.Exercise.Description)
	}
	if strings.TrimSpace(req.Code) != "" {
		fmt.Fprintf(&b, "受講者が今エディタに書いているコード:\n```go\n%s\n```\n\n", req.Code)
	}
	if recent := questions.recentInChapter(req.Lesson, req.Chapter, 3); len(recent) > 0 {
		b.WriteString("このチャプターでの直近の質疑:\n")
		for _, qa := range recent {
			fmt.Fprintf(&b, "Q: %s\nA: %s\n", qa.Question, qa.Answer)
		}
		b.WriteString("\n")
	}
	fmt.Fprintf(&b, "受講者の質問:\n%s\n\n", req.Question)
	b.WriteString(`回答のルール:
- です・ます調で、初心者向けに分かりやすく簡潔に答える(目安400字以内。コード例は字数に含めない)
- Python との対比が理解を助けるなら1つ入れる
- ファイル操作やツールは使わず、上記の文脈だけで答える
- 演習の答えそのものを聞かれた場合は、答えを書かずヒントに留める
- Markdown で書く(コードは ` + "```go" + ` フェンス)`)

	answer, err := runClaude(b.String(), 180*time.Second)
	if err != nil {
		http.Error(w, err.Error(), 502)
		return
	}

	qa := QA{
		ID:           fmt.Sprintf("qa-%d", time.Now().UnixNano()),
		Time:         time.Now().Format("2006-01-02 15:04"),
		Lesson:       req.Lesson,
		Chapter:      req.Chapter,
		ChapterTitle: chapterTitle,
		Question:     req.Question,
		Answer:       answer,
	}
	questions.add(qa)
	writeJSON(w, qa)
}

func handleQAList(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, map[string]any{"items": questions.all()})
}

func handleQADelete(w http.ResponseWriter, r *http.Request) {
	if !questions.remove(r.PathValue("id")) {
		http.Error(w, "not found", 404)
		return
	}
	writeJSON(w, map[string]bool{"ok": true})
}

// ---- 疑問点まとめ ----

func notesPath() string { return filepath.Join(appDir, "notes.md") }

func handleSummaryGet(w http.ResponseWriter, r *http.Request) {
	b, _ := os.ReadFile(notesPath())
	writeJSON(w, map[string]string{"summary": string(b)})
}

func handleSummaryCreate(w http.ResponseWriter, r *http.Request) {
	all := questions.all()
	if len(all) == 0 {
		http.Error(w, "保存された質問がまだありません", 400)
		return
	}
	var b strings.Builder
	b.WriteString("あなたは「Go入門編」（Goの基礎を学ぶ入門講座） の講師です。\n")
	b.WriteString("以下は受講者(Python 経験ありの Go 初心者)がこれまでにした質問と回答の記録です。\n\n")
	for _, qa := range all {
		fmt.Fprintf(&b, "--- [%s] lesson%02d ch%02d %s\nQ: %s\nA: %s\n\n",
			qa.Time, qa.Lesson, qa.Chapter, qa.ChapterTitle, qa.Question, qa.Answer)
	}
	b.WriteString(`この記録から「疑問点まとめノート」を Markdown で作成してください。構成:
# 疑問点まとめ
## つまずきの傾向 (質問をテーマ別に分類し、それぞれ2〜3文で本質を説明)
## 要復習ポイント (優先度順の箇条書き。各項目に関連レッスン番号を付ける)
## 理解できている点 (質問の仕方から判断できるもの)
## 次の一歩 (具体的な復習アドバイス2〜3個)
ルール: です・ます調。ツールは使わない。Markdown 本文だけを出力する。`)

	summary, err := runClaude(b.String(), 300*time.Second)
	if err != nil {
		http.Error(w, err.Error(), 502)
		return
	}
	os.WriteFile(notesPath(), []byte(summary+"\n"), 0o644)
	writeJSON(w, map[string]string{"summary": summary})
}

// ---- 復習問題の生成 (lesson90/chNN.json) ----

func handleReviewCreate(w http.ResponseWriter, r *http.Request) {
	all := questions.all()
	if len(all) == 0 {
		http.Error(w, "保存された質問がまだありません", 400)
		return
	}

	var b strings.Builder
	b.WriteString("あなたは「Go入門編」（Goの基礎を学ぶ入門講座） の教材作成者です。\n")
	b.WriteString("以下は受講者(Python 経験ありの Go 初心者)の質問記録です。この受講者の弱点を突く復習チャプターを 1〜3 個作ってください。\n\n")
	for _, qa := range all {
		fmt.Fprintf(&b, "--- lesson%02d ch%02d %s\nQ: %s\nA: %s\n\n", qa.Lesson, qa.Chapter, qa.ChapterTitle, qa.Question, qa.Answer)
	}
	b.WriteString("出力は次のスキーマの JSON 配列 **のみ** を ```json フェンスで囲んで出力してください。前後に説明文を書かないでください。\n\n")
	b.WriteString("```\n[{\n" +
		`  "lesson": 90, "chapter": 1,
  "title": "復習: <弱点のテーマ>",
  "priority": "🔴",
  "text": "説明 Markdown。受講者が過去に質問した誤解のポイントを踏まえる。Python ではこう書いた:/Go ではこう書く: の対比コードを1組含める",
  "sample": "動く完全な Go コード (package main + func main、10〜25行)",
  "exercise": {
    "description": "問題文 Markdown。末尾に **期待出力** と ` + "```" + `text フェンスの期待出力を含める",
    "starter": "package main\n\nimport \"fmt\"\n\nfunc main() {\n\t// ここに書く\n\t_ = fmt.Println\n}\n",
    "stdin": "",
    "expected_stdout": "期待出力 (末尾改行付き)",
    "answer": "解答コード (なぜこう書くかのコメント2〜3行付き)"
  }
}]` + "\n```\n\n")
	b.WriteString(`ルール:
- chapter は 1 から連番 (保存時にこちらで振り直すので仮でよい)
- 演習は標準入力を使わず固定値で出力する問題にする (stdin は "")
- answer を go run すると expected_stdout と完全一致すること (行末の空白なし)
- float の丸め誤差が出る値・乱数・時刻など非決定的な出力は禁止
- priority は 🔴/🟡/⚪ のいずれか。説明は です・ます調
- ツールやファイル操作は使わない。JSON だけを出力する`)

	raw, err := runClaude(b.String(), 420*time.Second)
	if err != nil {
		http.Error(w, err.Error(), 502)
		return
	}
	chapters, err := parseChapterArray(raw)
	if err != nil {
		http.Error(w, "生成結果の JSON を解釈できませんでした: "+err.Error(), 502)
		return
	}

	type created struct {
		Lesson  int    `json:"lesson"`
		Chapter int    `json:"chapter"`
		Title   string `json:"title"`
	}
	var ok []created
	var failed []string

	dir := filepath.Join(courseRoot, fmt.Sprintf("lesson%02d", reviewLesson))
	os.MkdirAll(dir, 0o755)
	next := nextChapterNum(dir)

	for _, c := range chapters {
		if err := validateGenerated(&c); err != nil {
			// 1 回だけ修正を試みる
			if fixed, ferr := repairChapter(c, err); ferr == nil {
				c = *fixed
			} else {
				failed = append(failed, fmt.Sprintf("%s (%v)", c.Title, err))
				continue
			}
		}
		c.Lesson = reviewLesson
		c.Chapter = next
		buf, _ := json.MarshalIndent(&c, "", "  ")
		path := filepath.Join(dir, fmt.Sprintf("ch%02d.json", next))
		if err := os.WriteFile(path, append(buf, '\n'), 0o644); err != nil {
			failed = append(failed, c.Title+" (書き込み失敗)")
			continue
		}
		ok = append(ok, created{reviewLesson, next, c.Title})
		next++
	}
	if len(ok) > 0 {
		ensureReviewLessonInCourse()
	}
	writeJSON(w, map[string]any{"created": ok, "failed": failed})
}

// parseChapterArray は Claude の出力から JSON 配列を取り出す。
func parseChapterArray(raw string) ([]Chapter, error) {
	if m := regexp.MustCompile("(?s)```(?:json)?\\s*(\\[.*?\\])\\s*```").FindStringSubmatch(raw); m != nil {
		raw = m[1]
	} else if i := strings.Index(raw, "["); i >= 0 {
		if j := strings.LastIndex(raw, "]"); j > i {
			raw = raw[i : j+1]
		}
	}
	var chapters []Chapter
	if err := json.Unmarshal([]byte(raw), &chapters); err != nil {
		return nil, err
	}
	if len(chapters) == 0 {
		return nil, errors.New("チャプターが空です")
	}
	return chapters, nil
}

// validateGenerated は生成チャプターを実際に go run して検証する。
func validateGenerated(c *Chapter) error {
	if c.Title == "" || c.Sample == "" || c.Exercise.Answer == "" || c.Exercise.Starter == "" {
		return errors.New("必須フィールドが空です")
	}
	tests := c.tests()
	if strings.TrimSpace(tests[0].Expected) == "" {
		return errors.New("期待出力が空です")
	}
	if stdout, stderr, timedOut, err := runGo(c.Sample, tests[0].Stdin); err != nil || timedOut || stderr != "" {
		_ = stdout
		return fmt.Errorf("sample が動きません: %s", firstNonEmpty(stderr, "タイムアウトまたは内部エラー"))
	}
	for i, t := range tests {
		stdout, stderr, timedOut, err := runGo(c.Exercise.Answer, t.Stdin)
		if err != nil || timedOut || stderr != "" {
			return fmt.Errorf("answer がテスト%dで動きません: %s", i+1, firstNonEmpty(stderr, "タイムアウトまたは内部エラー"))
		}
		if normalize(stdout) != normalize(t.Expected) {
			return fmt.Errorf("answer の出力がテスト%dで不一致 (got %q / want %q)", i+1, normalize(stdout), normalize(t.Expected))
		}
	}
	return nil
}

// repairChapter は検証エラーを Claude に渡して 1 回だけ修正させる。
func repairChapter(c Chapter, cause error) (*Chapter, error) {
	orig, _ := json.MarshalIndent(&c, "", "  ")
	prompt := fmt.Sprintf(`以下の Go 学習教材チャプター JSON に問題があります。
問題: %s

JSON:
%s

問題を修正した完全な JSON を、同じスキーマの **JSON 配列(要素1個)** として `+"```json フェンスだけで"+`出力してください。説明文は書かないでください。answer を go run した出力が expected_stdout と完全一致するようにしてください。`, cause, string(orig))
	raw, err := runClaude(prompt, 300*time.Second)
	if err != nil {
		return nil, err
	}
	chapters, err := parseChapterArray(raw)
	if err != nil {
		return nil, err
	}
	fixed := chapters[0]
	if err := validateGenerated(&fixed); err != nil {
		return nil, err
	}
	return &fixed, nil
}

func nextChapterNum(dir string) int {
	next := 1
	re := regexp.MustCompile(`^ch(\d{2})\.json$`)
	entries, _ := os.ReadDir(dir)
	var nums []int
	for _, e := range entries {
		if m := re.FindStringSubmatch(e.Name()); m != nil {
			var n int
			fmt.Sscanf(m[1], "%d", &n)
			nums = append(nums, n)
		}
	}
	if len(nums) > 0 {
		sort.Ints(nums)
		next = nums[len(nums)-1] + 1
	}
	return next
}

// ensureReviewLessonInCourse は course.json に復習レッスンの行を追加する。
func ensureReviewLessonInCourse() {
	path := filepath.Join(courseRoot, "course.json")
	var lessons []courseLesson
	if b, err := os.ReadFile(path); err == nil {
		json.Unmarshal(b, &lessons)
	}
	for _, l := range lessons {
		if l.Lesson == reviewLesson {
			return
		}
	}
	lessons = append(lessons, courseLesson{Lesson: reviewLesson, Title: "復習問題（質問から自動生成）"})
	b, _ := json.MarshalIndent(lessons, "", "  ")
	os.WriteFile(path, append(b, '\n'), 0o644)
}

func firstNonEmpty(s, fallback string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return fallback
	}
	if lines := strings.SplitN(s, "\n", 2); len(lines) > 0 {
		return lines[0]
	}
	return s
}
