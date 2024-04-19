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
	count := 0
	for i := 1; i <=100 ; i++ {
		if i%3 == 0 {
			fmt.Println(i)
			count++
		}
	}
	fmt.Println("The number of multiples of 3 from 1 to 100 is:", count)

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
	// 	{Name: "John", Age: 30, Address: "123 Main Street"},
	// 	{Name: "Jane", Age: 25, Address: "456 Elm Street"},
	// 	{Name: "Bob", Age: 40, Address: "789 Oak Street"},
	// 	{Name: "Alice", Age: 35, Address: "1011 Pine Street"},
	// }
	// for _, person := range people {
	// 	fmt.Printf("Name -: %v (Age: %v)\n", person.Name, person.Age)
	// 	fmt.Println("Address -:", person.Address)
	// 	fmt.Println()
	// }

	// ข้อ 6
	// var c Company
	// c2 := new(Company)
	
	// c.Name = "ABC Company"
	// c.Address = "123 Main Street, City, Country"
	// c.PhoneNumber = "123-456-7890"
	// c.Website = "www.abccompany.com"

	// c2.Name = "ABC Company"
	// c2.Address = "123 Main Street, City, Country"
	// c2.PhoneNumber = "123-456-7890"
	// c2.Website = "www.abccompany.com"
	
	// fmt.Println("Company Name:", c.Name)
	// fmt.Println("Address:", c.Address)
	// fmt.Println("Phone Number:", c.PhoneNumber)
	// fmt.Println("Website:", c.Website)
	// fmt.Println()
	
	// fmt.Printf("\nCompany Name: %v \nAddress: %v \nPhone Number: %v \nWebsite: %v\n", c2.Name, c2.Address, c2.PhoneNumber, c2.Website)
	// fmt.Println("Company Name:", c2.Name)
	// fmt.Println("Address:", c2.Address)
	// fmt.Println("Phone Number:", c2.PhoneNumber)
	// fmt.Println("Website:", c2.Website)
	// fmt.Println()

	// company := Company{
	// 	Name:       "ABC Company",
	// 	Address:    "123 Main Street, City, Country",
	// 	PhoneNumber: "123-456-7890",
	// 	Website:    "www.abccompany.com",
	// }

	// fmt.Println("Company Name:", company.Name)
	// fmt.Println("Address:", company.Address)
	// fmt.Println("Phone Number:", company.PhoneNumber)
	// fmt.Println("Website:", company.Website)
	
	// special 
	// for i := 1; i <= 6; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }
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
