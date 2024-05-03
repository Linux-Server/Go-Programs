package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("hello go essentials!")
	var investmentAmount float64
	var expectedReturnRate float64 // Short assignment operator
	var years float64
	fmt.Println("Enter the investment amount ?")
	fmt.Scan(&investmentAmount)

	fmt.Println("enter the return rate ? ")
	fmt.Scan(&expectedReturnRate)

	fmt.Println("enter the years ? ")
	fmt.Scan(&years)

	futureValue := investmentAmount * (math.Pow(1+expectedReturnRate/100, years))
	fmt.Println("The value is : ", futureValue)
	// variableDeclararion()
}

func variableDeclararion() {
	name, age := "sam", 12
	var catagory, status string = "sam", "ray"
	fmt.Printf("The name is %v and age is %v\n", name, age)
	fmt.Printf("The name is %v and age is %v\n", catagory, status)

}
