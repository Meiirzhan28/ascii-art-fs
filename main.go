package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"output/src"
)

func main() {
	if len(os.Args) == 3 {
		str := os.Args[1]
		banner := os.Args[2]
		if src.Validstr(str, banner) == nil {
			if src.Banner(banner) == nil {
				file, err := os.Open(banner + ".txt")
				if err != nil {
					log.Fatal(err)
				}
				defer file.Close()
				scanner := bufio.NewScanner(file)
				ascii := []string{}
				for scanner.Scan() {
					s := strings.ReplaceAll(scanner.Text(), "/n", "")
					ascii = append(ascii, s)
				}
				if err := scanner.Err(); err != nil {
					log.Fatal(err)
				}
				src.ReadOut(ascii, str)
			} else {
				fmt.Println(src.Banner(banner))
			}
		} else {
			fmt.Println(src.Validstr(str, banner))
		}
	}
	fmt.Println(src.Hash("shadow"))
}
