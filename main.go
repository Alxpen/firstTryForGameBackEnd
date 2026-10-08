package main

import "awesomeProject/world"

func main() {
	world.MM = world.NewMgrMgr()
	world.MM.Pm.Run()
}
