package base64

import (
	"errors"
	"io"

	"github.com/reiver/go-ascii"
	"sourcecode.social/reiver/go-erorr"
	"sourcecode.social/reiver/go-utf8"
)

// readRune returns the next rune that isn't a LF, CR, or SP.
//
// The decoding of something base64-encoded ignores LF, CR. and SP.
// So, readRune keeps reading runes until it receives a rune that isn't
// a LF, CR, or SP. I.e., it ignores LF, CR, or SP.
func readRune(reader io.Reader) (rune, error) {
	if nil == reader {
		return 0, errNilReader
	}

	var r rune
	loop: for {
		var size int
		var err error

		r, size, err = utf8.ReadRune(reader)
		if errors.Is(err, io.EOF) {
			return 0, io.EOF
		}
		if nil != err {
			return 0, erorr.Errorf("base64: problem reading rune: %w", err)
		}
		if size <= 0 {
			return 0, erorr.Errorf("base64: expected size of read rune to be greater-than 0 but actually was %d", size)
		}

		switch r {
		case ascii.LF:
			continue loop
		case ascii.CR:
			continue loop
		case ascii.SP:
			continue loop
		default:
			break loop
		}
	}

	return r, nil
}
