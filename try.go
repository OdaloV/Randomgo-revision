package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// User struct to hold registration data
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Phone    string `json:"phone"`
	Password string `json:"-"` // "-" means don't show in JSON
}

// Global variable to store user data
var currentUser User

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Registration Form ===")
	
	// Username validation
	for {
		fmt.Print("Enter user name (minimum 3 characters): ")
		username, _ := reader.ReadString('\n')
		currentUser.Username = strings.TrimSpace(username)  // ✅ Store directly in currentUser
		
		if currentUser.Username == "" {
			fmt.Println("username cannot be empty")
			continue
		}
		if len(currentUser.Username) < 3 {
			fmt.Println("Username must have a minimum of 3 characters")
			continue
		}
		break
	}

	// Email validation
	for {
		fmt.Print("Enter your email (johnd@mail.com): ")
		email, _ := reader.ReadString('\n')
		currentUser.Email = strings.TrimSpace(email)  // ✅ Store directly in currentUser
		
		if currentUser.Email == "" {
			fmt.Println("email cannot be empty")
			continue
		}
		if !strings.Contains(currentUser.Email, "@") {
			fmt.Println("must contain @")
			continue
		}
		containsat := strings.Index(currentUser.Email, "@")
		containsdot := strings.LastIndex(currentUser.Email, ".")

		if containsdot <= containsat || containsdot >= len(currentUser.Email)-1 {
			fmt.Println("enter a valid email address, johnd@mail.com")
			continue
		}
		break
	}
	
	// Age validation
	for {
		fmt.Print("Enter your age: ")
		ageInput, _ := reader.ReadString('\n')
		ageInput = strings.TrimSpace(ageInput)

		if ageInput == "" {
			fmt.Println("age cannot be empty")
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
		currentUser.Age = ageValue  // ✅ Store directly in currentUser
		break
	}

	// Phone validation
	for {
		fmt.Print("Enter your phone number (10-15 digits): ")
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
			if c < '0' || c > '9' {
				isValid = false
				break
			}
		}
		if !isValid {
			fmt.Println("Phone number should contain only digits")
			continue
		}
		if len(cleanNumber) < 10 || len(cleanNumber) > 15 {
			fmt.Println("Phone number should have 10-15 digits")
			continue
		}
		currentUser.Phone = cleanNumber  // ✅ Store directly in currentUser
		break
	}

	// Password validation
	for {
		fmt.Print("Enter password: ")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(password)

		if password == "" {
			fmt.Println("password cannot be empty")
			continue
		}
		if len(password) < 8 {
			fmt.Println("password must be at least 8 characters")
			continue
		}
		fmt.Print("Confirm password: ")
		pwdverify, _ := reader.ReadString('\n')
		pwdverify = strings.TrimSpace(pwdverify)

		if password != pwdverify {
			fmt.Println("Passwords do not match, try again")
			continue
		}
		currentUser.Password = password  // ✅ Store directly in currentUser
		break
	}

	fmt.Println("\n✅ Registration successful!")
	fmt.Println("Here's what you registered:")
	fmt.Printf("Username: %s\n", currentUser.Username)
	fmt.Printf("Age: %d\n", currentUser.Age)
	fmt.Printf("Email: %s\n", currentUser.Email)
	fmt.Printf("Phone: %s\n", currentUser.Phone)
	fmt.Println("\n🚀 Starting server on http://localhost:8080")
	fmt.Println("Press Ctrl+C to stop")

	// Handle root path - show JSON
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(currentUser)
	})
	
	// Start server
	http.ListenAndServe(":8080", nil)
}