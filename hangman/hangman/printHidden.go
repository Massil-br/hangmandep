package hangmandep

// printHidden affiche les caractères cachés du mot choisi.
func PrintHidden(wordChosen string) []string {
	var wordHidden []string
	// Ajoute un "_" pour chaque caractère du mot choisi
	for range wordChosen {
		wordHidden = append(wordHidden, "_")
	}
	return wordHidden
}
