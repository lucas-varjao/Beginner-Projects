package main

import (
	"sync"
	"testing"
	"time"
)

func TestIsBinaryString(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"0", true},
		{"1", true},
		{"01", true},
		{"1010", true},
		{"2", false},
		{"a", false},
		{"01a", false},
		{"", true},
		{"11111111", true},
		{"10203", false},
	}

	for _, test := range tests {
		result := isBinaryString(test.input)
		if result != test.expected {
			t.Errorf("isBinaryString(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestProcessChar(t *testing.T) {
	tests := []struct {
		char     string
		position int
		expected int
	}{
		{"1", 0, 1},
		{"1", 1, 2},
		{"1", 2, 4},
		{"0", 0, 0},
		{"0", 1, 0},
		{"1", 3, 8},
	}

	for _, test := range tests {
		valuesChan := make(chan int, 1)
		go processChar(test.char, test.position, valuesChan)

		select {
		case result := <-valuesChan:
			if result != test.expected {
				t.Errorf("processChar(%q, %d) = %d; want %d",
					test.char, test.position, result, test.expected)
			}
		case <-time.After(time.Second):
			t.Errorf("processChar(%q, %d) timed out", test.char, test.position)
		}
	}
}

func TestSumDecimal(t *testing.T) {
	tests := []struct {
		values   []int
		expected int
	}{
		{[]int{1}, 1},
		{[]int{1, 2, 4, 8}, 15},
		{[]int{0, 0, 0}, 0},
		{[]int{}, 0},
		{[]int{16, 32, 64}, 112},
	}

	for _, test := range tests {
		valuesChan := make(chan int)
		resultChan := make(chan int)

		go sumDecimal(valuesChan, resultChan)

		go func() {
			for _, v := range test.values {
				valuesChan <- v
			}
			close(valuesChan)
		}()

		select {
		case result := <-resultChan:
			if result != test.expected {
				t.Errorf("sumDecimal(%v) = %d; want %d",
					test.values, result, test.expected)
			}
		case <-time.After(time.Second):
			t.Errorf("sumDecimal(%v) timed out", test.values)
		}
	}
}

func TestBinaryConversion(t *testing.T) {
	tests := []struct {
		binary   string
		expected int
	}{
		{"1", 1},
		{"10", 2},
		{"11", 3},
		{"100", 4},
		{"1000", 8},
		{"1010", 10},
		{"1111", 15},
		{"10000", 16},
		{"11111111", 255},
		{"10101010", 170},
		{"11001100", 204},
		{"10011001", 153},
		{"11110000", 240},
		{"00001111", 15},
		{"01010101", 85},
		{"11000011", 195},
	}

	for _, test := range tests {
		valuesChan := make(chan int)
		resultChan := make(chan int)
		var wg sync.WaitGroup

		go sumDecimal(valuesChan, resultChan)
		wg.Add(len(test.binary))
		inversePosition := 0
		for i := len(test.binary) - 1; i >= 0; i-- {
			char := string(test.binary[i])
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
		if result != test.expected {
			t.Errorf("Binary conversion of %s = %d; want %d",
				test.binary, result, test.expected)
		}
	}
}
