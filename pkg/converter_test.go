package pkg

import (
	"reflect"
	"testing"
)

var singleTransaction = []Transaction{
	{
		Date:              "01.12.2021",
		Text:              "Einzahlung USD Noten",
		InternalReference: "A012B34567CDE89F",
		DepositAmount:     "500",
		ValueDate:         "01.12.2021",
		Balance:           "1000",
	},
}

var inputMultipleTransactions = []Transaction{
	singleTransaction[0],
	{},
	singleTransaction[0],
	{
		Text: "Transaction details",
	},
}

var expectedMultipleTransactions = []Transaction{
	singleTransaction[0],
	singleTransaction[0],
}

var inputCombinedTransactions = []Transaction{
	{
		Date:              "01.12.2021",
		Text:              "Belastungen eBanking Mobile (2)",
		InternalReference: "A012B34567CDE89F",
		WithdrawalAmount:  "500",
		ValueDate:         "01.12.2021",
		Balance:           "300",
	},
	{
		Text:         "Swisscom",
		Currency:     "CHF",
		AmountDetail: "100",
	},
	{
		Text:         "EKZ",
		Currency:     "CHF",
		AmountDetail: "200",
	},
}

var expectedCombinedTransactions = []Transaction{
	{
		Date:              "01.12.2021",
		Text:              "Swisscom",
		InternalReference: "A012B34567CDE89F-1",
		WithdrawalAmount:  "100",
		ValueDate:         "01.12.2021",
		Balance:           "",
	},
	{
		Date:              "01.12.2021",
		Text:              "EKZ",
		InternalReference: "A012B34567CDE89F-2",
		WithdrawalAmount:  "200",
		ValueDate:         "01.12.2021",
		Balance:           "",
	},
}

func TestConvertTransactionsWithSingleTransactions(t *testing.T) {
	value := ConvertTransactions(singleTransaction)
	if !reflect.DeepEqual(singleTransaction, value) {
		t.Fatalf("Single transaction not converted as expected.")
	}
}

func TestConvertTransactionsWithMultipleTransactions(t *testing.T) {
	value := ConvertTransactions(inputMultipleTransactions)
	if !reflect.DeepEqual(expectedMultipleTransactions, value) {
		t.Fatalf("Multiple transaction not converted as expected.")
	}
}

func TestConvertTransactionsWithCombinedTransactions(t *testing.T) {
	value := ConvertTransactions(inputCombinedTransactions)
	if !reflect.DeepEqual(expectedCombinedTransactions, value) {
		t.Fatalf("Combined transactions not converted as expected.")
	}
}
