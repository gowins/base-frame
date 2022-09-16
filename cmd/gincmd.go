package cmd

import (
	"os"
	"sync"

	"github.com/gowins/dionysus/ginx"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	WebServerAddr = "GAPI_ADDR"
	addrFlagName  = "addr"
)

var (
	defaultWebServerAddr = ":8080"
)

type ginCommand struct {
	g   *ginx.GinRouter
	cmd *cobra.Command

	addr string

	once sync.Once
}

func NewGinCommand() *ginCommand {
	return &ginCommand{
		g:   ginx.NewZeroGinRouter(),
		cmd: &cobra.Command{Use: "gin", Short: "Run as go-zero server"},
	}
}

func (t *ginCommand) GetCmd() *cobra.Command {
	t.once.Do(func() {
		if envAddr := os.Getenv(WebServerAddr); envAddr != "" {
			defaultWebServerAddr = envAddr
		}
		t.cmd.Flags().StringVarP(&t.addr, addrFlagName, "a", defaultWebServerAddr, "the http server address")
	})

	t.cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return t.g.Run(t.addr)
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

func (t *ginCommand) RegPreRunFunc(value string, f func() error) error {
	return nil
}

func (t *ginCommand) RegPostRunFunc(value string, f func() error) error {
	t.cmd.PostRunE = func(cmd *cobra.Command, args []string) error {
		return t.g.Shutdown()
	}
	return nil
}

func (t *ginCommand) GetEngine() ginx.ZeroGinRouter {
	return t.g
}
