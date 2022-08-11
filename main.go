package main

import "fmt"

type Number interface{
	int64 | float64
}

func main(){
	int := map[string]int64{
		"first": 24,
		"second": 50,
	}
	float := map[string]float64{
		"first": 24.29,
		"second": 500.60,
	}
	fmt.Printf("Sum of Ints: %v \n Sums of Floats: %v \n", sumIntOrFloat(int), sumIntOrFloat(float))

}

func sumIntOrFloat[K comparable, V Number](m map[K]V) V{
	var sum V
	for _, i := range m {
		sum += i
	}
	return sum
}