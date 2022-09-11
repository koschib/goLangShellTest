package main

// create a shell with os package and execute the command

import (
	"fmt"
	"os/exec"
)

// shell function ls -l, ls, mkdir, cat, exit

func shell() {

	for {

		fmt.Print("koschishell> ")

		var cmd string

		fmt.Scanln(&cmd)

		switch cmd {

		case "ls":

			out, err := exec.Command("ls").Output()

			if err != nil {

				fmt.Printf("\033[31m%s\033[0m ", err)

			}

			fmt.Println(string(out))

		case "ls -l":

			out, err := exec.Command("ls", "-l").Output()

			if err != nil {

				fmt.Printf("\033[31m%s\033[0m ", err)

			}

			fmt.Println(string(out))

		case "mkdir":

			fmt.Print("Enter directory name: ")

			var dir string

			fmt.Scanln(&dir)

			out, err := exec.Command("mkdir", dir).Output()

			if err != nil {

				fmt.Printf("\033[31m%s\033[0m ", err)

			}

			fmt.Println(string(out))

		case "cat":

			fmt.Print("Enter file name: ")

			var file string

			fmt.Scanln(&file)

			out, err := exec.Command("cat", file).Output()

			if err != nil {

				// fmt print error in red color
				fmt.Printf("\033[31m%s\033[0m ", err)

			}

			fmt.Println(string(out))

		case "exit":

			return

		default:

			fmt.Println("Invalid command")

		}

	}

}

func main() {

	shell()

}
