package list

import (
	"context"

	"github.com/loft-sh/api/v4/pkg/product"
	"github.com/loft-sh/log"
	"github.com/loft-sh/log/table"
	"github.com/loft-sh/vcluster/pkg/cli/config"
	"github.com/loft-sh/vcluster/pkg/cli/flags"
	"github.com/loft-sh/vcluster/pkg/platform"
	"github.com/loft-sh/vcluster/pkg/platform/clihelper"
	"github.com/loft-sh/vcluster/pkg/platform/helper"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// TeamsCmd holds the cmd flags
type TeamsCmd struct {
	*flags.GlobalFlags

	log log.Logger
	cfg *config.CLI
}

func newTeamsCmd(globalFlags *flags.GlobalFlags, cfg *config.CLI) *cobra.Command {
	cmd := &TeamsCmd{
		GlobalFlags: globalFlags,
		log:         log.GetInstance(),
		cfg:         cfg,
	}
	description := product.ReplaceWithHeader("list teams", `
List the vCluster platform teams you are a member of

Example:
vcluster platform list teams
########################################################
	`)
	clustersCmd := &cobra.Command{
		Use:   "teams",
		Short: product.Replace("Lists the vCluster platform teams you are member of"),
		Long:  description,
		Args:  cobra.NoArgs,
		RunE: func(cobraCmd *cobra.Command, _ []string) error {
			return cmd.Run(cobraCmd.Context())
		},
	}

	return clustersCmd
}

func (cmd *TeamsCmd) Run(ctx context.Context) error {
	platformClient, err := platform.NewClientFromConfig(ctx, cmd.cfg)
	if err != nil {
		return err
	}

	client, err := platformClient.Management()
	if err != nil {
		return err
	}

	userName, teamName, err := helper.GetCurrentUser(ctx, client)
	if err != nil {
		return err
	} else if teamName != nil {
		return errors.New("logged in as a team")
	}

	header := []string{
		"Name",
	}
	values := [][]string{}
	for _, team := range userName.Teams {
		values = append(values, []string{
			clihelper.DisplayName(team),
		})
	}

	table.PrintTable(cmd.log, header, values)
	return nil
}
