package min_stack

import "math"

type MinStack struct {
	min  int
	vals []int
}

func Constructor() MinStack {
	return MinStack{min: math.MaxInt32}
}

func (this *MinStack) Push(val int) {
	this.vals = append(this.vals, val)
	if val < this.min {
		this.min = val
	}
}

func (this *MinStack) Pop() {
	find := this.vals[len(this.vals)-1] == this.min
	this.vals = this.vals[:len(this.vals)-1]

	if find {
		if len(this.vals) == 0 {
			this.min = math.MaxInt32
		} else {
			this.min = this.vals[0]
			for _, val := range this.vals {
				if val < this.min {
					this.min = val
				}
			}
		}
	}
}

func (this *MinStack) Top() int {
	return this.vals[len(this.vals)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}
