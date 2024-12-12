package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	handleError(err)
	defer f.Close()

	rules := map[string]map[string]any{}
	updates := [][]string{}
	handleRules := true
	// get the rules and updates
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case len(line) == 0:
			handleRules = false
		case handleRules:
			pageRule := strings.Split(line, "|")
			if _, ok := rules[pageRule[0]]; !ok {
				rules[pageRule[0]] = map[string]any{}
			}
			rules[pageRule[0]][pageRule[1]] = nil
		default:
			pages := strings.Split(line, ",")
			updates = append(updates, pages)
		}
	}

	// get the valid row and add the middle number
	midSum := 0
	fixedUpdates := [][]string{}
	for _, update := range updates {
		valid := true
		for i, page := range update {
			after := rules[page]
			if len(after) == 0 {
				continue
			}
			for j := 0; j < i; j++ {
				if _, ok := after[update[j]]; ok {
					valid = false
					break
				}
			}
			if !valid {
				break
			}
		}
		if !valid {
			// fix the update pages based on the rules
			fix(rules, update)
			fixedUpdates = append(fixedUpdates, update)
			continue
		}
		m := len(update) / 2
		midSum += atoi(update[m])
	}
	fmt.Printf("Mid sum of valid updates: %d\n", midSum)

	// get the fixed rule sum
	fixedSum := 0
	for _, update := range fixedUpdates {
		m := len(update) / 2
		fixedSum += atoi(update[m])
	}
	fmt.Printf("Mid sum of fixed update: %d\n", fixedSum)

}

func fix(rules map[string]map[string]any, update []string) {
	for i, page := range update {
		after := rules[page]
		if len(after) == 0 {
			continue
		}
		for j := 0; j < i; j++ {
			if _, ok := after[update[j]]; ok {
				move(update, j, i)
				break
			}
		}
	}
}

func move(update []string, start, end int) {

	for end > start {
		update[end-1], update[end] = update[end], update[end-1]
		end--
	}
}

func atoi(str string) int {
	num := 0
	for _, n := range str {
		num *= 10
		num += int(n - '0')
	}
	return num
}
func handleError(err error) {
	if err == nil {
		return
	}
	panic(err.Error())
}
