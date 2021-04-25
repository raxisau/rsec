package rsec

import "math/rand"

const (
	LOWER_ALPHA = 1
	UPPER_ALPHA = 2
	NUMBERS     = 4
	SYMBOLS     = 8
	SIMPLE      = 3  // 1 | 2;
	MEDIUM      = 7  // 1 | 2 | 4
	COMPLEX     = 15 // 1 | 2 | 4 | 8;
)

// These letter groups are chosen to remove the ambigious letters
// In this case G l I 1 0 O o are all a little ambigious
// Also the letter w is removed. In some fonts it is hard to read
var (
	LowercaseCharacters = []rune("abcdefghijkmnpqrstuvxyz") // l and o removed
	UppercaseCharacters = []rune("ABCDEFHJKLMNPRSTUVXYZ")   // G, I, O, Q removed
	NumberCharacters    = []rune("23456789")                //1 and 0 removed
	SymbolCharacters    = []rune("!@#$%^&*+?{}[]<>")

	// This is the ratio of the different groups. Good readability but maintaining security
	letterGroupIndexes = []int{
		LOWER_ALPHA,
		LOWER_ALPHA,
		UPPER_ALPHA,
		UPPER_ALPHA,
		NUMBERS,
		SYMBOLS,
	}

	letterGroups = map[int][]rune{
		LOWER_ALPHA: LowercaseCharacters,
		UPPER_ALPHA: UppercaseCharacters,
		NUMBERS:     NumberCharacters,
		SYMBOLS:     SymbolCharacters,
	}
)

// PassGen This is the array of group indexes. The reason that theere are 2 alpha
// for upper and lower is that this gives a higher density of alpha
// For example passGen ( 6, self::COMPLEX ) will give 2 upper case,
// 2 lowercase one number and one symbol. This ratio seems to be the most readable
// while keeping the password strength
func PassGen(desiredLen, complexity int) string {

	password := make([]rune, desiredLen)
	passwordIdx := 0

	for i := 0; passwordIdx < desiredLen; i++ {

		// Skip this group is the password is too complex
		letterGroupIndex := letterGroupIndexes[i%len(letterGroupIndexes)]
		if (letterGroupIndex & complexity) == 0 {
			continue
		}

		letters := letterGroups[letterGroupIndex]
		password[passwordIdx] = letters[rand.Intn(len(letters))]
		passwordIdx++
	}

	rand.Shuffle(desiredLen, func(i, j int) { password[i], password[j] = password[j], password[i] })
	return string(password)
}
