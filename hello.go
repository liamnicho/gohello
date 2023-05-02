package main

import "fmt"

func main() {
	totalname := 10
	namenumber := 1
	for namenumber <= totalname {
		fmt.Println("Liam")
		namenumber++
	}
	{
		sum := 1
		for i := 1; i <= 10; i++ {
			sum *= i
		}
		fmt.Println(sum)
	}
}
