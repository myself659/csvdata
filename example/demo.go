package main

import (
	"fmt"
	"github.com/myself659/csvdata"
	"io"
	"os"
)

type hzhouseitem struct {
	//Id   int64  `gorm:"primary_key`
	Date    string  `gorm:"primary_key"`
	Name    string  `gorm:"primary_key"`
	Area    string  `gorm:"type:varchar(128)"`
	Wc      int     `sql:"wc"`
	Wd      int     `sql:"wd"`
	Areanum float64 `sql: "Areanum"`
	Id      uint
}

func main() {
	var i uint
	fw, _ := os.Create("writer.csv")
	enc := csvdata.NewEncoder(fw)
	for i = 0; i < 10; i++ {
		item := hzhouseitem{"abc", "def", "west", 1, 2, 112.2, i}
		err := enc.Encode(item)
		fmt.Println("encode:", item, err)
	}
	fw.Close()
	fr, _ := os.Open("writer.csv")
	dec := csvdata.NewDecoder(fr)
	di := hzhouseitem{}
	for {
		err := dec.Decode(&di)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(di)
	}
	fr.Close()
	fmt.Print("end")
}
