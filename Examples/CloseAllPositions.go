package main

import "github.com/zLeki/DxTrade-Api-Go"

func main() {
	identity := dx.Identity{
		Username: "Username",
		Password: "Password",
		Server:   "ftmo",
	}
	identity.Login()
	identity.CloseAllPositions()
}
