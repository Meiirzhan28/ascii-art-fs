package src

import (
	"fmt"
	"strings"
)

func ReadOut(s []string, banner string) {
	if len(banner) > 0 {
		e := map[rune][]string{}
		var b string
		var q rune = 32
		count := 0
		for i := 1; i < len(s); i += 9 {
			e[q] = s[i : i+8]
			q++
		}
		for _, v := range banner {
			if string(v) == "\\" {
				count++
			}
		}
		if count*2 == len(banner) {
			for i := 0; i < count; i++ {
				fmt.Println()
			}
			return
		}
		k := strings.ReplaceAll(banner, "\\n", "\n")
		l := strings.Split(k, "\n")

		for _, w := range l {
			if w == "" {
				b += "\n"
			} else {
				for i := 0; i < 8; i++ {
					for t := 0; t < len(w); t++ {
						if w[t] >= 32 && w[t] <= 126 {
							b += e[rune(w[t])][i]
						}
					}
					b += "\n"
				}
			}
		}
		fmt.Print(b)
	}
}
