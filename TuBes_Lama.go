package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var clear map[string]func()

type datamember struct {
	userid, balance, rpoint, bill, jt int
	username, phonum                  string
	pesan                             reservation
}

type reservation struct {
	hour   float64
	day    string
	movies [200]string
}

var (
	data                                [1000]datamember
	h, m, s, waktu                      float64
	un, d, pn                           string
	id, jdata, jtiket, bl, rp, c, movie int
)

func main() {
	fmt.Println("\n=====================================================")
	fmt.Println("          Cinema12 Movie Ticket Reservation        ")
	fmt.Println("=====================================================\n")
	imd()
}

func menu() {
	fmt.Println("-----------------------------------------------------")
	fmt.Print("Menu:\n1. RESERVATION\n2. PRINT ORDER\n3. CHECK BALANCE AND REWARD POINT\n4. QUIT MENU\n\nChoose: ")
	fmt.Scan(&c)
	for c != 1 && c != 2 && c != 3 && c != 4 {
		fmt.Println("\n*Maaf, pilihan tidak tersedia.\n")
		fmt.Print("Menu:\n1. RESERVATION\n2. PRINT ORDER\n3. CHECK BALANCE AND REWARD POINT\n4. QUIT MENU\n\nChoose: ")
		fmt.Scan(&c)
	}
	if c == 1 {
		rsv()
	} else if c == 2 {
		printorder()
	} else if c == 3 {
		cetakblrp()
	} else if c == 4 {
		fmt.Println("-----------------------------------------------------")
		fmt.Println("\n*Terimakasih telah menggunakan layanan kami.\n")
		QuitProgram()
	}
}

func ClearCMD() {
	cmd := exec.Command("cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func QuitProgram() {
	fmt.Println("Quitting program in 3 seconds...")
	time.Sleep(3 * time.Second)
	ClearCMD()
	fmt.Println("Done!")
	ClearCMD()
	fmt.Println("Done!")
	ClearCMD()
	fmt.Println("Done!")
	ClearCMD()
	fmt.Println("Done!")
	ClearCMD()
	fmt.Println("Done!")
	ClearCMD()
	fmt.Println("Done!")
	ClearCMD()
	fmt.Println("Done!")
}

func imd() {
	fmt.Println("- INPUT MEMBER DATA -")
	fmt.Print("\nBerapa data member yang ingin diinput? ")
	fmt.Scan(&jdata)
	fmt.Println("\n*Harap input User ID sesuai urutan (1-10).\n")
	for jdata <= 0 {
		fmt.Print("*Maaf, jumlah input data member tidak bisa bernilai 0 atau kurang.\n")
		fmt.Print("\nBerapa data member yang ingin diinput? ")
		fmt.Scan(&jdata)
	}
	i := 0
	for i < jdata {
		fmt.Print("User ID       : ")
		fmt.Scan(&id)
		data[i].userid = id
		fmt.Print("Username      : ")
		fmt.Scan(&un)
		data[i].username = un
		fmt.Print("Phone Number  : ")
		fmt.Scan(&pn)
		data[i].phonum = pn
		fmt.Print("Balance       : Rp")
		fmt.Scan(&bl)
		data[i].balance = bl
		for bl <= 0 {
			fmt.Print("*Maaf, input saldo tidak bisa kosong atau kurang dari 0 rupiah.\n")
			fmt.Print("Balance       : Rp")
			fmt.Scan(&bl)
			data[i].balance = bl
		}
		fmt.Print("Reward Point  : ")
		fmt.Scan(&rp)
		data[i].rpoint = rp
		for rp < 0 {
			fmt.Print("*Maaf, input reward point tidak bisa kurang dari 0.\n")
			fmt.Print("Reward Point  : ")
			fmt.Scan(&rp)
			data[i].rpoint = rp
		}
		fmt.Printf("*Data member %v telah terdaftar.\n\n", un)
		createFile()
		i++
	}
}