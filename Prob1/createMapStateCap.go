package main

import "fmt"

func main() {
	stateCapital := map[string]string{"Uttarpradesh": "Lucknow", "Uttarakhand": "Dheradun", "Punjab": "Chandigarh", "Bihar": "Patna", "Gujarat": "Gandhinagar", "Haryana": "Chandigarh", "Himachal Pradesh": "Shimla", "Sikkim": "Gangtok", "Rajasthan": "Jaipur", "Tamil Nadu": "Chennai"}
	for k, value := range stateCapital {
		fmt.Printf("data type of value is %T .. data type of key is %T..\n", k, value)
	}
}
