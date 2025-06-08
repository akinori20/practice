package lintfail

func EmptyIf() {
    // 無意味な空のif文でlintエラーtgrg
    if true {
    }
}
