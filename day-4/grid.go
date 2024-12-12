package main

// D
// R
// O
// W
func matchUp(r, c int, grid [][]byte, word string) bool {
	switch {
	case r-len(word)+1 < 0:
		return false
	default:
	}
	for i := 0; i < len(word); i++ {
		if grid[r][c] != word[i] {
			return false
		}
		r--
	}
	return true
}

// W
// O
// R
// D
func matchDown(r, c int, grid [][]byte, word string) bool {
	switch {
	case r+len(word) > len(grid):
		return false
	default:
	}
	for i := 0; i < len(word); i++ {
		if grid[r][c] != word[i] {
			return false
		}
		r++
	}
	return true
}

// D R O W
func matchLeft(r, c int, grid [][]byte, word string) bool {
	switch {
	case c-len(word)+1 < 0:
		return false
	default:
	}
	for i := 0; i < len(word); i++ {
		if grid[r][c] != word[i] {
			return false
		}
		c--
	}
	return true
}

// W O R D
func matchRight(r, c int, grid [][]byte, word string) bool {
	switch {
	case c+len(word) > len(grid[0]):
		return false
	default:
	}
	for i := 0; i < len(word); i++ {
		if grid[r][c] != word[i] {
			return false
		}
		c++
	}
	return true
}

// D
//
//	R
//	 O
//	  W
func matchUpLeft(r, c int, grid [][]byte, word string) bool {
	switch {
	case r-len(word)+1 < 0:
		return false
	case c-len(word)+1 < 0:
		return false
	default:
	}
	for i := 0; i < len(word); i++ {
		if grid[r][c] != word[i] {
			return false
		}
		r--
		c--
	}
	return true
}

//	  D
//	 R
//	O
//
// W
func matchUpRight(r, c int, grid [][]byte, word string) bool {
	switch {
	case r-len(word)+1 < 0:
		return false
	case c+len(word) > len(grid[0]):
		return false
	default:
	}
	for i := 0; i < len(word); i++ {
		if grid[r][c] != word[i] {
			return false
		}
		r--
		c++
	}
	return true
}

//	  W
//	 O
//	R
//
// D
func matchDownLeft(r, c int, grid [][]byte, word string) bool {
	switch {
	case r+len(word) > len(grid):
		return false
	case c-len(word)+1 < 0:
		return false
	default:
	}
	for i := 0; i < len(word); i++ {
		if grid[r][c] != word[i] {
			return false
		}
		r++
		c--
	}
	return true
}

// W
//
//	O
//	 R
//	  D
func matchDownRight(r, c int, grid [][]byte, word string) bool {
	switch {
	case r+len(word) > len(grid):
		return false
	case c+len(word) > len(grid[0]):
		return false
	default:
	}
	for i := 0; i < len(word); i++ {
		if grid[r][c] != word[i] {
			return false
		}
		r++
		c++
	}
	return true
}

func matchX(r, c int, grid [][]byte) bool {
	switch {
	case r-1 < 0 || r+1 == len(grid):
		return false
	case c-1 < 0 || c+1 == len(grid[0]):
		return false
	// M M
	//  A
	// S S
	case grid[r-1][c-1] == 'M' && grid[r-1][c+1] == 'M' && grid[r+1][c-1] == 'S' && grid[r+1][c+1] == 'S':
		return true
	// S M
	//  A
	// S M
	case grid[r-1][c-1] == 'S' && grid[r-1][c+1] == 'M' && grid[r+1][c-1] == 'S' && grid[r+1][c+1] == 'M':
		return true
	// S S
	//  A
	// M M
	case grid[r-1][c-1] == 'S' && grid[r-1][c+1] == 'S' && grid[r+1][c-1] == 'M' && grid[r+1][c+1] == 'M':
		return true
	// M S
	//  A
	// M S
	case grid[r-1][c-1] == 'M' && grid[r-1][c+1] == 'S' && grid[r+1][c-1] == 'M' && grid[r+1][c+1] == 'S':
		return true
	default:
	}
	return false
}
