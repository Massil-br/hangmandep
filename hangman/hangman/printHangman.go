package hangman

import "fmt"

// fonction qui déssine l'homme pendu
func PrintHangman(attemptsLeft int) {
	switch attemptsLeft {
	case 9:
		fmt.Print(`
         
         
         
         
         
  =========
`)
	case 8:
		fmt.Print(`
         
        |  
        |  
        |  
        |  
  =========
`)
	case 7:
		fmt.Print(`
    +---+  
        |  
        |  
        |  
        |  
  =========
  
`)
	case 6:
		fmt.Print(`
    +---+  
    |   |  
        |  
        |  
        |  
  =========
`)
	case 5:
		fmt.Print(`
    +---+  
    |   |  
    O   |  
        |  
        |  
  =========
`)
	case 4:
		fmt.Print(`
    +---+  
    |   |  
    O   |  
   /|   |  
        |  
  =========
`)
	case 3:
		fmt.Print(`
    +---+  
    |   |  
    O   |  
   /|\  |  
        |  
  =========
`)
	case 2:
		fmt.Print(`
    +---+  
    |   |  
    O   |  
   /|\  |  
   /    |  
  =========
`)
	case 1:
		fmt.Print(`
    +---+  
    |   |  
    O   |  
   /|\  |  
   / \  |  
  =========
`)
	case 0:
		fmt.Print(`
    +---+  
    |   |  
    O   |  
   /|\  |  
   / \  |  
  =========
	  
`)
	default:
		fmt.Print(`
		 ██████ ▄▄▄█████▓ ▄▄▄       ██▀███  ▄▄▄█████▓
		▒██    ▒ ▓  ██▒ ▓▒▒████▄   ▓██ ▒ ██▒▓  ██▒ ▓▒
		░ ▓██▄   ▒ ▓██░ ▒░▒██  ▀█▄ ▓██ ░▄█ ▒▒ ▓██░ ▒░
		  ▒   ██▒░ ▓██▓ ░ ░██▄▄▄▄██▒██▀▀█▄  ░ ▓██▓ ░ 
		▒██████▒▒  ▒██▒ ░ ▒▓█   ▓██░██▓ ▒██▒  ▒██▒ ░ 
		▒ ▒▓▒ ▒ ░  ▒ ░░   ░▒▒   ▓▒█░ ▒▓ ░▒▓░  ▒ ░░   
		░ ░▒  ░      ░    ░ ░   ▒▒   ░▒ ░ ▒     ░    
		░  ░  ░    ░        ░   ▒    ░░   ░   ░      
			  ░                 ░     ░              
`)
	}
}
