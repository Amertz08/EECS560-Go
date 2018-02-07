package main


import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"github.com/Amertz08/EECS560-go/Lab01/LinkedList"
)

func main() {
	list := LinkedList.NewLinkedList()

	choice := 0
	for choice != 5 {
		choice = usage()

		switch choice {
		case 1:
			value := getIntInput("Insert value: ")
			list.Insert(value)
		case 2:
			value := getIntInput("Delete value: ")
			list.Erase(value)
		case 3:
			value := getIntInput("Find value: ")
			if list.Find(value) {
				fmt.Println("Value found")
			} else {
				fmt.Println("Vale not found")
			}
		case 4:
			list.Print()
		case 5:
			fmt.Println("Exiting...")
		default:
			fmt.Println("Invalid input")
		}
	}
}

func getInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}

func getIntInput(prompt string) int {
	text := getInput(prompt)
	value, _ := strconv.Atoi(text)
	return value
}

func usage() int {
	fmt.Print(
		"1 - Insert\n" +
		"2 - Delete\n" +
		"3 - Find\n" +
		"4 - Print\n" +
		"5 - Exit\n")
	choice := getIntInput("Select a choice: ")
	return choice
}
