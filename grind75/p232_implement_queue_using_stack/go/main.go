package implement_queue_using_stacks

type MyQueue struct {
	input  []int
	output []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	if len(this.output) != 0 {
		for len(this.output) != 0 {
			this.input = append(this.input, this.output[len(this.output)-1])
			this.output = this.output[:len(this.output)-1]
		}
	}

	this.input = append(this.input, x)
}

func (this *MyQueue) Pop() int {
	if len(this.input) != 0 {
		for len(this.input) != 0 {
			this.output = append(this.output, this.input[len(this.input)-1])
			this.input = this.input[:len(this.input)-1]
		}
	}

	temp := this.output[len(this.output)-1]
	this.output = this.output[:len(this.output)-1]

	return temp
}

func (this *MyQueue) Peek() int {
	if len(this.input) != 0 {
		for len(this.input) != 0 {
			this.output = append(this.output, this.input[len(this.input)-1])
			this.input = this.input[:len(this.input)-1]
		}
	}

	return this.output[len(this.output)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.input)+len(this.output) == 0
}