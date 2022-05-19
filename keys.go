package tuitest

// Key sequences taken from https://github.com/charmbracelet/bubbletea/blob/master/key.go
const (
	keyNUL byte = 0   // null, \0
	keySOH byte = 1   // start of heading
	keySTX byte = 2   // start of text
	keyETX byte = 3   // break, ctrl+c
	keyEOT byte = 4   // end of transmission
	keyENQ byte = 5   // enquiry
	keyACK byte = 6   // acknowledge
	keyBEL byte = 7   // bell, \a
	keyBS  byte = 8   // backspace
	keyHT  byte = 9   // horizontal tabulation, \t
	keyLF  byte = 10  // line feed, \n
	keyVT  byte = 11  // vertical tabulation \v
	keyFF  byte = 12  // form feed \f
	keyCR  byte = 13  // carriage return, \r
	keySO  byte = 14  // shift out
	keySI  byte = 15  // shift in
	keyDLE byte = 16  // data link escape
	keyDC1 byte = 17  // device control one
	keyDC2 byte = 18  // device control two
	keyDC3 byte = 19  // device control three
	keyDC4 byte = 20  // device control four
	keyNAK byte = 21  // negative acknowledge
	keySYN byte = 22  // synchronous idle
	keyETB byte = 23  // end of transmission block
	keyCAN byte = 24  // cancel
	keyEM  byte = 25  // end of medium
	keySUB byte = 26  // substitution
	keyESC byte = 27  // escape, \e
	keyFS  byte = 28  // file separator
	keyGS  byte = 29  // group separator
	keyRS  byte = 30  // record separator
	keyUS  byte = 31  // unit separator
	keyDEL byte = 127 // delete. on most systems this is mapped to backspace, I hear
)

// Control key aliases.
const (
	KeyNull      byte = keyNUL
	KeyBreak     byte = keyETX
	KeyEnter     byte = keyCR
	KeyBackspace byte = keyDEL
	KeyTab       byte = keyHT
	KeyEsc       byte = keyESC
	KeyEscape    byte = keyESC

	KeyCtrlAt           byte = keyNUL // ctrl+@
	KeyCtrlA            byte = keySOH
	KeyCtrlB            byte = keySTX
	KeyCtrlC            byte = keyETX
	KeyCtrlD            byte = keyEOT
	KeyCtrlE            byte = keyENQ
	KeyCtrlF            byte = keyACK
	KeyCtrlG            byte = keyBEL
	KeyCtrlH            byte = keyBS
	KeyCtrlI            byte = keyHT
	KeyCtrlJ            byte = keyLF
	KeyCtrlK            byte = keyVT
	KeyCtrlL            byte = keyFF
	KeyCtrlM            byte = keyCR
	KeyCtrlN            byte = keySO
	KeyCtrlO            byte = keySI
	KeyCtrlP            byte = keyDLE
	KeyCtrlQ            byte = keyDC1
	KeyCtrlR            byte = keyDC2
	KeyCtrlS            byte = keyDC3
	KeyCtrlT            byte = keyDC4
	KeyCtrlU            byte = keyNAK
	KeyCtrlV            byte = keySYN
	KeyCtrlW            byte = keyETB
	KeyCtrlX            byte = keyCAN
	KeyCtrlY            byte = keyEM
	KeyCtrlZ            byte = keySUB
	KeyCtrlOpenBracket  byte = keyESC // ctrl+[
	KeyCtrlBackslash    byte = keyFS  // ctrl+\
	KeyCtrlCloseBracket byte = keyGS  // ctrl+]
	KeyCtrlCaret        byte = keyRS  // ctrl+^
	KeyCtrlUnderscore   byte = keyUS  // ctrl+_
	KeyCtrlQuestionMark byte = keyDEL // ctrl+?
)

// Other keys.
const (
	KeyUp    string = "\x1b[A"
	KeyDown  string = "\x1b[B"
	KeyRight string = "\x1b[C"
	KeyLeft  string = "\x1b[D"
)
