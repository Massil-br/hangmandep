package hangman

import (
	"bufio"
	"fmt"
	"os"
)

// readWords lit un fichier de mots et retourne une liste de mots.
func ReadWords(wordsFile string) ([]string, error) {
	var wordsList []string
	// Ouvre le fichier
	words, err := os.Open(wordsFile)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de l'ouverture du fichier: %v", err)
	}
	defer words.Close()
	// Crée un scanner pour lire le fichier ligne par ligne
	scanner := bufio.NewScanner(words)
	// Lit le contenu du fichier ligne par ligne
	for scanner.Scan() {
		wordsList = append(wordsList, scanner.Text())
	}
	// Vérifie s'il y a des erreurs pendant la lecture du fichier
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("erreur lors de la lecture du fichier: %v", err)
	}
	return wordsList, nil
}
