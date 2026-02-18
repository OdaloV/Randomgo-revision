// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main(){

// 	reader := bufio.NewReader(os.Stdin)
// 	var ID string
// 	var name string
// 	var grade int
// 	var GPA float64
// 	//var classes []string

// 	for {
// 		fmt.Print("Enter student ID:")
// 		ID, _ = reader.ReadString('\n')
// 		ID = strings.TrimSpace(ID)
// 		if len(ID) != 6 {
// 			fmt.Println("ID characters must be 6")
// 			continue
// 		}
// 		if !strings.HasPrefix(ID, "S") {
// 			fmt.Println("ID must start with S")
// 			continue
// 		}
// 		digits := ID[1:]
// 		isDigits := true

// 		for i := 0; i < len(digits); i++ {
// 			c := digits[i]
// 			if c < '0' || c > '9' {
// 				isDigits = false
// 				break
// 			}
// 			if !isDigits {
// 				fmt.Println("The remaining characters must be digits")
// 				continue
// 			}
// 		}
// 		break
// 	}
// 	for {
// 		fmt.Print("Enter your 2 names:")
// 		name, _ = reader.ReadString('\n')
// 		name = strings.TrimSpace(name)
// 		if name == ""{
// 			fmt.Println("name cannot be empty")
// 			continue
// 		}

// 		parts := strings.Split(name, " ")
// 		if len(parts) != 2 {
// 			fmt.Println("Enter first and last name")
// 			continue
// 		}
// 		break
// 	}
// 	for {
// 		fmt.Print("Enter your grade :")
// 		Gradesinput,_ := reader.ReadString('\n')
// 		Gradesinput = strings.TrimSpace(Gradesinput)

// 		Grades, err := strconv.Atoi(Gradesinput)
// 		if err != nil {
// 			fmt.Println("Enter a valid number")
// 			continue
// 		}

// 		if Grades < 9 || Grades > 12 {
// 			fmt.Println("Grade must be between 9 and 12")
// 			continue
// 		}
// 		grade = Grades
// 		break
// 	}	
// 	for {
// 		fmt.Print("Enter GPA:")
// 		GPAinput,_ := reader.ReadString('\n')
// 		GPAinput = strings.TrimSpace(GPAinput)

// 		GPAI,err := strconv.ParseFloat(GPAinput, 64)
// 		if err != nil {
// 			fmt.Println("Enter a valid number eg 2.5,3.54")
// 			continue
// 		}
// 		if GPAI < 0.0 || GPAI > 4.0{
// 			fmt.Println("GPA must be between 0.0 and 4,0")
// 			continue
// 		}
// 		GPA =GPAI
// 		break

// 	}
	

// 	fmt.Print(ID)	
// 	fmt.Print(name)
// 	fmt.Print(grade)
// 	fmt.Print(GPA)
// }