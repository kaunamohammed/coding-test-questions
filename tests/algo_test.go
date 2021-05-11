package tests

import (
	"testing"

	"github.com/kaunamohammed/godsa/algos"
	"github.com/kaunamohammed/godsa/algos/slidingwindow"
	"github.com/stretchr/testify/assert"
)

func TestMoveZerosToBack(t *testing.T) {

	result := slidingwindow.MoveZerosToBack([]int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0})
	expected := []int{4, 2, 4, 3, 5, 1, 0, 0, 0, 0}

	assert.Equal(t, result, expected)

}

func TestMoveZerosToFront(t *testing.T) {

	result := slidingwindow.MoveZerosToFront([]int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0})
	expected := []int{0, 0, 0, 0, 4, 2, 4, 3, 5, 1}

	assert.Equal(t, result, expected)

}

func TestLengthOfLongestSubStringWithDistinctCharacters(t *testing.T) {

	result := slidingwindow.LengthOfLongestSubStringWithDistinctCharacters("AAAHBBBBCCAA", 2)

	assert.Equal(t, 6, result)

}

func TestLengthOfLongestSubStringWithoutRepeatingCharacters(t *testing.T) {

	result := slidingwindow.LengthOfLongestSubStringWithoutRepeatingCharacters("abcabcbb")

	assert.Equal(t, 3, result)

}

func TestLengthOfSmallestSubArrayGreaterThanOrEqualToGivenSum(t *testing.T) {

	result := slidingwindow.LengthOfSmallestSubArrayGreaterThanOrEqualToGivenSum(8, []int{1, 2, 3, 4, 5, 6, 7, 8})

	assert.Equal(t, 1, result)

}

func TestMaxSumSubArray(t *testing.T) {

	result := slidingwindow.MaxSumSubArray([]int{4, 2, 1, 7, 8, 1, 2, 8, 1, 0}, 3)

	assert.Equal(t, 16, result)

}

func TestAlternatingSort(t *testing.T) {

	isSortedAscending, result := algos.AlternatingSort([]int{-92, -23, 0, 45, 89, 96, 99, 95, 89, 41, -17, -48})
	expected := []int{-92, -48, -23, -17, 0, 41, 45, 89, 89, 95, 96, 99}

	assert.Equal(t, expected, result)
	assert.True(t, isSortedAscending)

}

func TestFirstNonRepeatingCharacter(t *testing.T) {

	result := slidingwindow.FirstNonRepeatingCharacter("asdfsdafdasfjdfsafnnunl")

	assert.Equal(t, "j", result)

}

func TestRemoveElement(t *testing.T) {

	result := algos.RemoveElement([]int{0, 1, 2, 2, 3, 0, 5, 2}, 2)

	assert.Equal(t, 5, result, "expected %v, got %v", 5, result)

}

func TestNumSubArrayProductLessThanK(t *testing.T) {

	result := slidingwindow.NumSubArrayProductLessThanK([]int{10, 5, 2, 6}, 5)

	assert.Equal(t, 8, result)

}

func TestMaxSequenceWithContiniousOne(t *testing.T) {

	result := slidingwindow.MaxSequenceWithContiniousOne([]int{0, 1, 0, 1, 0, 0, 1, 1}, 3)

	assert.Equal(t, 7, result)

}

func TestNumOfSubarraysEqualToK(t *testing.T) {

	result := slidingwindow.NumOfSubarraysEqualToK([]int{1, 7, 4, 3, 1, 2, 1, 5, 1}, 7)
	assert.Equal(t, 4, result)

}

func TestDistinctElementsCountsInSubArray(t *testing.T) {

	result := slidingwindow.DistinctElementsCountsInSubArray( []int{2, 1, 2, 3, 2, 1, 4, 5}, 5)
	assert.Equal(t, "3 3 4 5", result)
	
}
