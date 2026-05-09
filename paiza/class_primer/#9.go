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

type Hero struct {
	l int
	h int
	a int
	d int
	s int
	c int
	f int
}

func NewHero(l, h, a, d, s, c, f int) *Hero {
	return &Hero{
		l: l,
		h: h,
		a: a,
		d: d,
		s: s,
		c: c,
		f: f,
	}
}

func (hero *Hero) levelup(h, a, d, s, c, f int) {
	hero.l++
	hero.h += h
	hero.a += a
	hero.d += d
	hero.s += s
	hero.c += c
	hero.f += f
}

func (hero *Hero) muscle_training(h, a int) {
	hero.h += h
	hero.a += a
}

func (hero *Hero) running(d, s int) {
	hero.d += d
	hero.s += s
}

func (hero *Hero) study(c int) {
	hero.c += c
}

func (hero *Hero) pray(f int) {
	hero.f += f
}

func main() {
	defer w.Flush()
	var x, y int
	fmt.Fscan(r, &x, &y)
	heroMap := make(map[int]*Hero, x)

	for i := 1; i <= x; i++ {
		var l, h, a, d, s, c, f int
		fmt.Fscan(r, &l, &h, &a, &d, &s, &c, &f)
		heroMap[i] = NewHero(l, h, a, d, s, c, f)
	}

	for i := 0; i < y; i++ {
		var heroNum int
		var medhods string

		fmt.Fscan(r, &heroNum, &medhods)
		hero := heroMap[heroNum]
		switch medhods {
		case "levelup":
			var h, a, d, s, c, f int
			fmt.Fscan(r, &h, &a, &d, &s, &c, &f)
			hero.levelup(h, a, d, s, c, f)
		case "muscle_training":
			var h, a int
			fmt.Fscan(r, &h, &a)
			hero.muscle_training(h, a)
		case "running":
			var d, s int
			fmt.Fscan(r, &d, &s)
			hero.running(d, s)
		case "study":
			var c int
			fmt.Fscan(r, &c)
			hero.study(c)
		case "pray":
			var f int
			fmt.Fscan(r, &f)
			hero.pray(f)
		}
	}

	for i := 1; i <= len(heroMap); i++ {
		fmt.Fprintln(w, heroMap[i].l, heroMap[i].h, heroMap[i].a, heroMap[i].d, heroMap[i].s, heroMap[i].c, heroMap[i].f)
	}
}
