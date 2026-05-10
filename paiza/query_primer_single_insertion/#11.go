package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

type timelineAuthor struct {
	name string
	year int
}

func main() {
	defer w.Flush()
	var x, y int
	var timelineAuthorList []timelineAuthor
	fmt.Fscan(r, &x, &y)
	// 特にこのパートは何もしない
	for i := 0; i < x; i++ {
		var name string
		fmt.Fscan(r, &name)
	}

	for i := 0; i < y; i++ {
		var t timelineAuthor
		fmt.Fscan(r, &t.year, &t.name)
		timelineAuthorList = append(timelineAuthorList, t)
	}

	sort.Slice(timelineAuthorList, func(i, j int) bool {
		if timelineAuthorList[i].year != timelineAuthorList[j].year {
			return timelineAuthorList[i].year < timelineAuthorList[j].year
		}
		return timelineAuthorList[i].name < timelineAuthorList[j].name
	})

	for _, v := range timelineAuthorList {
		fmt.Fprintln(w, v.name)
	}
}
