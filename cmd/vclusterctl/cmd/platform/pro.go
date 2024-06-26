package platform

import (
	"github.com/loft-sh/vcluster/cmd/vclusterctl/cmd/platform/connect"
	"github.com/loft-sh/vcluster/pkg/cli/flags"
	"github.com/spf13/cobra"
)

func NewProCmd(globalFlags *flags.GlobalFlags) (*cobra.Command, error) {
	proCmd := &cobra.Command{
		Use:   "pro",
		Short: "vCluster platform subcommands",
		Long: `#######################################################
#################### vcluster pro #####################
#######################################################

Deprecated, please use vcluster platform instead
		`,
		Args: cobra.NoArgs,
	}

	startCmd := NewStartCmd(globalFlags)

	proCmd.AddCommand(startCmd)
	proCmd.AddCommand(NewResetCmd(globalFlags))
	proCmd.AddCommand(connect.NewConnectCmd(globalFlags))
	proCmd.AddCommand(NewAccessKeyCmd(globalFlags))

	return proCmd, nil
}
