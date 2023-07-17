package main()

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


fucn main() {
	fmt.Println("Thanks for rating:", input)
	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating,", input)
	numRatig , err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nill {
		fmt.Println("err")
	} else {
		fmt.Println("Added 1 to your rating:", numRatig)
	}
}