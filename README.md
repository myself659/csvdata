# Introduction

csvdata has the following functions:
1. convert  struct in golang to csv  
2. convert  csv to struct in golang


# Usage  

code:

```
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

```

output:

```
start encode------
    encode: {abc def west 1 2 112.2 0}
    encode: {abc def west 1 2 112.2 1}
    encode: {abc def west 1 2 112.2 2}
    encode: {abc def west 1 2 112.2 3}
    encode: {abc def west 1 2 112.2 4}
    encode: {abc def west 1 2 112.2 5}
    encode: {abc def west 1 2 112.2 6}
    encode: {abc def west 1 2 112.2 7}
    encode: {abc def west 1 2 112.2 8}
    encode: {abc def west 1 2 112.2 9}
    start decode------
    decode: {abc def west 1 2 112.2 0}
    decode: {abc def west 1 2 112.2 1}
    decode: {abc def west 1 2 112.2 2}
    decode: {abc def west 1 2 112.2 3}
    decode: {abc def west 1 2 112.2 4}
    decode: {abc def west 1 2 112.2 5}
    decode: {abc def west 1 2 112.2 6}
    decode: {abc def west 1 2 112.2 7}
    decode: {abc def west 1 2 112.2 8}
    decode: {abc def west 1 2 112.2 9}
```


### TODO 

1. support tag like `csv: tag1` 
