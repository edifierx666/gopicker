package output

import (
  "os/exec"

  "github.com/charmbracelet/bubbles/list"
  "github.com/charmbracelet/lipgloss"
  "github.com/edifierx666/gopicker/pkg/fetcher"

  tea "github.com/charmbracelet/bubbletea"
)

type it struct {
  *fetcher.Result
}

func (i it) FilterValue() string {
  return i.Name
}

func (i it) Title() string {
  return i.Name
}

func (i it) Description() string {
  return i.Result.Snippet
}

type model struct {
  list     list.Model
  chooice  *it
  showList bool
}

func (m *model) ShowList(b bool) {
  m.showList = b
}

func (m *model) Init() tea.Cmd {
  return nil
}

var docStyle = lipgloss.NewStyle()

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.KeyMsg:
    if msg.String() == "ctrl+c" {
      return m, tea.Quit
    }
    if msg.String() == "enter" {
      i, ok := m.list.SelectedItem().(*it)
      if ok {
        m.chooice = i
      }
      command := exec.Command("go", "get", m.chooice.PKGURL)
      process := tea.ExecProcess(command, func(err error) tea.Msg {
        return tea.Quit()
      })
      return m, process
    }
  case tea.WindowSizeMsg:
    h, v := docStyle.GetFrameSize()
    m.list.SetSize(msg.Width-h, msg.Height-v)
  case tea.MouseMsg:
    return m, nil
  }
  var cmd tea.Cmd

  m.list, cmd = m.list.Update(msg)

  return m, cmd
}

func (m *model) View() string {

  return docStyle.Render(m.list.View())
}
