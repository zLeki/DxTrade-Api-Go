package main

import "github.com/zLeki/DxTrade-Api-Go"

func main() {
	identity := dx.Identity{
		Username: "1210003069",
		Password: "2K2=WJ3^6rj5",
		Server:   "ftmo", // Or your desired prop firm
	}
	identity.Login()
	// Do Buy or Sell for your desired position
	identity.Buy(0.01, dx.MARKET, "US30.cash", dx.US30) // You must enter both the symbol and the designated id
}
