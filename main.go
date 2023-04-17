package main

import (
	"fmt"
)

func main() {
	err := fromCSV("./spell_full.csv")
	if err != nil {
		fmt.Println(err)
	}
}
