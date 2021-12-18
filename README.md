# ZKB Converter

Converts grouped transactions in a ZKB transaction CSV (incl. details) to single transactions.

## Usage

```
go run main.go "data/export-detailed.csv" > data/export.csv
```

## Example

Grouped transactions like:
```
"03.12.2021";"Belastungen eBanking Mobile (2)";"";"";"A012B34567CDE89H";"";"300";"";"03.12.2021";"600";""
"";"Swisscom";"CHF";"100";"";"";"";"";"";"";""
"";"EKZ";"CHF";"200";"";"";"";"";"";"";""
```

get converted to:

```
"03.12.2021";"Swisscom";"";"";"A012B34567CDE89H-1";"";"100";"";"03.12.2021";"";""
"03.12.2021";"EKZ";"";"";"A012B34567CDE89H-2";"";"200";"";"03.12.2021";"";""
```
