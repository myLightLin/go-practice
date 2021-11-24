package main

import "fmt"

/*
* 实现差分数组
* obj := Constructor(test)
* obj.increment(1, 3, 2)
* obj.result()
 */
func main() {
	test := []int{1, 2, 3, 4, 5}
	obj := Constructor(test)
	obj.increment(1, 3, 4)
	res := obj.result()
	fmt.Printf("result is %v\n", res)
}

type Difference struct {
	diff []int
}

func Constructor(nums []int) Difference {
	diff := make([]int, len(nums))
	diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return Difference{diff: diff}
}

func (this *Difference) increment(i, j, val int) {
	this.diff[i] += val
	if j+1 < len(this.diff) {
		this.diff[j+1] -= val
	}
}

func (this *Difference) result() []int {
	res := make([]int, len(this.diff))
	res[0] = this.diff[0]
	for i := 1; i < len(this.diff); i++ {
		res[i] = res[i-1] + this.diff[i]
	}
	return res
}
