package year2020

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Dimension struct {
	length int
	width  int
	height int
}

func check(err error) {
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func ElvesTotalWrappingPaper() {
	// Read test file
	path, err := filepath.Abs("year2020/testFiles/2020_02_test_input.txt")
	check(err)
	f, err := os.Open(path)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	// Add total wrapping paper. Initialized to 0
	var total = 0
	// Todo: See how we can turn this to a go routine
	for scanner.Scan() {
		stringEquation := scanner.Text()
		// Look into each string and break it into l w h by using x as separator
		arr := strings.Split(stringEquation, "x")
		if len(arr) > 3 || len(arr) < 3 {
			log.Fatalf("Expected array to be 3 but got %d", len(arr))
		}
		l, err := strconv.Atoi(arr[0])
		check(err)
		w, err := strconv.Atoi(arr[1])
		check(err)
		h, err := strconv.Atoi(arr[2])
		check(err)
		// Create a dimension object
		// Call GetTotalSquareFeetNeeded function giving it a dimension
		// Add the result to total
		fmt.Printf("Calculating l: %d, w: %d, h %d \n", l, w, h)
		total = total + getTotalSquareFeetNeeded(Dimension{l, w, h})
		fmt.Printf("New Total: %d \n", total)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// Print total
	fmt.Printf("Total needed %d", total)
}

func getTotalSquareFeetNeeded(dimension Dimension) int {
	var total = 0
	lwArea := dimension.width * dimension.length
	whArea := dimension.width * dimension.height
	lhArea := dimension.length * dimension.height
	minArea := math.Min(float64(lhArea), math.Min(float64(whArea), float64(lwArea)))

	total = (lwArea * 2) + (lhArea * 2) + (whArea * 2) + int(minArea)

	return total
}
