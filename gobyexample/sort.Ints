package main

import (
	//	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

const file = "perf_sort.log"

var arr []int
var size int

var info string

func init() {

	var err error
	const default_size = 1000

	if len(os.Args) == 1 {
		size = default_size
	} else {
		size, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	}

}

func create_arr() {

	for i := 0; i < size; i++ {
		arr = append(arr, rand.Intn(100000))
	}

}

func sorting(arr []int) {

	//fmt.Println("sort.Slice before: ", arr)

	start := time.Now()

	sort.Ints(arr)

	end := time.Since(start).Nanoseconds()

	//fmt.Println("sort.Ints execution time: ", time.Since(start).Nanoseconds(),"ns")

	//fmt.Println("sort.Slice after: ", arr)

	info = "sort.Ints " + "n=" + strconv.Itoa(size) + " first=" + strconv.Itoa(arr[0]) + " last=" + strconv.Itoa(arr[len(arr)-1]) +
			" # " + strconv.FormatInt(end, 10) + "ns"

	Log(file, info)

}

func Log(file string, info string) {

	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	d := []byte(info + "\n")
	if _, err := f.Write([]byte(d)); err != nil {
		f.Close() // ignore error; Write error takes precedence
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func main() {

	create_arr()
	sorting(arr)

}
