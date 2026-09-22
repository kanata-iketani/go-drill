// coursegen は lessonNN/chMM.json（教材の正）を検証し、
// README.md / chMM_sample.go / chMM_exercise.md / answer/chMM_answer.go を生成する。
//
// 使い方:
//
//	go run ./tools/coursegen              # 全レッスンを検証 + 生成
//	go run ./tools/coursegen -only 6-9    # lesson06〜09 のみ
//	go run ./tools/coursegen -no-gen      # 検証だけ行う
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Test struct {
	Stdin    string `json:"stdin"`
	Expected string `json:"expected_stdout"`
}

type Exercise struct {
	Description string `json:"description"`
	Starter     string `json:"starter"`
	Stdin       string `json:"stdin,omitempty"`
	Expected    string `json:"expected_stdout,omitempty"`
	Answer      string `json:"answer"`
	Tests       []Test `json:"tests,omitempty"`
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

var (
	root   = flag.String("root", ".", "講座ルート")
	only   = flag.String("only", "", "対象レッスン番号（例: 6-9 や 1,3,5）。空なら全部")
	noGen  = flag.Bool("no-gen", false, "検証のみ行い、ファイル生成をしない")
	noFix  = flag.Bool("no-fix", false, "gofmt 差分があっても JSON を書き換えない")
	failed bool
)

func main() {
	flag.Parse()
	targets := parseOnly(*only)

	titles := map[int]string{}
	if b, err := os.ReadFile(filepath.Join(*root, "course.json")); err == nil {
		var cl []courseLesson
		json.Unmarshal(b, &cl)
		for _, c := range cl {
			titles[c.Lesson] = c.Title
		}
	}

	type result struct {
		lesson, chapter int
		pass            bool
		errs            []string
	}
	var (
		mu      sync.Mutex
		results []result
		wg      sync.WaitGroup
		sem     = make(chan struct{}, 8)
	)

	chapters := map[int][]*Chapter{}
	// lesson90 以降はアプリが自動生成する復習レッスンなので対象外
	for n := 1; n <= 89; n++ {
		if targets != nil && !targets[n] {
			continue
		}
		dir := filepath.Join(*root, fmt.Sprintf("lesson%02d", n))
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		re := regexp.MustCompile(`^ch(\d{2})\.json$`)
		for _, e := range entries {
			m := re.FindStringSubmatch(e.Name())
			if m == nil {
				continue
			}
			path := filepath.Join(dir, e.Name())
			ch, errs := loadAndFix(path)
			if ch != nil {
				chapters[n] = append(chapters[n], ch)
			}
			if len(errs) > 0 {
				results = append(results, result{n, atoi(m[1]), false, errs})
				failed = true
				continue
			}
			wg.Add(1)
			go func(n int, ch *Chapter) {
				defer wg.Done()
				sem <- struct{}{}
				defer func() { <-sem }()
				errs := validateRun(ch)
				mu.Lock()
				results = append(results, result{n, ch.Chapter, len(errs) == 0, errs})
				if len(errs) > 0 {
					failed = true
				}
				mu.Unlock()
			}(n, ch)
		}
	}
	wg.Wait()

	// 生成
	if !*noGen && !failed {
		for n, chs := range chapters {
			sort.Slice(chs, func(i, j int) bool { return chs[i].Chapter < chs[j].Chapter })
			if err := generate(n, titles[n], chs); err != nil {
				fmt.Printf("生成 NG lesson%02d: %v\n", n, err)
				failed = true
			}
		}
	}

	// 集計表
	sort.Slice(results, func(i, j int) bool {
		if results[i].lesson != results[j].lesson {
			return results[i].lesson < results[j].lesson
		}
		return results[i].chapter < results[j].chapter
	})
	byLesson := map[int][2]int{} // [チャプター数, 合格数]
	for _, r := range results {
		v := byLesson[r.lesson]
		v[0]++
		if r.pass {
			v[1]++
		}
		byLesson[r.lesson] = v
		for _, e := range r.errs {
			fmt.Printf("NG lesson%02d/ch%02d: %s\n", r.lesson, r.chapter, e)
		}
	}
	var nums []int
	for n := range byLesson {
		nums = append(nums, n)
	}
	sort.Ints(nums)
	fmt.Println("| レッスン | チャプター数 | answer 合格数 |")
	fmt.Println("|---|---|---|")
	for _, n := range nums {
		v := byLesson[n]
		fmt.Printf("| lesson%02d %s | %d | %d |\n", n, titles[n], v[0], v[1])
	}
	if failed {
		os.Exit(1)
	}
	fmt.Println("ALL OK")
}

func atoi(s string) int { n, _ := strconv.Atoi(strings.TrimLeft(s, "0")); return n }

func parseOnly(s string) map[int]bool {
	if s == "" {
		return nil
	}
	m := map[int]bool{}
	for _, part := range strings.Split(s, ",") {
		if lo, hi, ok := strings.Cut(part, "-"); ok {
			a, _ := strconv.Atoi(strings.TrimSpace(lo))
			b, _ := strconv.Atoi(strings.TrimSpace(hi))
			for i := a; i <= b; i++ {
				m[i] = true
			}
		} else {
			a, _ := strconv.Atoi(strings.TrimSpace(part))
			m[a] = true
		}
	}
	return m
}

// loadAndFix は JSON を読み、静的な検証と gofmt 整形（自動修正）を行う。
func loadAndFix(path string) (*Chapter, []string) {
	var errs []string
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, []string{err.Error()}
	}
	var c Chapter
	if err := json.Unmarshal(b, &c); err != nil {
		return nil, []string{"JSON parse error: " + err.Error()}
	}
	if c.Title == "" || c.Text == "" || c.Sample == "" ||
		c.Exercise.Description == "" || c.Exercise.Starter == "" || c.Exercise.Answer == "" {
		errs = append(errs, "必須フィールドが空です (title/text/sample/description/starter/answer)")
	}
	switch c.Priority {
	case "🔴", "🟡", "⚪":
	default:
		errs = append(errs, "priority は 🔴/🟡/⚪ のいずれかにしてください")
	}
	if !strings.Contains(c.Text, "Python") {
		errs = append(errs, "text に Python との対比がありません")
	}
	if c.Lesson >= 15 && len(c.Exercise.Tests) < 3 {
		errs = append(errs, "lesson15 以降は tests を最低 3 ケース入れてください")
	}
	for i, t := range c.tests() {
		if strings.TrimSpace(t.Expected) == "" {
			errs = append(errs, fmt.Sprintf("テスト %d の expected_stdout が空です", i+1))
		}
	}

	// gofmt 自動整形（sample / answer / starter）
	changed := false
	for _, p := range []*string{&c.Sample, &c.Exercise.Answer, &c.Exercise.Starter} {
		fixed, err := format.Source([]byte(*p))
		if err != nil {
			errs = append(errs, "gofmt できないコードがあります: "+err.Error())
			continue
		}
		if string(fixed) != *p {
			*p = string(fixed)
			changed = true
		}
	}
	if changed && !*noFix && len(errs) == 0 {
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.SetEscapeHTML(false)
		enc.SetIndent("", "  ")
		enc.Encode(&c)
		os.WriteFile(path, buf.Bytes(), 0o644)
	}
	return &c, errs
}

