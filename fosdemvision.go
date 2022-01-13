package fosdem2022sog

import (
	"fmt"
	"time"
)

func GiveTwelvePointsTo(name string) {
	fmt.Println(name, ": 12 points")
}

func FOSDEMStartVotingNOW() {
	fmt.Println("Voting starts now!")
	time.Sleep(time.Second)
}

func GiveVoteTo(country string, amount int) {
	if country == "United Kingdom" {
		fmt.Println("I'm sorry that cannot be true")
		return
	}
	fmt.Println(country, ":", amount)
}
