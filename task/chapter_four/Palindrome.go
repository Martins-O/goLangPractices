package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var original string

	for len(original) != 5 {
		fmt.Println("Enter a five digit number: ")
		original = scanner.Text()
		scanner.Scan()
		//_, err := fmt.Scanf("%s", &original)
		//if err != nil {
		//	log.Fatalln("err:: ", err)
		//}
		if len(original) == 5 {
			if original[0] == original[4] && original[1] == original[3] {
				fmt.Println("This is a Palindrome")
			} else {
				fmt.Println("This is not a Palindrome")
			}
		} else {
			fmt.Println("Invalid input")
		}
	}
}
