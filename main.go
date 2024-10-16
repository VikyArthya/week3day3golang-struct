package main

import "fmt"

// //buat struct kemudian lakukan init sebanyak 3 data useappend
// //buat func filter object struct berdasarkan tahun
// //kondisi : tampilkan 2 data jika tahun lebih dari 2005
// func main() {
// 	var person []Person_data

// 	person = append(person, Person_data{name: "Viky1", year: 2002})
// 	person = append(person, Person_data{name: "Viky2", year: 2007})
// 	person = append(person, Person_data{name: "Viky3", year: 2008})

// 	print_data(person)
// }

// type Person_data struct {
// 	name string
// 	year int
// }

// func print_data(person []Person_data) {
// 	for _, person_ := range person {
// 		if person_.year > 2005 {
// 			fmt.Println("Tahun : ", person_.year)
// 		}
// 	}
// }

type Motor struct {
	Kecepatan int
}

type Mobil struct {
	Kecepatan int
}

type Bajaj struct {
	Kecepatan int
}

func (m Motor) JarakPerLiter() float64 {
	return 3.0
}

func (m Mobil) JarakPerLiter() float64 {
	return 1.0
}

func (b Bajaj) JarakPerLiter() float64 {
	return 4.0
}

func kendaraanPalingEfisien(bensin float64, kendaraan ...interface{}) string {
	maxJarak := 0.0
	namaKendaraan := ""

	for _, k := range kendaraan {
		switch v := k.(type) {
		case Motor:
			jarak := bensin * v.JarakPerLiter()
			if jarak > maxJarak {
				maxJarak = jarak
				namaKendaraan = "Motor"
			}
		case Mobil:
			jarak := bensin * v.JarakPerLiter()
			if jarak > maxJarak {
				maxJarak = jarak
				namaKendaraan = "Mobil"
			}
		case Bajaj:
			jarak := bensin * v.JarakPerLiter()
			if jarak > maxJarak {
				maxJarak = jarak
				namaKendaraan = "Bajaj"
			}
		}
	}

	return namaKendaraan
}

func main() {
	motor := Motor{Kecepatan: 60}
	mobil := Mobil{Kecepatan: 100}
	bajaj := Bajaj{Kecepatan: 50}

	bensin := 10.0

	kendaraanTerhemat := kendaraanPalingEfisien(bensin, motor, mobil, bajaj)
	fmt.Printf("Kendaraan paling efisien dengan %v liter bensin adalah: %v\n", bensin, kendaraanTerhemat)
}
