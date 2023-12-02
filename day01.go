package main

func day1a() int {
	lines := mustReadInput(1)
	return day1aWork(lines)
}

func day1aWork(lines []string) int {
	var total int
	for _, line := range lines {
		total += day1aExtractDigits(line)
	}
	return total
}

func day1aExtractDigits(input string) int {
	firstLast := [2]rune{0, 0}

	for _, chr := range input {
		if chr < '0' || chr > '9' {
			continue
		}

		if firstLast[0] == 0 {
			firstLast[0] = chr
		}

		firstLast[1] = chr
	}

	return 10*int(firstLast[0]-'0') + int(firstLast[1]-'0')
}

func day1b() int {
	lines := mustReadInput(1)
	return day1bWork(lines)
}

func day1bWork(lines []string) int {
	var total int
	for _, line := range lines {
		total += day1bExtractDigits(line)
	}
	return total
}

func day1bExtractDigits(input string) int {
	letters := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	firstLast := [2]int{0, 0}
	bytes := []byte(input)

	for i, chr := range bytes {
		digit := -1

		if chr >= '0' && chr <= '9' {
			digit = int(chr - '0')
		} else {
			for d, letter := range letters {
				if (len(bytes)-i+1) > len(letter) && string(bytes[i:i+len(letter)]) == letter {
					digit = d + 1
					break
				}
			}
		}

		if digit == -1 {
			continue
		}

		if firstLast[0] == 0 {
			firstLast[0] = digit
		}

		firstLast[1] = digit
	}

	return 10*firstLast[0] + firstLast[1]
}
