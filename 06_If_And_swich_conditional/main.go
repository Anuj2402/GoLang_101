package main

import (
	"fmt"
	"time"
)

func main(){
	// If-else conditional statement 
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age) // Read user input for age

	if age >= 18 {
		fmt.Println("you are an adult and can vote ")
	}
	if age < 18 {
		fmt.Println("you are not an adult and cannot vote ")
	} else {
		fmt.Println("you are an adult and can vote ")
	}
	// If-else if-else conditional statement
	if age < 13 {
		fmt.Println("you are a child")
	}
	if age >= 13 && age < 18 {
		fmt.Println("you are a teenager")
	}
	if age >= 18 && age < 60 {
		fmt.Println("you are an adult")
	} else {
		fmt.Println("you are a senior citizen")
	}

	//we can also declair variable inside the if statement
	if name := "Anuj Kumar"; name == "Anuj Kumar" {
		fmt.Println("Hello, Anuj Kumar!")
	} else {
		fmt.Println("Hello, stranger!")
	}

	// Switch statement -> why we use switch statement?
	// Switch statement is used to execute one block of code among many alternatives.
	var day int
	fmt.Print("Enter a number (1-7) for the day of the week: ")
	fmt.Scanln(&day) // Read user input for day
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid input! Please enter a number between 1 and 7.")
 }

 // swich with multiple conditions

  switch time.Now().Weekday(){
  case time.Monday, time.Sunday:
	fmt.Println("It's Monday or Sunday, time to relax!")
	  case time.Tuesday, time.Wednesday, time.Thursday:
	fmt.Println("It's a weekday, time to work!")
	  case time.Friday:
	fmt.Println("It's Friday, the weekend is near!")
	  default:
	fmt.Println("It's a day off, enjoy your time!")

  }


  //Type switch statement
  whoAmI : = func(i interface{}){
	switch i.(type) {
	case string:
		fmt.Println("I am a string")
	case int:
		fmt.Println("I am an int ")
	case bool:
		fmt.Println("I am a bool")
	default:
		fmt.Println("I am of a different type")
		
  }
  

}