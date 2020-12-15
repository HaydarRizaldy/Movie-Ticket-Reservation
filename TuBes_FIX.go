package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

type datamember struct {
	userid, balance, rpoint, bill, jt int
	username, phonum                  string
	movies                            [100]string
}

var (
	data                                                                              [100]datamember
	day                                                                               [7]string
	h, m, s, waktu                                                                    float64
	un, d, pn                                                                         string
	id, jtiket, bl, rp, choose, movie, harga, i_terakhir, i_terakhir_rsv, i_tertinggi int
)

func main() {
	Opening()
	QuitProgram()
}

func Opening() {
	var jawab string

	fmt.Println("\n==============================================================================")
	fmt.Println()
	fmt.Println("        ██████╗██╗███╗   ██╗███████╗███╗   ███╗ █████╗      ██╗██████╗          ")
	fmt.Println("       ██╔════╝██║████╗  ██║██╔════╝████╗ ████║██╔══██╗    ███║╚════██╗         ")
	fmt.Println("       ██║     ██║██╔██╗ ██║█████╗  ██╔████╔██║███████║    ╚██║ █████╔╝         ")
	fmt.Println("       ██║     ██║██║╚██╗██║██╔══╝  ██║╚██╔╝██║██╔══██║     ██║██╔═══╝          ")
	fmt.Println("       ╚██████╗██║██║ ╚████║███████╗██║ ╚═╝ ██║██║  ██║     ██║███████╗         ")
	fmt.Println("        ╚═════╝╚═╝╚═╝  ╚═══╝╚══════╝╚═╝     ╚═╝╚═╝  ╚═╝     ╚═╝╚══════╝         ")
	fmt.Println("                         - Movie Ticket Reservation -                         \n")
	fmt.Println("==============================================================================\n")
	fmt.Println("- REGISTRATION -")
	fmt.Println("\n> Input")
	i := 0
	i_terakhir = i
	i_tertinggi = i_terakhir
	data[i].userid = i
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
	data[i].rpoint = 0
	createFile()
	fmt.Print("\nApakah mau registrasi member baru lagi (Ya/Tidak)? ")
	fmt.Scan(&jawab)
	for jawab != "Ya" && jawab != "Tidak" {
		fmt.Println("\n*Maaf, kata kunci yang ada masukkan salah. Coba ulangi lagi.\n")
		fmt.Print("Apakah mau registrasi member baru lagi (Ya/Tidak)? ")
		fmt.Scan(&jawab)
	}
	for jawab == "Ya" {
		i = i + 1
		i_terakhir = i
		i_tertinggi = i_terakhir
		data[i].userid = i
		fmt.Print("\nUsername      : ")
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
		data[i].rpoint = 0
		createFile()
		fmt.Print("\nApakah mau registrasi member baru lagi (Ya/Tidak)? ")
		fmt.Scan(&jawab)
		for jawab != "Ya" && jawab != "Tidak" {
			fmt.Println("\n*Maaf, kata kunci yang ada masukkan salah. Coba ulangi lagi.\n")
			fmt.Print("Apakah mau registrasi member baru lagi (Ya/Tidak)? ")
			fmt.Scan(&jawab)
		}
	}
	fmt.Println()
	menu()
}

func ClearCMD() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func QuitProgram() {
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("\n*Terimakasih telah menggunakan layanan kami.")
	time.Sleep(1 * time.Second)
	fmt.Println("\n==============================================================================")
	fmt.Println()
	fmt.Println("        ██████╗██╗███╗   ██╗███████╗███╗   ███╗ █████╗      ██╗██████╗          ")
	fmt.Println("       ██╔════╝██║████╗  ██║██╔════╝████╗ ████║██╔══██╗    ███║╚════██╗         ")
	fmt.Println("       ██║     ██║██╔██╗ ██║█████╗  ██╔████╔██║███████║    ╚██║ █████╔╝         ")
	fmt.Println("       ██║     ██║██║╚██╗██║██╔══╝  ██║╚██╔╝██║██╔══██║     ██║██╔═══╝          ")
	fmt.Println("       ╚██████╗██║██║ ╚████║███████╗██║ ╚═╝ ██║██║  ██║     ██║███████╗         ")
	fmt.Println("        ╚═════╝╚═╝╚═╝  ╚═══╝╚══════╝╚═╝     ╚═╝╚═╝  ╚═╝     ╚═╝╚══════╝         ")
	fmt.Println("                         - Movie Ticket Reservation -                         \n")
	fmt.Println("==============================================================================\n")
	time.Sleep(1 * time.Second)
	fmt.Println("Deleting DataMember.txt...\n")
	time.Sleep(1 * time.Second)
	deleteFile()
	fmt.Println("*DataMember.txt has been deleted.\n")
	time.Sleep(1 * time.Second)
	fmt.Println("Quitting program in 3 seconds...")
	time.Sleep(3 * time.Second)
	ClearCMD()
}

