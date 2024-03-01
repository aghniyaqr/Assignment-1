package main

import (
	"fmt"
	"os"
)

type Biodata struct {
	name    string
	address string
	job     string
	reason  string
}

func printBiodata(p Biodata) {
	fmt.Println("Nama        :", p.name)
	fmt.Println("Alamat      :", p.address)
	fmt.Println("Pekerjaan   :", p.job)
	fmt.Println("Alasan      :", p.reason)
}

func main() {
	args := os.Args
	fmt.Println("Nomor absen :", args[1:])

	if len(os.Args) < 2 {
		fmt.Println("Masukan nomor absen 1-8")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "1":
		biodata1 := Biodata{
			name:    "Aghniya Qinthara",
			address: "Pangalengan",
			job:     "Guru Privat",
			reason:  "Switch Career",
		}
		printBiodata(biodata1)
	case "2":
		biodata2 := Biodata{
			name:    "Allisha Febian",
			address: "Dago",
			job:     "PCB Designer",
			reason:  "Tambah Skill",
		}
		printBiodata(biodata2)
	case "3":
		biodata3 := Biodata{
			name:    "Aprilia Lestari",
			address: "Padalarang",
			job:     "Mahasiswa",
			reason:  "Isi Waktu libur",
		}
		printBiodata(biodata3)
	case "4":
		biodata4 := Biodata{
			name:    "Damayanti Afiani",
			address: "Soreang",
			job:     "Instrument Engineer",
			reason:  "Tambah Skill",
		}
		printBiodata(biodata4)
	case "5":
		biodata5 := Biodata{
			name:    "Lintang Cahaya",
			address: "Karawang",
			job:     "Finance",
			reason:  "Tambah Skill",
		}
		printBiodata(biodata5)
	case "6":
		biodata6 := Biodata{
			name:    "Novianti",
			address: "Tasikmalaya",
			job:     "Mahasiswa",
			reason:  "Isi Waktu Libur",
		}
		printBiodata(biodata6)
	case "7":
		biodata7 := Biodata{
			name:    "Tarrisa Fitri",
			address: "Cimahi",
			job:     "Wiraswasta",
			reason:  "Tambah Skill",
		}
		printBiodata(biodata7)
	case "8":
		biodata8 := Biodata{
			name:    "Yuspa Sarah",
			address: "Garut",
			job:     "Electronic Engineer",
			reason:  "Tambah Skill",
		}
		printBiodata(biodata8)
	default:
		fmt.Println("Biodata tidak ada. Hanya masukan nomor absen 1-8")
		os.Exit(1)
	}
}
