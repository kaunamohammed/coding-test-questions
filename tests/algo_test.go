package tests

import (
	"reflect"
	"testing"

	"github.com/kaunamohammed/godsa/algos"
)

func TestMoveZerosToBack(t *testing.T) {

	result := algos.MoveZerosToBack([]int{4,2,4,0,0,3,0,5,1,0})
	expected:= []int {4,2,4,3,5,1,0,0,0,0}
	if reflect.DeepEqual(result, expected) {
		t.Logf("MoveZerosToBack(\"\") sucess, expected %v, got %v", expected, result)
	} else {
		t.Errorf("MoveZerosToBack(\"\") failed, expected %v, got %v", expected, result)
	}

}

func TestMoveZerosToFront(t *testing.T) {

	result := algos.MoveZerosToFront([]int{4,2,4,0,0,3,0,5,1,0})
	expected:= []int {0,0,0,0,4,2,4,3,5,1}
	if reflect.DeepEqual(result, expected) {
		t.Logf("MoveZerosToBack(\"\") sucess, expected %v, got %v", expected, result)
	} else {
		t.Errorf("MoveZerosToBack(\"\") failed, expected %v, got %v", expected, result)
	}

}

func TestLengthOfLongestSubStringWithDistinctCharacters(t *testing.T) {

	result := algos.LengthOfLongestSubStringWithDistinctCharacters("AAAHBBBBCCAA", 2)

	if result == 6 {
		t.Logf("LengthOfLongestSubStringWithDistinctCharacters(\"\") sucess, expected %v, got %v", 6, result)
	} else {
		t.Errorf("LengthOfLongestSubStringWithDistinctCharacters(\"\") failed, expected %v, got %v", 6, result)
	}

}

func TestLengthOfLongestSubStringWithoutRepeatingCharacters(t *testing.T) {

	result := algos.LengthOfLongestSubStringWithoutRepeatingCharacters("abcabcbb")

	if result == 3 {
		t.Logf("LengthOfLongestSubStringWithoutRepeatingCharacters(\"\") sucess, expected %v, got %v", 3, result)
	} else {
		t.Errorf("LengthOfLongestSubStringWithoutRepeatingCharacters(\"\") failed, expected %v, got %v", 3, result)
	}

}

func TestLengthOfSmallestSubArrayGivenSum(t *testing.T) {

	result := algos.LengthOfSmallestSubArrayGivenSum(8, []int {4,2,2,7,8,1,2,8,1,0})

	if result == 1 {
		t.Logf("LengthOfSmallestSubArrayGivenSum(\"\") sucess, expected %v, got %v", 1, result)
	} else {
		t.Errorf("LengthOfSmallestSubArrayGivenSum(\"\") failed, expected %v, got %v", 1, result)
	}

}

func TestMaxSumSubArray(t *testing.T) {

	result := algos.MaxSumSubArray([]int {4, 2, 1, 7, 8, 1, 2, 8, 1, 0}, 3)

	if result == 16 {
		t.Logf("MaxSumSubArray(\"\") sucess, expected %v, got %v", 16, result)
	} else {
		t.Errorf("MaxSumSubArray(\"\") failed, expected %v, got %v", 16, result)
	}

}

func TestAlternatingSort(t *testing.T) {

	result := algos.AlternatingSort([]int { -92, -23, 0, 45, 89, 96, 99, 95, 89, 41, -17, -48 })

	if result {
		t.Logf("AlternatingSort(\"\") sucess, expected %v, got %v", true, result)
	} else {
		t.Errorf("AlternatingSort(\"\") failed, expected %v, got %v", true, result)
	}

}

func TestFirstNonRepeatingCharacter(t *testing.T) {

	result := algos.FirstNonRepeatingCharacter("asdfsdafdasfjdfsafnnunl")

	if result == "j" {
		t.Logf("FirstNonRepeatingCharacter(\"\") sucess, expected %v, got %v", "j", result)
	} else {
		t.Errorf("FirstNonRepeatingCharacter(\"\") failed, expected %v, got %v", "j", result)
	}

}

func TestRemoveElement(t *testing.T) {

	result := algos.RemoveElement([]int {0,1,2,2,3,0,5,2}, 2)

	if result == 5 {
		t.Logf("RemoveElement(\"\") sucess, expected %v, got %v", 5, result)
	} else {
		t.Errorf("RemoveElement(\"\") failed, expected %v, got %v", 5, result)
	}

}
