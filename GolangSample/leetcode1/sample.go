package main

import (
	"fmt"
	"sort"
)

func main() {

}

type MedianFinder struct {
	a []int
}

func Constructor() MedianFinder {

	m := MedianFinder{}
	return m

}

func (this *MedianFinder) AddNum(num int) {
	this.a = append(this.a, num)
}

func (this *MedianFinder) FindMedian() float64 {
	sum := 0

	sort.Ints(this.a)

	fmt.Println(this.a, len(this.a))

	if len(this.a)%2 == 0 {
		sum = this.a[len(this.a)/2] + this.a[len(this.a)/2-1]
		fmt.Println("even ", sum, float64(sum/2))
		return float64(sum / 2)
	} else if len(this.a) > 1 {
		// sum = this.a[len(this.a)/2 + 1] + this.a[len(this.a)/2 - 1]
		// fmt.Println("odd ", len(this.a), this.a, sum)
		// return float64(sum/2)
		fmt.Println("Odd ", this.a[len(this.a)/2])
		return float64(this.a[len(this.a)/2])
	} else {
		return float64(this.a[0])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
