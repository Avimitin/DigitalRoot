package main

func DigitalRoot(n int) int {
	for cm := countNum(n); cm > 1; cm = countNum(n) {
		digits := insertDigit(n, cm)
		n = sum(digits)
	}
	return n
}

func countNum(n int) int {
	var digitCount int
	for n > 0 {
		n /= 10
		digitCount++
	}
	return digitCount
}

func sum(digits []int) int {
	var i int
	for _, d := range digits {
		i += d
	}
	return i
}

func insertDigit(n int, amount int) []int {
	var digits []int
	for amount > 0 {
		i := getBit(n, amount)
		n = getPart(n, amount)
		amount--
		digits = append(digits, i)
	}
	return digits
}

func getPart(n int, bit int) int {
	return n % pow(bit-1)
}

func getBit(n int, bit int) int {
	bit -= 1
	pow := pow(bit)
	return n / pow
}

func pow(p int) int {
	var pow = 10
	if p == 0 {
		return 1
	}
	for p > 1 {
		pow *= 10
		p--
	}
	return pow
}
