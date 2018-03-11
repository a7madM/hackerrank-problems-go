package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	actualResult := ClacChoclateCount(10, 2, 5)
	var expectedResult = 6
	if actualResult != expectedResult {
		t.Fatal("Expected", expectedResult, " Got", actualResult)
	}
}

func TestCase2(t *testing.T) {
	actualResult := ClacChoclateCount(12, 4, 4)

	var expectedResult = 3
	if actualResult != expectedResult {
		t.Fatal("Expected", expectedResult, " Got", actualResult)
	}
}

func TestCase3(t *testing.T) {
	actualResult := ClacChoclateCount(6, 2, 2)
	var expectedResult = 5
	if actualResult != expectedResult {
		t.Fatal("Expected", expectedResult, " Got", actualResult)
	}
	// fmt.Println(ClacChoclateCount(43203, 60, 5))
}

func TestCase4(t *testing.T) {
	actualResult := ClacChoclateCount(43203, 60, 5)
	var expectedResult = 899
	if actualResult != expectedResult {
		t.Fatal("Expected", expectedResult, " Got", actualResult)
	}
}
