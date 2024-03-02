package main

import (
	"fmt"
	"os"
)

func main() {

	allureResultsDir := os.Getenv("ALLURE_RESULTS")
	fmt.Println("Results are here:", allureResultsDir)

}
