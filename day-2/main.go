package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("reports.txt")
	handleError(err)
	defer f.Close()

	reports := [][]int{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		records := strings.Split(line, " ")
		report := make([]int, len(records))
		for i, rec := range records {
			report[i] = atoi(rec)
		}
		reports = append(reports, report)
	}

	validReports := 0
	var validator func([]int, bool) bool
	for _, report := range reports {
		if d := report[0] - report[len(report)-1]; d > 0 {
			validator = validateDecreasingReport
		} else {
			validator = validateIncreasingReport
		}
		if validator(report, false) {
			validReports++
		}
	}
	fmt.Printf("valid reports %d\n", validReports)

	validReports = 0
	for _, report := range reports {
		if d := report[0] - report[len(report)-1]; d > 0 {
			validator = validateDecreasingReport
		} else {
			validator = validateIncreasingReport
		}
		if validator(report, true) {
			validReports++
		}
	}
	fmt.Printf("valid reports deux %d\n", validReports)
}

func remove(nums []int, idx int) []int {
	switch {
	case idx <= 0:
		return nums[1:]
	case idx >= len(nums)-1:
		return nums[:len(nums)-1]
	default:
		n := append([]int{}, nums[:idx]...)
		n = append(n, nums[idx+1:]...)
		return n
	}

}
func atoi(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}

func validateIncreasingReport(nums []int, allowRemoval bool) bool {
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff <= 0 || diff > 3 {
			if !allowRemoval {
				return false
			}
			if validateIncreasingReport(remove(nums, i-1), false) {
				return true
			}
			if validateIncreasingReport(remove(nums, i), false) {
				return true
			}
			if validateIncreasingReport(remove(nums, i+1), false) {
				return true
			}
			return false
		}
	}
	return true
}

func validateDecreasingReport(nums []int, allowRemoval bool) bool {
	for i := 1; i < len(nums); i++ {
		diff := nums[i-1] - nums[i]
		if diff <= 0 || diff > 3 {
			if !allowRemoval {
				return false
			}
			if validateDecreasingReport(remove(nums, i-1), false) {
				return true
			}
			if validateDecreasingReport(remove(nums, i), false) {
				return true
			}
			if validateDecreasingReport(remove(nums, i+1), false) {
				return true
			}
			return false
		}
	}
	return true
}

func handleError(err error) {
	if err == nil {
		return
	}
	panic(err.Error())
}
