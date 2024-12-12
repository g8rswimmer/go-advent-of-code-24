package main

import "strings"

type parser struct {
	mem       []byte
	idx       int
	allowSkip bool
}

func (p *parser) Next() (int, int, bool) {
	if !p.nextInstruction() {
		return -1, -1, false
	}
	// check to make sure there is an instruction
	if string(p.mem[p.idx:p.idx+4]) != "mul(" {
		p.idx++
		return p.Next()
	}
	// move to the number
	p.idx += 4
	end := p.idx
	hasComma := false
	// get the end and make sure there is a comma
	for end < len(p.mem) {
		if p.mem[end] == ')' {
			break
		}
		if p.mem[end] == ',' {
			hasComma = true
		}
		end++
	}
	if end >= len(p.mem) {
		return -1, -1, false
	}
	if !hasComma {
		p.idx++
		return p.Next()
	}
	// get the numbers to multiply
	nums := strings.Split(string(p.mem[p.idx:end]), ",")
	if len(nums) != 2 {
		p.idx++
		return p.Next()
	}
	// get the numbers between 1-3 digits
	n1, ok := atoi(nums[0])
	if !ok {
		p.idx++
		return p.Next()
	}
	n2, ok := atoi(nums[1])
	if !ok {
		p.idx++
		return p.Next()
	}
	p.idx = end
	return n1, n2, true
}

// get the next valid mul instruction
// if filtering is on, then account for that
func (p *parser) nextInstruction() bool {
	enable := true
	done := false
	for p.idx < len(p.mem) {
		switch {
		case p.mem[p.idx] == 'm':
			if !p.allowSkip || enable {
				done = true
			}
		case p.idx < len(p.mem)-len("do()") && string(p.mem[p.idx:p.idx+len("do()")]) == "do()":
			enable = true
		case p.idx < len(p.mem)-len("don't()") && string(p.mem[p.idx:p.idx+len("don't()")]) == "don't()" && p.allowSkip:
			enable = false
		}
		if done {
			break
		}
		p.idx++
	}
	return p.idx < len(p.mem)-8
}

func atoi(str string) (int, bool) {
	if len(str) > 3 || len(str) == 0 {
		return -1, false
	}
	num := 0
	for _, n := range str {
		if n < '0' || n > '9' {
			return -1, false
		}
		num *= 10
		num += int(n - '0')
	}
	return num, true
}
