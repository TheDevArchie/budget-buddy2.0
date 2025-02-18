package cmd

import (
    "fmt"
	"github.com/spf13/cobra"

	"budget_buddy/budget/utils"
)

var listCmd = &cobra.Command{
    Use: "list",
    Short: "List items",
    Run: func(cmd *cobra.Command, args []string) {
        user, _ := cmd.Flags().GetString("user")
        userDirectory := fmt.Sprintf("%s/%s", utils.DataFilesDirectory, user)
        fileName := fmt.Sprintf("%s/%s", userDirectory, utils.GetCurrentMonthFile())

        if !utils.DataFileExists(userDirectory) || !utils.DataFileExists(fileName) {
            panic("Data directory doesn't exist! Pls run the \"setup\" command first to set up user data")
        }
        expenses, _ := utils.GatherExpensesFromFile(fileName)

        utils.PrintExpenses(expenses.GetExpenseList())

    },
}

func init() {
    listCmd.Flags().StringP("user", "u", "", "User associated with expenses")
}
