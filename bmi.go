package main

import "fmt"

func main() {
	var bmi, tinggiBadan, beratBadan float64

	fmt.Print("Masukkan nilai BMI: ")
	fmt.Scan(&bmi)
	fmt.Print("Masukkan tinggi badan (m): ")
	fmt.Scan(&tinggiBadan)
	
	beratBadan = bmi * (tinggiBadan * tinggiBadan)

	fmt.Printf("Berat badan anda adalah: %.f kg\n", beratBadan)
}
