package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for indx, _ := range pic {
		pic[indx] = make([]uint8, dx)
	}
	// prepare an image
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			if i*i < dy*dy-j*j && (dx-j)*(dx-j) < dx*dx-(dy-i)*(dy-i) {
				pic[i][j] = 255
			}
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
