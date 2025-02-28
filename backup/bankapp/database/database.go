package database

import (
	"fmt"
	"os"
	"strconv"

	"bankapp/appcfg"
)

func ReadDataFromFile() float64 {
	data, err := os.ReadFile(appcfg.Datafile)
	if err != nil {
		fmt.Printf("fails to read file - %v\n", err)
	}
	bstring := string(data)
	bfloat64, err := strconv.ParseFloat(bstring, 64)
	if err != nil {
		fmt.Println("Error -", err)
	}
	return bfloat64
}

func SaveDataToFile(balance float64) {
	b := fmt.Sprint(balance)
	os.WriteFile(appcfg.Datafile, []byte(b), 0644)
}
