package base64

import (
	"errors"
	"io"

	"sourcecode.social/reiver/go-erorr"
)

// Decode reads base64-encoded data from 'reader', decodes it, and then writes the decoded data to 'writer'.
//
// Decode handles multiple different base64 styles.
//
// I.e., Decode can handle base64-encoded data that does and doesn't use the '=' padding-character, and
// Decode can handle both '+' and '-' for the symbol 63, and can handle '/', ',', '_' for symbol 64.
func Decode(writer io.Writer, reader io.Reader) error {
	if nil == writer {
		return errNilWriter
	}
	if nil == reader {
		return errNilReader
	}

	var eof bool
	var done bool
	for !eof || done {
		var r1 rune
		{
			var err error

			r1, err = readRune(reader)
			if errors.Is(err, io.EOF) {
				return nil
			}
			if nil != err {
				return err
			}
		}

		var r2 rune
		{
			var err error

			r2, err = readRune(reader)
			if errors.Is(err, io.EOF) {
				return err2ndRuneEOF
			}
			if nil != err {
				return err
			}
		}

		var r3 rune = '='
		{
			r, err := readRune(reader)
			eof = errors.Is(err, io.EOF)
			if nil != err && !eof {
				return err
			}
			if !eof {
				r3 = r
			}
		}

		var r4 rune = '='
		if !eof {
			r, err := readRune(reader)
			eof = errors.Is(err, io.EOF)
			if nil != err && !eof {
				return err
			}
			if '=' == r3 && eof {
				return err4thRuneEOFBut3rdRuneEqual
			}
			done = '=' == r
			if !eof {
				r4 = r
			}
		}

		{
			b1,b2,b3, length, err := decodeRunesToBytes(r1,r2,r3,r4)
			if nil != err {
				return err
			}

			var buffer = [3]byte{b1,b2,b3}

			n, err := writer.Write(buffer[:length])
			if nil != err {
				return erorr.Errorf("base64: problem writing %d bytes to writer: %w", length, err)
			}
			if n != length {
				return erorr.Errorf("base64: problem writing %d bytes to writer — writer reported that only %d bytes were written", length, n)
			}
		}
	}

	return nil
}
