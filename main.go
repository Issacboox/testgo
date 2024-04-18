package main

import (
	"fmt"
	// "os"
	// "strconv"
	// "strings"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

type Company struct {
	Name       string
	Address    string
	PhoneNumber string
	Website    string
}

func main() {
	// fmt.Println("Hello world")

	// //ข้อ1
	// for i := 0; i <=100 ; i++ {
	// 	if i%3 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	//ข้อ 1.2
	// num2()

	//ข้อ2
	// x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	// min, max := findMinMax(x)
	// fmt.Println("Minimum:", min)
	// fmt.Println("Maximum:", max)

	// //ข้อ3
	// count := countNumber(1000)

	// ข้อ 3.1
	// count := countNumber(10000)
	// fmt.Println("The number of 9s in the numbers from 1 to 10000 is:", count)

	// ข้อ 4
	// myWords := "AW SOME GO!"
	// newWords := removeSpaces(myWords)
	// fmt.Println(newWords)

	// ข้อ 4.1
	// myWords := "ine t"
	// newWords := removeSpaces(myWords)
	// fmt.Println(newWords)

	// ข้อ 5
	// people := []Person{
	// 	{Name: "Mr. Tim Carry", Age: 25, Address: "142rd roads, Virginir, 22202"},
	// 	{Name: "Ms. Laura McCoy", Age: 22, Address: "123/western Street, New York, 12304"},
	// 	{Name: "Nura", Age: 22, Address: "Saraburi"},
	// }

	// fmt.Println("Enter 0 to print all , 1 to print person 2, or any other key to exit:")
	// var choice int
	// _, err := fmt.Scan(&choice)
	// if err != nil {
	// 	fmt.Println("Error reading input:", err)
	// 	os.Exit(1)
	// }

	// switch choice {
	// case 1:
	// 	if len(people) > 1 {
	// 		p := people[1]
	// 		fmt.Printf("Name -: %s (Age: %d)\n", p.Name, p.Age)
	// 		fmt.Printf("Address -: %s\n", p.Address)
	// 	} else {
	// 		fmt.Println("No person 2 data available")
	// 	}
	// case 0:
	// 	fmt.Println("Result -:")
	// 	for _, p := range people {
	// 		fmt.Printf("Name -: %s (Age: %d)\n", p.Name, p.Age)
	// 		fmt.Printf("Address -: %s\n", p.Address)
	// 	}
	// default:
	// 	fmt.Println("No data found")
	// }
	
	// ข้อ 6
	company := Company{
		Name:       "ABC Company",
		Address:    "123 Main Street, City, Country",
		PhoneNumber: "123-456-7890",
		Website:    "www.abccompany.com",
	}

	fmt.Println("Company Name:", company.Name)
	fmt.Println("Address:", company.Address)
	fmt.Println("Phone Number:", company.PhoneNumber)
	fmt.Println("Website:", company.Website)
	
}
// func num2() {
// 	x := 4
// 	y := 3
// 	res := powerNumber(x, y)
// 	fmt.Println("Number is == > ", x)
// 	fmt.Println("PowerNumber is == > ", res)
// }

// func num3() {
// 	//
// }

// func powerNumber(x, y int) int {
// 	result := 1
// 	for i := 0; i < y; i++ {
// 		result *= x
// 	}
// 	return result
// }

// func findMinMax(numbers []int) (min, max int) {
// 	min = numbers[0]
// 	max = numbers[0]
// 	for _, number := range numbers {
// 		if number < min {
// 			min = number
// 		}
// 		if number > max {
// 			max = number
// 		}
// 	}
// 	return
// }

// func countNumber(n int) int {
// 	count := 0
// 	for i := 1; i <= n; i++ {
// 		str := strconv.Itoa(i)
// 		for _, c := range str {
// 			if c == '9' {
// 				count++
// 			}
// 		}
// 	}
// 	return count
// }

// func removeSpaces(s string) string {
// 	var newText string
// 	for _, c := range s {
// 		if c != ' ' {
// 			newText += string(c)
// 		}
// 	}
// 	return newText
// }
