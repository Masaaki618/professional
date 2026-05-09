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

type ability struct {
	frame       int
	attack      int
	magicAttack bool
}

type hero struct {
	hp         int
	abilityMap map[int]*ability
}

func (h *hero) dead() bool {
	return h.hp < 1
}

func (h *hero) applyMagicBuff() {
	for i := 1; i <= 3; i++ {
		a := h.abilityMap[i]
		if a.magicAttack {
			continue
		}
		a.attack += 5
		if a.frame-3 < 1 {
			a.frame = 1
		} else {
			a.frame -= 3
		}
	}
}

func NewHero(hp int, attacks, frames [3]int) *hero {
	abilityMap := make(map[int]*ability)
	for i := 0; i < 3; i++ {
		abilityMap[i+1] = &ability{
			frame:       frames[i],
			attack:      attacks[i],
			magicAttack: frames[i] == 0 && attacks[i] == 0,
		}
	}
	return &hero{
		hp:         hp,
		abilityMap: abilityMap,
	}
}

func main() {
	defer w.Flush()
	var N, K int
	fmt.Fscan(r, &N, &K)
	heroMap := make(map[int]*hero)

	for i := 1; i <= N; i++ {
		var hp, f1, a1, f2, a2, f3, a3 int
		fmt.Fscan(r, &hp, &f1, &a1, &f2, &a2, &f3, &a3)
		heroMap[i] = NewHero(hp, [3]int{a1, a2, a3}, [3]int{f1, f2, f3})
	}

	for i := 0; i < K; i++ {
		var p1, t1, p2, t2 int
		fmt.Fscan(r, &p1, &t1, &p2, &t2)
		h1, h2 := heroMap[p1], heroMap[p2]
		a1, a2 := h1.abilityMap[t1], h2.abilityMap[t2]

		if h1.dead() || h2.dead() {
			continue
		}

		if a1.magicAttack {
			h1.applyMagicBuff()
			if !a2.magicAttack {
				h1.hp -= a2.attack
				continue
			}
		}

		if a2.magicAttack {
			h2.applyMagicBuff()
			if !a1.magicAttack {
				h2.hp -= a1.attack
				continue
			}
		}

		if a1.frame < a2.frame {
			h2.hp -= a1.attack
		} else if a1.frame > a2.frame {
			h1.hp -= a2.attack
		}
	}

	var aliveCount int
	for _, v := range heroMap {
		if !v.dead() {
			aliveCount++
		}
	}

	fmt.Fprintln(w, aliveCount)
}
