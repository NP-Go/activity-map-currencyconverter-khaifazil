package main

import (
	"fmt"
	"strconv"
	"strings"
)

//Declare Stuct here

type currency struct {
	currencyName string
	rate         float64
}

//Declare Map here

func main() {
	//Insert Main Code here
	rateTable := map[string]currency{
		"USD": currency{"US Dollar", 1.1318},
		"JYP": currency{"Japanese Yen", 121.05},
		"GBP": currency{"Pound Sterling", 0.90630},
		"CNY": currency{"Chinese Yuan Renminbi", 7.9944},
		"SGD": currency{"Singapore Dollar", 1.5743},
		"MYR": currency{"Malaysian Ringgit", 4.8390},
	}

	// fmt.Println(rateTable["USD"].currencyName, rateTable["USD"].rate)

	var userCurrencyFrom string
	var userAmount string
	var userCurrencyTo string

	fmt.Println("Enter your currency: ")
	fmt.Scanln(&userCurrencyFrom)
	fmt.Println("Enter amount: ")
	fmt.Scanln(&userAmount)
	fmt.Println("Enter currency to convert to: ")
	fmt.Scanln(&userCurrencyTo)

	currecyFrom := strings.ToUpper(userCurrencyFrom)
	currecyTo := strings.ToUpper(userCurrencyTo)

	amountValue, _ := strconv.ParseInt(userAmount, 10, 0)

	result := (float64(amountValue) / rateTable[currecyFrom].rate) * rateTable[currecyTo].rate

	fmt.Printf("You will recive: %.2f %v", result, currecyTo)

}
