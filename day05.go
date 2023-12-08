package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

func day5a() int {
	return day5aMinLocation(mustReadInput(5))
}

func day5b() int {
	return day5bMinLocation(mustReadInput(5))
}

type day5route struct {
	dst, src [2]int
}

type theMap [][]day5route

type subject int

const (
	seed subject = iota
	soil
	fertilizer
	water
	light
	temperature
	humidity
	location
)

func day5ParseLines(lines []string) (seeds []int, themap theMap) {
	themap = make(theMap, location+1)

	i := 2
	from := seed
outer:
	for {
		i++

		for {
			if i == len(lines)-1 {
				break outer
			}
			if lines[i] == "" {
				i++
				break
			}

			nums := lineToNums(lines[i])
			delta := nums[2] - 1
			themap[from+1] = append(themap[from+1], day5route{
				dst: [2]int{nums[0], nums[0] + delta},
				src: [2]int{nums[1], nums[1] + delta},
			})

			i++
		}

		from++
	}

	for i := range themap {
		slices.SortFunc(themap[i], func(a, b day5route) int {
			return cmp.Compare(a.src[0], b.src[0])
		})
	}

	_, sseeds, _ := strings.Cut(lines[0], ": ")
	seeds = lineToNums(sseeds)

	return seeds, themap
}

func findDest(routes []day5route, item int) (int, int) {
	for _, route := range routes {
		if item >= route.src[0] && item <= route.src[1] {
			return route.dst[0] + item - route.src[0], route.dst[1]
		}
	}

	return item, item
}

func day5aMinLocation(lines []string) int {
	seeds, themap := day5ParseLines(lines)

	var result []int

	var follow func(item int, typ subject)
	follow = func(item int, typ subject) {
		dst, _ := findDest(themap[typ], item)

		if (typ + 1) > location {
			result = append(result, dst)
			return
		}
		follow(dst, typ+1)
	}

	for _, seed := range seeds {
		follow(seed, soil)
	}

	return min2(result...)
}

func day5bMinLocation(lines []string) int {
	seeds, themap := day5ParseLines(lines)

	var follow func(start, stop int, typ subject, result *int)
	follow = func(start, stop int, typ subject, result *int) {
		croutes := themap[typ]

		var dst, to int
		for i := start; i <= stop; i++ {
			dst, to = findDest(croutes, i)
			to = min(dst+stop-i, to)

			if (typ + 1) > location {
				if *result == -1 || dst < *result {
					*result = dst
				}
			} else {
				follow(dst, to, typ+1, result)
			}

			i += to - dst
		}
	}

	results := make(chan int, 1)

	count := 0
	for i := 0; i < len(seeds); i += 2 {
		i := i
		count++
		go func() {
			result := -1
			follow(seeds[i], seeds[i]+seeds[i+1]-1, soil, &result)

			results <- result
		}()
	}

	min := -1
	for i := 0; i < count; i++ {
		result := <-results
		if min == -1 || result < min {
			min = result
		}
		fmt.Println(i+1, "of", count, ":", result)
	}

	return min
}
