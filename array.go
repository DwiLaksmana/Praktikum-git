package main

import "fmt"

func main() {
//    buat array
    // var nilai [4]int
	// nilai[0] = 88
	// nilai[1] = 76
	
	// suhu := [3]float64{24.5, 25, 23.8}
	// hari := []string{"SEN", "SEL", "RAB"}
	// fmt.Println(nilai, "panjangnya ", len(nilai))
	// fmt.Println(suhu[2])
	// fmt.Println(len(hari))

	// ===========================

// ubah nilai array
	// nilai := [4]int{88, 76, 95, 81}
	// fmt.Println(nilai[3])

	// //print dengan variabel
	// // i:=3
	// // fmt.Println(nilai[i])

	// //mengganti nilai didalam array
	// nilai[0] = 90
	// fmt.Println(nilai[0])
	// fmt.Println(nilai)
	//============================

//  slice
	// angka := make ([]int, 0, 3)
	// fmt.Println(angka, len(angka), cap(angka))
	// angka = append(angka, 10, 20, 30)
	// fmt.Println(angka, len(angka), cap(angka))
	// angka = append(angka, 40)
	// fmt.Println(angka, len(angka), cap(angka))
	//cap akan terus dikali dua jika sudah melebihi kapasitas sebelumnya

//latihan 1
	// nilai := [] int{80, 75, 90}
	// fmt.Println(nilai, len(nilai), cap(nilai))
	// nilai = append(nilai, 85, 95)
	// fmt.Println(nilai, len(nilai), cap(nilai))
	// nilai[0] = 100
	// fmt.Println(nilai, len(nilai), cap(nilai))

// for array
	// nilai := []int{70, 85, 90, 65, 88}

	// for i := 0; i < len(nilai); i++ {
	// 	fmt.Println(i, nilai[i])
	// }

	// for i, v := range nilai {
	// 	fmt.Println(i, v)
	// }

	// for _, v := range nilai {
	// 	fmt.Println(v)
	// }

// agregasikan nilai dalam variabel (menjumlahkan nilai)

	nilai := []float64{78.5, 82, 91.5, 64, 88}
	total := float64(0)
	lulus := 0

	for _, v := range nilai {
		total += v
		if v >= 75 {
			lulus++
		}
	}
	rata := total / float64(len(nilai))
	fmt.Println("Total nilai:", total)
	fmt.Println("Jumlah siswa lulus:", lulus)
	fmt.Println("Rata-rata nilai:", rata)

}