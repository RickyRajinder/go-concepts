package subarray_sum_equals_k


func subarraySum(nums []int, k int) int {
	var count = 0
	var sum = 0
	var set  = map[int]int{}
	set[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if _, ok := set[sum - k]; ok {
			count += set[sum - k]
		}
		get, ok := set[sum]
		if !ok {
			get = 0
		}
		set[sum] = get + 1
	}
	return count
}
