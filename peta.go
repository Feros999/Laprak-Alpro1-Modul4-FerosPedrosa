package main

import (
	"fmt"
	"math"
)

func main() {
	var ax, ay, bx, by, cx, cy float64

	fmt.Print("Masukkan koordinat titik A (x y): ")
	fmt.Scan(&ax, &ay)
	fmt.Print("Masukkan koordinat titik B (x y): ")
	fmt.Scan(&bx, &by)
	fmt.Print("Masukkan koordinat titik C (x y): ")
	fmt.Scan(&cx, &cy)

	AB := math.Sqrt(math.Pow(bx-ax, 2) + math.Pow(by-ay, 2))
	BC := math.Sqrt(math.Pow(cx-bx, 2) + math.Pow(cy-by, 2))
	CA := math.Sqrt(math.Pow(ax-cx, 2) + math.Pow(ay-cy, 2))

	sisiTerpanjang := math.Max(AB, math.Max(BC, CA))

	fmt.Printf("Sisi terpanjang dari segitiga tersebut adalah: %.2f\n", sisiTerpanjang)
}
