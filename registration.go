package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var username string
	var email string
	var age int
	var Phonenumber string
	var password string
	var pwdverify string

	for {
		fmt.Print("Enter user name ,a minimum of 3 characters :")
		username, _ = reader.ReadString('\n')
		username = strings.TrimSpace(username)
		if username == "" {
			fmt.Println("username cannot be empty")
			continue
		}
		if len(username) < 3 {
			fmt.Println("Username must have a minimum of 3 characters")
			continue
		}

		break
	}

	for {
		fmt.Print("Enter your email,johd@mail.com :")
		email, _ = reader.ReadString('\n')
		email = strings.TrimSpace(email)
		if email == "" {
			fmt.Println("email cannot be empty")
			continue
		}
		if !strings.Contains(email, "@") {
			fmt.Println("must contain @")
			continue
		}
		containsat := strings.Index(email, "@")
		containsdot := strings.LastIndex(email, ".")

		if containsdot <= containsat || containsdot >= len(email)-1 {
			fmt.Println("enter a valid emai address, johnd@mail.com")
			continue
		}
		break

	}
	for {
		fmt.Print("Enter your age :")
		ageInput, _ := reader.ReadString('\n')
		ageInput = strings.TrimSpace(ageInput)

		if ageInput == "" {
			fmt.Println(" age cannot be empty")
			continue
		}
		ageValue, err := strconv.Atoi(ageInput)
		if err != nil {
			fmt.Println("enter a valid number")
			continue
		}

		if ageValue < 1 || ageValue > 120 {
			fmt.Println("Enter a valid age")
			continue
		}
		age = ageValue
		break

	}

	for {
		fmt.Print("Enter your phone number 10 -15 digits :")
		PhoneInput, _ := reader.ReadString('\n')
		PhoneInput = strings.TrimSpace(PhoneInput)

		if PhoneInput == "" {
			fmt.Println("Phone number cannot be empty")
			continue
		}
		cleanNumber := strings.ReplaceAll(PhoneInput, "-", "")
		cleanNumber = strings.ReplaceAll(cleanNumber, "(", "")
		cleanNumber = strings.ReplaceAll(cleanNumber, ")", "")
		cleanNumber = strings.ReplaceAll(cleanNumber, " ", "")

		isValid := true
		for i := 0; i < len(cleanNumber); i++ {
			c := cleanNumber[i]
			if c < '0' || c >'9' {
				isValid = false
				break
			}
		}
		if !isValid {
			fmt.Println("Phone number should contain only digits")
			continue
		}
		if len(cleanNumber) < 10 || len(cleanNumber) > 15 {
			fmt.Println("Phone number should have a minimum of 10 digits and max 15")
			continue
		}
		Phonenumber = cleanNumber
		break
	}

	for {
		fmt.Print("Enter password :")
		password, _ = reader.ReadString('\n')
		password = strings.TrimSpace(password)

		if password == "" {
			fmt.Println("password cannot be empty")
			continue
		}
		if len(password) < 8 {
			fmt.Println("password must contain more than 8 characters")
			continue
		}
		fmt.Print("Confirm password :")
		pwdverify, _ = reader.ReadString('\n')
		pwdverify = strings.TrimSpace(pwdverify)

		if password != pwdverify {
			fmt.Println("Passwords do not match,try again")
			continue
		}
		break
	}

	fmt.Println("Here's what you registered:")
	fmt.Printf("Username: %s\n", username)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Email: %s\n", email)
	fmt.Printf("Phone: %s\n", Phonenumber)
}
