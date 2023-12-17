package main

import "testing"

func TestDay15a(t *testing.T) {
	expect := 510273
	result := day15a()

	AssertEqual(t, expect, result)
}

func TestDay15aHash(t *testing.T) {
	input := "HASH"
	expect := 52
	result := day15aHash(input)

	AssertEqual(t, expect, result)
}

func TestDay15aTotal(t *testing.T) {
	input := "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"
	expect := 1320
	result := day15aTotal(input)

	AssertEqual(t, expect, result)
}
