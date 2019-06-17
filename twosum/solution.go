package main

func twosum(nums []int, target int) []int {
	for i:= 0; i < len(nums); i++ {
		num := target - nums[i]
		for j:=i+1; j < len(nums); j++ {
			if num - nums[j] == 0 {
				res := []int{i, j}
				return res
			}
		}
	}
	return nil
}
