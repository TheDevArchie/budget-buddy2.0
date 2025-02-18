package models

import (
	"encoding/json"
    "fmt"
    "os"

)
type Expense struct {
    Name string `json:"name,omitempty"`
    Amount float64 `json:"amount,omitempty"`
}

type Expenses struct {
    Expenses []Expense
}

func (x *Expenses) CreateExpenses(expenses []Expense) {
    x.populateExpenses(expenses)
}

func (x *Expenses) length() (int) {
    return len(x.Expenses)
}

func (x *Expenses) AddExpense(name string, amount float64) {
    newExpense := Expense{Name: name, Amount: amount}
    x.Expenses = append(x.Expenses, newExpense)
}

func (x *Expenses) WriteToFile(fileName string) {
    b, err := json.Marshal(x.Expenses)

    if err != nil {
        fmt.Println("Failed to write to file")
        return
    }

    err = os.WriteFile(fileName, b, 0644)

    if err != nil {
        fmt.Println("Data failed to write to file!") // return error back instead
    }
}

func (x *Expenses) totalAmount() float64 {
    var sum float64 = 0

    for _, expsense := range x.Expenses {
        sum += expsense.Amount
    }
    return sum
}

func (x *Expenses) populateExpenses(expenses []Expense) {
    for _, expense := range expenses {
        if expense.Amount != 0 || expense.Name != "" {
            x.Expenses = append(x.Expenses, expense)
        }
    }
}

func (x *Expenses) GetExpenseList() []Expense {
    return x.Expenses
}
