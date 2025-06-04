package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Type a number and press Enter to beep that many times.")

	for{
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		input = strings.TrimSpace(input)
		if input == "exit" {
			fmt.Println("Exiting...")
			break
		}
		if input == "" {
			fmt.Println("Please enter a valid number or type 'exit' to quit.")
			continue
		}
		n,err:= strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number or type 'exit' to quit.")
			continue
		}

		if n < 0 {
			fmt.Println("Please enter a non-negative number.")
			continue
		}
		for i := 0; i < n; i++ {

			// in cli man page run:
			// man ascii
			// then you can see the ASCII code for bell character is 7 
			// and which is in hexadecimal format.
			os.Stdout.Write([]byte("\x07"))
			time.Sleep(1 * time.Second)
			}
	}

	
}
