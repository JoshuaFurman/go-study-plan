package main

import (
	"fmt"
	"math"
	"sort"
)

// TwoSum finds two numbers that add up to target (sorted array)
func TwoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left, right}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}

// RemoveDuplicates removes duplicates from sorted array in-place
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}

// RemoveElement removes all instances of val in-place
func RemoveElement(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

// MaxArea finds maximum area of water container
func MaxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		width := right - left
		h := min(height[left], height[right])
		area := width * h
		maxArea = max(maxArea, area)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

// TrappingRainWater calculates trapped rainwater
func TrappingRainWater(height []int) int {
	if len(height) == 0 {
		return 0
	}

	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	totalWater := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				totalWater += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				totalWater += rightMax - height[right]
			}
			right--
		}
	}
	return totalWater
}

// ThreeSum finds all unique triplets that sum to zero
func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		// Skip duplicates
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		target := -nums[i]

		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Skip duplicates
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

// IsPalindrome checks if string is palindrome using two pointers
func IsPalindrome(s string) bool {
	runes := []rune(s)
	left, right := 0, len(runes)-1

	for left < right {
		// Skip non-alphanumeric characters
		for left < right && !isAlphanumeric(runes[left]) {
			left++
		}
		for left < right && !isAlphanumeric(runes[right]) {
			right--
		}

		if toLower(runes[left]) != toLower(runes[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

// ReverseString reverses string in-place
func ReverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// MoveZeroes moves all zeros to end while maintaining relative order
func MoveZeroes(nums []int) {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}

// MinSubArrayLen finds minimum length subarray with sum >= target
func MinSubArrayLen(target int, nums []int) int {
	minLength := math.MaxInt32
	sum := 0
	left := 0

	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		for sum >= target && left <= right {
			minLength = min(minLength, right-left+1)
			sum -= nums[left]
			left++
		}
	}

	if minLength == math.MaxInt32 {
		return 0
	}
	return minLength
}

// LengthOfLongestSubstring finds length of longest substring without repeating characters
func LengthOfLongestSubstring(s string) int {
	charIndex := make(map[byte]int)
	maxLength := 0
	left := 0

	for right := 0; right < len(s); right++ {
		if idx, exists := charIndex[s[right]]; exists && idx >= left {
			left = idx + 1
		}
		charIndex[s[right]] = right
		maxLength = max(maxLength, right-left+1)
	}

	return maxLength
}

// CharacterReplacement finds longest substring with at most k replacements
func CharacterReplacement(s string, k int) int {
	count := make(map[byte]int)
	maxCount := 0
	left := 0
	result := 0

	for right := 0; right < len(s); right++ {
		count[s[right]]++
		maxCount = max(maxCount, count[s[right]])

		// Window size - max frequency > k means we need to shrink
		for (right-left+1)-maxCount > k {
			count[s[left]]--
			left++
		}

		result = max(result, right-left+1)
	}

	return result
}

// Helper functions
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isAlphanumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + 32
	}
	return r
}

func main() {
	fmt.Println("=== Two Pointers & Sliding Window Demo ===\n")

	// 1. Two Sum
	fmt.Println("1. Two Sum (sorted array):")
	nums := []int{2, 7, 11, 15}
	target := 9
	result := TwoSum(nums, target)
	fmt.Printf("Array: %v, Target: %d\n", nums, target)
	fmt.Printf("Indices: %v (values: %d, %d)\n\n", result, nums[result[0]], nums[result[1]])

	// 2. Remove Duplicates
	fmt.Println("2. Remove Duplicates:")
	dupArr := []int{1, 1, 2, 2, 3, 3, 3}
	fmt.Printf("Original: %v\n", dupArr)
	newLength := RemoveDuplicates(dupArr)
	fmt.Printf("After removal: %v (length: %d)\n\n", dupArr[:newLength], newLength)

	// 3. Remove Element
	fmt.Println("3. Remove Element:")
	elemArr := []int{3, 2, 2, 3}
	val := 3
	fmt.Printf("Original: %v, Remove: %d\n", elemArr, val)
	newLength = RemoveElement(elemArr, val)
	fmt.Printf("After removal: %v (length: %d)\n\n", elemArr[:newLength], newLength)

	// 4. Container With Most Water
	fmt.Println("4. Container With Most Water:")
	heights := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Printf("Heights: %v\n", heights)
	maxWater := MaxArea(heights)
	fmt.Printf("Maximum area: %d\n\n", maxWater)

	// 5. Trapping Rain Water
	fmt.Println("5. Trapping Rain Water:")
	rainHeights := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Printf("Heights: %v\n", rainHeights)
	water := TrappingRainWater(rainHeights)
	fmt.Printf("Trapped water: %d units\n\n", water)

	// 6. 3Sum
	fmt.Println("6. 3Sum (sum to zero):")
	sumArr := []int{-1, 0, 1, 2, -1, -4}
	fmt.Printf("Array: %v\n", sumArr)
	triplets := ThreeSum(sumArr)
	fmt.Printf("Triplets: %v\n\n", triplets)

	// 7. Valid Palindrome
	fmt.Println("7. Valid Palindrome:")
	palindromeStr := "A man, a plan, a canal: Panama"
	fmt.Printf("String: %s\n", palindromeStr)
	isPal := IsPalindrome(palindromeStr)
	fmt.Printf("Is palindrome: %t\n\n", isPal)

	// 8. Reverse String
	fmt.Println("8. Reverse String:")
	strBytes := []byte("hello")
	fmt.Printf("Original: %s\n", string(strBytes))
	ReverseString(strBytes)
	fmt.Printf("Reversed: %s\n\n", string(strBytes))

	// 9. Move Zeroes
	fmt.Println("9. Move Zeroes:")
	zeroArr := []int{0, 1, 0, 3, 12}
	fmt.Printf("Original: %v\n", zeroArr)
	MoveZeroes(zeroArr)
	fmt.Printf("After moving: %v\n\n", zeroArr)

	// 10. Minimum Size Subarray Sum
	fmt.Println("10. Minimum Size Subarray Sum:")
	subArr := []int{2, 3, 1, 2, 4, 3}
	subTarget := 7
	fmt.Printf("Array: %v, Target: %d\n", subArr, subTarget)
	minLen := MinSubArrayLen(subTarget, subArr)
	fmt.Printf("Minimum length: %d\n\n", minLen)

	// 11. Longest Substring Without Repeating Characters
	fmt.Println("11. Longest Substring Without Repeating:")
	subStr := "abcabcbb"
	fmt.Printf("String: %s\n", subStr)
	longest := LengthOfLongestSubstring(subStr)
	fmt.Printf("Longest length: %d\n\n", longest)

	// 12. Character Replacement
	fmt.Println("12. Character Replacement:")
	replaceStr := "AABABBA"
	k := 1
	fmt.Printf("String: %s, K: %d\n", replaceStr, k)
	longestReplace := CharacterReplacement(replaceStr, k)
	fmt.Printf("Longest with %d replacements: %d\n\n", k, longestReplace)

	fmt.Println("=== Demo Complete ===")
}
