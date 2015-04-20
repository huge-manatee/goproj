package main

import (
	"golang.org/x/tour/pic"
	//"math"
)

func Pic(dx, dy int) [][]uint8 {

	// allocate the two dimentional array
	a := make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		a[i] = make([]uint8, dx)
	}
	
	
	// calculate vales and store in array
	
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			a[i][j] = uint8( (i + j)/2 )
			//a[i][j] = uint8( i * j )
			//a[i][j] = uint8( math.Pow(float64(i),float64(j) ))
		}
	}
	
	
	return a
}

func main() {
	pic.Show(Pic)
}
