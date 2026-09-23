// ローカルで動く Go 学習アプリ（教材・コードエディタ・自動採点の3ペイン）。
// 左に教材、右上にエディタ、右下に出力と合否を表示する 3 ペイン構成。
// 教材データは lessonNN/chMM.json を正とする。
//
// 起動: go run ./app  →  http://127.0.0.1:8080
package main

import (
	"context"
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"log"
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

//go:embed static
var staticFS embed.FS

var (
	courseRoot string
	appDir     string // progress.json / work/ の置き場所（app ディレクトリ）
)

// ---- 教材データ ----

type Test struct {
	Stdin    string `json:"stdin"`
	Expected string `json:"expected_stdout"`
}

type Exercise struct {
	Description string `json:"description"`
	Starter     string `json:"starter"`
	Stdin       string `json:"stdin"`
	Expected    string `json:"expected_stdout"`
	Answer      string `json:"answer"`
	Tests       []Test `json:"tests"`
}

type Chapter struct {
	Lesson   int      `json:"lesson"`
	Chapter  int      `json:"chapter"`
	Title    string   `json:"title"`
	Priority string   `json:"priority"`
	Text     string   `json:"text"`
	Sample   string   `json:"sample"`
	Exercise Exercise `json:"exercise"`
}

// tests は単一 stdin/expected 形式と tests 配列形式を統一して返す。
func (c *Chapter) tests() []Test {
	if len(c.Exercise.Tests) > 0 {
		return c.Exercise.Tests
	}
	return []Test{{Stdin: c.Exercise.Stdin, Expected: c.Exercise.Expected}}
}

type courseLesson struct {
	Lesson int    `json:"lesson"`
	Title  string `json:"title"`
}

func chapterPath(lesson, ch int) string {
	return filepath.Join(courseRoot, fmt.Sprintf("lesson%02d", lesson), fmt.Sprintf("ch%02d.json", ch))
}

func loadChapter(lesson, ch int) (*Chapter, error) {
	b, err := os.ReadFile(chapterPath(lesson, ch))
	if err != nil {
		return nil, err
	}
	var c Chapter
	if err := json.Unmarshal(b, &c); err != nil {
		return nil, fmt.Errorf("%s: %w", chapterPath(lesson, ch), err)
	}
	return &c, nil
}

// ---- 進捗（app/progress.json）----

type progressStore struct {
	mu   sync.Mutex
	path string
	m    map[string]string // "lesson01/ch01" -> "ran" | "passed"
}

var progress *progressStore

func newProgressStore(path string) *progressStore {
	s := &progressStore{path: path, m: map[string]string{}}
	if b, err := os.ReadFile(path); err == nil {
		json.Unmarshal(b, &s.m)
	}
	return s
}

func (s *progressStore) get(key string) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.m[key]
}

// set は状態を前進方向にのみ更新する（passed を ran に戻さない）。
func (s *progressStore) set(key, status string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.m[key] == "passed" && status != "passed" {
		return
	}
	if s.m[key] == status {
		return
	}
	s.m[key] = status
	b, _ := json.MarshalIndent(s.m, "", "  ")
	os.WriteFile(s.path, b, 0o644)
}

func progressKey(lesson, ch int) string {
	return fmt.Sprintf("lesson%02d/ch%02d", lesson, ch)
}

// ---- 書きかけコード（app/work/）----

func workPath(lesson, ch int) string {
	return filepath.Join(appDir, "work", fmt.Sprintf("lesson%02d_ch%02d.go", lesson, ch))
}

func saveWork(lesson, ch int, code string) {
	if strings.TrimSpace(code) == "" {
		return
	}
	os.MkdirAll(filepath.Join(appDir, "work"), 0o755)
	os.WriteFile(workPath(lesson, ch), []byte(code), 0o644)
}

func loadWork(lesson, ch int) string {
	b, err := os.ReadFile(workPath(lesson, ch))
	if err != nil {
		return ""
	}
	return string(b)
}

// ---- go run 実行 ----

func runGo(code, stdin string) (stdout, stderr string, timedOut bool, err error) {
	tmp, err := os.MkdirTemp("", "go-drill-run-*")
	if err != nil {
		return "", "", false, err
	}
	defer os.RemoveAll(tmp)

	if err := os.WriteFile(filepath.Join(tmp, "main.go"), []byte(code), 0o644); err != nil {
		return "", "", false, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "go", "run", "main.go")
	cmd.Dir = tmp
	cmd.Stdin = strings.NewReader(stdin)
	var out, errBuf strings.Builder
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	runErr := cmd.Run()
	if ctx.Err() == context.DeadlineExceeded {
		return out.String(), errBuf.String(), true, nil
	}
	_ = runErr // 非 0 終了（コンパイルエラー等）は stderr で伝わるのでエラー扱いにしない
	return out.String(), errBuf.String(), false, nil
}

