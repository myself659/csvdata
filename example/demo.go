package main

import (
	"fmt"
	"io"
	"os"

	"github.com/myself659/csvdata"
)

type hzhouseitem struct {
	//Id   int64  `gorm:"primary_key`
	Date    string
	Name    string
	Area    string
	Wc      int
	Wd      int
	Areanum float64
	ID      uint
}

func main() {
	var i uint
	fw, _ := os.Create("demo.csv")
	enc := csvdata.NewEncoder(fw)
	fmt.Println("start encode------")
	for i = 0; i < 10; i++ {
		item := hzhouseitem{"abc", "def", "west", 1, 2, 112.2, i}
		err := enc.Encode(item)
		if err == nil {
			fmt.Println("encode:", item)
		} else {
			fmt.Println("encode:", err)
		}
	}
	fw.Close()
	fr, _ := os.Open("demo.csv")
	dec := csvdata.NewDecoder(fr)
	di := hzhouseitem{}
	fmt.Println("start decode------")
	for {
		err := dec.Decode(&di)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("decode:", di)
	}
	fr.Close()
}
