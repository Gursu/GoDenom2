package canvas

import "fmt"

const MEAT = 34
const Etler = 35

func FillRect(w int, h int) {
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Print("@")
		}
		fmt.Print("\n")
	}
}

func Add(x int) int {
	return x + thefatrat
}
