package base64

import (
	"testing"
)

func TestTranslateRune(t *testing.T) {

	tests := []struct{
		Rune rune
		Expected byte
	}{
		{
			Rune: 'A',
			Expected: 0,
		},
		{
			Rune: 'B',
			Expected: 1,
		},
		{
			Rune: 'C',
			Expected: 2,
		},
		{
			Rune: 'D',
			Expected: 3,
		},
		{
			Rune: 'E',
			Expected: 4,
		},
		{
			Rune: 'F',
			Expected: 5,
		},
		{
			Rune: 'G',
			Expected: 6,
		},
		{
			Rune: 'H',
			Expected: 7,
		},
		{
			Rune: 'I',
			Expected: 8,
		},
		{
			Rune: 'J',
			Expected: 9,
		},
		{
			Rune: 'K',
			Expected:10 ,
		},
		{
			Rune: 'L',
			Expected: 11,
		},
		{
			Rune: 'M',
			Expected: 12,
		},
		{
			Rune: 'N',
			Expected: 13,
		},
		{
			Rune: 'O',
			Expected: 14,
		},
		{
			Rune: 'P',
			Expected: 15,
		},
		{
			Rune: 'Q',
			Expected: 16,
		},
		{
			Rune: 'R',
			Expected: 17,
		},
		{
			Rune: 'S',
			Expected: 18,
		},
		{
			Rune: 'T',
			Expected: 19,
		},
		{
			Rune: 'U',
			Expected: 20,
		},
		{
			Rune: 'V',
			Expected: 21,
		},
		{
			Rune: 'W',
			Expected: 22,
		},
		{
			Rune: 'X',
			Expected: 23,
		},
		{
			Rune: 'Y',
			Expected: 24,
		},
		{
			Rune: 'Z',
			Expected: 25,
		},



		{
			Rune: 'a',
			Expected: 26,
		},
		{
			Rune: 'b',
			Expected: 27,
		},
		{
			Rune: 'c',
			Expected: 28,
		},
		{
			Rune: 'd',
			Expected: 29,
		},
		{
			Rune: 'e',
			Expected: 30,
		},
		{
			Rune: 'f',
			Expected: 31,
		},
		{
			Rune: 'g',
			Expected: 32,
		},
		{
			Rune: 'h',
			Expected: 33,
		},
		{
			Rune: 'i',
			Expected: 34,
		},
		{
			Rune: 'j',
			Expected: 35,
		},
		{
			Rune: 'k',
			Expected: 36,
		},
		{
			Rune: 'l',
			Expected: 37,
		},
		{
			Rune: 'm',
			Expected: 38,
		},
		{
			Rune: 'n',
			Expected: 39,
		},
		{
			Rune: 'o',
			Expected: 40,
		},
		{
			Rune: 'p',
			Expected: 41,
		},
		{
			Rune: 'q',
			Expected: 42,
		},
		{
			Rune: 'r',
			Expected: 43,
		},
		{
			Rune: 's',
			Expected: 44,
		},
		{
			Rune: 't',
			Expected: 45,
		},
		{
			Rune: 'u',
			Expected: 46,
		},
		{
			Rune: 'v',
			Expected: 47,
		},
		{
			Rune: 'w',
			Expected: 48,
		},
		{
			Rune: 'x',
			Expected: 49,
		},
		{
			Rune: 'y',
			Expected: 50,
		},
		{
			Rune: 'z',
			Expected: 51,
		},



		{
			Rune: '0',
			Expected: 52,
		},
		{
			Rune: '1',
			Expected: 53,
		},
		{
			Rune: '2',
			Expected: 54,
		},
		{
			Rune: '3',
			Expected: 55,
		},
		{
			Rune: '4',
			Expected: 56,
		},
		{
			Rune: '5',
			Expected: 57,
		},
		{
			Rune: '6',
			Expected: 58,
		},
		{
			Rune: '7',
			Expected: 59,
		},
		{
			Rune: '8',
			Expected: 60,
		},
		{
			Rune: '9',
			Expected: 61,
		},



		{
			Rune: '+',
			Expected: 62,
		},
		{
			Rune: '-',
			Expected: 62,
		},



		{
			Rune: '/',
			Expected: 63,
		},
		{
			Rune: ',',
			Expected: 63,
		},
		{
			Rune: '_',
			Expected: 63,
		},
	}

	for testNumber, test := range tests {

		actual, ok := translateRune(test.Rune)
		if !ok {
			t.Errorf("For test #%d, expected to be able to translate the rune but actually wasn't.", testNumber)
			t.Logf("RUNE: %q (%U)", test.Rune, test.Rune)
			continue
		}

		{
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual translated-rune is not what was expected.", testNumber)
				t.Errorf("EXPECTED: 0x%02X", expected)
				t.Errorf("ACTUAL:   0x%02X", actual)
				t.Logf("RUNE: %q (%U)", test.Rune, test.Rune)
				continue
			}
		}
	}
}