// validateRun は sample / answer を実際に go vet + go run して検証する。
func validateRun(c *Chapter) []string {
	var errs []string
	tests := c.tests()

	if _, stderr, timedOut, err := execGo(c.Sample, tests[0].Stdin, true); err != nil || timedOut || stderr != "" {
		errs = append(errs, "sample が動きません: "+firstLine(stderr, err, timedOut))
	}
	for i, t := range tests {
		stdout, stderr, timedOut, err := execGo(c.Exercise.Answer, t.Stdin, i == 0)
		if err != nil || timedOut || stderr != "" {
			errs = append(errs, fmt.Sprintf("answer がテスト %d で動きません: %s", i+1, firstLine(stderr, err, timedOut)))
			break
		}
		if normalize(stdout) != normalize(t.Expected) {
			errs = append(errs, fmt.Sprintf("answer がテスト %d で不一致: got %q want %q",
				i+1, normalize(stdout), normalize(t.Expected)))
		}
	}
	return errs
}

func firstLine(stderr string, err error, timedOut bool) string {
	if timedOut {
		return "タイムアウト"
	}
	if stderr != "" {
		lines := strings.SplitN(strings.TrimSpace(stderr), "\n", 2)
		return lines[0]
	}
	if err != nil {
		return err.Error()
	}
	return "不明"
}

