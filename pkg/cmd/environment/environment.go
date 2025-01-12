package environment

import (
	"github.com/MakeNowJust/heredoc/v2"
	cmdDelete "github.com/OctopusDeploy/cli/pkg/cmd/environment/delete"
	cmdList "github.com/OctopusDeploy/cli/pkg/cmd/environment/list"
	"github.com/OctopusDeploy/cli/pkg/constants"
	"github.com/OctopusDeploy/cli/pkg/constants/annotations"
	"github.com/OctopusDeploy/cli/pkg/factory"
	"github.com/spf13/cobra"
)

func NewCmdEnvironment(f factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "environment <command>",
		Short: "Manage environments",
		Long:  "Manage environments in Octopus Deploy",
		Example: heredoc.Docf(`
			$ %[1]s environment list
			$ %[1]s environment ls
		`, constants.ExecutableName),
		Annotations: map[string]string{
			annotations.IsInfrastructure: "true",
		},
	}

	cmd.AddCommand(cmdList.NewCmdList(f))
	cmd.AddCommand(cmdDelete.NewCmdDelete(f))
	return cmd
}
