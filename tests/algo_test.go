package tests

import (
	"reflect"
	"testing"

	"github.com/kaunamohammed/godsa/algos"
)

func TestMoveZerosToBack(t *testing.T) {

	result := algos.MoveZerosToBack([]int{4,2,4,0,0,3,0,5,1,0})
	expected:= []int {4,2,4,3,5,1,0,0,0,0}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MoveZerosToBack(\"\") failed, expected %v, got %v", expected, result)
	}

}

func TestMoveZerosToFront(t *testing.T) {

	result := algos.MoveZerosToFront([]int{4,2,4,0,0,3,0,5,1,0})
	expected:= []int {0,0,0,0,4,2,4,3,5,1}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MoveZerosToBack(\"\") failed, expected %v, got %v", expected, result)
	}

}

func TestLengthOfLongestSubStringWithDistinctCharacters(t *testing.T) {

	result := algos.LengthOfLongestSubStringWithDistinctCharacters("AAAHBBBBCCAA", 2)

	if result != 6 {
		t.Errorf("LengthOfLongestSubStringWithDistinctCharacters(\"\") failed, expected %v, got %v", 6, result)
	}

}

func TestLengthOfLongestSubStringWithoutRepeatingCharacters(t *testing.T) {

	result := algos.LengthOfLongestSubStringWithoutRepeatingCharacters("abcabcbb")

	if result != 3 {
		t.Errorf("LengthOfLongestSubStringWithoutRepeatingCharacters(\"\") failed, expected %v, got %v", 3, result)
	}

}

func TestLengthOfSmallestSubArrayGivenSum(t *testing.T) {

	result := algos.LengthOfSmallestSubArrayGivenSum(8, []int {4,2,2,7,8,1,2,8,1,0})

	if result != 1 {
		t.Errorf("LengthOfSmallestSubArrayGivenSum(\"\") failed, expected %v, got %v", 1, result)
	}

}

func TestMaxSumSubArray(t *testing.T) {

	result := algos.MaxSumSubArray([]int {4, 2, 1, 7, 8, 1, 2, 8, 1, 0}, 3)

	if result != 16 {
		t.Errorf("MaxSumSubArray(\"\") failed, expected %v, got %v", 16, result)
	}

}

func TestAlternatingSort(t *testing.T) {

	isSortedAscending, arr := algos.AlternatingSort([]int { -92, -23, 0, 45, 89, 96, 99, 95, 89, 41, -17, -48 })
	expected := []int { -92, -48, -23, -17, 0, 41, 45, 89, 89, 95, 96, 99 }

	if !isSortedAscending {
		t.Errorf("AlternatingSort(\"\") failed, expected %v, got %v", expected, arr)
	}

}

func TestFirstNonRepeatingCharacter(t *testing.T) {

	result := algos.FirstNonRepeatingCharacter("asdfsdafdasfjdfsafnnunl")

	if result != "j" {
		t.Errorf("FirstNonRepeatingCharacter(\"\") failed, expected %v, got %v", "j", result)
	}

}

func TestRemoveElement(t *testing.T) {

	result := algos.RemoveElement([]int {0,1,2,2,3,0,5,2}, 2)

	if result != 5 {
		t.Errorf("RemoveElement(\"\") failed, expected %v, got %v", 5, result)
	}

}
