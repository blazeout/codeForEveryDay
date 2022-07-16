package _0_滑动窗口的平均值

type MovingAverage struct {
	arr      []int
	sum      int
	length   int
	capacity int
}

// Constructor /** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		arr:      make([]int, size),
		sum:      0,
		length:   0,
		capacity: size,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if this.length < this.capacity {
		// 那么说明没有放满
		this.arr[this.length] = val
		this.sum += val
		this.length++
	} else if this.length == this.capacity {
		this.arr = append(this.arr, val)
		valueOfOne := this.arr[0]
		this.arr = this.arr[1:]
		this.sum -= valueOfOne
		this.sum += val
	}
	return float64(this.sum) / float64(this.length)
}
