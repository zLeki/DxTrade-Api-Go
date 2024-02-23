package main

import "github.com/zLeki/DxTrade-Api-Go"

func main() {
	identity := dx.Identity{
		Username: "1210003069",
		Password: "2K2=WJ3^6rj5",
		Server:   "ftmo",
	}
	identity.Login()
	identity.CloseAllPositions()
}
