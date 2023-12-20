package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/xigxog/fox/internal/log"
	"github.com/xigxog/fox/internal/repo"
)

var releaseCmd = &cobra.Command{
	Use:    "release (NAME | COMMIT | SHORT COMMIT | VERSION | TAG | BRANCH)",
	Args:   cobra.ExactArgs(1),
	PreRun: setup,
	Run:    release,
	Short:  "Release specified AppDeployment and VirtualEnv",
	Long: strings.TrimSpace(`
The release command activates the routes of the components belonging to the 
specified AppDeployment. This causes genesis events matching components' routes
to be automatically sent to the component with the specified environment being 
injected.

The AppDeployment can be identified by its name, commit, short-commit (first 7 
characters), version, Git tag, or Git branch. 🦊 Fox will inspect the Kubernetes
cluster to find a matching AppDeployment. If more than one AppDeployment is
found you will be prompted to select the desired AppDeployment.
`),
	Example: strings.TrimSpace(`
# Release the AppDeployment named 'main' using the 'dev' Virtual Environment.
fox release main --virtual-env dev

# Release the AppDeployment with version 'v1.2.3' using the 'prod' 
# VirtualEnv, creating an VirtualEnvSnapshot if needed.
fox release v1.2.3 --virtual-env prod --create-snapshot
`),
}

func init() {
	releaseCmd.Flags().StringVarP(&cfg.Flags.VirtEnv, "virtual-env", "e", "", "name of VirtualEnv or VirtualEnvSnapshot to use")
	releaseCmd.Flags().BoolVarP(&cfg.Flags.CreateVirtEnv, "create-snapshot", "c", false, "create an immutable snapshot of environment and use for release")

	addCommonDeployFlags(releaseCmd)

	releaseCmd.MarkFlagRequired("virtual-env")

	rootCmd.AddCommand(releaseCmd)
}

func release(cmd *cobra.Command, args []string) {
	appDep := args[0]
	checkCommonDeployFlags(cfg.Flags.VirtEnv)

	env := repo.New(cfg).Release(appDep)

	// Makes output less cluttered.
	env.ManagedFields = nil
	log.Marshal(env)
}
