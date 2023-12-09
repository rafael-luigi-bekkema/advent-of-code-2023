package main

func day9a() int {
	return sumtotal(mustReadInput(9), day9aExtrapolate)
}

func day9b() int {
	return sumtotal(mustReadInput(9), day9bExtrapolate)
}

func allzero(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}

func day9aParse(line string) [][]int {
	nums := lineToNums(line)
	extras := [][]int{nums}

	for {
		nums := extras[len(extras)-1]
		newnums := make([]int, len(nums)-1)
		for i := 0; i < len(nums)-1; i++ {
			diff := nums[i+1] - nums[i]
			newnums[i] = diff
		}
		extras = append(extras, newnums)

		if allzero(newnums) {
			break
		}
	}

	return extras
}

func day9aExtrapolate(line string) int {
	extras := day9aParse(line)

	var lastAdded int
	for i := len(extras) - 1; i >= 0; i-- {
		if i == len(extras)-1 {
			continue
		}

		lastAdded += extras[i][len(extras[i])-1]
	}

	return lastAdded
}

func day9bExtrapolate(line string) int {
	extras := day9aParse(line)

	var lastAdded int
	for i := len(extras) - 1; i >= 0; i-- {
		if i == len(extras)-1 {
			continue
		}

		lastAdded = extras[i][0] - lastAdded
	}

	return lastAdded
}
