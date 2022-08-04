package src

import (
	"errors"
)

func Validstr(s string, banner string) error {
	switch banner {
	case "standard", "shadow", "thinkertoy":
		for _, i := range s {
			if (i < 32 || i > 126) && i != 10 {
				return errors.New("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
			}
		}
	default:
		return errors.New("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
	}
	return nil
}
