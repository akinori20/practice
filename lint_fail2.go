package lintfail

// コメントがないexported関数でlintエラー
func Add(a int, b int) int {
    return a + b
}
