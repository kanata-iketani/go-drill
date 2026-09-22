// Go入門編 学習アプリ フロントエンド
// CodeMirror 6 を CDN (esm.sh) から読み込み、失敗時は textarea にフォールバックする。

const $ = (id) => document.getElementById(id);

// ---- 状態 ----
let course = { lessons: [] };
let cur = { lesson: 0, chapter: 0 }; // 現在表示中
let curData = null;                  // /api/chapter のレスポンス
let editor = null;                   // {get, set, focus}
let saveTimer = null;

// ---- テーマ ----
function initTheme() {
  const saved = localStorage.getItem('godrill-theme');
  const dark = saved ? saved === 'dark' : matchMedia('(prefers-color-scheme: dark)').matches;
  document.body.classList.toggle('dark', dark);
  $('themeToggle').onclick = () => {
    const nowDark = document.body.classList.toggle('dark');
    localStorage.setItem('godrill-theme', nowDark ? 'dark' : 'light');
  };
}

// ---- 簡易 Markdown レンダラ ----
function esc(s) {
  return s.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
}
function inlineMd(s) {
  s = esc(s);
  s = s.replace(/`([^`]+)`/g, '<code>$1</code>');
  s = s.replace(/\*\*([^*]+)\*\*/g, '<strong>$1</strong>');
  return s;
}
function mdToHtml(md) {
  const lines = (md || '').replace(/\r\n/g, '\n').split('\n');
  let html = '', i = 0, list = null;
  const closeList = () => { if (list) { html += `</${list}>`; list = null; } };
  while (i < lines.length) {
    const line = lines[i];
    if (/^```/.test(line)) {
      closeList();
      const code = [];
      i++;
      while (i < lines.length && !/^```/.test(lines[i])) { code.push(lines[i]); i++; }
      i++;
      html += `<pre class="code"><code>${esc(code.join('\n'))}</code></pre>`;
      continue;
    }
    const h = line.match(/^(#{1,4})\s+(.*)/);
    if (h) { closeList(); const lv = Math.min(h[1].length + 2, 6); html += `<h${lv}>${inlineMd(h[2])}</h${lv}>`; i++; continue; }
    const ul = line.match(/^[-*]\s+(.*)/);
    if (ul) { if (list !== 'ul') { closeList(); html += '<ul>'; list = 'ul'; } html += `<li>${inlineMd(ul[1])}</li>`; i++; continue; }
    const ol = line.match(/^\d+\.\s+(.*)/);
    if (ol) { if (list !== 'ol') { closeList(); html += '<ol>'; list = 'ol'; } html += `<li>${inlineMd(ol[1])}</li>`; i++; continue; }
    if (line.trim() === '') { closeList(); i++; continue; }
    const para = [line];
    i++;
    while (i < lines.length && lines[i].trim() !== '' && !/^(#{1,4}\s|```|[-*]\s|\d+\.\s)/.test(lines[i])) {
      para.push(lines[i]); i++;
    }
    closeList();
    html += `<p>${para.map(inlineMd).join('<br>')}</p>`;
  }
  closeList();
  return html;
}

// ---- エディタ（CodeMirror 6 / textarea フォールバック）----
async function initEditor() {
  const host = $('editorHost');
  try {
    const [cm, langGo, view, commands, theme] = await Promise.all([
      import('https://esm.sh/codemirror@6.0.1'),
      import('https://esm.sh/@codemirror/lang-go@6.0.1'),
      import('https://esm.sh/@codemirror/view@6.34.1'),
      import('https://esm.sh/@codemirror/commands@6.7.1'),
      import('https://esm.sh/@codemirror/theme-one-dark@6.1.2'),
    ]);
    const runKey = view.keymap.of([
      { key: 'Ctrl-Enter', run: () => { runCode(); return true; }, preventDefault: true },
      { key: 'Cmd-Enter', run: () => { runCode(); return true; }, preventDefault: true },
    ]);
    const updateListener = cm.EditorView.updateListener.of((u) => {
      if (u.docChanged) scheduleSave();
    });
    const ev = new cm.EditorView({
      doc: '',
      parent: host,
      extensions: [
        runKey, // basicSetup より先に置いて Enter 系のキーを先取りする
        cm.basicSetup,
        langGo.go(),
        view.keymap.of([commands.indentWithTab]),
        theme.oneDark,
        cm.EditorView.theme({ '&': { height: '100%' } }),
        updateListener,
      ],
    });
    editor = {
      get: () => ev.state.doc.toString(),
      set: (v) => ev.dispatch({ changes: { from: 0, to: ev.state.doc.length, insert: v } }),
      focus: () => ev.focus(),
    };
  } catch (e) {
    console.warn('CodeMirror の読み込みに失敗したため textarea を使います:', e);
    const ta = document.createElement('textarea');
    ta.id = 'fallbackEditor';
    ta.spellcheck = false;
    ta.addEventListener('keydown', (ev) => {
      if (ev.key === 'Tab') {
        ev.preventDefault();
        const { selectionStart: s, selectionEnd: en } = ta;
        ta.value = ta.value.slice(0, s) + '\t' + ta.value.slice(en);
        ta.selectionStart = ta.selectionEnd = s + 1;
        scheduleSave();
      }
      if ((ev.ctrlKey || ev.metaKey) && ev.key === 'Enter') {
        ev.preventDefault();
        runCode();
      }
    });
    ta.addEventListener('input', scheduleSave);
    host.appendChild(ta);
    editor = {
      get: () => ta.value,
      set: (v) => { ta.value = v; },
      focus: () => ta.focus(),
    };
  }
}

