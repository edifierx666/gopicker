package output

import (
  "fmt"

  "github.com/charmbracelet/bubbles/list"
  tea "github.com/charmbracelet/bubbletea"

  "go.uber.org/fx"
  "gopicker/pkg/fetcher"
)

func Module() fx.Option {
  return fx.Module("output", fx.Invoke(func(res []*fetcher.Result) {
    var items []list.Item
    for _, re := range res {
      items = append(items, &it{Result: re})
    }
    initModel := &model{
      list:    list.New(items, list.NewDefaultDelegate(), 0, 2),
      chooice: nil,
    }

    p := tea.NewProgram(initModel)
    if err := p.Start(); err != nil {
      fmt.Println("Error running program:", err)
    }
    fmt.Println(initModel.chooice.Link)
  }))
}