// execGo は一時モジュールで（vet が必要なら go vet も）go run する。
func execGo(code, stdin string, vet bool) (stdout, stderr string, timedOut bool, err error) {
	tmp, err := os.MkdirTemp("", "coursegen-*")
	if err != nil {
		return "", "", false, err
	}
	defer os.RemoveAll(tmp)
	os.WriteFile(filepath.Join(tmp, "go.mod"), []byte("module tmp\n\ngo 1.22\n"), 0o644)
	os.WriteFile(filepath.Join(tmp, "main.go"), []byte(code), 0o644)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if vet {
		cmd := exec.CommandContext(ctx, "go", "vet", ".")
		cmd.Dir = tmp
		var vetErr strings.Builder
		cmd.Stderr = &vetErr
		if err := cmd.Run(); err != nil {
			return "", "go vet NG: " + vetErr.String(), false, nil
		}
	}

	runCtx, cancelRun := context.WithTimeout(ctx, 10*time.Second)
	defer cancelRun()
	cmd := exec.CommandContext(runCtx, "go", "run", ".")
	cmd.Dir = tmp
	cmd.Stdin = strings.NewReader(stdin)
	var out, errBuf strings.Builder
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	cmd.Run()
	if runCtx.Err() == context.DeadlineExceeded {
		return out.String(), errBuf.String(), true, nil
	}
	return out.String(), errBuf.String(), false, nil
}

func normalize(s string) string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	lines := strings.Split(s, "\n")
	for i := range lines {
		lines[i] = strings.TrimRight(lines[i], " \t")
	}
	return strings.TrimRight(strings.Join(lines, "\n"), "\n")
}

// generate は JSON からレッスンの README / sample / exercise / answer を書き出す。
func generate(lesson int, title string, chs []*Chapter) error {
	dir := filepath.Join(*root, fmt.Sprintf("lesson%02d", lesson))

	var readme strings.Builder
	fmt.Fprintf(&readme, "# Lesson %02d %s\n", lesson, title)
	if intro, err := os.ReadFile(filepath.Join(dir, "intro.md")); err == nil {
		readme.WriteString("\n" + strings.TrimSpace(string(intro)) + "\n")
	}
	for _, c := range chs {
		fmt.Fprintf(&readme, "\n## ch%02d %s\n\n%s\n", c.Chapter, c.Title, strings.TrimSpace(c.Text))
	}
	if err := os.WriteFile(filepath.Join(dir, "README.md"), []byte(readme.String()), 0o644); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Join(dir, "answer"), 0o755); err != nil {
		return err
	}
	for _, c := range chs {
		chName := fmt.Sprintf("ch%02d", c.Chapter)
		if err := os.WriteFile(filepath.Join(dir, chName+"_sample.go"), []byte(c.Sample), 0o644); err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(dir, "answer", chName+"_answer.go"), []byte(c.Exercise.Answer), 0o644); err != nil {
			return err
		}

		var ex strings.Builder
		fmt.Fprintf(&ex, "# lesson%02d %s 演習: %s\n\n%s\n", lesson, chName, c.Title, strings.TrimSpace(c.Exercise.Description))
		if len(c.Exercise.Tests) > 0 {
			ex.WriteString("\n## テストケース\n")
			for i, t := range c.Exercise.Tests {
				fmt.Fprintf(&ex, "\n### ケース %d\n\n入力:\n\n```text\n%s\n```\n\n期待出力:\n\n```text\n%s\n```\n",
					i+1, strings.TrimRight(t.Stdin, "\n"), strings.TrimRight(t.Expected, "\n"))
			}
		}
		if err := os.WriteFile(filepath.Join(dir, chName+"_exercise.md"), []byte(ex.String()), 0o644); err != nil {
			return err
		}
	}
	return nil
}