function scheduleSave() {
  clearTimeout(saveTimer);
  saveTimer = setTimeout(() => {
    if (!curData) return;
    fetch('/api/save', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ lesson: cur.lesson, chapter: cur.chapter, code: editor.get() }),
    }).catch(() => {});
  }, 800);
}

// ---- コースナビゲーション ----
async function loadCourse() {
  course = await (await fetch('/api/course')).json();
  const sel = $('lessonSelect');
  sel.innerHTML = '';
  for (const l of course.lessons) {
    const passed = l.chapters.filter((c) => c.status === 'passed').length;
    const opt = document.createElement('option');
    opt.value = l.lesson;
    opt.textContent = `Lesson ${String(l.lesson).padStart(2, '0')} ${l.title || ''} (${passed}/${l.chapters.length})`;
    sel.appendChild(opt);
  }
  sel.onchange = () => {
    const l = course.lessons.find((x) => x.lesson === Number(sel.value));
    if (l) openChapter(l.lesson, l.chapters[0].chapter);
  };
  updateProgressBadge();
}

function lessonOf(n) {
  return course.lessons.find((x) => x.lesson === n);
}

function renderChapterTabs() {
  const l = lessonOf(cur.lesson);
  const nav = $('chapterTabs');
  nav.innerHTML = '';
  if (!l) return;
  for (const c of l.chapters) {
    const b = document.createElement('button');
    b.className = 'chTab' + (c.chapter === cur.chapter ? ' active' : '') + (c.status === 'passed' ? ' passed' : '');
    b.textContent = `${String(c.chapter).padStart(2, '0')} ${c.title}`;
    b.onclick = () => openChapter(cur.lesson, c.chapter);
    nav.appendChild(b);
  }
}

function updateProgressBadge() {
  let total = 0, passed = 0;
  for (const l of course.lessons) {
    total += l.chapters.length;
    passed += l.chapters.filter((c) => c.status === 'passed').length;
  }
  $('progressBadge').textContent = `合格 ${passed}/${total}`;
}

async function openChapter(lesson, chapter) {
  const res = await fetch(`/api/chapter/${lesson}/${chapter}`);
  if (!res.ok) return;
  curData = await res.json();
  cur = { lesson, chapter };
  location.hash = `#${lesson}/${chapter}`;

  $('lessonSelect').value = lesson;
  renderChapterTabs();

  const l = lessonOf(lesson);
  $('chTitle').textContent =
    `${String(lesson).padStart(2, '0')}-${String(chapter).padStart(2, '0')} ${curData.title} ${curData.priority || ''}`;
  $('chText').innerHTML = mdToHtml(curData.text);
  $('chSample').textContent = curData.sample;
  $('chExercise').innerHTML = mdToHtml(curData.description);
  $('answerBox').classList.add('hidden');
  $('answerCode').textContent = '';
  $('stdinBox').value = curData.stdin || '';
  $('outputBox').textContent = '';
  $('outputBox').classList.remove('err');
  $('judgeResult').innerHTML = '';

  editor.set(curData.savedCode || curData.starter || defaultStarter());
  $('lessonPane').scrollTop = 0;
  qaSetContext();
}

