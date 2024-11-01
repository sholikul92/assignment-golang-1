package main

import (
	"a21hc3NpZ25tZW50/helper"
	"fmt"
	"strings"
	// "a21hc3NpZ25tZW50/helper"
)

var Students string = "A1234_Aditira_TI, B2131_Dito_TK, A3455_Afis_MI"
var StudentStudyPrograms string = "TI_Teknik Informatika, TK_Teknik Komputer, SI_Sistem Informasi, MI_Manajemen Informasi"

func Login(id string, name string) string {
	if id == "" || name == "" {
		return "ID or Name is undefined!"
	}

	if len(id) != 5 {
		return "ID must be 5 characters long!"
	}

	students := strings.Split(Students, ", ")

	var nameStudent string
	var prody string

	for _, student := range students {
		idStudentAvailable := strings.Contains(student, id)
		nameStudentAvailable := strings.Contains(student, name)

		if idStudentAvailable && nameStudentAvailable {
			nameStudent = strings.Split(student, "_")[1]
			prody = strings.Split(student, "_")[2]
			break
		}
	}

	if nameStudent == "" && prody == "" {
		return "Login gagal: data mahasiswa tidak ditemukan"
	}

	result := fmt.Sprintf("Login berhasil: %s (%s)", nameStudent, prody)

	return result // TODO: replace this
}

func Register(id string, name string, major string) string {
	return "" // TODO: replace this
}

func GetStudyProgram(code string) string {
	return "" // TODO: replace this
}

func main() {
	fmt.Println("Selamat datang di Student Portal!")

	for {
		helper.ClearScreen()
		fmt.Println("Students: ", Students)
		fmt.Println("Student Study Programs: ", StudentStudyPrograms)

		fmt.Println("\nPilih menu:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Program Study")
		fmt.Println("4. Keluar")

		var pilihan string
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			helper.ClearScreen()
			var id, name string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)

			fmt.Println(Login(id, name))

			helper.Delay(5)
		case "2":
			helper.ClearScreen()
			var id, name, jurusan string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)
			fmt.Print("Masukkan jurusan: ")
			fmt.Scan(&jurusan)
			fmt.Println(Register(id, name, jurusan))

			helper.Delay(5)
		case "3":
			helper.ClearScreen()
			var kode string
			fmt.Print("Masukkan kode: ")
			fmt.Scan(&kode)

			fmt.Println(GetStudyProgram(kode))
			helper.Delay(5)
		case "4":
			fmt.Println("Terima kasih telah menggunakan Student Portal.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// func main for debugging
// func main() {
// 	fmt.Println(Login("B2131", "Dito"))
// }
