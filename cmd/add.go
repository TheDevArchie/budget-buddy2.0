package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"budget_buddy/budget/models"
	"budget_buddy/budget/utils"
	// "budget_buddy/budget/utils"
)

var expenseName string
var expenseAmount float64
var user string

var addCmd = &cobra.Command{
    Use: "add",
    Short: "Add items",
    Run: func(cmd *cobra.Command, args []string) {
        user, _ := cmd.Flags().GetString("user")
        expenseAmount, _ := cmd.Flags().GetFloat64("amount")
        expenseName, _ := cmd.Flags().GetString("name")

        userDirectory := fmt.Sprintf("%s/%s", utils.DataFilesDirectory, user)
        fileName := fmt.Sprintf("%s/%s", userDirectory, utils.GetCurrentMonthFile())

        if !utils.DataFileExists(userDirectory) || !utils.DataFileExists(fileName) {
            panic("Data directory doesn't exist! Pls run the \"setup\" command first to set up user data")
        }
        expenses, _ := utils.GatherExpensesFromFile(fileName)

        addExpense(expenseName, expenseAmount, fileName, expenses)
    },
}

func init() {
    addCmd.Flags().StringP("name", "n", "","Name of Expense")
    addCmd.Flags().StringP("user", "u", "", "User")
    addCmd.Flags().Float64P("amount", "a",0, "Amount of Expense")

    addCmd.MarkFlagsRequiredTogether("name", "amount", "user")
}

func addExpense(name string, amount float64, fileName string, expenses models.Expenses) {
    expenses.AddExpense(name, amount)
    expenses.WriteToFile(fileName)
}
