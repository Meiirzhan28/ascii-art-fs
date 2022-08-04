package src

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"os"
)

func Banner(banner string) error {
	switch banner {
	case "standard":
		if Hash(banner) == "a51f800619146db0c42d26db3114c99f" {
			return nil
		}
	case "shadow":
		if Hash(banner) == "d44671e556d138171774efbababfc135" {
			return nil
		}

	case "thinkertoy":
		if Hash(banner) == "8efd138877a4b281312f6dd1cbe84add" {
			return nil
		}
	}
	return errors.New("Banner changed")
}

func Hash(s string) string {
	h := md5.New()
	f, err := os.Open(s + ".txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = io.Copy(h, f)
	if err != nil {
		panic(err)
	}
	a := fmt.Sprintf("%x", h.Sum(nil))
	return a
}
