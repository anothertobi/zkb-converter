package pkg

import (
	"os"
	"reflect"
	"testing"
)

var expectedInputTransactions = []Transaction{
	{
		Date:              "03.12.2021",
		Text:              "Belastungen eBanking Mobile (2)",
		InternalReference: "A012B34567CDE89H",
		WithdrawalAmount:  "300",
		ValueDate:         "03.12.2021",
		Balance:           "600",
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
	{
		Date:              "02.12.2021",
		Text:              "Belastung eBill: Viseca Payment Services SA",
		InternalReference: "A012B34567CDE89G",
		WithdrawalAmount:  "100",
		ValueDate:         "02.12.2021",
		Balance:           "900",
	},
	{},
	{
		Date:              "01.12.2021",
		Text:              "Einzahlung USD Noten",
		InternalReference: "A012B34567CDE89F",
		DepositAmount:     "500",
		ValueDate:         "01.12.2021",
		Balance:           "1000",
	},
}

var outputTransactions = []Transaction{
	{
		Date:              "03.12.2021",
		Text:              "Swisscom",
		InternalReference: "A012B34567CDE89H-1",
		WithdrawalAmount:  "100",
		ValueDate:         "03.12.2021",
		Balance:           "",
	},
	{
		Date:              "03.12.2021",
		Text:              "EKZ",
		InternalReference: "A012B34567CDE89H-2",
		WithdrawalAmount:  "200",
		ValueDate:         "03.12.2021",
		Balance:           "",
	},
	{
		Date:              "02.12.2021",
		Text:              "Belastung eBill: Viseca Payment Services SA",
		InternalReference: "A012B34567CDE89G",
		WithdrawalAmount:  "100",
		ValueDate:         "02.12.2021",
		Balance:           "900",
	},
	{
		Date:              "01.12.2021",
		Text:              "Einzahlung USD Noten",
		InternalReference: "A012B34567CDE89F",
		DepositAmount:     "500",
		ValueDate:         "01.12.2021",
		Balance:           "1000",
	},
}

func TestCSVtoTransactions(t *testing.T) {
	data, err := os.ReadFile("../testdata/input.csv")
	if err != nil {
		t.Fatalf("Failed reading file with %s", err)
	}
	transactions := CSVtoTransactions(data)

	if len(transactions) == 0 {
		t.Fatalf("Returned empty slice")
	}

	if !reflect.DeepEqual(expectedInputTransactions, transactions) {
		t.Fatalf("Transactions don't match expected transactions")
	}
}

func TestTransactionstoCSV(t *testing.T) {
	expectedData, err := os.ReadFile("../testdata/output.csv")
	if err != nil {
		t.Fatalf("Failed reading file with %s", err)
	}

	data := TransactionstoCSV(outputTransactions)

	if !reflect.DeepEqual(expectedData, data) {
		t.Fatalf("CSV data doesn't match expected data")
	}
}
