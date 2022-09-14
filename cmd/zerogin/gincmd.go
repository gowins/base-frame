package zerogin

import (
	"base-frame/cmd"
	"base-frame/ginhelper"
	"os"

	"github.com/spf13/pflag"

	"github.com/zeromicro/go-zero/rest"

	"github.com/spf13/cobra"
)

type ginCommand struct {
	g *ginhelper.GinRouter

	cmd *cobra.Command
}

func NewGinCommand(rc rest.RestConf) *ginCommand {
	return &ginCommand{
		g:   ginhelper.NewZeroGinRouter(rc),
		cmd: &cobra.Command{Use: "zero", Short: "Run as go-zero server"},
	}
}

func (t *ginCommand) GetCmd() *cobra.Command {
	t.cmd.Run = func(cmd *cobra.Command, args []string) {
		t.g.Run()
	}
	return t.cmd
}

func (t *ginCommand) Start() {
	if err := t.cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func (t *ginCommand) RegFlagSet(set *pflag.FlagSet) {
	t.cmd.Flags().AddFlagSet(set)
}

func (t *ginCommand) Flags() *pflag.FlagSet {
	return t.cmd.Flags()
}

func (t *ginCommand) RegPreRunFunc(value string, priority cmd.Priority, f func() error) error {
	t.cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		return f()
	}
	return nil
}

func (t *ginCommand) RegPostRunFunc(value string, priority cmd.Priority, f func() error) error {
	t.cmd.PostRun = func(cmd *cobra.Command, args []string) {
		t.g.Shutdown()
	}
	return nil
}

func (t *ginCommand) GetZeroGinRouter() ginhelper.ZeroGinRouter {
	return t.g
}
