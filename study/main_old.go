// package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"sync"
// 	"time"
// )

// type User struct {
// 	Id    int
// 	Name  string
// 	Email string
// }

// func main() {
// 	http.HandleFunc("/", handleRequest)
// 	http.ListenAndServe(":8080", nil)
// }

// func handleRequest(w http.ResponseWriter, r *http.Request) {
// 	// Thực hiện công việc trước khi sử dụng Goroutine
// 	fmt.Println("Handling request...")

// 	// Sử dụng Goroutine để thực hiện công việc xử lý yêu cầu
// 	go func() {
// 		// Đặt delay để giả lập công việc tốn thời gian
// 		time.Sleep(6 * time.Second)
// 		fmt.Println("Finished processing request")
// 	}()

// 	// Phản hồi ngay lập tức cho client
// 	fmt.Fprintln(w, "Request received and being processed...")
// }
// func getTodos(w http.ResponseWriter, r *http.Request) {
// 	userChan := make(chan User)
// 	wg := sync.WaitGroup{}

// 	wg.Add(2)
// 	go readExcelFile(userChan, &wg)
// 	go saveDataToDB(userChan)
// 	wg.Wait()
// 	w.Header().Set("Content-Type", "application/json")
// 	// fmt.Printf("got /hello request\n")
// 	io.WriteString(w, "Hello, HTTP!\n")
// }
// func readExcelFile(userChan chan<- User, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	// Gửi dữ liệu vào channel dataChan
// 	for i := 0; i < 10; i++ {
// 		time.Sleep(4 * time.Second)
// 		userChan <- User{
// 			i,
// 			"thanh",
// 			"Thanhnv@gmail.com",
// 		}
// 	}

// }
// func saveDataToDB(userChan <-chan User) {
// 	for user := range userChan {
// 		log.Println("Error:", user.Id)
// 	}
// }

// // func main() {
// // 	randoms := []int{}
// // 	for i := 0; i < 100; i++ {
// // 		randoms = append(randoms, i)
// // 	}
// // 	inputChan := generatePipeline(randoms)
// // 	c1 := fanOut(inputChan)
// // 	c2 := fanOut(inputChan)
// // 	c3 := fanOut(inputChan)

// // 	c := fanIn(c1, c2, c3)
// // 	// fan-int
// // 	sum := 0
// // 	for i := 0; i < len(randoms); i++ {
// // 		sum += <-c
// // 	}
// // 	fmt.Println(sum)
// // }

// func generatePipeline(number []int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		for i := 0; i < len(number); i++ {
// 			out <- number[i]
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// func fanOut(in <-chan int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		for n := range in {
// 			out <- n * n
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// func fanIn(inputChan ...<-chan int) <-chan int {
// 	in := make(chan int)
// 	go func() {
// 		for _, c := range inputChan {
// 			for n := range c {
// 				in <- n
// 			}
// 		}
// 	}()
// 	return in
// }
