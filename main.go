package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(generateIdPengajuanSb())
}

func generateIdPengajuanSb() (idPengajuanSb string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomNumber := r.Intn(90000) + 10000 // Rentang angka 10000-99999

	// Format ID dengan angka random dan tanggal
	idPengajuanSb = fmt.Sprintf("SB-%05d-%s", randomNumber, time.Now().Format("20060102"))
	return idPengajuanSb
}