package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
)

func map_storage(UserName string, address string, phone string, email_address string){
	array := map[string]string{"Name": UserName, "address": address, "PhoneNumber": phone, "email": email_address}
	for key, value := range array{
		fmt.Printf("%v: %v\n", key, value)
	}
}

func Contact_book(){
	var UserName string;
	var address string;
	var phone_number string;
	var email_adress string;

	fmt.Printf("Enter your name: %v\n", UserName)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan(){
		UserName = scanner.Text()
	}
	fmt.Printf("Enter your address: %v\n", address)
	scanner2 := bufio.NewScanner(os.Stdin);
	if scanner2.Scan(){
		address = scanner.Text()
	}
	fmt.Printf("Enter your phoneNumber: +234-%v", phone_number)
	fmt.Scan(&phone_number)
	fmt.Printf("Enter your email address: %v\n", email_adress)
	for {
		fmt.Scan(&email_adress)
		if strings.Contains(email_adress, "@") {
			map_storage(UserName, address, phone_number, email_adress)
			break
		}else{
			fmt.Println("sorry your email do not contain @")
			continue
		}
	}
}
func main(){
	Contact_book()
}