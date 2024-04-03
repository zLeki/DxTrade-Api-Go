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
	accountData := identity.GetAccountMetrix()
	fmt.Println("Available funds", accountData.Body.AllMetrics.AvailableFunds)
}
