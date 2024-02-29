package main

import (
	"github.com/zLeki/DxTrade-Api-Go"
)

func main() {
	identity := dx.Identity{
		Username: "USERNAME",
		Password: "PASSWORD",
		Server:   "ftmo",
	}
	identity.Login()
	identity.Buy(1500, dx.MARKET, 0, 0, "USDJPY", dx.USDJPY)
}
