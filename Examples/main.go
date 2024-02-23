package main

import (
	"fmt"
	"github.com/zLeki/DxTrade-Api-Go"
)

func main() {
	identity := dx.Identity{
		Username: "1210003069",
		Password: "2K2=WJ3^6rj5",
		Server:   "ftmo",
	}
	identity.Login()
	positions := identity.GetTransactions()
	for _, v := range positions.Body {
		if v.Quantity > 0 {
			fmt.Println("Buy position with", v.Quantity, " quantity")
			fmt.Println("Closing all buy positions!")
			identity.ClosePosition(v.PositionKey.PositionCode, v.Quantity, 0, v.PositionKey.PositionCode, v.PositionKey.InstrumentId)
		} else if v.Quantity < 0 {
			fmt.Println("Sell position with ", v.Quantity, " quantity")
		}
	}
}
