package bai

import "fmt"

// Bài 2: Sắp xếp nổi bọt (Bubble Sort)
// Mô tả: Hãy cài đặt thuật toán sắp xếp nổi bọt để sắp xếp một mảng số nguyên.
// Input: Một mảng số nguyên chưa sắp xếp.
// Output: Mảng đã được sắp xếp theo thứ tự tăng dần.
type Bai2 struct{}

func (b *Bai2) Sapxepnoibot() {
	data := []int{
		123, 232, 124, 213, 1234, 2312, 342, 1234, 2342, 213, 126,
	}
	for i := 0; i < len(data); i++ {
		temp := data[i]
		for j := i + 1; j < len(data); j++ {
			if temp > data[j] {
				data[i] = data[j]
				data[j] = temp
				temp = data[i]
			}
		}
		fmt.Println(data[i])
	}
}
