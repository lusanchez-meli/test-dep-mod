package main

import (
	"fmt"

	"github.com/StudioSol/set"
	"github.com/leekchan/accounting"
)

func main() {
	// set uses dep
	i64 := set.NewLinkedHashSetINT64(12345, 49999, 12345, 1289, 1289)

	// accounting uses go mod
	ac := accounting.Accounting{
		Symbol:    "ARS",
		Thousand:  ".",
		Decimal:   ",",
		Precision: 2,
	}

	for v := range i64.Iter() {
		fmt.Println(ac.FormatMoney(float64(v) / 100))
	}
}
