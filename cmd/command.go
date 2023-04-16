package cmd

import (
  "errors"
  "fmt"
  "os"

  "github.com/spf13/cobra"
  "go.uber.org/fx"
  cfg2 "gopicker/cfg"
  "gopicker/pkg/fetcher"
  "gopicker/pkg/output"
)

var cfg = &cfg2.Cfg{}
var rootCmd = &cobra.Command{
  Use:   "gopicker 包名",
  Short: "gopicker 快速获取pkg.go.dev上的包",
  Long:  `gopicker 快速获取pkg.go.dev上的包`,
  Args: func(cmd *cobra.Command, args []string) error {
    if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
      return errors.New("giegie 给个包名")
    }
    return nil
  },
  Version: "0.0.1",
  Run: func(cmd *cobra.Command, args []string) {
    cfg.Name = args[0]
    fx.New(fx.NopLogger, fx.Supply(cfg), fetcher.Module(), output.Module())
  },
}

func init() {
  rootCmd.PersistentFlags().IntVarP(&cfg.Limit, "limit", "l", 20, "")
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}
