package main
import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var clear map[string]func()

func ClearCMD() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func QuitProgram() {
	fmt.Println("Quitting program in 3 seconds...")
	time.Sleep(3 * time.Second)
	ClearCMD()
	fmt.Println("Done!")
}

func main () {
	var n int

	fmt.Print("Input command: "); fmt.Scan(&n)
	for n != 0 {
		fmt.Println("*Wrong command. Please try again.\n")
		fmt.Print("Input command: "); fmt.Scan(&n)
	}
	if n == 0 {
		QuitProgram()
	}
}