package main

func CanWin(stones int) bool {
	return !(stones%4 == 0)
}
