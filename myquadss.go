package piscine

import (
	"github.com/01-edu/z01"
)

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 1; row <= y; row++ {
		for column := 1; column <= x; column++ {
			if (row == 1 || row == y) && (column == 1 || column == x) {
				z01.PrintRune('o')
			} else if (row == 1 || row == y) && !(column == 1 || column == x) {
				z01.PrintRune('-')
			} else if !(row == 1 || row == y) && (column == 1 || column == x) {
				z01.PrintRune('|')
			} else if !(row == 1 || row == y) && !(column == 1 || column == x) {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 0; row < y; row++ {
		for column := 0; column < x; column++ {
			if (row == 0 || row == y-1) && !(column == 0 || column == x-1) {
				z01.PrintRune('*')
			} else if !(row == 0 || row == y-1) {
				if column == 0 || column == x-1 {
					z01.PrintRune('*')
				} else if !(column == 0 || column == x-1) {
					z01.PrintRune(' ')
				}
			} else if (row == 0 || row == y-1) && (column == 0 || column == x-1) {
				if row == 0 && column == 0 {
					z01.PrintRune('/')
				} else if row == 0 && column == x-1 {
					z01.PrintRune('\\')
				} else if row == y-1 && column == 0 {
					z01.PrintRune('\\')
				} else if row == y-1 && column == x-1 {
					z01.PrintRune('/')
				}
			}
		}
		z01.PrintRune('\n')
	}
}

func QuadC(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 0; row < y; row++ {
		for column := 0; column < x; column++ {
			if row == 0 || row == y-1 {
				if column == 0 || column == x-1 {
					if row == 0 {
						z01.PrintRune('A')
					} else {
						z01.PrintRune('C')
					}
				} else {
					z01.PrintRune('B')
				}
			} else {
				if column == 0 || column == x-1 {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}

func QuadD(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 0; row < y; row++ {
		for column := 0; column < x; column++ {
			if row == 0 || row == y-1 {
				if column == 0 {
					z01.PrintRune('A')
				} else if column == x-1 {
					z01.PrintRune('C')
				} else if !(column == 0 || column == x-1) {
					z01.PrintRune('B')
				}
			} else {
				if column == 0 || column == x-1 {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}

func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 0; row < y; row++ {
		for column := 0; column < x; column++ {
			if (row == 0 || row == y-1) && !(column == 0 || column == x-1) {
				z01.PrintRune('B')
			} else if !(row == 0 || row == y-1) {
				if column == 0 || column == x-1 {
					z01.PrintRune('B')
				} else if !(column == 0 || column == x-1) {
					z01.PrintRune(' ')
				}
			} else if (row == 0 || row == y-1) && (column == 0 || column == x-1) {
				if row == 0 && column == 0 {
					z01.PrintRune('A')
				} else if row == 0 && column == x-1 {
					z01.PrintRune('C')
				} else if row == y-1 && column == 0 {
					z01.PrintRune('C')
				} else if row == y-1 && column == x-1 {
					z01.PrintRune('A')
				}
			}
		}
		z01.PrintRune('\n')
	}
}
