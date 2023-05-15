package main

import "fmt"

func main() {
	var numbers [5]int   //immutable
	fmt.Println(numbers) // [0,0,0,0,0]

	numbers[0] = 20
	numbers[4] = 10

	fmt.Println(numbers)
	fmt.Println(numbers[4])

	names := [3]string{"Name1", "Name2", "Name3"}
	fmt.Println(names, len(names))

	colors := [...]string{
		"Color 1",
		"Color 2",
		"Color 3",
		"Color 4",
	}
	fmt.Println(colors, len(colors)) // .... len = 4

	coins := [...]string{
		0: "Coin 1",
		2: "Coin 3",
		3: "Coin 4",
		5: "Coin 6",
	}
	fmt.Println(coins, len(coins)) // .... len = 6

	// TODO - Sub Array
	subCoins := coins[2:4]
	fmt.Println(subCoins, len(subCoins))
	subCoins = coins[:3]
	fmt.Println(subCoins, len(subCoins))
	subCoins = coins[3:]
	fmt.Println(subCoins, len(subCoins))

}
