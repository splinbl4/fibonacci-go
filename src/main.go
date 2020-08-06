package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	start := r.FormValue("start")
	end := r.FormValue("end")
	if start != "" && end != "" {
		start, err := strconv.Atoi(start)
		end, err := strconv.Atoi(end)

		if err != nil {
			log.Print(err)
		}

		for num := range Count(fibi, start, end) {
			fmt.Fprintf(w, "Number = %d\n", num)
		}
	}
}

type fibfunc func(int) int

func Count(fib fibfunc, start int, end int) chan int {
	ch := make(chan int)
	go func(ch chan int) {
		for i := start; i <= end; i++ {
			ch <- fib(i)
		}

		close(ch)
	}(ch)

	return ch
}

func fibi(n int) int {
	var a, b int = 1, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}
