package commands

import (
	"github.com/spf13/cobra"

	"github.com/barnbridge/barnbrdige-yieldfarm-cli/epochs"
)

var epochsCmd = &cobra.Command{
	Use: "epochs",
	PreRun: func(cmd *cobra.Command, args []string) {
		bindViperToContractsFlags(cmd)
		bindViperToTokensFlags(cmd)
	},
	Run: func(cmd *cobra.Command, args []string) {
		epochs.Run(getTokens())
	},
}

func init() {
	RootCmd.AddCommand(epochsCmd)

	addContractsFlags(epochsCmd)
	addTokensFlags(epochsCmd)
}
