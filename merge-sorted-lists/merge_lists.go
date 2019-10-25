package merge

func mergeLists(arr1, arr2 []int) []int {
	merged := make([]int, len(arr1)+len(arr2))
	i, j, k := 0,0,0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			merged[k] = arr1[i]
			k++
			i++
		} else {
			merged[k] = arr2[j]
			k++
			j++
		}
	}
	for i < len(arr1) {
		merged[k] = arr1[i]
		k++
		i++
	}
	for j < len(arr2) {
		merged[k] = arr2[j]
		k++
		j++
	}
	return merged
}
