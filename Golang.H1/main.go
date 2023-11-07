package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var U int
	var P int
	var entered_password string
	var count int
	Admin_password := "admin"
	Student_password := "user"
	log.SetPrefix("LOG DATE: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	logDosyam, _ := os.OpenFile("kayit.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0777)
	defer logDosyam.Close()
	log.SetOutput(logDosyam)

	fmt.Println("Please select your access type! \n 1-Admin \n 2-Student")
	fmt.Scan(&U)
	if U == 1 {
		fmt.Println("Enter password: ")
		fmt.Scan(&entered_password)
		if entered_password == Admin_password {

			log.Println("Admin Succesfully logged in!")
			fmt.Println("Admin Succesfully logged in!")

			fmt.Println("Which option would you like to choose?\n1-Log\n2-Exit")
			fmt.Scan(&P)
			if P == 1 {

				file, err := os.Open("kayit.txt")
				if err != nil {
					fmt.Println("Error! The file didn't open.")
				}
				data, err := io.ReadAll(file)
				if err != nil {
					fmt.Println("Error! The file didn't open.")
				}
				logs := string(data)
				fmt.Print(logs)

			} else if P == 2 {
				return
			} else {
				fmt.Println("Enter a valid number.")
				return
			}

			return
		} else {
			for count = 5; count > 0; count-- {

				log.Printf("Username: Admin.\n Admin remaining attempts: %d\nAccess:DENIED", count)
				fmt.Printf("Invalid Password for Admin.\n Your remaining attempts: %d\n", count)
				fmt.Println("Enter Password: ")
				fmt.Scan(&entered_password)
				if entered_password == Admin_password {
					log.Println("Admin Succesfully logged in")
					fmt.Println("Admin Succesfully logged in")
					fmt.Println("Which option would you like to choose?\n1-Log\n2-Exit")
					fmt.Scan(&P)
					if P == 1 {
						os.Open("kayit.txt")
					} else if P == 2 {
						return
					} else {
						fmt.Println("Enter a valid number.")
						return
					}
					return
				}
				if count == 0 {

					log.Println("Admin Banned.")
					fmt.Println("Admin Banned.")
				}
			}
		}
	}
	if U == 2 {
		fmt.Println("Enter password: ")
		fmt.Scan(&entered_password)
		if entered_password == Student_password {
			fmt.Println("Student succesfully logged in!")
			log.Println("Student succesfully logged in!")

			return
		} else {
			for count = 5; count > 0; count-- {

				log.Printf("Username: Student.\n Student remaining attempts: %d\n Access:DENIED", count)
				fmt.Printf("Invalid Password for student.\n Student remaining attempts: %d\n", count)

				fmt.Println("Enter Password: ")
				fmt.Scan(&entered_password)
				if entered_password == Student_password {

					log.Println("Login status: successfully!")
					fmt.Println("Login status: successfully!")
					return
				}
				if count == 0 {

					log.Println("Student banned.")
					fmt.Println("Student Banned.")
				}
			}
		}
	}
	if !(U == 1) || !(U == 2) {
		fmt.Println("Enter valid number.")
	}

}
