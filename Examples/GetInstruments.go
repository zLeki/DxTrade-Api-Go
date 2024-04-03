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
	isntruments := identity.GetInstruments()
	for _, instrument := range isntruments.Body {
		println(instrument.Symbol)
	}
}
