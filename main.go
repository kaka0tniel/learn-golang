package main

import (
	"fmt"
	"os"
)

type Biodata struct {
	ID        int
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

type BiodataService struct {
	biodatas []Biodata
}

func (s *BiodataService) PrintBiodatas() {
	fmt.Println("Daftar Biodata:")
	for _, biodata := range s.biodatas {
		fmt.Printf("ID: %d, Nama: %s, alamat: %s, pekerjaan: %s, alasan: %s\n", biodata.ID, biodata.nama, biodata.alamat, biodata.pekerjaan, biodata.alasan)
	}
}

func (s *BiodataService) AddBiodata(newBiodata Biodata) {
	newBiodata.ID = len(s.biodatas) + 1
	s.biodatas = append(s.biodatas, newBiodata)
	fmt.Println("Biodata berhasil ditambahkan.")
}

func main() {
	biodataService := BiodataService{}

	biodataService.AddBiodata(Biodata{nama: "Messi", alamat: "Argentina", pekerjaan: "Pemain bola", alasan: "to be GOAT"})
	biodataService.AddBiodata(Biodata{nama: "Taufik Hidayat", alamat: "Indonesia", pekerjaan: "pemain badminton", alasan: "untuk menjadi bintang iklan sonice"})

	biodataService.PrintBiodatas()

	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go <Nama> <Alamat> <Pekerjaan> <Alasan>")
		return
	}

	namaStr := os.Args[1]
	alamatStr := os.Args[2]
	pekerjaanStr := os.Args[3]
	alasanStr := os.Args[4]

	biodataService.AddBiodata(Biodata{nama: namaStr, alamat: alamatStr, pekerjaan: pekerjaanStr, alasan: alasanStr})

	biodataService.PrintBiodatas()

}
