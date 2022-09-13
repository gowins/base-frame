package zerocmd

import (
	"base-frame/cmd"
	"fmt"
	"github.com/spf13/pflag"
	"sync"

	"github.com/spf13/cobra"
)

type zeroCmd struct {
	cmd  *cobra.Command
	once sync.Once
}

func New() *zeroCmd {
	return &zeroCmd{
		cmd: &cobra.Command{Use: "zero", Short: "Run as go-zero server"},
	}
}

func (t *zeroCmd) GetCmd() *cobra.Command {
	t.cmd.Run = func(cmd *cobra.Command, args []string) {
		t.start()
	}
	return t.cmd
}

func (t *zeroCmd) start() {
	fmt.Println(1111111)
}

func (t *zeroCmd) RegFlagSet(set *pflag.FlagSet) {
	t.cmd.Flags().AddFlagSet(set)
}

func (t *zeroCmd) Flags() *pflag.FlagSet {
	return t.cmd.Flags()
}

func (t *zeroCmd) RegPreRunFunc(value string, priority cmd.Priority, f func() error) error {
	return nil
}

func (t *zeroCmd) RegPostRunFunc(value string, priority cmd.Priority, f func() error) error {
	return nil
}
