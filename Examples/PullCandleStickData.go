package main

import (
	"fmt"
	"github.com/zLeki/DxTrade-Api-Go"
)

func main() {
	identity := dx.Identity{
		Username: "USERNAME",
		Password: "PASSWORD",
		Server:   "ftmo",
	}
	identity.Login()
	data := identity.GetCandleStickData("USDJPY")
	// Print oldest available price
	fmt.Println("Available data", len(data.Body.Data))
	fmt.Println(data.Body.Data[0].Open)
}
