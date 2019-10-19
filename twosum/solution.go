package main

func twosum(nums []int, target int) []int {
	var seen map[int]struct{}
	var res []int
	for i := range nums {
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