function defaultStarter() {
  return 'package main\n\nimport "fmt"\n\nfunc main() {\n\t// ここに書く\n\t_ = fmt.Println\n}\n';
}

// ---- 実行・採点 ----
async function runCode() {
  const btn = $('runBtn');
  btn.disabled = true;
  $('outputBox').textContent = '実行中...';
  $('outputBox').classList.remove('err');
  try {
    const res = await (await fetch('/api/run', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ lesson: cur.lesson, chapter: cur.chapter, code: editor.get(), stdin: $('stdinBox').value }),
    })).json();
    let out = res.stdout || '';
    if (res.stderr) {
      out += (out ? '\n' : '') + res.stderr;
      $('outputBox').classList.add('err');
    }
    if (res.timedOut) out += '\n（タイムアウト: 10秒以内に終了しませんでした）';
    $('outputBox').textContent = out || '（出力なし）';
    markLocalStatus('ran');
  } catch (e) {
    $('outputBox').textContent = '実行に失敗しました: ' + e;
    $('outputBox').classList.add('err');
  } finally {
    btn.disabled = false;
  }
}

async function judgeCode() {
  const btn = $('judgeBtn');
  btn.disabled = true;
  $('judgeResult').innerHTML = '<div class="banner">採点中...</div>';
  try {
    const res = await (await fetch('/api/judge', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ lesson: cur.lesson, chapter: cur.chapter, code: editor.get() }),
    })).json();
    renderJudge(res);
    markLocalStatus(res.pass ? 'passed' : 'ran');
  } catch (e) {
    $('judgeResult').innerHTML = `<div class="banner fail">採点に失敗しました: ${esc(String(e))}</div>`;
  } finally {
    btn.disabled = false;
  }
}

function renderJudge(res) {
  let html = res.pass
    ? '<div class="banner pass">✅ 合格です！おめでとうございます</div>'
    : '<div class="banner fail">❌ 不合格です。期待出力と比べてみましょう</div>';
  res.results.forEach((r, idx) => {
    const head = `テストケース ${idx + 1}: ${r.pass ? '✅ 一致' : r.timedOut ? '⏱ タイムアウト' : '❌ 不一致'}`;
    let body = '';
    if (r.stderr) {
      body = `<div class="caseCol"><div class="lbl">エラー出力（stderr）</div><pre class="stderrBox">${esc(r.stderr)}</pre></div>`;
    } else if (!r.pass) {
      body =
        `<div class="caseCol"><div class="lbl">期待した出力</div><pre>${esc(r.expected)}</pre></div>` +
        `<div class="caseCol"><div class="lbl">あなたの出力</div><pre>${esc(r.stdout || '（出力なし）')}</pre></div>`;
      if (r.stdin) body = `<div class="caseCol"><div class="lbl">入力</div><pre>${esc(r.stdin)}</pre></div>` + body;
    }
    html += `<div class="caseCard"><div class="caseHead">${head}</div>` +
      (body ? `<div class="caseBody">${body}</div>` : '') + '</div>';
  });
  $('judgeResult').innerHTML = html;
}

// ローカルの course 情報とタブ表示を更新（サーバの progress.json は API 側で更新済み）
function markLocalStatus(status) {
  const l = lessonOf(cur.lesson);
  if (!l) return;
  const c = l.chapters.find((x) => x.chapter === cur.chapter);
  if (c && c.status !== 'passed') c.status = status;
  renderChapterTabs();
  updateProgressBadge();
  // レッスンセレクトの (n/m) 表示も更新
  const sel = $('lessonSelect');
  const opt = [...sel.options].find((o) => Number(o.value) === cur.lesson);
  if (opt) {
    const passed = l.chapters.filter((x) => x.status === 'passed').length;
    opt.textContent = `Lesson ${String(l.lesson).padStart(2, '0')} ${l.title || ''} (${passed}/${l.chapters.length})`;
  }
}

async function showAnswer() {
  if (!confirm('まず自分で解いてみましたか？\n模範解答を表示します。')) return;
  const res = await fetch(`/api/answer/${cur.lesson}/${cur.chapter}`);
  if (!res.ok) return;
  const data = await res.json();
  $('answerCode').textContent = data.answer;
  $('answerBox').classList.remove('hidden');
  $('answerBox').scrollIntoView({ behavior: 'smooth' });
}

