package kClosest

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func kClosest(points [][]int, K int) [][]int {

	var removeDups = func(arr [][]int) [][]int{
		return nil
	}
	var sort = func(i , j, k int){}
	var partition = func(i, j int) int{
		return 0
	}
	var distance = func(i int) int{
		return 0
	}
	var swap = func(i, j int){}

	removeDups = func(arr [][]int) [][]int {
		var set = map[string]bool{}
		for i := 0; i < len(points); i++ {
			set[strconv.Itoa(points[i][0]) + " " + strconv.Itoa(points[i][1])] = true
		}
		points = make([][]int, len(set))
		k := 0
		for s, _ := range set {
			sarr := strings.Split(s, " ")
			a, _ := strconv.Atoi(sarr[0])
			b, _ := strconv.Atoi(sarr[1])
			points[k] = []int{a, b}
			k++
		}
		return points
	}
	points = removeDups(points)
	sort = func(i , j, k int){
		if i >= j {
			return
		}
		rand.Seed(time.Now().UnixNano())
		rand := rand.Intn(j - i) + i
		swap(i, rand)
		mid := partition(i, j)
		leftLen := mid - i + 1
		if k < leftLen {
			sort(i, mid -1, k)
		} else if k > leftLen {
			sort(mid + 1, j, k - leftLen)
		}
	}

	partition = func(i, j int) int {
		oi := i
		pivot := distance(i)
		i++

		for {
			for i < j && distance(i) < pivot {
				i++
			}
			for i <= j && distance(j) > pivot {
				j--
			}
			if i >= j {
				break
			}
			swap(i, j)
		}
		swap(oi, j)
		return j
	}

	distance = func(i int) int {
		return points[i][0] * points[i][0] + points[i][1] * points[i][1]
	}

	swap = func(i, j int) {
		t0 := points[i][0]
		t1 := points[i][1]
		points[i][0] = points[j][0]
		points[i][1] = points[j][1]
		points[j][0] = t0
		points[j][1] = t1
	}

	sort(0, len(points) -1, K)
	var res = make([][]int, K)
	for i := 0; i < K; i++ {
		res[i] = points[i]
	}
	return res
}
