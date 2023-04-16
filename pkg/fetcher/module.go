package fetcher

import "go.uber.org/fx"

func Module() fx.Option {
  return fx.Module("fetcher", fx.Provide(ParseResult))
}
