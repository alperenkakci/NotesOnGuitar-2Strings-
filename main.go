package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	var i int
	m := make(map[int]string)
	m[1] = "C-Do"
	m[2] = "D-Re"
	m[3] = "İnce E-Mi"
	m[4] = "F-Fa"
	m[5] = "G-Sol"
	m[6] = "A-La"
	m[7] = "B-Si"
	m[8] = "Kalın E-Mi"

	n := make(map[int]string)
	n[1] = "C-Do"
	n[2] = "D-Re"
	n[3] = "E-Mi"
	n[4] = "F-Fa"
	n[5] = "G-Sol"
	n[6] = "İnce A-La"
	n[7] = "B-Si"
	n[8] = "Kalın A-La"

	finish := false
	// Uyku süresini alıyoruz
	fmt.Println("Notalar arası süreyi belirleyiniz (saniye cinsinden): ")
	fmt.Scan(&i)
	uykuSuresi := time.Duration(i) * time.Second

	for finish == false {
		sayi := rand.IntN(7) + 1
		tel := rand.IntN(2) + 1
		//fmt.Println(sayi, tel)
		if tel == 1 {
			fmt.Println(tel, ": ", m[sayi])
			time.Sleep(uykuSuresi)
		} else if tel == 2 {
			fmt.Println(tel, ": ", n[sayi])
			time.Sleep(uykuSuresi)
		}

	}
}
