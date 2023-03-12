package quad

import "fmt"

func QuadE(width, height int) {
	if width <= 0 || height <= 0 {
		return
	} else {
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if (i == 0 && j == 0) || (i == height-1 && j == width-1) {
					if (height == 1 && i == 0 && j != 0) || (width == 1 && i == height-1 && j == 0 && height != 1) {
						fmt.Print("C")
					} else {
						fmt.Print("A")
					}
				} else if (i == 0 && j == width-1) || (i == height-1 && j == 0) {
					fmt.Print("C")
				} else {
					if (i != 0 && (j == 0 || j == width-1)) || ((i == 0 || i == height-1) && (j != 0 && j != width-1)) {
						fmt.Print("B")
					} else {
						fmt.Print(" ")
					}
				}
			}
			fmt.Println()
		}
	}
}
