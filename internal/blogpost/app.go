package blogpost

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type focus int

const (
	focusTitle focus = iota
	focusTags
	focusContent
)

var (
	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("39")).
			MarginBottom(1)

	labelActiveStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("118"))

	labelIdleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("245"))

	boxActiveStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("118")).
			Padding(0, 1)

	boxIdleStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("238")).
			Padding(0, 1)

	suggestHeaderStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("245")).
				Italic(true)

	suggestActiveStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("229")).
				Background(lipgloss.Color("236")).
				PaddingLeft(1)

	suggestIdleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("220")).
				PaddingLeft(1)

	statusBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("245")).
			Background(lipgloss.Color("235")).
			Padding(0, 1)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))
)

type App struct {
	repoRoot     string
	defaultDraft bool
	allTags      []string
}

func NewApp(repoRoot string, defaultDraft bool) (*App, error) {
	tags, err := LoadExistingTags(repoRoot)
	if err != nil {
		return nil, fmt.Errorf("wczytywanie tagów: %w", err)
	}
	return &App{repoRoot: repoRoot, defaultDraft: defaultDraft, allTags: tags}, nil
}

func (a *App) Run() error {
	m := a.newModel()
	p := tea.NewProgram(m, tea.WithAltScreen())
	final, err := p.Run()
	if err != nil {
		return err
	}
	if result, ok := final.(model); ok && result.publishedPath != "" {
		fmt.Printf("Opublikowano: %s\n", result.publishedPath)
	}
	return nil
}

type model struct {
	repoRoot     string
	defaultDraft bool
	allTags      []string

	active   focus
	title    textinput.Model
	tags     textinput.Model
	content  textarea.Model
	status   string
	width    int
	height   int

	suggestions   []string
	suggestionIdx int

	publishedPath string
}

func (a *App) newModel() model {
	title := textinput.New()
	title.Placeholder = "np. Backup Schrödingera"
	title.Prompt = ""
	title.CharLimit = 240
	title.Width = 72
	title.Focus()

	tags := textinput.New()
	tags.Placeholder = "tag1, tag2 — podpowiedzi poniżej"
	tags.Prompt = ""
	tags.CharLimit = 512
	tags.Width = 72

	content := textarea.New()
	content.Placeholder = "Treść w Markdown…"
	content.SetWidth(72)
	content.SetHeight(12)
	content.ShowLineNumbers = false
	content.CharLimit = 0

	return model{
		repoRoot:     a.repoRoot,
		defaultDraft: a.defaultDraft,
		allTags:      a.allTags,
		active:       focusTitle,
		title:        title,
		tags:         tags,
		content:      content,
		status:       "Tab · następne pole   Shift+Tab · poprzednie   Ctrl+S · zapis   Ctrl+P · GitHub   Ctrl+Q · wyjście",
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		inner := max(20, m.width-6)
		m.title.Width = inner
		m.tags.Width = inner
		m.content.SetWidth(inner)
		m.content.SetHeight(m.contentHeight())
		return m, nil

	case tea.KeyMsg:
		key := msg.String()

		switch key {
		case "ctrl+c", "ctrl+q":
			return m, tea.Quit

		case "ctrl+s":
			return m.save(m.defaultDraft)

		case "ctrl+p":
			return m.save(false)

		case "tab":
			if m.active == focusTags && len(m.suggestions) > 0 {
				return m.applySuggestion()
			}
			return m.focusNext()

		case "shift+tab":
			return m.focusPrev()

		case "f2":
			if m.active == focusTags && len(m.suggestions) > 0 {
				return m.applySuggestion()
			}
		}

		if m.active == focusTags {
			switch key {
			case "up", "ctrl+p":
				if len(m.suggestions) > 0 {
					m.suggestionIdx = max(0, m.suggestionIdx-1)
					return m, nil
				}
			case "down", "ctrl+n":
				if len(m.suggestions) > 0 {
					m.suggestionIdx = min(len(m.suggestions)-1, m.suggestionIdx+1)
					return m, nil
				}
			case "enter":
				if len(m.suggestions) > 0 {
					return m.applySuggestion()
				}
				return m.focusNext()
			}
		}
	}

	var cmd tea.Cmd
	switch m.active {
	case focusTitle:
		m.title, cmd = m.title.Update(msg)
	case focusTags:
		m.tags, cmd = m.tags.Update(msg)
		m = m.refreshSuggestions()
	case focusContent:
		m.content, cmd = m.content.Update(msg)
	}
	return m, cmd
}

func (m model) View() string {
	if m.width == 0 {
		return "Ładowanie…\n"
	}

	var b strings.Builder
	b.WriteString(headerStyle.Render("✎ blogpost — edytor wpisów"))
	b.WriteString("\n\n")

	b.WriteString(m.renderField("Tytuł", focusTitle, m.title.View()))
	b.WriteString("\n")
	b.WriteString(m.renderField("Tagi", focusTags, m.tags.View()))
	b.WriteString("\n")

	if m.active == focusTags && len(m.suggestions) > 0 {
		b.WriteString(m.renderSuggestions())
		b.WriteString("\n")
	}

	b.WriteString(m.renderField("Treść", focusContent, m.content.View()))
	b.WriteString("\n\n")
	b.WriteString(helpStyle.Render("Tab/Shift+Tab · pola   ↑↓/F2/Enter · tag   Ctrl+S · zapis lokalny   Ctrl+P · wyślij na GitHub"))
	b.WriteString("\n")
	b.WriteString(statusBarStyle.Width(max(20, m.width-2)).Render(m.status))

	return b.String()
}

