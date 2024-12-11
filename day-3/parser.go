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
	if string(p.mem[p.idx:p.idx+4]) != "mul(" {
		p.idx++
		return p.Next()
	}
	p.idx += 4
	end := p.idx
	hasComma := false
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
	nums := strings.Split(string(p.mem[p.idx:end]), ",")
	if len(nums) != 2 {
		p.idx++
		return p.Next()
	}
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
