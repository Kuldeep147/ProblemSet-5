package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var (
		gen     int
		bucket1 int
		bucket2 int
		bucket3 int
		bucket4 int
		bucket5 int
	)
	for i := 0; i < 1000; i++ {
		gen = rand.Intn(25)
		if gen <= 5 {
			bucket1++
		} else if gen > 5 && gen <= 10 {
			bucket2++
		} else if gen > 10 && gen <= 15 {
			bucket3++
		} else if gen > 15 && gen <= 20 {
			bucket4++
		} else if gen > 20 && gen <= 25 {
			bucket5++
		}
	}
	numberDistribution := map[string]int{
		">0 && <= 5":      bucket1,
		"> 5 && <= 10":    bucket2,
		"> 10 &&  <= 15":  bucket3,
		" > 15 &&  <= 20": bucket4,
		" > 20 &&  <= 25": bucket5}
	for k, v := range numberDistribution {
		fmt.Printf("%s : %d\n", k, v)
	}

}