func TestTranslateRune_128(t *testing.T) {

	tests := []struct{
		Rune rune
		ExpectedByte byte
		ExpectedOK bool
	}{
		{
			Rune: 0x00,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x01,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x02,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x03,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x04,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x05,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x06,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x07,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x08,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x09,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x0A,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x0B,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x0C,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x0D,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x0E,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x0F,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x10,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x11,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x12,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x13,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x14,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x15,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x16,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x17,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x18,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x19,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x1A,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x1B,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x1C,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x1D,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x1E,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x1F,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x20,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x21,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x22,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x23,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x24,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x25,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x26,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x27,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x28,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x29,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x2A,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x2B, // '+'
			ExpectedByte: 62,
			ExpectedOK: true,
		},
		{
			Rune: 0x2C, // ','
			ExpectedByte: 63,
			ExpectedOK: true,
		},
		{
			Rune: 0x2D, // '-'
			ExpectedByte: 62,
			ExpectedOK: true,
		},
		{
			Rune: 0x2E,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x2F, // '/'
			ExpectedByte: 63,
			ExpectedOK: true,
		},
		{
			Rune: 0x30, // '0'
			ExpectedByte: 52,
			ExpectedOK: true,
		},
		{
			Rune: 0x31, // '1'
			ExpectedByte: 53,
			ExpectedOK: true,
		},
		{
			Rune: 0x32, // '2'
			ExpectedByte: 54,
			ExpectedOK: true,
		},
		{
			Rune: 0x33, // '3'
			ExpectedByte: 55,
			ExpectedOK: true,
		},
		{
			Rune: 0x34, // '4'
			ExpectedByte: 56,
			ExpectedOK: true,
		},
		{
			Rune: 0x35, // '5'
			ExpectedByte: 57,
			ExpectedOK: true,
		},
		{
			Rune: 0x36, // '6'
			ExpectedByte: 58,
			ExpectedOK: true,
		},
		{
			Rune: 0x37, // '7'
			ExpectedByte: 59,
			ExpectedOK: true,
		},
		{
			Rune: 0x38, // '8'
			ExpectedByte:60,
			ExpectedOK: true,
		},
		{
			Rune: 0x39, // '9'
			ExpectedByte: 61,
			ExpectedOK: true,
		},
		{
			Rune: 0x3A,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x3B,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x3C,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x3D,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x3E,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x3F,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x40,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x41, // 'A'
			ExpectedByte: 0,
			ExpectedOK: true,
		},
		{
			Rune: 0x42, // 'B'
			ExpectedByte: 1,
			ExpectedOK: true,
		},
		{
			Rune: 0x43, // 'C'
			ExpectedByte: 2,
			ExpectedOK: true,
		},
		{
			Rune: 0x44, // 'D'
			ExpectedByte: 3,
			ExpectedOK: true,
		},
		{
			Rune: 0x45, // 'E'
			ExpectedByte: 4,
			ExpectedOK: true,
		},
		{
			Rune: 0x46, // 'F'
			ExpectedByte: 5,
			ExpectedOK: true,
		},
		{
			Rune: 0x47, // 'G'
			ExpectedByte: 6,
			ExpectedOK: true,
		},
		{
			Rune: 0x48, // 'H'
			ExpectedByte: 7,
			ExpectedOK: true,
		},
		{
			Rune: 0x49, // 'I'
			ExpectedByte: 8,
			ExpectedOK: true,
		},
		{
			Rune: 0x4A, // 'J'
			ExpectedByte: 9,
			ExpectedOK: true,
		},
		{
			Rune: 0x4B, // 'K'
			ExpectedByte: 10,
			ExpectedOK: true,
		},
		{
			Rune: 0x4C, // 'L'
			ExpectedByte: 11,
			ExpectedOK: true,
		},
		{
			Rune: 0x4D, // 'M'
			ExpectedByte: 12,
			ExpectedOK: true,
		},
		{
			Rune: 0x4E, // 'N'
			ExpectedByte: 13,
			ExpectedOK: true,
		},
		{
			Rune: 0x4F, // 'O'
			ExpectedByte: 14,
			ExpectedOK: true,
		},
		{
			Rune: 0x50, // 'P'
			ExpectedByte: 15,
			ExpectedOK: true,
		},
		{
			Rune: 0x51, // 'Q'
			ExpectedByte: 16,
			ExpectedOK: true,
		},
		{
			Rune: 0x52, // 'R'
			ExpectedByte: 17,
			ExpectedOK: true,
		},
		{
			Rune: 0x53, // 'S'
			ExpectedByte: 18,
			ExpectedOK: true,
		},
		{
			Rune: 0x54, // 'T'
			ExpectedByte: 19,
			ExpectedOK: true,
		},
		{
			Rune: 0x55, // 'U'
			ExpectedByte: 20,
			ExpectedOK: true,
		},
		{
			Rune: 0x56, // 'V'
			ExpectedByte: 21,
			ExpectedOK: true,
		},
		{
			Rune: 0x57, // 'W'
			ExpectedByte: 22,
			ExpectedOK: true,
		},
		{
			Rune: 0x58, // 'X'
			ExpectedByte: 23,
			ExpectedOK: true,
		},
		{
			Rune: 0x59, // 'Y'
			ExpectedByte: 24,
			ExpectedOK: true,
		},
		{
			Rune: 0x5A, // 'Z'
			ExpectedByte: 25,
			ExpectedOK: true,
		},
		{
			Rune: 0x5B,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x5C,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x5D,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x5E,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x5F, // '_'
			ExpectedByte: 63,
			ExpectedOK: true,
		},
		{
			Rune: 0x60,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x61, // 'a'
			ExpectedByte: 26,
			ExpectedOK: true,
		},
		{
			Rune: 0x62, // 'b'
			ExpectedByte: 27,
			ExpectedOK: true,
		},
		{
			Rune: 0x63, // 'c'
			ExpectedByte: 28,
			ExpectedOK: true,
		},
		{
			Rune: 0x64, // 'd'
			ExpectedByte: 29,
			ExpectedOK: true,
		},
		{
			Rune: 0x65, // 'e'
			ExpectedByte: 30,
			ExpectedOK: true,
		},
		{
			Rune: 0x66, // 'f'
			ExpectedByte: 31,
			ExpectedOK: true,
		},
		{
			Rune: 0x67, // 'g'
			ExpectedByte: 32,
			ExpectedOK: true,
		},
		{
			Rune: 0x68, // 'h'
			ExpectedByte: 33,
			ExpectedOK: true,
		},
		{
			Rune: 0x69, // 'i'
			ExpectedByte: 34,
			ExpectedOK: true,
		},
		{
			Rune: 0x6A, // 'j'
			ExpectedByte: 35,
			ExpectedOK: true,
		},
		{
			Rune: 0x6B, // 'k'
			ExpectedByte: 36,
			ExpectedOK: true,
		},
		{
			Rune: 0x6C, // 'l'
			ExpectedByte: 37,
			ExpectedOK: true,
		},
		{
			Rune: 0x6D, // 'm'
			ExpectedByte: 38,
			ExpectedOK: true,
		},
		{
			Rune: 0x6E, // 'n'
			ExpectedByte: 39,
			ExpectedOK: true,
		},
		{
			Rune: 0x6F, // 'o'
			ExpectedByte: 40,
			ExpectedOK: true,
		},
		{
			Rune: 0x70, // 'p'
			ExpectedByte: 41,
			ExpectedOK: true,
		},
		{
			Rune: 0x71, // 'q'
			ExpectedByte: 42,
			ExpectedOK: true,
		},
		{
			Rune: 0x72, // 'r'
			ExpectedByte: 43,
			ExpectedOK: true,
		},
		{
			Rune: 0x73, // 's'
			ExpectedByte: 44,
			ExpectedOK: true,
		},
		{
			Rune: 0x74, // 't'
			ExpectedByte: 45,
			ExpectedOK: true,
		},
		{
			Rune: 0x75, // 'u'
			ExpectedByte: 46,
			ExpectedOK: true,
		},
		{
			Rune: 0x76, // 'v'
			ExpectedByte: 47,
			ExpectedOK: true,
		},
		{
			Rune: 0x77, // 'w'
			ExpectedByte: 48,
			ExpectedOK: true,
		},
		{
			Rune: 0x78, // 'x'
			ExpectedByte: 49,
			ExpectedOK: true,
		},
		{
			Rune: 0x79, // 'y'
			ExpectedByte: 50,
			ExpectedOK: true,
		},
		{
			Rune: 0x7A, // 'z'
			ExpectedByte: 51,
			ExpectedOK: true,
		},
		{
			Rune: 0x7B,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x7C,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x7D,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x7E,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
		{
			Rune: 0x7F,
			ExpectedByte: 0,
			ExpectedOK: false,
		},
	}

	for testNumber, test := range tests {

		actualByte, actualOK := translateRune(test.Rune)

		{
			expected := test.ExpectedOK
			actual   := actualOK

			if expected != actual {
				t.Errorf("For test #%d, the actual 'ok' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:   %t", actual)
				t.Logf("RUNE: %q (%U)", test.Rune, test.Rune)
				continue
			}
		}

		{
			expected := test.ExpectedByte
			actual   := actualByte

			if expected != actual {
				t.Errorf("For test #%d, the actual translated-byte (where only the least-significant 6 bits are relevant) is not what was expected.", testNumber)
				t.Logf("EXPECTED: %d (0x%02X)", expected, expected)
				t.Logf("ACTUAL:   %d (0x%02X)", actual, actual)
				t.Logf("RUNE: %q (%U)", test.Rune, test.Rune)
				continue
			}
		}
	}
}