function nextChapter() {
  const l = lessonOf(cur.lesson);
  if (!l) return;
  const idx = l.chapters.findIndex((c) => c.chapter === cur.chapter);
  if (idx >= 0 && idx + 1 < l.chapters.length) {
    openChapter(cur.lesson, l.chapters[idx + 1].chapter);
    return;
  }
  const li = course.lessons.findIndex((x) => x.lesson === cur.lesson);
  if (li >= 0 && li + 1 < course.lessons.length) {
    const nl = course.lessons[li + 1];
    openChapter(nl.lesson, nl.chapters[0].chapter);
  } else {
    alert('最後のチャプターです。おつかれさまでした！');
  }
}

// ---- 質問対話（Claude）----
const qa = {
  open: false,
  history: [], // 表示中チャプターの会話ログ {who, text}
};

function qaSetContext() {
  if (!curData) return;
  $('qaContext').textContent =
    `文脈: lesson${String(cur.lesson).padStart(2, '0')} ch${String(cur.chapter).padStart(2, '0')} ${curData.title}`;
}

function qaAppend(who, text, asMd) {
  const div = document.createElement('div');
  div.className = 'qaMsg ' + who;
  const label = who === 'user' ? 'あなた' : 'Claude 先生';
  div.innerHTML = `<div class="who">${label}</div><div class="bubble"></div>`;
  const bubble = div.querySelector('.bubble');
  if (asMd) bubble.innerHTML = mdToHtml(text);
  else bubble.textContent = text;
  $('qaLog').appendChild(div);
  $('qaLog').scrollTop = $('qaLog').scrollHeight;
  return div;
}

async function qaAsk() {
  const q = $('qaInput').value.trim();
  if (!q || !curData) return;
  $('qaInput').value = '';
  $('qaSend').disabled = true;
  qaAppend('user', q, false);
  const thinking = document.createElement('div');
  thinking.className = 'qaThinking';
  thinking.textContent = 'Claude が考えています…';
  $('qaLog').appendChild(thinking);
  $('qaLog').scrollTop = $('qaLog').scrollHeight;
  try {
    const res = await fetch('/api/ask', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        lesson: cur.lesson,
        chapter: cur.chapter,
        question: q,
        code: $('qaAttachCode').checked ? editor.get() : '',
      }),
    });
    thinking.remove();
    if (!res.ok) {
      qaAppend('claude', 'エラー: ' + (await res.text()), false);
      return;
    }
    const data = await res.json();
    qaAppend('claude', data.answer, true);
  } catch (e) {
    thinking.remove();
    qaAppend('claude', '通信エラー: ' + e, false);
  } finally {
    $('qaSend').disabled = false;
  }
}

async function qaLoadNotes() {
  try {
    const [listRes, sumRes] = await Promise.all([fetch('/api/qa'), fetch('/api/qa/summary')]);
    const list = (await listRes.json()).items || [];
    const summary = (await sumRes.json()).summary || '';
    $('qaSummary').innerHTML = summary ? mdToHtml(summary) : '';
    const box = $('qaList');
    box.innerHTML = '';
    if (list.length === 0) {
      box.innerHTML = '<div class="qaThinking">まだ質問の記録がありません。「質問する」タブから Claude に質問してみましょう。</div>';
      return;
    }
    for (const item of [...list].reverse()) {
      const div = document.createElement('div');
      div.className = 'qaItem';
      div.innerHTML =
        `<div class="meta"><span>${esc(item.time)} · lesson${String(item.lesson).padStart(2, '0')} ch${String(item.chapter).padStart(2, '0')} ${esc(item.chapterTitle || '')}</span>` +
        `<button data-id="${esc(item.id)}">削除</button></div>` +
        `<div class="q">Q. ${esc(item.question)}</div>` +
        `<div class="a">${mdToHtml(item.answer)}</div>`;
      div.querySelector('button').onclick = async (ev) => {
        if (!confirm('この質問記録を削除しますか？')) return;
        await fetch('/api/qa/' + ev.target.dataset.id, { method: 'DELETE' });
        qaLoadNotes();
      };
      box.appendChild(div);
    }
  } catch (e) {
    $('qaNotesStatus').textContent = '読み込みに失敗しました: ' + e;
  }
}

