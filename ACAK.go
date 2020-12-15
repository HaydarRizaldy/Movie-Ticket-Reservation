func menu() {
	fmt.Println("-----------------------------------------------------")
	fmt.Print("Menu:\n1. REGISTRATION\n2. RESERVATION\n3. PRINT ORDER\n4. CHECK BALANCE AND REWARD POINT\n5. SORT REWARD POINT\n*Enter any number to quit program\n\nChoose: ")
	fmt.Scan(&choose)
	for choose == 1 || choose == 2 || choose == 3 || choose == 4 {//|| choose == 5 {
		switch choose {
		case 1 : Registration()
		case 2 : Reservation()
		case 3 : PrintOrder()
		case 4 : CetakBlRp()
		//case 5 : SortRewardPoint()
		}
	}
}

}
		if choose == 1 {
			fmt.Printf("\nSelected Movie %d: Action", j+1)
			data[i].pesan.movies[j] = "Action"
		} else if choose == 2 {
			fmt.Printf("\nSelected Movie %d: Comedy", j+1)
			data[i].pesan.movies[j] = "Comedy"
		} else if choose == 3 {
			fmt.Printf("\nSelected Movie %d: Romance", j+1)
			data[i].pesan.movies[j] = "Romance"
		} else if choose == 4 {
			fmt.Printf("\nSelected Movie %d: Sci-Fi", j+1)
			data[i].pesan.movies[j] = "Sci-Fi"
		} else if choose == 5 {
			fmt.Printf("\nSelected Movie %d: Horror", j+1)
			data[i].pesan.movies[j] = "Horror"
		}

for i := 0; i < i_tertinggi; i++ {
		fmt.Print(peti[i].berat, " ")


func SortAscending() {
	for i := 0; i < banyakdata; i++ {
		j := i
		for j>0 && (*array)[j].travel < (*array)[j-1].travel {
			(*array)[j], (*array)[j-1] = (*array)[j-1], (*array)[j]
			j = j-1
		}
	}	
}

func SortAscending() {
	for i := 0; i <= i_tertinggi; i++ {
		minIndex := i
		for j := i + 1; j <= i_tertinggi; j++ {
			if (*data)[minIndex].rpoint > (*data)[j].rpoint {
				minIndex = j
			}
		}
    temp := (*data)[i].rpoint
    (*data)[i] = (*data)[minIndex].rpoint
    (*data)[minIndex].rpoint = temp
	}
}