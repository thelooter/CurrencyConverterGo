package main

import (
	"fmt"
	"log"
	"strconv"
	"thelooter.de/currencyconverter/internal/messages"
	"thelooter.de/currencyconverter/internal/request"
)

var (
	startingCurrency string
	targetCurrency   string
	amount           float64
)

func main() {
	messages.SendInitialMessage()

	var err error

	_, err = fmt.Scan(&startingCurrency)
	if err != nil {
		log.Fatal("Could not scan starting currency", err)
		return
	}

	fmt.Println("Enter the target currency: ")
	_, err = fmt.Scan(&targetCurrency)

	if err != nil {
		log.Fatal("Could not scan target currency", err)
		return
	}

	var amountString string
	fmt.Println("How much money do you want to convert?")
	_, err = fmt.Scan(&amountString)

	if err != nil {
		log.Fatal("Could not scan amount", err)
		return
	}

	for {
		amount, err = strconv.ParseFloat(amountString, 64)
		if err != nil {
			fmt.Println("Please enter a valid number")
			_, err = fmt.Scan(&amountString)
			if err != nil {
				log.Fatal("Could not scan amount", err)
				return
			}
		} else {
			break
		}
	}

	conversion := request.GetRate(startingCurrency, targetCurrency)

	fmt.Println("You will get " + strconv.FormatFloat(amount*conversion, 'f', 2, 64) + " " + targetCurrency)
}
