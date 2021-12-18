package pkg

import (
	"fmt"
	"strings"
)

func ConvertTransactions(transactions []Transaction) []Transaction {
	var parentTransactionIndex, internalReferenceIndex int
	var convertedTransactions []Transaction

	for i := 0; i < len(transactions); i++ {
		if strings.HasPrefix(transactions[i].Text, "Belastungen") {
			parentTransactionIndex = i
			internalReferenceIndex = 1
			continue
		}

		if len(transactions[i].Date) == 0 {
			if parentTransactionIndex >= 0 && len(transactions[i].Text) > 0 {
				tempTransaction := transactions[parentTransactionIndex]
				tempTransaction.Text = transactions[i].Text
				tempTransaction.WithdrawalAmount = transactions[i].AmountDetail
				tempTransaction.Balance = ""

				tempTransaction.InternalReference += fmt.Sprintf("-%d", internalReferenceIndex)
				internalReferenceIndex++

				convertedTransactions = append(convertedTransactions, tempTransaction)
				continue
			} else {
				continue
			}
		}

		parentTransactionIndex = -1

		convertedTransactions = append(convertedTransactions, transactions[i])
	}

	return convertedTransactions
}
