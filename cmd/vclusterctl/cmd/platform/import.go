package platform

import (
	"context"

	"github.com/loft-sh/log"
	"github.com/loft-sh/vcluster/pkg/cli"
	"github.com/loft-sh/vcluster/pkg/cli/config"
	"github.com/loft-sh/vcluster/pkg/cli/flags"

	loftctlUtil "github.com/loft-sh/vcluster/pkg/platform/loftutils"
	"github.com/spf13/cobra"
)

type ImportCmd struct {
	*flags.GlobalFlags
	cli.ActivateOptions

	Log log.Logger
}

func NewImportCmd(globalFlags *flags.GlobalFlags) *cobra.Command {
	cmd := &ImportCmd{
		GlobalFlags: globalFlags,
		Log:         log.GetInstance(),
	}

	description := `########################################################
############### vcluster platform import ###############
########################################################
Imports a vCluster into a vCluster platform project.

Example:
vcluster platform import my-vcluster --cluster connected-cluster \
--namespace vcluster-my-vcluster --project my-project --import-name my-vcluster
#######################################################
	`

	importCmd := &cobra.Command{
		Use:   "import" + loftctlUtil.VClusterNameOnlyUseLine,
		Short: "Imports a vCluster into a vCluster platform project",
		Long:  description,
		Args:  loftctlUtil.VClusterNameOnlyValidator,
		RunE: func(cobraCmd *cobra.Command, args []string) error {
			return cmd.Run(cobraCmd.Context(), args)
		},
	}

	importCmd.Flags().StringVar(&cmd.Manager, "manager", "", "The manager to use for managing the virtual cluster, can be either helm or platform.")
	importCmd.Flags().StringVar(&cmd.ClusterName, "cluster", "", "Cluster name of the cluster the virtual cluster is running on")
	importCmd.Flags().StringVar(&cmd.Project, "project", "", "The project to import the vCluster into")
	importCmd.Flags().StringVar(&cmd.ImportName, "import-name", "", "The name of the vCluster under projects. If unspecified, will use the vcluster name")

	return importCmd
}

// Run executes the functionality
func (cmd *ImportCmd) Run(ctx context.Context, args []string) error {
	cfg := cmd.LoadedConfig(cmd.Log)
	// check if we should create a platform vCluster
	if cfg.Manager.Type == config.ManagerPlatform {
		return cli.ActivatePlatform(ctx, &cmd.ActivateOptions, cmd.GlobalFlags, args[0], cmd.Log)
	}

	return cli.ActivateHelm(ctx, &cmd.ActivateOptions, cmd.GlobalFlags, args[0], cmd.Log)
}
