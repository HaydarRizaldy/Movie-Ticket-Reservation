package main

import (
	"fmt"
	"os"
	"os/exec"
)

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

func QuitProgram() {
	deleteFile()
	ClearCMD()
	fmt.Println("*DataMember.txt has been deleted.")
}

func ClearCMD() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	QuitProgram()
}