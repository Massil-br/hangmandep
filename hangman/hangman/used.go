package hangman

// fonction permettant de check si un string est compris dans un tableau
func Used(s []string, val string) bool {
	for _, item := range s {
		if item == val {
			return true
		}
	}
	return false
}
