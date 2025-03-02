package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"budget_buddy/budget/models"
)

var DataFilesDirectory string = "data"
var ExpenseDataFileTemplate string = "%4d_%02d_expenses.txt"

func GatherExpensesFromFile(fileName string) (models.Expenses, error){
    data, err := os.ReadFile(fileName) // Should we handle error?

    if err != nil {
        return models.Expenses{}, fmt.Errorf("failed to read: %w", err)
    }

    var expensesSlice []models.Expense

    err = json.Unmarshal(data, &expensesSlice)

    if err != nil {
        fmt.Errorf("Failed to Unmarshal the json, %w", err)
    }

    var expenses models.Expenses

    expenses.CreateExpenses(expensesSlice)

    return expenses, nil

}

func RunCmd(name string) {
    cmd := exec.Command(name)
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func PrintExpenses(expenses []models.Expense) {
    var total float64 = 0

    fmt.Println("***** EXPENSES *****")
    fmt.Println("--------------------")

    for i := range expenses {
        fmt.Printf("%s | Amount: %.2f \n", expenses[i].Name, expenses[i].Amount)
        total += expenses[i].Amount
    }

    fmt.Printf("\nTotal Amount: %.2f", total)
}

//Retrieve the appropriate file name
func GetCurrentMonthFile() string {
    currentDate := time.Now()

    month := currentDate.Month()
    year := currentDate.Year()

    fileName := fmt.Sprintf(ExpenseDataFileTemplate, year, month)
    return fileName
}

func GetExpenseFileWithMonthYear(dir string, month int, year int) (string, error) {
    fileName := fmt.Sprintf(ExpenseDataFileTemplate, year, month)
    fullFilePath := fmt.Sprintf("%s/%s", dir, fileName)
    fmt.Println(fullFilePath)

    if !DataFileExists(fullFilePath) {
        return "", fmt.Errorf("file %s does not exist", fileName)
    }

    return fullFilePath, nil
}

func CreateCurrentMonthFile(directory string) string{
    currentMonthFileName := fmt.Sprintf("%s/%s", directory, GetCurrentMonthFile())

    if DataFileExists(currentMonthFileName) {
        return ""
    }

    _, err := os.Create(currentMonthFileName)

    if err != nil {
        log.Fatal("Error creating current month file")
    }
    return currentMonthFileName

}

func DataFileExists(dataDirectory string) bool {
    if _, err := os.Stat(dataDirectory); os.IsNotExist(err) {
        return false
    }
    return true

}

func CreateDataDirectory(dataDirectory string) {
    doesExist := DataFileExists(dataDirectory)

    if !doesExist {
        err := os.MkdirAll(dataDirectory, 0775)

        if err != nil {
            log.Fatalf("error: %s", err)
        }
    }

}

func WriteToFile(obj interface{}, fileName string) {
    data, err := json.Marshal(obj)

    if err != nil {
        log.Fatal("Error converting struct to json")
    }

    // Get file contents 
    // add 

    file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0655)

    if err != nil {
        log.Fatal("Error opening file")
    }

    defer file.Close()

    if _, err := file.Write(data); err != nil {
        log.Fatal("Error appending struct to file")
        return
    }

    fmt.Println("Successfully appended to file")
}
