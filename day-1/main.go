package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("inputs.txt")
	handleError(err)
	defer f.Close()

	right, left := []int{}, []int{}

	// seperate the input to right and left
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		cols := strings.Split(line, " ")
		atoi := func(n string) int {
			i, _ := strconv.Atoi(n)
			return i
		}
		left = append(left, atoi(cols[0]))
		right = append(right, atoi(cols[len(cols)-1]))
	}
	// sort them so you can get the absolute distance bewteen values
	sort.Ints(left)
	sort.Ints(right)

	ans := 0
	abs := func(x int) int {
		if x < 0 {
			return x * -1
		}
		return x
	}

	// get the distance and sum them, also get the seen values and frequency
	seen := map[int]int{}
	for i := range left {
		ans += abs(left[i] - right[i])
		seen[right[i]]++
	}
	fmt.Printf("The answer is %d\n", ans)

	// get the similarity values and the frequency
	similarity := 0
	for _, num := range left {
		similarity += (num * seen[num])
	}
	fmt.Printf("similarity score: %d\n", similarity)
}

func handleError(err error) {
	if err == nil {
		return
	}
	panic(err.Error())
}