func (m model) renderField(label string, f focus, body string) string {
	labelStyle := labelIdleStyle
	boxStyle := boxIdleStyle
	prefix := "  "
	if m.active == f {
		labelStyle = labelActiveStyle
		boxStyle = boxActiveStyle
		prefix = "▸ "
	}
	return labelStyle.Render(prefix+label) + "\n" + boxStyle.Width(max(20, m.width-4)).Render(body)
}

func (m model) renderSuggestions() string {
	maxShow := min(6, len(m.suggestions))
	var lines []string
	lines = append(lines, suggestHeaderStyle.Render("  Podpowiedzi tagów"))
	for i := 0; i < maxShow; i++ {
		tag := m.suggestions[i]
		if i == m.suggestionIdx {
			lines = append(lines, suggestActiveStyle.Render("▸ "+tag))
		} else {
			lines = append(lines, suggestIdleStyle.Render("  "+tag))
		}
	}
	if len(m.suggestions) > maxShow {
		lines = append(lines, suggestIdleStyle.Render(fmt.Sprintf("  … +%d więcej", len(m.suggestions)-maxShow)))
	}
	return strings.Join(lines, "\n")
}

func (m model) contentHeight() int {
	reserved := 18
	if m.active == focusTags && len(m.suggestions) > 0 {
		reserved += min(6, len(m.suggestions)) + 2
	}
	h := m.height - reserved
	return max(6, h)
}

func (m model) focusNext() (model, tea.Cmd) {
	m = m.blurAll()
	switch m.active {
	case focusTitle:
		m.active = focusTags
		m.tags.Focus()
		m = m.refreshSuggestions()
	case focusTags:
		m.active = focusContent
		m.content.Focus()
	case focusContent:
		m.active = focusTitle
		m.title.Focus()
	}
	m.content.SetHeight(m.contentHeight())
	m.status = fmt.Sprintf("Aktywne pole: %s", m.fieldName())
	return m, textinput.Blink
}

func (m model) focusPrev() (model, tea.Cmd) {
	m = m.blurAll()
	switch m.active {
	case focusTitle:
		m.active = focusContent
		m.content.Focus()
	case focusTags:
		m.active = focusTitle
		m.title.Focus()
	case focusContent:
		m.active = focusTags
		m.tags.Focus()
		m = m.refreshSuggestions()
	}
	m.content.SetHeight(m.contentHeight())
	m.status = fmt.Sprintf("Aktywne pole: %s", m.fieldName())
	return m, textinput.Blink
}

func (m model) blurAll() model {
	m.title.Blur()
	m.tags.Blur()
	m.content.Blur()
	return m
}

func (m model) fieldName() string {
	switch m.active {
	case focusTitle:
		return "Tytuł"
	case focusTags:
		return "Tagi"
	default:
		return "Treść"
	}
}

func (m model) refreshSuggestions() model {
	selected := SelectedTagSet(ParseSelectedTags(m.tags.Value()))
	m.suggestions = FilterTags(m.allTags, CurrentTagFragment(m.tags.Value()), selected)
	if m.suggestionIdx >= len(m.suggestions) {
		m.suggestionIdx = 0
	}
	return m
}

func (m model) applySuggestion() (model, tea.Cmd) {
	if len(m.suggestions) == 0 {
		return m, nil
	}
	if m.suggestionIdx >= len(m.suggestions) {
		m.suggestionIdx = 0
	}
	tag := m.suggestions[m.suggestionIdx]
	m.tags.SetValue(ApplyTagSuggestion(m.tags.Value(), tag))
	m.suggestionIdx = 0
	m = m.refreshSuggestions()
	m.status = fmt.Sprintf("Dodano tag: %s", tag)
	return m, nil
}

func (m model) collectPost(draft bool) (Post, error) {
	title := strings.TrimSpace(m.title.Value())
	if title == "" {
		return Post{}, fmt.Errorf("tytuł jest wymagany")
	}
	return Post{
		Title:   title,
		Content: m.content.Value(),
		Tags:    ParseSelectedTags(m.tags.Value()),
		Draft:   draft,
	}, nil
}

func (m model) save(draft bool) (tea.Model, tea.Cmd) {
	post, err := m.collectPost(draft)
	if err != nil {
		m.status = "Błąd: " + err.Error()
		return m, nil
	}

	path, err := WritePost(m.repoRoot, post)
	if err != nil {
		m.status = "Błąd zapisu: " + err.Error()
		return m, nil
	}

	if draft {
		m.status = "Zapisano lokalnie: " + path
		return m, nil
	}

	m.status = "Wysyłanie na GitHub…"
	if err := CommitAndPush(m.repoRoot, path, post.Title); err != nil {
		m.status = "Błąd git: " + err.Error()
		return m, nil
	}

	m.publishedPath = path
	return m, tea.Quit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
