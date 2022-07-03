package tuitest

// Key sequences taken from https://github.com/charmbracelet/bubbletea/blob/master/key.go

const (
	keyNUL string = string(rune(0))   // null, \0
	keySOH string = string(rune(1))   // start of heading
	keySTX string = string(rune(2))   // start of text
	keyETX string = string(rune(3))   // break, ctrl+c
	keyEOT string = string(rune(4))   // end of transmission
	keyENQ string = string(rune(5))   // enquiry
	keyACK string = string(rune(6))   // acknowledge
	keyBEL string = string(rune(7))   // bell, \a
	keyBS  string = string(rune(8))   // backspace
	keyHT  string = string(rune(9))   // horizontal tabulation, \t
	keyLF  string = string(rune(10))  // line feed, \n
	keyVT  string = string(rune(11))  // vertical tabulation \v
	keyFF  string = string(rune(12))  // form feed \f
	keyCR  string = string(rune(13))  // carriage return, \r
	keySO  string = string(rune(14))  // shift out
	keySI  string = string(rune(15))  // shift in
	keyDLE string = string(rune(16))  // data link escape
	keyDC1 string = string(rune(17))  // device control one
	keyDC2 string = string(rune(18))  // device control two
	keyDC3 string = string(rune(19))  // device control three
	keyDC4 string = string(rune(20))  // device control four
	keyNAK string = string(rune(21))  // negative acknowledge
	keySYN string = string(rune(22))  // synchronous idle
	keyETB string = string(rune(23))  // end of transmission block
	keyCAN string = string(rune(24))  // cancel
	keyEM  string = string(rune(25))  // end of medium
	keySUB string = string(rune(26))  // substitution
	keyESC string = string(rune(27))  // escape, \e
	keyFS  string = string(rune(28))  // file separator
	keyGS  string = string(rune(29))  // group separator
	keyRS  string = string(rune(30))  // record separator
	keyUS  string = string(rune(31))  // unit separator
	keyDEL string = string(rune(127)) // delete. on most systems this is mapped to backspace, I hear
)

// Control key aliases.
const (
	KeyNull      string = keyNUL
	KeyBreak     string = keyETX
	KeyEnter     string = keyCR
	KeyBackspace string = keyDEL
	KeyTab       string = keyHT
	KeyEsc       string = keyESC
	KeyEscape    string = keyESC

	KeyCtrlAt           string = keyNUL // ctrl+@
	KeyCtrlA            string = keySOH
	KeyCtrlB            string = keySTX
	KeyCtrlC            string = keyETX
	KeyCtrlD            string = keyEOT
	KeyCtrlE            string = keyENQ
	KeyCtrlF            string = keyACK
	KeyCtrlG            string = keyBEL
	KeyCtrlH            string = keyBS
	KeyCtrlI            string = keyHT
	KeyCtrlJ            string = keyLF
	KeyCtrlK            string = keyVT
	KeyCtrlL            string = keyFF
	KeyCtrlM            string = keyCR
	KeyCtrlN            string = keySO
	KeyCtrlO            string = keySI
	KeyCtrlP            string = keyDLE
	KeyCtrlQ            string = keyDC1
	KeyCtrlR            string = keyDC2
	KeyCtrlS            string = keyDC3
	KeyCtrlT            string = keyDC4
	KeyCtrlU            string = keyNAK
	KeyCtrlV            string = keySYN
	KeyCtrlW            string = keyETB
	KeyCtrlX            string = keyCAN
	KeyCtrlY            string = keyEM
	KeyCtrlZ            string = keySUB
	KeyCtrlOpenBracket  string = keyESC // ctrl+[
	KeyCtrlBackslash    string = keyFS  // ctrl+\
	KeyCtrlCloseBracket string = keyGS  // ctrl+]
	KeyCtrlCaret        string = keyRS  // ctrl+^
	KeyCtrlUnderscore   string = keyUS  // ctrl+_
	KeyCtrlQuestionMark string = keyDEL // ctrl+?
)

// Other keys.
const (
	KeyUp    string = "\x1b[A"
	KeyDown  string = "\x1b[B"
	KeyRight string = "\x1b[C"
	KeyLeft  string = "\x1b[D"
)
