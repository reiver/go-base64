package base64

import (
	"testing"
)

func TestDecodeRunesToBytes(t *testing.T) {

	tests := []struct{
		Rune1 rune
		Rune2 rune
		Rune3 rune
		Rune4 rune
		ExpectedLength int
		ExpectedByte1 byte
		ExpectedByte2 byte
		ExpectedByte3 byte
	}{
		{
			Rune1: 'A', // -> 000000
			Rune2: 'A', // -> 00 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 0,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'A', // -> 000000
			Rune2: 'Q', // -> 01 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 1,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'A', // -> 000000
			Rune2: 'g', // -> 10 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 2,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'A', // -> 000000
			Rune2: 'w', // -> 11 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 3,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},

		{
			Rune1: 'B', // -> 000001
			Rune2: 'A', // -> 00 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 4,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'B', // -> 000001
			Rune2: 'Q', // -> 01 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 5,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'B', // -> 000001
			Rune2: 'g', // -> 10 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 6,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'B', // -> 000001
			Rune2: 'w', // -> 11 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 7,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},

		{
			Rune1: 'C', // -> 000010
			Rune2: 'A', // -> 00 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 8,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'C', // -> 000010
			Rune2: 'Q', // -> 01 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 9,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'C', // -> 000010
			Rune2: 'g', // -> 10 000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 10,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'C', // -> 000010
			Rune2: 'w', // -> 11 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 11,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},

		{
			Rune1: 'D', // -> 000011
			Rune2: 'A', // -> 00 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 12,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'D', // -> 000011
			Rune2: 'Q', // -> 01 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 13,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'D', // -> 000011
			Rune2: 'g', // -> 10 000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 14,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'D', // -> 000011
			Rune2: 'w', // -> 11 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 15,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},

		{
			Rune1: 'E', // -> 000100
			Rune2: 'A', // -> 00 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 16,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'E', // -> 000100
			Rune2: 'Q', // -> 01 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 17,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'E', // -> 000100
			Rune2: 'g', // -> 10 000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 18,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'E', // -> 000100
			Rune2: 'w', // -> 11 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 19,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},

		{
			Rune1: 'F', // -> 000101
			Rune2: 'A', // -> 00 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 20,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'F', // -> 000101
			Rune2: 'Q', // -> 01 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 21,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'F', // -> 000101
			Rune2: 'g', // -> 10 000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 22,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'F', // -> 000101
			Rune2: 'w', // -> 11 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 23,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},

		{
			Rune1: 'G', // -> 000110
			Rune2: 'A', // -> 00 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 24,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'G', // -> 000110
			Rune2: 'Q', // -> 01 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 25,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'G', // -> 000110
			Rune2: 'g', // -> 10 000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 26,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'G', // -> 000110
			Rune2: 'w', // -> 11 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 27,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},

		{
			Rune1: 'H', // -> 000111
			Rune2: 'A', // -> 00 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 28,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'H', // -> 000111
			Rune2: 'Q', // -> 01 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 29,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'H', // -> 000111
			Rune2: 'g', // -> 10 000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 30,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'H', // -> 000111
			Rune2: 'w', // -> 11 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 31,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},






		{
			Rune1: 'a', // -> 011010
			Rune2: 'A', // -> 00 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 104,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'a', // -> 011010
			Rune2: 'Q', // -> 01 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 105,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'a', // -> 011010
			Rune2: 'g', // -> 10 000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 106,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'a', // -> 011010
			Rune2: 'w', // -> 11 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 107,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},

		{
			Rune1: 'b', // -> 011011
			Rune2: 'A', // -> 00 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 108,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'b', // -> 011011
			Rune2: 'Q', // -> 01 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 109,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'b', // -> 011011
			Rune2: 'g', // -> 10 000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 110,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'b', // -> 011011
			Rune2: 'w', // -> 11 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 111,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},

		{
			Rune1: 'c', // -> 011100
			Rune2: 'A', // -> 00 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 112,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'c', // -> 011100
			Rune2: 'Q', // -> 01 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 113,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'c', // -> 011100
			Rune2: 'g', // -> 10 000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 114,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'c', // -> 011100
			Rune2: 'w', // -> 11 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 115,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},

		{
			Rune1: 'd', // -> 011101
			Rune2: 'A', // -> 00 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 116,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'd', // -> 011101
			Rune2: 'Q', // -> 01 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 117,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'd', // -> 011101
			Rune2: 'g', // -> 10 000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 118,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'd', // -> 011101
			Rune2: 'w', // -> 11 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 119,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},

		{
			Rune1: 'e', // -> 011110
			Rune2: 'A', // -> 00 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 120,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'e', // -> 011110
			Rune2: 'Q', // -> 01 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 121,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'e', // -> 011110
			Rune2: 'g', // -> 10 000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 122,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'e', // -> 011110
			Rune2: 'w', // -> 11 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 123,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},






		{
			Rune1: '0', // -> 110100
			Rune2: 'A', // -> 00 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 208,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: '0', // -> 110100
			Rune2: 'Q', // -> 01 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 209,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: '0', // -> 110100
			Rune2: 'g', // -> 10 000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 210,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: '0', // -> 110100
			Rune2: 'w', // -> 11 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 211,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},

		{
			Rune1: '1', // -> 110101
			Rune2: 'A', // -> 00 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 212,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: '1', // -> 110101
			Rune2: 'Q', // -> 01 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 213,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: '1', // -> 110101
			Rune2: 'g', // -> 10 000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 214,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: '1', // -> 110101
			Rune2: 'w', // -> 11 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 215,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},

		{
			Rune1: '2', // -> 110110
			Rune2: 'A', // -> 00 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 216,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: '2', // -> 110110
			Rune2: 'Q', // -> 01 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 217,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: '2', // -> 110110
			Rune2: 'g', // -> 10 000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 218,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: '2', // -> 110110
			Rune2: 'w', // -> 11 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 219,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},

		{
			Rune1: '3', // -> 110111
			Rune2: 'A', // -> 00 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 220,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: '3', // -> 110111
			Rune2: 'Q', // -> 01 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 221,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: '3', // -> 110111
			Rune2: 'g', // -> 10 000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 222,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: '3', // -> 110111
			Rune2: 'w', // -> 11 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 223,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},

		{
			Rune1: '4', // -> 111000
			Rune2: 'A', // -> 00 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 224,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: '4', // -> 111000
			Rune2: 'Q', // -> 01 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 225,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: '4', // -> 111000
			Rune2: 'g', // -> 10 000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 226,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: '4', // -> 111000
			Rune2: 'w', // -> 11 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 227,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},

		{
			Rune1: '5', // -> 111001
			Rune2: 'A', // -> 00 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 228,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: '5', // -> 111001
			Rune2: 'Q', // -> 01 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 229,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: '5', // -> 111001
			Rune2: 'g', // -> 10 000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 230,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: '5', // -> 111001
			Rune2: 'w', // -> 11 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 231,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},






		{
			Rune1: '+', // -> 111110
			Rune2: 'A', // -> 00 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 248,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: '+', // -> 111110
			Rune2: 'Q', // -> 01 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 249,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: '+', // -> 111110
			Rune2: 'g', // -> 10 000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 250,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: '+', // -> 111110
			Rune2: 'w', // -> 11 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 251,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},



		{
			Rune1: '-', // -> 111110
			Rune2: 'A', // -> 00 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 248,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: '-', // -> 111110
			Rune2: 'Q', // -> 01 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 249,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: '-', // -> 111110
			Rune2: 'g', // -> 10 000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 250,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: '-', // -> 111110
			Rune2: 'w', // -> 11 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 251,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},






		{
			Rune1: '/', // -> 111111
			Rune2: 'A', // -> 00 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 252,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: '/', // -> 111111
			Rune2: 'Q', // -> 01 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 253,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: '/', // -> 111111
			Rune2: 'g', // -> 10 000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 254,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: '/', // -> 111111
			Rune2: 'w', // -> 11 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 255,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},



		{
			Rune1: ',', // -> 111111
			Rune2: 'A', // -> 00 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 252,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: ',', // -> 111111
			Rune2: 'Q', // -> 01 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 253,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: ',', // -> 111111
			Rune2: 'g', // -> 10 000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 254,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: ',', // -> 111111
			Rune2: 'w', // -> 11 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 255,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},



		{
			Rune1: '_', // -> 111111
			Rune2: 'A', // -> 00 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 252,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: '_', // -> 111111
			Rune2: 'Q', // -> 01 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 253,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: '_', // -> 111111
			Rune2: 'g', // -> 10 000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 254,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: '_', // -> 111111
			Rune2: 'w', // -> 11 0000
			Rune3: '=',
			Rune4: '=',
			ExpectedLength: 1,
			ExpectedByte1: 255,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},



		{
			Rune1: 'A', // -> 000000
			Rune2: 'A', // -> 00 0000
			Rune3: 'A', // -> 0000 00
			Rune4: '=',
			ExpectedLength: 2,
			ExpectedByte1: 0,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'A', // -> 000000
			Rune2: 'Q', // -> 01 0000
			Rune3: 'A', // -> 0000 00
			Rune4: '=',
			ExpectedLength: 2,
			ExpectedByte1: 1,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'A', // -> 000000
			Rune2: 'g', // -> 10 0000
			Rune3: 'A', // -> 0000 00
			Rune4: '=',
			ExpectedLength: 2,
			ExpectedByte1: 2,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'A', // -> 000000
			Rune2: 'w', // -> 11 0000
			Rune3: 'A', // -> 0000 00
			Rune4: '=',
			ExpectedLength: 2,
			ExpectedByte1: 3,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},

		{
			Rune1: 'B', // -> 000001
			Rune2: 'A', // -> 00 0000
			Rune3: 'A', // -> 0000 00
			Rune4: '=',
			ExpectedLength: 2,
			ExpectedByte1: 4,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'B', // -> 000001
			Rune2: 'Q', // -> 01 0000
			Rune3: 'A', // -> 0000 00
			Rune4: '=',
			ExpectedLength: 2,
			ExpectedByte1: 5,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'B', // -> 000001
			Rune2: 'g', // -> 10 0000
			Rune3: 'A', // -> 0000 00
			Rune4: '=',
			ExpectedLength: 2,
			ExpectedByte1: 6,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'B', // -> 000001
			Rune2: 'w', // -> 11 0000
			Rune3: 'A', // -> 0000 00
			Rune4: '=',
			ExpectedLength: 2,
			ExpectedByte1: 7,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},

		{
			Rune1: 'C', // -> 000010
			Rune2: 'A', // -> 00 0000
			Rune3: 'A', // -> 0000 00
			Rune4: '=',
			ExpectedLength: 2,
			ExpectedByte1: 8,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'C', // -> 000010
			Rune2: 'Q', // -> 01 0000
			Rune3: 'A', // -> 0000 00
			Rune4: '=',
			ExpectedLength: 2,
			ExpectedByte1: 9,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'C', // -> 000010
			Rune2: 'g', // -> 10 0000
			Rune3: 'A', // -> 0000 00
			Rune4: '=',
			ExpectedLength: 2,
			ExpectedByte1: 10,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'C', // -> 000010
			Rune2: 'w', // -> 11 0000
			Rune3: 'A', // -> 0000 00
			Rune4: '=',
			ExpectedLength: 2,
			ExpectedByte1: 11,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},






		{
			Rune1: 'A', // -> 000000
			Rune2: 'A', // -> 00 0000
			Rune3: 'A', // -> 0000 00
			Rune4: 'A', // -> 000000
			ExpectedLength: 3,
			ExpectedByte1: 0,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'A', // -> 000000
			Rune2: 'Q', // -> 01 0000
			Rune3: 'A', // -> 0000 00
			Rune4: 'A', // -> 000000
			ExpectedLength: 3,
			ExpectedByte1: 1,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'A', // -> 000000
			Rune2: 'g', // -> 10 0000
			Rune3: 'A', // -> 0000 00
			Rune4: 'A', // -> 000000
			ExpectedLength: 3,
			ExpectedByte1: 2,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'A', // -> 000000
			Rune2: 'w', // -> 11 0000
			Rune3: 'A', // -> 0000 00
			Rune4: 'A', // -> 000000
			ExpectedLength: 3,
			ExpectedByte1: 3,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},

		{
			Rune1: 'B', // -> 000001
			Rune2: 'A', // -> 00 0000
			Rune3: 'A', // -> 0000 00
			Rune4: 'A', // -> 000000
			ExpectedLength: 3,
			ExpectedByte1: 4,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'B', // -> 000001
			Rune2: 'Q', // -> 01 0000
			Rune3: 'A', // -> 0000 00
			Rune4: 'A', // -> 000000
			ExpectedLength: 3,
			ExpectedByte1: 5,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'B', // -> 000001
			Rune2: 'g', // -> 10 0000
			Rune3: 'A', // -> 0000 00
			Rune4: 'A', // -> 000000
			ExpectedLength: 3,
			ExpectedByte1: 6,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'B', // -> 000001
			Rune2: 'w', // -> 11 0000
			Rune3: 'A', // -> 0000 00
			Rune4: 'A', // -> 000000
			ExpectedLength: 3,
			ExpectedByte1: 7,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},

		{
			Rune1: 'C', // -> 000010
			Rune2: 'A', // -> 00 0000
			Rune3: 'A', // -> 0000 00
			Rune4: 'A', // -> 000000
			ExpectedLength: 3,
			ExpectedByte1: 8,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'C', // -> 000010
			Rune2: 'Q', // -> 01 0000
			Rune3: 'A', // -> 0000 00
			Rune4: 'A', // -> 000000
			ExpectedLength: 3,
			ExpectedByte1: 9,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'C', // -> 000010
			Rune2: 'g', // -> 10 0000
			Rune3: 'A', // -> 0000 00
			Rune4: 'A', // -> 000000
			ExpectedLength: 3,
			ExpectedByte1: 10,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},
		{
			Rune1: 'C', // -> 000010
			Rune2: 'w', // -> 11 0000
			Rune3: 'A', // -> 0000 00
			Rune4: 'A', // -> 000000
			ExpectedLength: 3,
			ExpectedByte1: 11,
			ExpectedByte2: 0,
			ExpectedByte3: 0,
		},









		{
			Rune1: 'A', // -> 000000
			Rune2: 'Q', // -> 01 0000
			Rune3: 'E', // -> 0001 00
			Rune4: 'B', // -> 000001
			ExpectedLength: 3,
			ExpectedByte1: 1,
			ExpectedByte2: 1,
			ExpectedByte3: 1,
		},






		{
			Rune1: 'g', // -> 100000
			Rune2: 'I', // -> 00 1000
			Rune3: 'C', // -> 0000 10
			Rune4: 'A', // -> 000000
			ExpectedLength: 3,
			ExpectedByte1: 128,
			ExpectedByte2: 128,
			ExpectedByte3: 128,
		},






		{
			Rune1: 'h', // -> 100001
			Rune2: 'h', // -> 10 0001
			Rune3: 'h', // -> 1000 01
			Rune4: 'h', // -> 100001
			ExpectedLength: 3,
			ExpectedByte1: 134,
			ExpectedByte2: 24,
			ExpectedByte3: 97,
		},
	}

	for testNumber, test := range tests {

		actualByte1, actualByte2, actualByte3, actualLength, err := decodeRunesToBytes(test.Rune1, test.Rune2, test.Rune3, test.Rune4)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("RUNE-1: %q (%U)", test.Rune1, test.Rune1)
			t.Logf("RUNE-2: %q (%U)", test.Rune2, test.Rune2)
			t.Logf("RUNE-3: %q (%U)", test.Rune3, test.Rune3)
			t.Logf("RUNE-4: %q (%U)", test.Rune4, test.Rune4)
			continue
		}

		{
			expected := test.ExpectedLength
			actual := actualLength

			if expected != actual {
				t.Errorf("For test #%d, the number of (valid) bytes returs (i.e., 'length') is not what was expected." , testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				t.Logf("RUNE-1: %q (%U)", test.Rune1, test.Rune1)
				t.Logf("RUNE-2: %q (%U)", test.Rune2, test.Rune2)
				t.Logf("RUNE-3: %q (%U)", test.Rune3, test.Rune3)
				t.Logf("RUNE-4: %q (%U)", test.Rune4, test.Rune4)
				t.Logf("EXPECTED-BYTE_1: %#08b (0x%02X) (%d)", test.ExpectedByte1, test.ExpectedByte1, test.ExpectedByte1)
				t.Logf("EXPECTED-BYTE-2: %#08b (0x%02X) (%d)", test.ExpectedByte2, test.ExpectedByte2, test.ExpectedByte2)
				t.Logf("EXPECTED-BYTE-3: %#08b (0x%02X) (%d)", test.ExpectedByte3, test.ExpectedByte3, test.ExpectedByte3)
				t.Logf("ACTUAL-BYTE-1:   %#08b (0x%02X) (%d)", actualByte1, actualByte1, actualByte1)
				t.Logf("ACTUAL-BYTE-2:   %#08b (0x%02X) (%d)", actualByte2, actualByte2, actualByte2)
				t.Logf("ACTUAL-BYTE-3:   %#08b (0x%02X) (%d)", actualByte3, actualByte3, actualByte3)
				continue
			}
		}

		{
			expected := test.ExpectedByte1
			actual := actualByte1

			if expected != actual {
				t.Errorf("For test #%d, the actual value for byte №1 is not what was expected." , testNumber)
				t.Logf("EXPECTED: %#08b (0x%02X) (%d)", expected, expected, expected)
				t.Logf("ACTUAL:   %#08b (0x%02X) (%d)", actual, actual, actual)
				t.Logf("RUNE-1: %q (%U)", test.Rune1, test.Rune1)
				t.Logf("RUNE-2: %q (%U)", test.Rune2, test.Rune2)
				t.Logf("RUNE-3: %q (%U)", test.Rune3, test.Rune3)
				t.Logf("RUNE-4: %q (%U)", test.Rune4, test.Rune4)
				t.Logf("EXPECTED-BYTE_1: %#08b (0x%02X) (%d)", test.ExpectedByte1, test.ExpectedByte1, test.ExpectedByte1)
				t.Logf("EXPECTED-BYTE-2: %#08b (0x%02X) (%d)", test.ExpectedByte2, test.ExpectedByte2, test.ExpectedByte2)
				t.Logf("EXPECTED-BYTE-3: %#08b (0x%02X) (%d)", test.ExpectedByte3, test.ExpectedByte3, test.ExpectedByte3)
				t.Logf("ACTUAL-BYTE-1:   %#08b (0x%02X) (%d)", actualByte1, actualByte1, actualByte1)
				t.Logf("ACTUAL-BYTE-2:   %#08b (0x%02X) (%d)", actualByte2, actualByte2, actualByte2)
				t.Logf("ACTUAL-BYTE-3:   %#08b (0x%02X) (%d)", actualByte3, actualByte3, actualByte3)
				continue
			}
		}

		{
			expected := test.ExpectedByte2
			actual := actualByte2

			if expected != actual {
				t.Errorf("For test #%d, the actual value for byte №2 is not what was expected." , testNumber)
				t.Logf("EXPECTED: %#08b (0x%02X) (%d)", expected, expected, expected)
				t.Logf("ACTUAL:   %#08b (0x%02X) (%d)", actual, actual, actual)
				t.Logf("RUNE-1: %q (%U)", test.Rune1, test.Rune1)
				t.Logf("RUNE-2: %q (%U)", test.Rune2, test.Rune2)
				t.Logf("RUNE-3: %q (%U)", test.Rune3, test.Rune3)
				t.Logf("RUNE-4: %q (%U)", test.Rune4, test.Rune4)
				t.Logf("EXPECTED-BYTE_1: %#08b (0x%02X) (%d)", test.ExpectedByte1, test.ExpectedByte1, test.ExpectedByte1)
				t.Logf("EXPECTED-BYTE-2: %#08b (0x%02X) (%d)", test.ExpectedByte2, test.ExpectedByte2, test.ExpectedByte2)
				t.Logf("EXPECTED-BYTE-3: %#08b (0x%02X) (%d)", test.ExpectedByte3, test.ExpectedByte3, test.ExpectedByte3)
				t.Logf("ACTUAL-BYTE-1:   %#08b (0x%02X) (%d)", actualByte1, actualByte1, actualByte1)
				t.Logf("ACTUAL-BYTE-2:   %#08b (0x%02X) (%d)", actualByte2, actualByte2, actualByte2)
				t.Logf("ACTUAL-BYTE-3:   %#08b (0x%02X) (%d)", actualByte3, actualByte3, actualByte3)
				continue
			}
		}

		{
			expected := test.ExpectedByte3
			actual := actualByte3

			if expected != actual {
				t.Errorf("For test #%d, the actual value for byte №3 is not what was expected." , testNumber)
				t.Logf("EXPECTED: %#08b (0x%02X) (%d)", expected, expected, expected)
				t.Logf("ACTUAL:   %#08b (0x%02X) (%d)", actual, actual, actual)
				t.Logf("RUNE-1: %q (%U)", test.Rune1, test.Rune1)
				t.Logf("RUNE-2: %q (%U)", test.Rune2, test.Rune2)
				t.Logf("RUNE-3: %q (%U)", test.Rune3, test.Rune3)
				t.Logf("RUNE-4: %q (%U)", test.Rune4, test.Rune4)
				t.Logf("EXPECTED-BYTE_1: %#08b (0x%02X) (%d)", test.ExpectedByte1, test.ExpectedByte1, test.ExpectedByte1)
				t.Logf("EXPECTED-BYTE-2: %#08b (0x%02X) (%d)", test.ExpectedByte2, test.ExpectedByte2, test.ExpectedByte2)
				t.Logf("EXPECTED-BYTE-3: %#08b (0x%02X) (%d)", test.ExpectedByte3, test.ExpectedByte3, test.ExpectedByte3)
				t.Logf("ACTUAL-BYTE-1:   %#08b (0x%02X) (%d)", actualByte1, actualByte1, actualByte1)
				t.Logf("ACTUAL-BYTE-2:   %#08b (0x%02X) (%d)", actualByte2, actualByte2, actualByte2)
				t.Logf("ACTUAL-BYTE-3:   %#08b (0x%02X) (%d)", actualByte3, actualByte3, actualByte3)
				continue
			}
		}
	}
}
