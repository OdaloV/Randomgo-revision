package main

// Fields:
// Member ID (string, format: "LIB" followed by 4 digits, e.g., "LIB1234")
// Full Name (string, at least 2 words)
// Age (integer, minimum 5 years, maximum 100 years)
// Member Type (string, must be one of: "Student", "Teacher", "Senior", "Regular")
// Borrowed Books (integer, 0-10)
// Email (valid email format)
// Validation Rules:
// Member ID must start with "LIB" and have exactly 4 digits after
// Name must have first and last name
// Age must be between 5 and 100
// Member Type must be from the list
// Borrowed books must be 0-10
// Email must contain @ and . after @



import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type member struct {
	MemberID      string
    FullName      string
    Age           int
    MemberType    string
    BorrowedBooks int
    Email         string
}
func main() {
    reader := bufio.NewReader(os.Stdin)
    
    var newMember member
    
    newMember.MemberID = getMemberID(reader)
    fmt.Println() // 
    
    newMember.FullName = getfullname(reader)
    fmt.Println()
    
    newMember.Age = getage(reader)
    fmt.Println()
    
    newMember.MemberType = getmemtype(reader)
    fmt.Println()
    
    newMember.BorrowedBooks = getborrowed(reader)
    fmt.Println()
    
    newMember.Email = getEmail(reader)
    
    displayMember(newMember)
}


func memid(ID string) bool{
	if len(ID) != 7 {
		return false
	}
	if !strings.HasPrefix(ID,"LIB"){
		return false
	}
	digits := ID[3:]
	for i := 0;i < len(digits);i++ {
		if digits[i]< '0'|| digits[i]> '9'{
			return false
		}
	}
	return true

}
func getMemberID(reader *bufio.Reader)string{
	for {
		fmt.Print("Enter your member ID")
		input,_ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if memid(input) {
			fmt.Println("")
			return input
		}
		fmt.Println("Invalid ID")
	}

}
func fullname(name string) bool {
	if name == ""{
		return false
	}
	nameparts := strings.Fields(name)
	if len(nameparts) != 2 {
		return false
	}
	return true
}
func getfullname(reader *bufio.Reader) string {
	for{
		fmt.Print("Enter your first and second name")
		input,_ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if fullname(input){
			fmt.Println("Name accepted")
			return input
		}
		fmt.Println("Enter only 2 names")
	}
}
func ageval(age int) bool {
	
	if age < 5 || age > 100  {
		return false
	}
	return true

}
func getage(reader *bufio.Reader) int {
	for {
		fmt.Print("Enter your age")
		input,_ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		ages,err := strconv.Atoi(input)
		if err != nil{
			fmt.Println("Enter your age")
			continue
		}
		if ageval(ages){
			fmt.Println("age accepted")
			return ages
		}
		fmt.Println("age must be between 45 and 100")
	}
}
func memtype (typer string) bool {
	validmem := []string {"Student","Teacher","Senior","Regular"}
	for _,r := range validmem{
		if typer == r {
			return true
		}
	}
	return false
}
func getmemtype(reader *bufio.Reader) string{
	for {
		fmt.Print("Enter your member type")
		input,_ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if memtype(input){
			fmt.Println("type accepted")
			return input
		}
		fmt.Println("Invalid type")
	}
}
func borrowed (amt int) bool {
	if amt < 0 || amt > 10 {
		return false
	}
	return true
}
func getborrowed(reader *bufio.Reader) int {
	for {
		fmt.Print("How many books have you borrowed?")
		input ,_ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		borroweded,err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Enter valid value")
			continue
		}
		if borrowed(borroweded) {
			fmt.Println("Accepted")
			return borroweded
		}
		fmt.Println("Invalid type")

	}
}
func valemail(email string)bool {
	if email == "" {
		return false
	}
	if !strings.Contains(email, "@"){
		return false
	}
	at := strings.Index(email,"@")
	dot := strings.LastIndex(email, ".")

	if dot <= at {
		return false
	}
	return true

}
func getEmail(reader *bufio.Reader) string {
    for {
        fmt.Print("Enter Email (must contain @ and . after @): ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        
        if valemail(input) {
            fmt.Println(" Email accepted!")
            return input
        }
        
        fmt.Println(" Invalid email format).")
    }
}

func displayMember(members member) {

    fmt.Printf("Member ID:        %s\n", members.MemberID)
    fmt.Printf("Full Name:        %s\n", members.FullName)
    fmt.Printf("Age:              %d\n", members.Age)
    fmt.Printf("Member Type:      %s\n", members.MemberType)
    fmt.Printf("Borrowed Books:   %d\n", members.BorrowedBooks)
    fmt.Printf("Email:            %s\n", members.Email)
}