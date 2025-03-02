package cmd

import (
	"budget_buddy/budget/utils"
	"fmt"
	// "os"

	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
    Use: "setup",
    Short: "Setup Budget Buddy",
    Run: func(cmd *cobra.Command, args []string) {
        //Create sqlite db with Users table
        //Create Data file directory
        //If add user, then create user directory

        user, _ := cmd.Flags().GetString("user")
        addUserFlag, _ := cmd.Flags().GetBool("adduser")

        if addUserFlag {
            SetupUser(user)
        }
    },
}

func init() {
    setupCmd.Flags().Bool("adduser", false, "Add user")
    setupCmd.Flags().StringP("user", "u", "", "User name")
    // setupCmd.Flags().StringP("filename", "f", "", "Name of file to process")

    // setupCmd.MarkFlagRequired("filename")
    setupCmd.MarkFlagsRequiredTogether("adduser", "user")
}

func SetupUser(user string) string{
    userDataFilesDirectory := fmt.Sprintf("%s/%s", utils.DataFilesDirectory, user)
    utils.CreateDataDirectory(userDataFilesDirectory)
    userFile := utils.CreateCurrentMonthFile(userDataFilesDirectory)

    return userFile
}
//TODO: Calculate remaining total amount of money needed for each expense category
//TODO: Calculate leftover money after expenses with income .... Add args for paycheck amount and for timeframe of paycheck
//TODO: Once there is 'data/<user>/', alter expense filename into 'data/<user>/expenses/...txt.
//TODO: Then there can be an income folder as well
    //Then modify add command so that a user can add expense/income with needed args
//TODO: List out expense files of user... force a user arg
    //TODO: 'Add user' arg to setup command
//TODO: Add notifications to setup commands to inform user how the process is going
