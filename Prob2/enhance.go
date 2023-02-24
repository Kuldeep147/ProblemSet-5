package main

import (
	"fmt"
	"sort"
)

func main() {
	stateCapital := map[string]string{"Uttarpradesh": "Lucknow", "Uttarakhand": "Dheradun", "Punjab": "Chandigarh", "Bihar": "Patna", "Gujarat": "Gandhinagar", "Haryana": "Chandigarh", "Himachal Pradesh": "Shimla", "Sikkim": "Gangtok", "Rajasthan": "Jaipur", "Tamil Nadu": "Chennai"}
	for k, value := range stateCapital {
		fmt.Printf("data type of value is %T .. data type of key is %T..\n", k, value)
	}
	for k, value := range stateCapital {
		fmt.Println(k, value)
	}
	keys := make([]string, 0, len(stateCapital))
	for k := range stateCapital {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return true
	})
	for _, l := range keys {
		fmt.Println(stateCapital[l])
	}
}
