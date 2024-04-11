package hangmandep

import "unicode"

// containsNumbersOrSpaces vérifie si la chaîne contient des chiffres ou des espaces.
func ContainsNumbersOrSpaces(s string) bool {
	for _, r := range s {
		if unicode.IsNumber(r) || unicode.IsSpace(r) {
			return true
		}
	}
	return false
}
