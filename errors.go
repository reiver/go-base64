package base64

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	err2ndRuneEOF                = erorr.Error("base64: could not read 2nd rune (in sequence of runes of length 2, 3, or 4) because received an end-of-file (EOF)")
	err4thRuneEOFBut3rdRuneEqual = erorr.Error("base64: could not read 4th rune (in sequenc of runes of length 4) — there must be a 4th rune because the 3rd rune was '=' (U+003D)")
	errNilReader                 = erorr.Error("base64: nil reader")
	errNilWriter                 = erorr.Error("base64: nil writer")
)
