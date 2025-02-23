package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
)

func isBinaryString(s string) bool {
	return strings.IndexFunc(s, func(r rune) bool {
		return r != '0' && r != '1'
	}) == -1
}

func processChar(char string, position int, valuesChan chan<- int) {
	num, err := strconv.Atoi(char)
	if err != nil {
		return
	}
	base := 2.0
	expoente := float64(position)
	result := num * int(math.Pow(base, expoente))
	valuesChan <- result
}

func sumDecimal(valuesChan <-chan int, resultChan chan<- int) {
	sum := 0
	for value := range valuesChan {
		sum += value
	}
	resultChan <- sum
	close(resultChan)

}

func main() {
	fmt.Println("Welcome to Bin2Dec! - A Simple Go Application to convert Binary to Decimal")
	for {
		fmt.Print("Please enter a binary number (max 8 digits) (q to exit):")
		var binaryInput string
		_, err := fmt.Scanln(&binaryInput)
		if err != nil {
			fmt.Println("Invalid Input, please try again, enter only the binary number.")
			continue
		}
		if binaryInput == "q" {
			break
		}
		if len(binaryInput) > 8 {
			fmt.Println("Input exceeds 8 digits. Please enter a valid binary number.")
			continue
		}
		if !isBinaryString(binaryInput) {
			fmt.Println("Input is not a valid binary number. Please enter only 0s and 1s.")
			continue
		}
		valuesChan := make(chan int)
		resultChan := make(chan int)
		var wg sync.WaitGroup
		go sumDecimal(valuesChan, resultChan)
		wg.Add(len(binaryInput))
		inversePosition := 0
		for i := len(binaryInput) - 1; i >= 0; i-- {
			char := string(binaryInput[i])
			go func(c string, p int) {
				defer wg.Done()
				processChar(c, p, valuesChan)
			}(char, inversePosition)
			inversePosition += 1
		}
		go func() {
			wg.Wait()
			close(valuesChan)
		}()
		result := <-resultChan
		fmt.Printf("Decimal value: %d\n", result)
	}
}
