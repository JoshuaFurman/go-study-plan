package main

import (
	"fmt"
	"sort"
)

func main() {
	// Arrays - fixed size
	var numbers [5]int
	numbers[0] = 1
	numbers[1] = 2
	fmt.Println("Array:", numbers)

	// Array literal
	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Println("Primes array:", primes)

	// Slices - dynamic arrays
	var slice []int
	fmt.Println("Empty slice:", slice, "len:", len(slice), "cap:", cap(slice))

	// Slice literal
	slice = []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", slice, "len:", len(slice), "cap:", cap(slice))

	// Slice from array
	array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice1 := array[2:5] // elements 2, 3, 4
	slice2 := array[:3]  // elements 0, 1, 2
	slice3 := array[7:]  // elements 7, 8, 9
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
	fmt.Println("slice3:", slice3)

	// Append to slice
	slice = append(slice, 6, 7, 8)
	fmt.Println("After append:", slice, "len:", len(slice), "cap:", cap(slice))

	// Append slice to slice
	moreNumbers := []int{9, 10, 11}
	slice = append(slice, moreNumbers...)
	fmt.Println("After append slice:", slice, "len:", len(slice), "cap:", cap(slice))

	// Make slice with capacity
	newSlice := make([]int, 3, 10) // len=3, cap=10
	fmt.Println("Make slice:", newSlice, "len:", len(newSlice), "cap:", cap(newSlice))

	// Copy slice
	source := []int{1, 2, 3, 4, 5}
	destination := make([]int, len(source))
	copy(destination, source)
	fmt.Println("Copied slice:", destination)

	// Slice operations
	numbersSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Remove element at index 3
	index := 3
	numbersSlice = append(numbersSlice[:index], numbersSlice[index+1:]...)
	fmt.Println("After removing index 3:", numbersSlice)

	// Insert element at index 2
	numbersSlice = append(numbersSlice[:2], append([]int{99}, numbersSlice[2:]...)...)
	fmt.Println("After inserting 99 at index 2:", numbersSlice)

	// Maps - key-value pairs
	var person map[string]string
	fmt.Println("Nil map:", person)

	// Make map
	person = make(map[string]string)
	person["name"] = "Alice"
	person["age"] = "25"
	person["city"] = "New York"
	fmt.Println("Person map:", person)

	// Map literal
	student := map[string]interface{}{
		"name":   "Bob",
		"grades": []int{85, 92, 78},
		"active": true,
	}
	fmt.Println("Student map:", student)

	// Access map values
	name := person["name"]
	fmt.Println("Name:", name)

	// Check if key exists
	age, exists := person["age"]
	if exists {
		fmt.Println("Age:", age)
	}

	// Delete from map
	delete(person, "city")
	fmt.Println("After deleting city:", person)

	// Iterate over map
	fmt.Println("Iterating over person map:")
	for key, value := range person {
		fmt.Printf("  %s: %s\n", key, value)
	}

	// Map of maps
	users := map[string]map[string]string{
		"user1": {"name": "Alice", "role": "admin"},
		"user2": {"name": "Bob", "role": "user"},
	}
	fmt.Println("Users:", users)

	// Sorting slices
	unsorted := []int{3, 1, 4, 1, 5, 9, 2, 6}
	fmt.Println("Unsorted:", unsorted)

	// Sort in ascending order
	sort.Ints(unsorted)
	fmt.Println("Sorted ascending:", unsorted)

	// Sort in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(unsorted)))
	fmt.Println("Sorted descending:", unsorted)

	// Sorting strings
	words := []string{"zebra", "apple", "banana", "cherry"}
	fmt.Println("Unsorted words:", words)
	sort.Strings(words)
	fmt.Println("Sorted words:", words)

	// Custom sorting with sort.Slice
	people := []struct {
		name string
		age  int
	}{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 20},
	}

	fmt.Println("People before sorting:")
	for _, p := range people {
		fmt.Printf("  %s: %d\n", p.name, p.age)
	}

	// Sort by age
	sort.Slice(people, func(i, j int) bool {
		return people[i].age < people[j].age
	})

	fmt.Println("People after sorting by age:")
	for _, p := range people {
		fmt.Printf("  %s: %d\n", p.name, p.age)
	}

	// Multi-dimensional slices
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Matrix:")
	for _, row := range matrix {
		fmt.Println(" ", row)
	}

	// Accessing elements
	fmt.Printf("matrix[1][2] = %d\n", matrix[1][2]) // 6
}
