package tuitest

// Key sequences taken from https://github.com/charmbracelet/bubbletea/blob/master/key.go

const (
	keyNUL = string(rune(0))   // null, \0
	keySOH = string(rune(1))   // start of heading
	keySTX = string(rune(2))   // start of text
	keyETX = string(rune(3))   // break, ctrl+c
	keyEOT = string(rune(4))   // end of transmission
	keyENQ = string(rune(5))   // enquiry
	keyACK = string(rune(6))   // acknowledge
	keyBEL = string(rune(7))   // bell, \a
	keyBS  = string(rune(8))   // backspace
	keyHT  = string(rune(9))   // horizontal tabulation, \t
	keyLF  = string(rune(10))  // line feed, \n
	keyVT  = string(rune(11))  // vertical tabulation \v
	keyFF  = string(rune(12))  // form feed \f
	keyCR  = string(rune(13))  // carriage return, \r
	keySO  = string(rune(14))  // shift out
	keySI  = string(rune(15))  // shift in
	keyDLE = string(rune(16))  // data link escape
	keyDC1 = string(rune(17))  // device control one
	keyDC2 = string(rune(18))  // device control two
	keyDC3 = string(rune(19))  // device control three
	keyDC4 = string(rune(20))  // device control four
	keyNAK = string(rune(21))  // negative acknowledge
	keySYN = string(rune(22))  // synchronous idle
	keyETB = string(rune(23))  // end of transmission block
	keyCAN = string(rune(24))  // cancel
	keyEM  = string(rune(25))  // end of medium
	keySUB = string(rune(26))  // substitution
	keyESC = string(rune(27))  // escape, \e
	keyFS  = string(rune(28))  // file separator
	keyGS  = string(rune(29))  // group separator
	keyRS  = string(rune(30))  // record separator
	keyUS  = string(rune(31))  // unit separator
	keyDEL = string(rune(127)) // delete. on most systems this is mapped to backspace, I hear
)

// Control key aliases.
const (
	KeyNull      = keyNUL
	KeyBreak     = keyETX
	KeyEnter     = keyCR
	KeyBackspace = keyDEL
	KeyTab       = keyHT
	KeyEsc       = keyESC
	KeyEscape    = keyESC

	KeyCtrlAt           = keyNUL // ctrl+@
	KeyCtrlA            = keySOH
	KeyCtrlB            = keySTX
	KeyCtrlC            = keyETX
	KeyCtrlD            = keyEOT
	KeyCtrlE            = keyENQ
	KeyCtrlF            = keyACK
	KeyCtrlG            = keyBEL
	KeyCtrlH            = keyBS
	KeyCtrlI            = keyHT
	KeyCtrlJ            = keyLF
	KeyCtrlK            = keyVT
	KeyCtrlL            = keyFF
	KeyCtrlM            = keyCR
	KeyCtrlN            = keySO
	KeyCtrlO            = keySI
	KeyCtrlP            = keyDLE
	KeyCtrlQ            = keyDC1
	KeyCtrlR            = keyDC2
	KeyCtrlS            = keyDC3
	KeyCtrlT            = keyDC4
	KeyCtrlU            = keyNAK
	KeyCtrlV            = keySYN
	KeyCtrlW            = keyETB
	KeyCtrlX            = keyCAN
	KeyCtrlY            = keyEM
	KeyCtrlZ            = keySUB
	KeyCtrlOpenBracket  = keyESC // ctrl+[
	KeyCtrlBackslash    = keyFS  // ctrl+\
	KeyCtrlCloseBracket = keyGS  // ctrl+]
	KeyCtrlCaret        = keyRS  // ctrl+^
	KeyCtrlUnderscore   = keyUS  // ctrl+_
	KeyCtrlQuestionMark = keyDEL // ctrl+?
)

