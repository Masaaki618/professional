package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()
	userMap := map[string]string{
		"Kyoko":        "B",
		"Rio":          "O",
		"Tsubame":      "AB",
		"KurodaSensei": "A",
		"NekoSensei":   "A",
	}
	user := []string{"Kyoko", "Rio", "Tsubame", "KurodaSensei", "NekoSensei"}
	for _, l := range user {
		fmt.Fprintln(w, l, userMap[l])
	}

}
