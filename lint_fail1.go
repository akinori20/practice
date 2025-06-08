package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
    var unused int // 使われていない変数でlintエラー
}
