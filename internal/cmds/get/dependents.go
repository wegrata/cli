package get

import (
	"fmt"

	"github.com/deps-cloud/api/v1alpha/tracker"
	"github.com/deps-cloud/cli/internal/writer"

	"github.com/spf13/cobra"
)

func DependentsCommand(
	dependencyClient tracker.DependencyServiceClient,
	writer writer.Writer,
) *cobra.Command {
	req := &tracker.DependencyRequest{}

	cmd := &cobra.Command{
		Use:     "dependents",
		Short:   "Get the list of modules that depend on the given module",
		Example: "depscloud-cli get dependents -l go -o github.com -m deps-cloud/api",
		RunE: func(cmd *cobra.Command, args []string) error {
			if req.Language == "" || req.Organization == "" || req.Module == "" {
				return fmt.Errorf("language, organization, and module must be provided")
			}

			ctx := cmd.Context()

			response, err := dependencyClient.ListDependents(ctx, req)
			if err != nil {
				return err
			}

			for _, dependent := range response.Dependents {
				_ = writer.Write(dependent)
			}

			return nil
		},
	}

	addDependencyRequestFlags(cmd, req)

	return cmd
}
