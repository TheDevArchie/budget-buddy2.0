package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"budget_buddy/budget/utils"
)

var listCmd = &cobra.Command{
    Use: "list",
    Short: "List items",
    Run: func(cmd *cobra.Command, args []string) {
        user, _ := cmd.Flags().GetString("user")
        currentMonth, _ := cmd.Flags().GetInt("month")
        currentYear, _ := cmd.Flags().GetInt("year")

        userDirectory := fmt.Sprintf("%s/%s", utils.DataFilesDirectory, user)

        fileName, err := utils.GetExpenseFileWithMonthYear(userDirectory,currentMonth, currentYear)

        if err != nil {
            panic("No Filename found")
        }

        if !utils.DataFileExists(userDirectory) || !utils.DataFileExists(fileName) {
            panic("Data directory doesn't exist! Pls run the \"setup\" command first to set up user data")
        }

        expenses, _ := utils.GatherExpensesFromFile(fileName)

        utils.PrintExpenses(expenses.GetExpenseList())

    },
}

func init() {
    currentDate := time.Now()
    currentMonth := int(currentDate.Month())
    currentYear := int(currentDate.Year())

    listCmd.Flags().StringP("user", "u", "", "User associated with expenses")
    listCmd.Flags().IntP("month", "m", currentMonth, "Which month of expenses to view")
    listCmd.Flags().IntP("year", "y", currentYear, "Which year of the month the expenses to view")

    listCmd.MarkFlagsRequiredTogether("month", "year")
}


//TODO: Fix output of expense list
