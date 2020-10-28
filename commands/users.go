package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/barnbridge/barnbrdige-yieldfarm-cli/users"
)

var usersCmd = &cobra.Command{
	Use: "users",
	Run: func(cmd *cobra.Command, args []string) {
		u := users.New(users.Config{
			Tokens: getTokens(),
			Pool:   viper.GetString("pool"),
		})

		u.Run()
	},
}

func init() {
	RootCmd.AddCommand(usersCmd)

	usersCmd.Flags().String("graph-url", "https://api.thegraph.com/subgraphs/name/barnbridge/barnbridge-contracts", "URL of the grapql endpoint from TheGraph")
	viper.BindPFlag("graph-url", usersCmd.Flag("graph-url"))

	usersCmd.Flags().Int("top-n", 10, "How many users to show in top")
	viper.BindPFlag("top-n", usersCmd.Flag("top-n"))

	usersCmd.Flags().String("pool", "all", "The pool to show stats for [stable,lp,all]")
	viper.BindPFlag("pool", usersCmd.Flag("pool"))
}
