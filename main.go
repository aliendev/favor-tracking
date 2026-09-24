package main

import (
    "flag"
    "fmt"
)

func main() {

    // define flags
    helpFlag := flag.Bool("help", false, "Show help message")
    versionFlag := flag.Bool("version", false, "Show version message")
    addTripFlag := flag.Bool("add-trip", false, "Add a trip")
    addExpenseFlag := flag.Bool("add-expense", false, "Add an expense")
    viewTripsFlag := flag.Bool("view-trips", false, "View trips")
    viewExpensesFlag := flag.Bool("view-expenses", false, "View expenses")
    viewReportsFlag := flag.Bool("view-reports", false, "View reports")


    // define initial help message
    oopsMessage := "Usage: ./favor [command]"

    flag.Parse()
    args := flag.Args()

    if len(args) == 0 {
        fmt.Println(oopsMessage)
    }


}
