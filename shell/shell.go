package shell

import "fmt"
import "bufio"
import "os"
import "strings"

import "./twil"

func Shell() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		words := strings.Fields(text)
		switch {
		case (words[0] == "quit" || words[0] == "exit"):
			os.Exit(3)
		case words[0] == "twil":
			fmt.Println(twil.SendText(words[1]))
		case words[0] == "clear":
			fmt.Println("\033[2J")
		default:
			fmt.Println(words)
		}
	}

	return "exit"
}
