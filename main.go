package main

import (
	"log"
	"os"

	"github.com/anothertobi/zkb-converter/pkg"
)

func main() {
	inputData, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	transactions := pkg.CSVtoTransactions(inputData)

	transactions = pkg.ConvertTransactions(transactions)

	outputData := pkg.TransactionstoCSV(transactions)

	os.Stdout.Write(outputData)
}