async function qaSummarize() {
  $('qaSummaryBtn').disabled = true;
  $('qaNotesStatus').textContent = 'Claude が疑問点をまとめています…（1分ほどかかります）';
  try {
    const res = await fetch('/api/qa/summary', { method: 'POST' });
    if (!res.ok) {
      $('qaNotesStatus').textContent = 'エラー: ' + (await res.text());
      return;
    }
    const data = await res.json();
    $('qaSummary').innerHTML = mdToHtml(data.summary);
    $('qaNotesStatus').textContent = 'まとめを更新しました（app/notes.md に保存済み）';
  } catch (e) {
    $('qaNotesStatus').textContent = '通信エラー: ' + e;
  } finally {
    $('qaSummaryBtn').disabled = false;
  }
}

async function qaMakeReview() {
  if (!confirm('保存された質問をもとに、復習問題チャプターを生成します（数分かかることがあります）。よろしいですか？')) return;
  $('qaReviewBtn').disabled = true;
  $('qaNotesStatus').textContent = 'Claude が復習問題を作成・検証しています…（数分かかります）';
  try {
    const res = await fetch('/api/qa/review', { method: 'POST' });
    if (!res.ok) {
      $('qaNotesStatus').textContent = 'エラー: ' + (await res.text());
      return;
    }
    const data = await res.json();
    const n = (data.created || []).length;
    let msg = n > 0 ? `復習問題を ${n} 問作成しました。` : '作成できた問題がありませんでした。';
    if ((data.failed || []).length > 0) msg += ` (検証に失敗: ${data.failed.length}件)`;
    $('qaNotesStatus').textContent = msg;
    if (n > 0) {
      await loadCourse();
      if (confirm(msg + '\n「復習問題」レッスンを開きますか？')) {
        $('qaDrawer').classList.add('hidden');
        openChapter(data.created[0].lesson, data.created[0].chapter);
      }
    }
  } catch (e) {
    $('qaNotesStatus').textContent = '通信エラー: ' + e;
  } finally {
    $('qaReviewBtn').disabled = false;
  }
}

function initQA() {
  $('qaToggle').onclick = () => {
    $('qaDrawer').classList.toggle('hidden');
    qaSetContext();
  };
  $('qaClose').onclick = () => $('qaDrawer').classList.add('hidden');
  document.querySelectorAll('.qaTab').forEach((tab) => {
    tab.onclick = () => {
      document.querySelectorAll('.qaTab').forEach((t) => t.classList.remove('active'));
      tab.classList.add('active');
      const isChat = tab.dataset.tab === 'chat';
      $('qaChatPane').classList.toggle('hidden', !isChat);
      $('qaNotesPane').classList.toggle('hidden', isChat);
      if (!isChat) qaLoadNotes();
    };
  });
  $('qaSend').onclick = qaAsk;
  $('qaInput').addEventListener('keydown', (ev) => {
    if ((ev.ctrlKey || ev.metaKey) && ev.key === 'Enter') {
      ev.preventDefault();
      qaAsk();
    }
  });
  $('qaSummaryBtn').onclick = qaSummarize;
  $('qaReviewBtn').onclick = qaMakeReview;
}

// ---- 起動 ----
async function main() {
  initTheme();
  initQA();
  await initEditor();
  await loadCourse();

  $('runBtn').onclick = runCode;
  $('judgeBtn').onclick = judgeCode;
  $('answerBtn').onclick = showAnswer;
  $('nextBtn').onclick = nextChapter;
  $('resetBtn').onclick = () => {
    if (confirm('エディタの内容をスターターコードに戻します。よろしいですか？')) {
      editor.set(curData?.starter || defaultStarter());
      scheduleSave();
    }
  };
  $('copySample').onclick = () => {
    if (confirm('サンプルコードをエディタにコピーします（現在の内容は上書き）。よろしいですか？')) {
      editor.set(curData.sample);
      scheduleSave();
      editor.focus();
    }
  };

  // Claude CLI がある環境でだけ「質問」機能を出す（読者環境では非表示）
  try {
    const cfg = await (await fetch('/api/config')).json();
    if (!cfg.claudeAvailable) {
      const qt = $('qaToggle'); if (qt) qt.style.display = 'none';
      const qd = $('qaDrawer'); if (qd) qd.classList.add('hidden');
    }
  } catch (_) {}

  // URL ハッシュ or 最初のチャプターを開く
  const m = location.hash.match(/^#(\d+)\/(\d+)$/);
  if (m && lessonOf(Number(m[1]))) {
    openChapter(Number(m[1]), Number(m[2]));
  } else if (course.lessons.length > 0) {
    const l = course.lessons[0];
    openChapter(l.lesson, l.chapters[0].chapter);
  }
}

main();