// normalize は行末の空白と末尾の改行を無視した比較用に文字列を揃える。
func normalize(s string) string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	lines := strings.Split(s, "\n")
	for i := range lines {
		lines[i] = strings.TrimRight(lines[i], " \t")
	}
	return strings.TrimRight(strings.Join(lines, "\n"), "\n")
}

// ---- HTTP ハンドラ ----

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(v)
}

var chFileRe = regexp.MustCompile(`^ch(\d{2})\.json$`)

func handleCourse(w http.ResponseWriter, r *http.Request) {
	var titles []courseLesson
	if b, err := os.ReadFile(filepath.Join(courseRoot, "course.json")); err == nil {
		json.Unmarshal(b, &titles)
	}
	titleOf := map[int]string{}
	for _, t := range titles {
		titleOf[t.Lesson] = t.Title
	}

	type chapterInfo struct {
		Chapter  int    `json:"chapter"`
		Title    string `json:"title"`
		Priority string `json:"priority"`
		Status   string `json:"status"`
	}
	type lessonInfo struct {
		Lesson   int           `json:"lesson"`
		Title    string        `json:"title"`
		Chapters []chapterInfo `json:"chapters"`
	}

	var lessons []lessonInfo
	for n := 1; n <= 99; n++ {
		dir := filepath.Join(courseRoot, fmt.Sprintf("lesson%02d", n))
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		li := lessonInfo{Lesson: n, Title: titleOf[n]}
		for _, e := range entries {
			m := chFileRe.FindStringSubmatch(e.Name())
			if m == nil {
				continue
			}
			var chNum int
			fmt.Sscanf(m[1], "%d", &chNum)
			c, err := loadChapter(n, chNum)
			if err != nil {
				continue
			}
			li.Chapters = append(li.Chapters, chapterInfo{
				Chapter:  chNum,
				Title:    c.Title,
				Priority: c.Priority,
				Status:   progress.get(progressKey(n, chNum)),
			})
		}
		if len(li.Chapters) > 0 {
			sort.Slice(li.Chapters, func(i, j int) bool { return li.Chapters[i].Chapter < li.Chapters[j].Chapter })
			lessons = append(lessons, li)
		}
	}
	writeJSON(w, map[string]any{"lessons": lessons})
}

func pathInts(r *http.Request) (lesson, ch int, ok bool) {
	if _, err := fmt.Sscanf(r.PathValue("lesson"), "%d", &lesson); err != nil {
		return 0, 0, false
	}
	if _, err := fmt.Sscanf(r.PathValue("ch"), "%d", &ch); err != nil {
		return 0, 0, false
	}
	return lesson, ch, lesson >= 1 && lesson <= 99 && ch >= 1 && ch <= 99
}

func handleChapter(w http.ResponseWriter, r *http.Request) {
	lesson, ch, ok := pathInts(r)
	if !ok {
		http.Error(w, "invalid id", 400)
		return
	}
	c, err := loadChapter(lesson, ch)
	if err != nil {
		http.Error(w, "chapter not found", 404)
		return
	}
	tests := c.tests()
	writeJSON(w, map[string]any{
		"lesson":      c.Lesson,
		"chapter":     c.Chapter,
		"title":       c.Title,
		"priority":    c.Priority,
		"text":        c.Text,
		"sample":      c.Sample,
		"description": c.Exercise.Description,
		"starter":     c.Exercise.Starter,
		"stdin":       tests[0].Stdin, // 「実行」の初期入力
		"testCount":   len(tests),
		"savedCode":   loadWork(lesson, ch),
		"status":      progress.get(progressKey(lesson, ch)),
	})
}

func handleAnswer(w http.ResponseWriter, r *http.Request) {
	lesson, ch, ok := pathInts(r)
	if !ok {
		http.Error(w, "invalid id", 400)
		return
	}
	c, err := loadChapter(lesson, ch)
	if err != nil {
		http.Error(w, "chapter not found", 404)
		return
	}
	writeJSON(w, map[string]string{"answer": c.Exercise.Answer})
}

type runReq struct {
	Lesson  int    `json:"lesson"`
	Chapter int    `json:"chapter"`
	Code    string `json:"code"`
	Stdin   string `json:"stdin"`
}

func handleSave(w http.ResponseWriter, r *http.Request) {
	var req runReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	saveWork(req.Lesson, req.Chapter, req.Code)
	writeJSON(w, map[string]bool{"ok": true})
}

