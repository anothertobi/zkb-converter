package pkg

import (
	"fmt"
	"reflect"
	"strings"
)

func CSVtoTransactions(data []byte) []Transaction {
	var transactions []Transaction

	dataString := string(data)

	records := strings.Split(dataString, "\r\n")

	for i := 1; i < len(records)-1; i++ { // remove leading and trailing line
		recordElements := strings.Split(records[i], ";")

		for j, recordElement := range recordElements {
			if len(recordElement) > 2 {
				recordElements[j] = recordElement[1 : len(recordElement)-1] // remove leading and trailing quote
			} else {
				recordElements[j] = ""
			}
		}

		transactions = append(transactions, Transaction{
			Date:              recordElements[0],
			Text:              recordElements[1],
			Currency:          recordElements[2],
			AmountDetail:      recordElements[3],
			InternalReference: recordElements[4],
			ReferenceNumber:   recordElements[5],
			WithdrawalAmount:  recordElements[6],
			DepositAmount:     recordElements[7],
			ValueDate:         recordElements[8],
			Balance:           recordElements[9],
			Purpose:           recordElements[10],
		})
	}

	return transactions
}

func TransactionstoCSV(transactions []Transaction) []byte {
	var records []string

	for _, transaction := range transactions {
		var record string

		v := reflect.ValueOf(transaction)

		for i := 0; i < v.NumField(); i++ {
			value := v.Field(i).Interface()

			if i == v.NumField()-1 {
				record += fmt.Sprintf(`"%s"`, value)
			} else {
				record += fmt.Sprintf(`"%s";`, value)
			}
		}

		records = append(records, record)
	}

	var outputString string

	for _, record := range records {
		outputString += fmt.Sprintf("%s\n", record)
	}

	return []byte(outputString)
}
