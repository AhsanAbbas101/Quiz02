
package Quiz02

import (
	"fmt"
	"bufio"
	"os"
)

// Displays option menu
func displayMenu() {
	fmt.Println()
	fmt.Println("Please select an option:")
	fmt.Println("1 – Print Covid19 cases in Pakistan")
	fmt.Println("2 – Print Covid19 cases in SouthKorea")
	fmt.Println("3 – Print Covid19 cases in France")
	fmt.Println("4 – Print my personalized message to Coronavirus")
	fmt.Println("0 – Exit")

	fmt.Println()
	
}


func Execute() {
	
	// Cases As of 29/3/2020-15:03hrs - Source : WHO
	var cases = []int{1526,9583,37575}

	execute := true
	for execute == true {

		displayMenu()

		var input int 
		validate_input := true

		// Validate Input
		for validate_input == true {

			fmt.Printf("Input >> ")
			_, err := fmt.Scan(&input)
			
			if err != nil {
				fmt.Println("Error: ", err)
			} else if input < 0 || input > 4 {
				fmt.Println("Invalid Input! Please Enter Again.")
			} else {
				validate_input = false
			}
			
			var discard string
			fmt.Scanln(&discard)		// Flush

		}

		switch input {

		case 0:
			fmt.Println("Exiting...")
			execute = false
		case 1:
			fmt.Printf("Covid19 cases in Pakistan: %d\n", cases[0])
		case 2:
			fmt.Printf("Covid19 cases in SouthKorea: %d\n", cases[1])
		case 3:
			fmt.Printf("Covid19 cases in France: %d\n", cases[2])
		case 4:
			fmt.Printf("Enter Message >> ")
			in := bufio.NewReader(os.Stdin)
			line, _ := in.ReadString('\n')			// Read input with spaces
			fmt.Printf("Your message: %s\n",line)
		}
	}
}
