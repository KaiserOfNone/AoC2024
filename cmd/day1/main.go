package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var extra = flag.Bool("extra", false, "an extra flag")
var path = flag.String("file", "", "a file path")

func main() {
	flag.Parse()
	f, err := os.Open(*path)
	if err != nil {
		fmt.Printf("Failed to open input: %v\n", err)
		os.Exit(-1)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	newFunction(scanner)
	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %v\n", err)
		os.Exit(-1)
	}
}

func newFunction(scanner *bufio.Scanner) {
	left, right := []int{}, []int{}
	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), "   ", 2)
		r, l := parts[0], parts[1]
		rn, err := strconv.Atoi(r)
		if err != nil {
			fmt.Printf("Failed to parse int [%s]: %v\n", r, err)
			continue
		}
		ln, err := strconv.Atoi(l)
		if err != nil {
			fmt.Printf("Failed to parse int [%s]: %v\n", l, err)
			continue
		}
		right = append(right, rn)
		left = append(left, ln)
	}
	sort.Ints(left)
	sort.Ints(right)
	diffs := []int{}
	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]
		if diff < 0 {
			diff *= -1
		}
		diffs = append(diffs, diff)
	}
	total := 0
	for _, n := range diffs {
		total += n
	}
	fmt.Printf("Total distance is %v\n", total)
	if *extra {
		solveExtra(right, left)
	}
}

func solveExtra(right []int, left []int) {
	counts := map[int]int{}
	for _, n := range right {
		counts[n] += 1
	}
	similarityScore := 0
	for _, n := range left {
		similarityScore += n * counts[n]
	}
	fmt.Printf("Similarity score: %v\n", similarityScore)
}
