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

// func Setup(user string) string{
//     userDataFilesDirectory := fmt.Sprintf("%s/%s", utils.DataFilesDirectory, user)
//     utils.CreateDataDirectory(userDataFilesDirectory)
//     userFile := utils.CreateCurrentMonthFile(userDataFilesDirectory)
//
//     return userFile
// }
//TODO: Calculate remaining total amount of money needed for each expense category
//TODO: Calculate leftover money after expenses with income .... Add args for paycheck amount and for timeframe of paycheck
//TODO: Add arg for which user to apply budget to
    //Can include the user in the filename. OR The files are stored in each 'User' directory
    //inside data directory
//TODO: Once there is 'data/<user>/', alter expense filename into 'data/<user>/expenses/...txt.
//TODO: Then there can be an income folder as well
    //Then modify add command so that a user can add expense/income with needed args
//TODO: List out expense files of user... force a user arg
//TODO: Add setup command that is required initially
    //TODO: 'Add user' arg to setup command
