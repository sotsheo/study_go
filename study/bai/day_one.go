package bai

import "fmt"

// Bài 1: Tìm kiếm nhị phân
// Mô tả: Cho một mảng đã sắp xếp theo thứ tự tăng dần, hãy viết thuật toán tìm kiếm nhị phân để tìm phần tử x trong mảng.
// Input: Một mảng đã sắp xếp và một số x.
// Output: Chỉ số (index) của x trong mảng, hoặc -1 nếu không tìm thấy

type Bai1 struct{}

func (n *Bai1) FindNumber(so int) {
	data := []int{
		12, 123, 12, 31231, 123, 123,
	}
	result := false
	for i, item := range data {
		if item == so {
			fmt.Printf("Số tìm thấy ở vị trí: %d", i)
			result = true
		}
	}
	if !result {
		fmt.Printf("Không tìm thấy số: %d", so)
	}
}
