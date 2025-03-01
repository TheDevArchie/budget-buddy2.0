package cmd

import (
	// "budget_buddy/budget/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use: "budget_buddy",
    Short: "budget_buddy is a cli tool for helping manage your budget",
    Run: func(cmd *cobra.Command, args []string) {
    },
}

func init() {
    rootCmd.AddCommand(setupCmd)
    rootCmd.AddCommand(listCmd)
    rootCmd.AddCommand(addCmd)
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Error occurred while Executing budget_buddy '%s'\n", err)
        os.Exit(1)
    }
}

//TODO: Calculate amount needed for income based on expenses
//TODO: Calculate leftover money after expenses with income .... Add args for paycheck amount and for timeframe of paycheck
//TODO: Once there is 'data/<user>/', alter expense filename into 'data/<user>/expenses/...txt.
//TODO: Then there can be an income folder as well
    //Then modify add command so that a user can add expense/income with needed args
//TODO: List out expense files of user... force a user arg
