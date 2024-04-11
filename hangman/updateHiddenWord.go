package hangmandep

// updateHiddenWord met à jour le mot caché avec la lettre devinée par l'utilisateur.
func UpdateHiddenWord(wordChosen string, wordHidden []string, guessedLetter rune) {
	for i, char := range wordChosen {
		if char == guessedLetter {
			wordHidden[i] = string(guessedLetter)
		}
	}
}