func createFile() {
	file, err := os.Create("DataMember.txt")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close()
	file.WriteString(fmt.Sprintln("------------------------------------------------------------------------------"))
	for i := 0; i <= i_tertinggi; i++ {
		file.WriteString(fmt.Sprintf("# Data User Ke-%d\n\n", i+1))
		file.WriteString(fmt.Sprintf("Username          : %s \n", data[i].username))
		file.WriteString(fmt.Sprintf("Phone Number      : %s \n", data[i].phonum))
		file.WriteString(fmt.Sprint("Selected Movie(s) : "))
		for j := 0; j < data[i].jt; j++ {
			file.WriteString(fmt.Sprintf("(%v) ", data[i].movies[j]))
		}
		file.WriteString(fmt.Sprintf("\nBalance           : %v \n", data[i].balance))
		file.WriteString(fmt.Sprintf("Reward Points     : %v \n", data[i].rpoint))
		file.WriteString(fmt.Sprintf("Bills             : %v \n", data[i].bill))
		file.WriteString(fmt.Sprintln("------------------------------------------------------------------------------"))
	}
}

var path = "DataMember.txt"

func deleteFile() {
	// delete file
	var err = os.Remove(path)
	if isError(err) {
		return
	}
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func menu() {
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Print("Menu:\n1. REGISTRATION\n2. RESERVATION\n3. PRINT ORDER\n4. CHECK BALANCE AND REWARD POINT\n5. SORT MEMBERS BY REWARD POINT\n*Enter any other number to quit program\n\nChoose: ")
	fmt.Scan(&choose)
	for choose == 1 || choose == 2 || choose == 3 || choose == 4 || choose == 5 {
		switch choose {
		case 1:
			Registration()
		case 2:
			Reservation()
		case 3:
			PrintOrder()
		case 4:
			CetakBlRp()
		case 5: 
			SortRewardPoint()
		}
	}
}

func Registration() {
	var jawab string

	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("\n- REGISTRATION -")
	fmt.Println("\n> Input")
	i := i_terakhir + 1
	i_terakhir = i
	i_tertinggi = i_terakhir
	data[i].userid = i
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
	data[i].rpoint = 0
	createFile()
	fmt.Print("\nApakah mau registrasi member baru lagi (Ya/Tidak)? ")
	fmt.Scan(&jawab)
	for jawab != "Ya" && jawab != "Tidak" {
		fmt.Println("\n*Maaf, kata kunci yang ada masukkan salah. Coba ulangi lagi.\n")
		fmt.Print("Apakah mau registrasi member baru lagi (Ya/Tidak)? ")
		fmt.Scan(&jawab)
	}
	for jawab == "Ya" {
		i = i + 1
		i_terakhir = i
		i_tertinggi = i_terakhir
		data[i].userid = i
		fmt.Print("\nUsername      : ")
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
		data[i].rpoint = 0
		createFile()
		fmt.Print("\nApakah mau registrasi member baru lagi (Ya/Tidak)? ")
		fmt.Scan(&jawab)
		for jawab != "Ya" && jawab != "Tidak" {
			fmt.Println("\n*Maaf, kata kunci yang ada masukkan salah. Coba ulangi lagi.\n")
			fmt.Print("Apakah mau registrasi member baru lagi (Ya/Tidak)? ")
			fmt.Scan(&jawab)
		}
	}
	fmt.Println()
	menu()
}

func CetakBlRp() {
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("\n- CHECK BALANCE AND REWARD POINT -\n")
	fmt.Println("> Check")
	fmt.Print("Username     : ")
	fmt.Scan(&un)
	for i := 0; i < len(data); i++ {
		//for un != data[i].username {
		//	fmt.Println("*Data tidak ditemukan. Silahkan coba lagi.\n")
		//	fmt.Print("Username     : ")
		//	fmt.Scan(&un)
		//}
		if un == data[i].username {
			fmt.Printf("\n# Data Member\n")
			fmt.Print("Username     : ")
			fmt.Println(data[i].username)
			fmt.Print("Balance      : Rp")
			fmt.Println(data[i].balance)
			fmt.Print("Reward Point : ")
			fmt.Println(data[i].rpoint)
			fmt.Println("\n*Data member a/n", data[i].username, "telah tercetak.\n")
		}
	}
	fmt.Println("*Cek saldo dan reward point selesai.\n")
	menu()
}

func Reservation() {

	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("\n- RESERVATION -")
	fmt.Println("\n> Validation")
	fmt.Print("Username      : ")
	fmt.Scan(&un)
	fmt.Print("Phone Number  : ")
	fmt.Scan(&pn)
	i_terakhir_rsv = 0
	for i := 0; i < len(data); i++ {
		if un == data[i].username && pn == data[i].phonum {
			fmt.Println("\n# Data Member")
			fmt.Print("Username     : ")
			fmt.Println(data[i].username)
			fmt.Print("Phone Number : ")
			fmt.Println(data[i].phonum)
			fmt.Print("Balance      : Rp")
			fmt.Println(data[i].balance)
			fmt.Print("Reward Point : ")
			fmt.Println(data[i].rpoint)
			fmt.Println("\n*Validasi berhasil. Silahkan melanjutkan ke reservasi tiket.")
			i_terakhir_rsv = i
		}
	}
	i := i_terakhir_rsv
	fmt.Println("\n> Reservation Time")
	fmt.Print("Hour (HH MM SS): ")
	fmt.Scan(&h, &m, &s)
	waktu = s + (m * 60) + (h * 3600)
	if (waktu > 0*3600) && (waktu <= 11*3600) {
		fmt.Print("Day            : ")
		fmt.Scan(&d)
		if (d == "Senin") || (d == "senin") {
			fmt.Print("\n*Reservasi bisa dilakukan.\n")
			day[i] = d
			pilihfilmHariBiasa()
			i_terakhir_rsv = i
		} else if (d == "Selasa") || (d == "selasa") {
			fmt.Print("\n*Reservasi bisa dilakukan.\n")
			day[i] = d
			pilihfilmHariBiasa()
			i_terakhir_rsv = i
		} else if (d == "Rabu") || (d == "rabu") {
			fmt.Print("\n*Reservasi bisa dilakukan.\n")
			day[i] = d
			pilihfilmHariBiasa()
			i_terakhir_rsv = i
		} else if (d == "Kamis") || (d == "kamis") {
			fmt.Print("\n*Reservasi bisa dilakukan.\n")
			day[i] = d
			pilihfilmHariBiasa()
			i_terakhir_rsv = i
		} else if (d == "Jumat") || (d == "jumat") {
			fmt.Print("\n*Reservasi bisa dilakukan. \n*Pada hari Jumat, tiket akan dikenakan biaya lebih sebesar Rp10000.\n")
			day[i] = d
			pilihfilmHariJSM()
			i_terakhir_rsv = i
		} else if (d == "Sabtu") || (d == "sabtu") {
			fmt.Print("\n*Reservasi bisa dilakukan. \n*Pada hari Sabtu, tiket akan dikenakan biaya lebih sebesar Rp10000.\n")
			day[i] = d
			pilihfilmHariJSM()
			i_terakhir_rsv = i
		} else if (d == "Minggu") || (d == "minggu") {
			fmt.Print("\n*Reservasi bisa dilakukan. \n*Pada hari Minggu, tiket akan dikenakan biaya lebih sebesar Rp10000.\n")
			day[i] = d
			pilihfilmHariJSM()
			i_terakhir_rsv = i
		} else {
			fmt.Print("*Maaf, input hari salah. [ERROR]\n")
		}
	} else if waktu == 0 {
		fmt.Print("*Maaf, input jam salah. [ERROR]\n")
	} else {
		fmt.Print("*Maaf, anda telah melewati batas waktu reservasi. Silakan reservasi di antara jam 00.00 - 11.00.\n")
	}
}

const harga_normal = 35000

func pilihfilmHariBiasa() {
	fmt.Print("\n*Anda berhak mendapatkan 10 reward points dari setiap tiket yang dipesan.\n")
	fmt.Print("Berapa tiket film yang ingin dipesan (Max. 2)? ")
	fmt.Scan(&jtiket)
	for jtiket <= 0 || jtiket > 2 {
		if jtiket <= 0 {
			fmt.Print("*Maaf, jumlah tiket film yang ingin dipesan tidak bisa 0 atau kurang.\n")
			fmt.Print("\nBerapa tiket film yang ingin dipesan?")
			fmt.Scan(&jtiket)
		} else if jtiket > 2 {
			fmt.Print("*Maaf, maksimal jumlah tiket yang dapat dipesan adalah 2.\n")
			fmt.Print("\nBerapa tiket film yang ingin dipesan?")
			fmt.Scan(&jtiket)
		}
	}
	i := i_terakhir_rsv
	data[i].jt = 0
	data[i].jt = jtiket
	harga = 0
	for j := 0; j < data[i].jt; j++ {
		fmt.Print("\nMovie:\n1. Action\n2. Comedy\n3. Romance\n4. Sci-Fi\n5. Horror\n\nSelect: ")
		fmt.Scan(&choose)
		for choose != 1 && choose != 2 && choose != 3 && choose != 4 && choose != 5 {
			fmt.Println("\n*Maaf, pilihan tidak tersedia.\n")
			fmt.Print("\nMovie:\n1. Action\n2. Comedy\n3. Romance\n4. Sci-Fi\n5. Horror\n\nSelect: ")
			fmt.Scan(&choose)
		}
		if choose == 1 {
			fmt.Printf("\nSelected Movie (%d): Action\n", j+1)
			data[i].movies[j] = "Action"
		} else if choose == 2 {
			fmt.Printf("\nSelected Movie (%d): Comedy\n", j+1)
			data[i].movies[j] = "Comedy"
		} else if choose == 3 {
			fmt.Printf("\nSelected Movie (%d): Romance\n", j+1)
			data[i].movies[j] = "Romance"
		} else if choose == 4 {
			fmt.Printf("\nSelected Movie (%d): Sci-Fi\n", j+1)
			data[i].movies[j] = "Sci-Fi"
		} else if choose == 5 {
			fmt.Printf("\nSelected Movie (%d): Horror\n", j+1)
			data[i].movies[j] = "Horror"
		}
		harga = harga + harga_normal
		data[i].rpoint = data[i].rpoint + 10
	}
	data[i].bill = harga
	data[i].balance = data[i].balance - data[i].bill
	fmt.Println("\n*Pemilihan tiket film selesai.\n")
	createFile()
	menu()
}

const harga_tambahan = 10000

func pilihfilmHariJSM() {
	fmt.Print("\n*Anda berhak mendapatkan 10 reward points dari setiap tiket yang dipesan.\n")
	fmt.Print("Berapa tiket film yang ingin dipesan (Max. 2)? ")
	fmt.Scan(&jtiket)
	for jtiket <= 0 || jtiket > 2 {
		if jtiket <= 0 {
			fmt.Print("*Maaf, jumlah tiket film yang ingin dipesan tidak bisa 0 atau kurang.\n")
			fmt.Print("\nBerapa tiket film yang ingin dipesan (Max. 2)? ")
			fmt.Scan(&jtiket)
		} else if jtiket > 2 {
			fmt.Print("*Maaf, maksimal jumlah tiket yang dapat dipesan adalah 2.\n")
			fmt.Print("\nBerapa tiket film yang ingin dipesan (Max. 2)? ")
			fmt.Scan(&jtiket)
		}
	}
	i := i_terakhir_rsv
	data[i].jt = 0
	data[i].jt = jtiket
	harga = 0
	for j := 0; j < data[i].jt; j++ {
		fmt.Print("\nMovie:\n1. Action\n2. Comedy\n3. Romance\n4. Sci-Fi\n5. Horror\n\nSelect: ")
		fmt.Scan(&choose)
		for choose != 1 && choose != 2 && choose != 3 && choose != 4 && choose != 5 {
			fmt.Println("\n*Maaf, pilihan tidak tersedia.\n")
			fmt.Print("\nMovie:\n1. Action\n2. Comedy\n3. Romance\n4. Sci-Fi\n5. Horror\n\nSelect: ")
			fmt.Scan(&choose)
		}
		if choose == 1 {
			fmt.Printf("\nSelected Movie (%d): Action\n", j+1)
			data[i].movies[j] = "Action"
		} else if choose == 2 {
			fmt.Printf("\nSelected Movie (%d): Comedy\n", j+1)
			data[i].movies[j] = "Comedy"
		} else if choose == 3 {
			fmt.Printf("\nSelected Movie (%d): Romance\n", j+1)
			data[i].movies[j] = "Romance"
		} else if choose == 4 {
			fmt.Printf("\nSelected Movie (%d): Sci-Fi\n", j+1)
			data[i].movies[j] = "Sci-Fi"
		} else if choose == 5 {
			fmt.Printf("\nSelected Movie (%d): Horror\n", j+1)
			data[i].movies[j] = "Horror"
		}
		harga = harga + harga_normal + harga_tambahan
		data[i].rpoint = data[i].rpoint + 10
	}
	data[i].bill = harga
	data[i].balance = data[i].balance - data[i].bill
	fmt.Println("\n*Pemilihan tiket film selesai.\n")
	createFile()
	menu()
}

func PrintOrder() {
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("\n- PRINT ORDER -\n")
	fmt.Print("Hari pemesanan tiket: ")
	fmt.Scan(&d)
	fmt.Printf("\nBon pemesanan tiket pada hari %s:\n\n", d)
	k := 1
	for i := 0; i < len(day); i++ {
		if d == day[i] {
			fmt.Printf("# Bon Pemesanan User %d\n", k)
			fmt.Print("Username       : ")
			fmt.Println(data[i].username)
			fmt.Print("Phone Number   : ")
			fmt.Println(data[i].phonum)
			fmt.Print("Selected Movies: ")
			for j := 0; j < data[i].jt; j++ {
				fmt.Printf("(%s) ", data[i].movies[j])
			}
			fmt.Print("\nBills          : ")
			fmt.Println(data[i].bill)
			fmt.Println()
		}
		k++
	}
	fmt.Println("*Bon pemesanan telah tercetak.\n")
	menu()
}

func SortRewardPoint() {
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("\n- SORT MEMBERS BY REWARD POINT -\n")
	fmt.Print("Sort:\n1. Ascending\n2. Descending\n\nChoose: "); fmt.Scan(&choose)
	switch choose {
	case 1:
		SortAscending()
		j := 1
		for i := 0; i <= i_tertinggi; i++ {
			fmt.Printf("\nNo. %d\n", j)
			fmt.Print("Username     : ")
			fmt.Println(data[i].username)
			fmt.Print("Reward Point : ")
			fmt.Println(data[i].rpoint)
			j++
		}
	case 2:
		SortDescending()
		j := 1
		for i := 0; i <= i_tertinggi; i++ {
			fmt.Printf("\nNo. %d\n", j)
			fmt.Print("Username     : ")
			fmt.Println(data[i].username)
			fmt.Print("Reward Point : ")
			fmt.Println(data[i].rpoint)
			j++
		}
	}
	fmt.Println("\n*Sorting reward point telah selesai.\n")
	menu()
}

func SortAscending() {
	for i := 0; i <= i_tertinggi; i++ {
		minIndex := i
		for j := i; j <= i_tertinggi; j++ {
			if data[minIndex].rpoint > data[j].rpoint {
				minIndex = j
			}
		}
	    a := data[i].userid
	    data[i].userid = data[minIndex].userid
	    data[minIndex].userid = a

	    b := data[i].username
	    data[i].username = data[minIndex].username
	    data[minIndex].username = b
		
	    c := data[i].phonum
	    data[i].phonum = data[minIndex].phonum
	    data[minIndex].phonum = c
		
	    d := data[i].balance
	    data[i].balance = data[minIndex].balance
	    data[minIndex].balance = d
		
	    e := data[i].rpoint
	    data[i].rpoint = data[minIndex].rpoint
	    data[minIndex].rpoint = e

	    for j := 0; j < data[i].jt; j++ {
		    f := data[i].movies[j]
		    data[i].movies[j] = data[minIndex].movies[j]
		    data[minIndex].movies[j] = f
		}

		g := data[i].bill
	    data[i].bill = data[minIndex].bill
	    data[minIndex].bill = g
	}
	createFile()
}

func SortDescending() {
	for i := 0; i <= i_tertinggi; i++ {
		maxIndex := i
		for j := i; j <= i_tertinggi; j++ {
			if data[maxIndex].rpoint < data[j].rpoint {
				maxIndex = j
			}
		}
	    a := data[i].userid
	    data[i].userid = data[maxIndex].userid
	    data[maxIndex].userid = a
		
	    b := data[i].username
	    data[i].username = data[maxIndex].username
	    data[maxIndex].username = b
		
	    c := data[i].phonum
	    data[i].phonum = data[maxIndex].phonum
	    data[maxIndex].phonum = c
		
	    d := data[i].balance
	    data[i].balance = data[maxIndex].balance
	    data[maxIndex].balance = d
		
	    e := data[i].rpoint
	    data[i].rpoint = data[maxIndex].rpoint
	    data[maxIndex].rpoint = e

	    for j := 0; j < data[i].jt; j++ {
		    f := data[i].movies[j]
		    data[i].movies[j] = data[maxIndex].movies[j]
		    data[maxIndex].movies[j] = f
		}

		g := data[i].bill
	    data[i].bill = data[maxIndex].bill
	    data[maxIndex].bill = g
	}
	createFile()
}