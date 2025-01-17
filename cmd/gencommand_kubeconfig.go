package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/budimanjojo/talhelper/pkg/config"
	"github.com/budimanjojo/talhelper/pkg/generate"
)

var gencommandKubeconfigCmd = &cobra.Command{
	Use:   "kubeconfig",
	Short: "Generate talosctl kubeconfig commands.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadAndValidateFromFile(gencommandCfgFile, gencommandEnvFile)
		if err != nil {
			log.Fatalf("failed to parse config file: %s", err)
		}

		err = generate.GenerateKubeconfigCommand(cfg, gencommandOutDir, gencommandNode, gencommandExtraFlags)
		if err != nil {
			log.Fatalf("failed to generate talosctl kubeconfig command: %s", err)
		}
	},
}

func init() {
	gencommandCmd.AddCommand(gencommandKubeconfigCmd)
}