func handleRun(w http.ResponseWriter, r *http.Request) {
	var req runReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	saveWork(req.Lesson, req.Chapter, req.Code)
	stdout, stderr, timedOut, err := runGo(req.Code, req.Stdin)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	progress.set(progressKey(req.Lesson, req.Chapter), "ran")
	writeJSON(w, map[string]any{"stdout": stdout, "stderr": stderr, "timedOut": timedOut})
}

func handleJudge(w http.ResponseWriter, r *http.Request) {
	var req runReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	c, err := loadChapter(req.Lesson, req.Chapter)
	if err != nil {
		http.Error(w, "chapter not found", 404)
		return
	}
	saveWork(req.Lesson, req.Chapter, req.Code)

	type caseResult struct {
		Stdin    string `json:"stdin"`
		Expected string `json:"expected"`
		Stdout   string `json:"stdout"`
		Stderr   string `json:"stderr"`
		TimedOut bool   `json:"timedOut"`
		Pass     bool   `json:"pass"`
	}
	var results []caseResult
	allPass := true
	for _, t := range c.tests() {
		stdout, stderr, timedOut, err := runGo(req.Code, t.Stdin)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		pass := !timedOut && stderr == "" && normalize(stdout) == normalize(t.Expected)
		allPass = allPass && pass
		results = append(results, caseResult{
			Stdin: t.Stdin, Expected: t.Expected,
			Stdout: stdout, Stderr: stderr, TimedOut: timedOut, Pass: pass,
		})
		if stderr != "" && !pass {
			// コンパイルエラーは全ケース同じ結果になるので 1 ケースで打ち切る
			break
		}
	}
	if allPass {
		progress.set(progressKey(req.Lesson, req.Chapter), "passed")
	} else {
		progress.set(progressKey(req.Lesson, req.Chapter), "ran")
	}
	writeJSON(w, map[string]any{"pass": allPass, "results": results})
}

// ---- 起動 ----

func detectRoot() string {
	for _, dir := range []string{".", ".."} {
		if _, err := os.Stat(filepath.Join(dir, "course.json")); err == nil {
			abs, _ := filepath.Abs(dir)
			return abs
		}
	}
	abs, _ := filepath.Abs(".")
	return abs
}

var (
	claudeAvailable bool
	enableAI        bool
)

func handleConfig(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, map[string]any{"claudeAvailable": claudeAvailable})
}

func main() {
	flag.StringVar(&courseRoot, "root", "", "講座ルート（lessonNN の親ディレクトリ。省略時は自動検出）")
	flag.StringVar(&claudeModel, "claude-model", "sonnet", "質問回答に使う Claude モデル（claude CLI の --model に渡す）")
	addr := flag.String("addr", "127.0.0.1:8080", "待ち受けアドレス（127.0.0.1 のみ推奨）")
	flag.BoolVar(&enableAI, "enable-ai", false, "AI質問機能を有効にする（要 Claude Code CLI。既定は無効）")
	flag.Parse()

	if enableAI {
		if _, err := exec.LookPath("claude"); err == nil {
			claudeAvailable = true
		} else {
			log.Println("警告: -enable-ai が指定されましたが claude CLI が見つかりません。AI質問機能は無効のままです。")
		}
	}

	if courseRoot == "" {
		courseRoot = detectRoot()
	}
	appDir = filepath.Join(courseRoot, "app")
	progress = newProgressStore(filepath.Join(appDir, "progress.json"))
	questions = newQAStore(filepath.Join(appDir, "questions.json"))

	static, _ := fs.Sub(staticFS, "static")
	http.Handle("/", http.FileServer(http.FS(static)))
	http.HandleFunc("GET /api/config", handleConfig)
	http.HandleFunc("GET /api/course", handleCourse)
	http.HandleFunc("GET /api/chapter/{lesson}/{ch}", handleChapter)
	http.HandleFunc("GET /api/answer/{lesson}/{ch}", handleAnswer)
	http.HandleFunc("POST /api/save", handleSave)
	http.HandleFunc("POST /api/run", handleRun)
	http.HandleFunc("POST /api/judge", handleJudge)
	http.HandleFunc("POST /api/ask", handleAsk)
	http.HandleFunc("GET /api/qa", handleQAList)
	http.HandleFunc("DELETE /api/qa/{id}", handleQADelete)
	http.HandleFunc("GET /api/qa/summary", handleSummaryGet)
	http.HandleFunc("POST /api/qa/summary", handleSummaryCreate)
	http.HandleFunc("POST /api/qa/review", handleReviewCreate)

	log.Printf("Go入門編を起動しました: http://%s （講座ルート: %s）", *addr, courseRoot)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