const (
	// Arrow keys
	KeyUp                = "\x1b[A"
	KeyDown              = "\x1b[B"
	KeyRight             = "\x1b[C"
	KeyLeft              = "\x1b[D"
	KeyShiftUp           = "\x1b[1;2A"
	KeyShiftDown         = "\x1b[1;2B"
	KeyShiftRight        = "\x1b[1;2C"
	KeyShiftLeft         = "\x1b[1;2D"
	KeyAltUp             = "\x1b[1;3A"
	KeyAltDown           = "\x1b[1;3B"
	KeyAltRight          = "\x1b[1;3C"
	KeyAltLeft           = "\x1b[1;3D"
	KeyAltShiftUp        = "\x1b[1;4A"
	KeyAltShiftDown      = "\x1b[1;4B"
	KeyAltShiftRight     = "\x1b[1;4C"
	KeyAltShiftLeft      = "\x1b[1;4D"
	KeyCtrlUp            = "\x1b[1;5A"
	KeyCtrlDown          = "\x1b[1;5B"
	KeyCtrlRight         = "\x1b[1;5C"
	KeyCtrlLeft          = "\x1b[1;5D"
	KeyCtrlShiftUp       = "\x1b[1;6A"
	KeyCtrlShiftDown     = "\x1b[1;6B"
	KeyCtrlShiftRight    = "\x1b[1;6C"
	KeyCtrlShiftLeft     = "\x1b[1;6D"
	KeyCtrlAltUp         = "\x1b[1;7A"
	KeyCtrlAltDown       = "\x1b[1;7B"
	KeyCtrlAltRight      = "\x1b[1;7C"
	KeyCtrlAltLeft       = "\x1b[1;7D"
	KeyCtrlAltShiftUp    = "\x1b[1;8A"
	KeyCtrlAltShiftDown  = "\x1b[1;8B"
	KeyCtrlAltShiftRight = "\x1b[1;8C"
	KeyCtrlAltShiftLeft  = "\x1b[1;8D"

	// Miscellaneous keys
	KeyShiftTab = "\x1b[Z"

	KeyInsert    = "\x1b[2~"
	KeyAltInsert = "\x1b[3;2~"

	KeyDelete    = "\x1b[3~"
	KeyAltDelete = "\x1b[3;3~"

	KeyPgUp    = "\x1b[5~"
	KeyAltPgUp = "\x1b[5;3~"

	KeyCtrlPgUp    = "\x1b[5;5~"
	KeyCtrlAltPgUp = "\x1b[5;7~"

	KeyPgDown        = "\x1b[6~"
	KeyAltPgDown     = "\x1b[6;3~"
	KeyCtrlPgDown    = "\x1b[6;5~"
	KeyCtrlAltPgDown = "\x1b[6;7~"

	KeyHome             = "\x1b[1~"
	KeyAltHome          = "\x1b[1;3H"
	KeyCtrlHome         = "\x1b[1;5H"
	KeyCtrlAltHome      = "\x1b[1;7H"
	KeyShiftHome        = "\x1b[1;2H"
	KeyShiftAltHome     = "\x1b[1;4H"
	KeyCtrlShiftHome    = "\x1b[1;6H"
	KeyCtrlAltShiftHome = "\x1b[1;8H"

	KeyEnd             = "\x1b[4~"
	KeyAltEnd          = "\x1b[1;3F"
	KeyCtrlEnd         = "\x1b[1;5F"
	KeyCtrlAltEnd      = "\x1b[1;7F"
	KeyShiftEnd        = "\x1b[1;2F"
	KeyAltShiftEnd     = "\x1b[1;4F"
	KeyCtrlShiftEnd    = "\x1b[1;6F"
	KeyCtrlAltShiftEnd = "\x1b[1;8F"

	KeyF1  = "\x1bOP"
	KeyF2  = "\x1bOQ"
	KeyF3  = "\x1bOR"
	KeyF4  = "\x1bOS"
	KeyF5  = "\x1b[15~"
	KeyF6  = "\x1b[17~"
	KeyF7  = "\x1b[18~"
	KeyF8  = "\x1b[19~"
	KeyF9  = "\x1b[20~"
	KeyF10 = "\x1b[21~"
	KeyF11 = "\x1b[23~"
	KeyF12 = "\x1b[24~"
	KeyF13 = "\x1b[25~"
	KeyF14 = "\x1b[26~"
	KeyF15 = "\x1b[28~"
	KeyF16 = "\x1b[29~"
	KeyF17 = "\x1b[31~"
	KeyF18 = "\x1b[32~"
	KeyF19 = "\x1b[33~"
	KeyF20 = "\x1b[34~"

	KeyAltF1  = "\x1b[1;3P"
	KeyAltF2  = "\x1b[1;3Q"
	KeyAltF3  = "\x1b[1;3R"
	KeyAltF4  = "\x1b[1;3S"
	KeyAltF5  = "\x1b[15;3~"
	KeyAltF6  = "\x1b[17;3~"
	KeyAltF7  = "\x1b[18;3~"
	KeyAltF8  = "\x1b[19;3~"
	KeyAltF9  = "\x1b[20;3~"
	KeyAltF10 = "\x1b[21;3~"
	KeyAltF11 = "\x1b[23;3~"
	KeyAltF12 = "\x1b[24;3~"
	KeyAltF13 = "\x1b[25;3~"
	KeyAltF14 = "\x1b[26;3~"
	KeyAltF15 = "\x1b[28;3~"
	KeyAltF16 = "\x1b[29;3~"
	KeyAltF17 = "\x1b\x1b[31~"
	KeyAltF18 = "\x1b\x1b[32~"
	KeyAltF19 = "\x1b\x1b[33~"
	KeyAltF20 = "\x1b\x1b[34~"
)
