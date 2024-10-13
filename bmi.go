package main

import "fmt"

func main() {
	var bmi, tinggiBadan, beratBadan float64

	// Meminta input BMI dan tinggi badan
	fmt.Print("Masukkan nilai BMI: ")
	fmt.Scan(&bmi)
	fmt.Print("Masukkan tinggi badan (m): ")
	fmt.Scan(&tinggiBadan)

	// Menghitung berat badan berdasarkan rumus BMI = beratBadan / (tinggiBadan * tinggiBadan)
	beratBadan = bmi * (tinggiBadan * tinggiBadan)

	// Menampilkan hasil berat badan
	fmt.Printf("Berat badan anda adalah: %.f kg\n", beratBadan)
}
