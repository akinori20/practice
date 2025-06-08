package lintfail

import "fmt"

func UnusedVar() {
    fmt.Println("Hello, world!")
    var unused int // 使われていない変数でlintエラー
}
