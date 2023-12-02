package main

func day1a() int {
	return sumtotal(mustReadInput(1), day1aExtractDigits)
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
	return sumtotal(mustReadInput(1), day1bExtractDigits)
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
