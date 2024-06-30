package base64

import (
	"sourcecode.social/reiver/go-erorr"
)

// decodeRunesToBytes decodes the next 3 output bytes by reading the next 4 runes from the io.Reader.
//
// 00_123456 <-- 1st read (decoded byte №1)
// 00_781234 <-- 2nd read (decoded byte №2)
// 00_567812 <-- 3rd read (decoded byte №3)
// 00_345678 <-- 4th read (decoded byte №4)
//
// I.e.,...
//
// [123456    <-- 1st read (decoded byte №1)
//  78] [1234 <-- 2nd read (decoded byte №2)
//  5678] [12 <-- 3rd read (decoded byte №3)
//  345678]   <-- 4th read (decoded byte №4)
func decodeRunesToBytes(r1, r2, r3, r4 rune) (b1 byte, b2 byte, b3 byte, length int, err error) {

	if '=' == r1 {
		return 0,0,0, 0, erorr.Errorf("base64: could not translate the 1st-rune (i.e., %q (%U)) (in the 4 rune sequence) because it is cannot be a padding-character", r1, r1)
	}
	if '=' == r2 {
		return 0,0,0, 0, erorr.Errorf("base64: could not translate the 2nd-rune (i.e., %q (%U)) (in the 4 rune sequence) because it is cannot be a padding-character", r2, r2)
	}
	if '=' == r3 && '=' != r4 {
		return 0,0,0, 0, erorr.Errorf("base64: could not translate because the 3rd-rune (i.e., %q (%U)) (in the 4 rune sequence) is a padding-character but the 4th-rune (i.e., %q (%U)) is not ad padding-character — if the 3rd-rune is a padding character then the 4th-rune must also be a padding character", r3, r3, r4, r4)
	}

	var x1,x2,x3,x4 byte
	var ok bool
	{
		x1, ok = translateRune(r1)
		if !ok {
			return 0,0,0, 0, erorr.Errorf("base64: could not translate the 1st-rune (i.e., %q (%U)) (in the 4 rune sequence) because it is not a valid symbol used in base64-encodings", r1, r1)
		}

		x2, ok = translateRune(r2)
		if !ok {
			return 0,0,0, 0, erorr.Errorf("base64: could not translate the 2nd-rune (i.e., %q (%U)) (in the 4 rune sequence) because it is not a valid symbol used in base64-encodings", r2, r2)
		}
		if '=' == r3 {
			//                          vvvv
			//                     00_781234
			var rest byte = x2 & 0b00_001111

			if 0 != rest {
				return 0,0,0, 0, erorr.Errorf("base64: the 2nd-rune (i.e., %q (%U)) (in the 4 rune sequence) cannot translate to something where the least-significant 4 bits are not 0b0000 because the 3rd-rune (i.e., %q (%U)) is a padding-character", r2, r2, r3, r3)
			}
		}
		length++

		if '=' != r3 {
			x3, ok = translateRune(r3)
			if !ok {
				return 0,0,0, 0, erorr.Errorf("base64: could not translate the 3rd-rune (i.e., %q (%U)) (in the 4 rune sequence) because it is not a valid symbol used in base64-encodings", r3, r3)
			}
			if '=' == r4 {
				//                            vv
				//                     00_567812
				var rest byte = x3 & 0b00_000011

				if 0 != rest {
					return 0,0,0, 0, erorr.Errorf("base64: the 3rd-rune (i.e., %q (%U)) (in the 4 rune sequence) cannot translate to something where the least-significant 2 bits are not 0b00 because the 4th-rune (i.e., %q (%U)) is a padding-character", r3, r3, r4, r4)
				}
			}
			length++
		}

		if '=' != r4 {
			x4, ok = translateRune(r4)
			if !ok {
				return 0,0,0, 0, erorr.Errorf("base64: could not translate the 4th-rune (i.e., %q (%U)) (in the 4 rune sequence) because it is not a valid symbol used in base64-encodings", r4, r4)
			}
			length++
		}
	}

	{
		b1 = (x1 << 2) | (x2 >> 4)
		b2 = (x2 << 4) | (x3 >> 2)
		b3 = (x3 << 6) | x4
	}

	return b1,b2,b3,length,nil
}
