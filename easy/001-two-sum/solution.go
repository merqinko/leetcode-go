package main

import "fmt"

// algorithm Two Sum
// Input: An array of integers nums and an integer target
// Output: Indices of the two numbers such that they add up to target

// 1. Create an empty map numMap to store numbers and their indices.
// 2. Iterate through the array nums using a loop:
//    a. For each number num at index i, calculate the complement as target - num.
//    b. Check if the complement exists in numMap:
// 	  i. If it exists, return the indices of the complement and the current number as a slice [index, i].
//    c. If it does not exist, store the current number and its index in numMap.
// 3. If no valid pair is found after iterating through the array, return nil.

func twoSum(nums []int, target int) []int {

	numMap := make(map[int]int) //map yg menyimpan semua angka dan index ||  A mapping to store numbers and their indices

	for i, num := range nums {
		complement := target - num //cari angka yang dibutuhkan untuk sampe ke jumlah target Find the required number to reach the target
		if index, found := numMap[complement]; found {
			return []int{index, i} //Mengembalikan indeks komplemen dan angka saat ini ||// Return indices of the complement and current number
		}
		numMap[num] = i //menyimpan angka dengan indexnya||store the number with its index
	}
	return nil //jika tidak ditemukan pasangan yang sesuai, kembalikan nil || If no valid pair is found, return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result) // Output: [0 1]
}
