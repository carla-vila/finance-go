package main

import (
	"fmt"

	"github.com/piquette/finance-go/quote"
)

func main() {
	smbl := "cldr"

	q, _ := quote.Get(smbl)

	fmt.Printf("------%v------\n", q.ShortName)

}
