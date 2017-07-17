package main

import (
	"fmt"
	"github.com/adskyfly/queue4go"
	// "math/rand"
	// "runtime"
	"net/http"
	"strconv"
	// "sync"
	// "time"
)

type mytype struct {
	name string
	age  int
}

// func test() {
// 	// runtime.GOMAXPROCS(1)
// 	queue := queue4go.Queue("test")
// 	start := time.Now()
// 	var wg sync.WaitGroup
// 	wg.Add(5)
// 	queue.SetMaxLength(2)
// 	// fmt.Println(queue.GetMaxLength())
// 	// rand.Seed(time.Now().UnixNano())
// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i < 1; i++ {
// 			queue.Push(1)
// 		}
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i < 1; i++ {
// 			queue.Push(2)
// 		}
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i < 1; i++ {
// 			queue.Push(3)
// 		}
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i < 1; i++ {
// 			queue.Push(4)
// 		}
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i < 1; i++ {
// 			queue.Push(5)
// 		}
// 	}()
// 	wg.Wait()
// 	delta := time.Since(start)
// 	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
// 	ql := queue.Length()
// 	for i := 0; i < ql; i++ {
// 		fmt.Println(queue.Pop())
// 	}
// }

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		name := query["name"]
		if len(name) > 1 {
			w.Write([]byte("name参数只能有一个"))
			return
		}
		opt := query["opt"]
		if len(opt) > 1 {
			w.Write([]byte("opt参数只能有一个"))
			return
		} else if len(opt) == 0 {
			w.Write([]byte("必须要有opt参数"))
			return
		}

		queue := queue4go.Queue(name[0])

		switch opt[0] {
		case "push":
			data := query["data"]
			if !(len(data) == 1) {
				w.Write([]byte("push模式下data参数必须有且只有1个"))
				return
			}
			if queue.Length() == queue.GetMaxLength() {
				w.Write([]byte("PUSH_END"))
				return
			}
			if !queue.Push(data[0]) {
				w.Write([]byte("PUSH_ERROR"))
			} else {
				w.Write([]byte("PUSH_OK"))
			}
			break
		case "pop":
			w.Write([]byte(queue.Pop().(string)))
			break
		case "view":
			pos := query["pos"]
			if !(len(pos) == 1) {
				w.Write([]byte("view模式下pos参数必须有且只有1个"))
				return
			}
			num, err := strconv.Atoi(pos[0])
			if err != nil {
				w.Write([]byte("pos参数必须为数字"))
				return
			}
			w.Write([]byte(queue.Pos(num).(string)))
			break
		case "reset":
			if queue.Reset() {
				w.Write([]byte("RESET_OK"))
			} else {
				w.Write([]byte("RESET_ERROR"))
			}
			break
		case "maxlength":
			num := query["num"]
			if !(len(num) == 1) {
				w.Write([]byte("maxlength模式下num参数必须有且只有1个"))
				return
			}
			length, err := strconv.Atoi(num[0])
			if err != nil {
				w.Write([]byte("num参数必须为数字"))
				return
			}
			if length < queue.Length() {
				w.Write([]byte("MAXQUEUE_CANCEL"))
				return
			}
			if queue.SetMaxLength(length) {
				w.Write([]byte("MAXQUEUE_OK"))
				return
			} else {
				w.Write([]byte("MAXQUEUE_ERROR"))
				return
			}
		}
	})

	http.ListenAndServe(":8080", nil)
}
