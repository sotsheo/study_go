package main

import (
	"fmt"
	"math/rand"
	"study/bai"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	randomInts := rand.Perm(100000000)
	start := time.Now()
	// fmt.Println(randomInts)
	result := bai.Bai4{}
	result.Main(randomInts, 0, len(randomInts)-1)
	// fmt.Println(randomInts)
	// Lấy thời gian kết thúc
	end := time.Now()

	// Tính toán thời gian thực hiện
	duration := end.Sub(start)

	fmt.Printf("Thời gian thực hiện: %v\n", duration)
}
