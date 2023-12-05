package main

import (
	"fmt"
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

type theMap map[string][]day5route
type rouTes map[string]string

func day5ParseLines(lines []string) (seeds []int, themap theMap, routes rouTes) {
	themap = theMap{}
	routes = rouTes{}

	i := 2
outer:
	for {
		sfromTo, _, _ := strings.Cut(lines[i], " ")
		from, to, _ := strings.Cut(sfromTo, "-to-")
		i++

		for {
			if i == len(lines)-1 {
				break outer
			}
			if lines[i] == "" {
				i++
				break
			}

			nums := toNums(lines[i])
			delta := nums[2] - 1
			routes[from] = to
			themap[to] = append(themap[to], day5route{
				dst: [2]int{nums[0], nums[0] + delta},
				src: [2]int{nums[1], nums[1] + delta},
			})

			i++
		}
	}

	_, sseeds, _ := strings.Cut(lines[0], ": ")
	seeds = toNums(sseeds)

	return seeds, themap, routes
}

func findDest(routes []day5route, item int) (int, int) {
	var dst int
	for _, route := range routes {
		if item >= route.src[0] && item <= route.src[1] {
			diff := item - route.src[0]
			dst = route.dst[0] + diff
			return dst, route.dst[1]
		}
	}

	return item, item
}

func day5aMinLocation(lines []string) int {
	seeds, themap, routes := day5ParseLines(lines)

	var result []int

	var follow func(item int, typ string)
	follow = func(item int, typ string) {
		dst, _ := findDest(themap[typ], item)

		dstTyp, ok := routes[typ]
		if !ok {
			result = append(result, dst)
			return
		}
		follow(dst, dstTyp)
	}

	for _, seed := range seeds {
		follow(seed, "soil")
	}

	return min(result...)
}

func day5bMinLocation(lines []string) int {
	seeds, themap, routes := day5ParseLines(lines)

	var follow func(start, stop int, typ string, result *int)
	follow = func(start, stop int, typ string, result *int) {
		croutes := themap[typ]
		dstTyp, ok := routes[typ]

		var dst, to int
		for i := start; i <= stop; i++ {
			dst, to = findDest(croutes, i)
			to = min(dst+stop-i, to)

			if !ok {
				if *result == -1 || dst < *result {
					*result = dst
				}
			} else {
				follow(dst, to, dstTyp, result)
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
			follow(seeds[i], seeds[i]+seeds[i+1]-1, "soil", &result)

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
