package hangmandep

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func NewGame() {
	// variable pour stocker les lettres rrentrées pendant le jeu
	var used_letter []string
	// Fichier contenant la liste des mots
	wordsFile := "words.txt"
	// Récupère la liste des mots
	wordsList, err := ReadWords(wordsFile)
	// Vérifie si la fonction a fonctionné correctement
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier des mots.")
		return
	}
	// Sélectionne un mot de la liste
	var wordChosen string
	if len(wordsList) > 0 {
		wordChosen = wordsList[rand.Intn(len(wordsList))]
	} else {
		fmt.Println("La liste des mots est vide.")
		return
	}
	// Initialise le mot caché
	wordHidden := PrintHidden(wordChosen)

	// Nombre maximal de tentatives
	maxAttempts := 10
	attemptsLeft := maxAttempts
	// Boucle pour permettre à l'utilisateur de deviner les lettres
	reader := bufio.NewReader(os.Stdin)
	for attemptsLeft > 0 {
		PrintHangman(attemptsLeft)
		fmt.Printf("Il vous reste %d tentatives.\n", attemptsLeft)
		fmt.Println(strings.Join(wordHidden, " "))
		fmt.Print("Devinez une lettre: ")
		//input est la lettre rentrée par l'utilisateur
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Erreur lors de la lecture de l'entrée.")
			continue
		}
		// Supprimez les espaces blancs en début ou en fin de chaîne
		input = strings.TrimSpace(input)

		// Vérifiez si l'entrée est vide ou contient des chiffres/des espaces
		if input == "" || ContainsNumbersOrSpaces(input) {
			fmt.Println("Veuillez entrer une lettre valide.")
			continue
		}

		// extraction de la lettre entrée en empechant la cassure (tout en minuscule meme si on rentre une maj)
		guessedLetter := []rune(strings.ToLower(input))[0]
		// Si le nombre de lettres est au dessus de 1 à l'entrée alors on doit réecrire 1 seule
		if len(input) > 1 {
			fmt.Println("Veuillez entrer qu'une seule lettre")
			continue
		}
		//si une lettre est déja utilisée alors on affiche un msg d erreure et on recommence
		if Used(used_letter, string(guessedLetter)) {
			fmt.Println("la lettre saisie est déja utilisée")
			continue
		}

		// mise a jour de la liste de lettres utilisée
		used_letter = append(used_letter, strings.ToLower(input))

		// Mettre à jour le mot caché avec la lettre devinée
		UpdateHiddenWord(wordChosen, wordHidden, guessedLetter)

		// Vérifie si la lettre devinée est correcte
		guessedCorrectly := false
		for _, char := range wordChosen {
			if char == guessedLetter {
				guessedCorrectly = true
				break
			}
		}

		fmt.Println("========================================")

		// Vérifiez si le mot est entièrement révélé
		if strings.Join(wordHidden, "") == wordChosen {
			fmt.Print(`
			▓██   ██▓ ▒█████   █    ██      █     █░  ██▓ ███▄    █ 
			 ▒██  ██▒▒██▒  ██▒ ██  ▓██▒    ▓█░ █ ░█░▒▓██▒ ██ ▀█   █ 
			  ▒██ ██░▒██░  ██▒▓██  ▒██░    ▒█░ █ ░█ ▒▒██▒▓██  ▀█ ██▒
			  ░ ▐██▓░▒██   ██░▓▓█  ░██░    ░█░ █ ░█ ░░██░▓██▒  ▐▌██▒
			  ░ ██▒▓░░ ████▓▒░▒▒█████▓     ░░██▒██▓ ░░██░▒██░   ▓██░
			   ██▒▒▒ ░ ▒░▒░▒░ ░▒▓▒ ▒ ▒     ░ ▓░▒ ▒   ░▓  ░ ▒░   ▒ ▒ 
			 ▓██ ░▒░   ░ ▒ ▒░ ░░▒░ ░ ░       ▒ ░ ░  ░ ▒ ░░ ░░   ░ ▒░
			 ▒ ▒ ░░  ░ ░ ░ ▒   ░░░ ░ ░       ░   ░  ░ ▒ ░   ░   ░ ░ 
			 ░ ░         ░ ░     ░             ░      ░           ░ 
			
		 `)
			fmt.Println("Félicitations ! Vous avez deviné le mot ", wordChosen, "!")
			break
		}
		// Décrémentez les tentatives restantes uniquement si la lettre devinée est incorrecte
		if !guessedCorrectly {
			attemptsLeft--
		}
	}
	if attemptsLeft == 0 {
		fmt.Println("Désolé, vous avez épuisé toutes vos tentatives. Le mot était :", wordChosen)
		fmt.Print(`
		  ▄████  ▄▄▄      ███▄ ▄███▓ ▓█████     ▒█████   ██▒   █▓ ▓█████ ██▀███  
		▒ ██▒ ▀█▒▒████▄   ▓██▒▀█▀ ██▒ ▓█   ▀    ▒██▒  ██▒▓██░   █▒ ▓█   ▀▓██ ▒ ██▒
		░▒██░▄▄▄░▒██  ▀█▄ ▓██    ▓██░ ▒███      ▒██░  ██▒ ▓██  █▒░ ▒███  ▓██ ░▄█ ▒
		░░▓█  ██▓░██▄▄▄▄██▒██    ▒██  ▒▓█  ▄    ▒██   ██░  ▒██ █░░ ▒▓█  ▄▒██▀▀█▄  
		░▒▓███▀▒░▒▓█   ▓██▒██▒   ░██▒▒░▒████    ░ ████▓▒░   ▒▀█░  ▒░▒████░██▓ ▒██▒
		 ░▒   ▒  ░▒▒   ▓▒█░ ▒░   ░  ░░░░ ▒░     ░ ▒░▒░▒░    ░ ▐░  ░░░ ▒░ ░ ▒▓ ░▒▓░
		  ░   ░  ░ ░   ▒▒ ░  ░      ░░ ░ ░        ░ ▒ ▒░    ░ ░░  ░ ░ ░    ░▒ ░ ▒ 
		░ ░   ░ ░  ░   ▒  ░      ░       ░      ░ ░ ░ ▒        ░      ░    ░░   ░ 
			  ░        ░         ░   ░   ░          ░ ░        ░  ░   ░     ░     
`)
	}
}
