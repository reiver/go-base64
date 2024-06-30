package base64

// translateRune returns the 6-bit translated rune.
//
// Since Go doesn't have a 6-bit type, the translated rune is returned as a byte,
// with the understanding that the most-significant 2 bits are to be ignored.
//
// translateRune handles a number of different variantions of base64.
// Includeing: RFC-1421, RFC-2045, RFC-2152, RFC-3501, RFC-4648 (base64 & base64url), RFC-4880.
//
// This 'handling' of all those different variations of base64 comes into play with the encoded
// form for 0b111110 (62) and 0b111111 (63).
//
// For accepted symbols for 0b111110 (62) are: '+', '-'.
//
// For accepted symbols for 0b111111 (63) are: '/', ',', '_'.
//
// I.e., translateRune does this:
//
//	RUNE   BINARY    DECIMAL
//	
//	'A' -> 000000     0
//	'B' -> 000001     1
//	'C' -> 000010     2
//	'D' -> 000011     3
//	'E' -> 000100     4
//	'F' -> 000101     5
//	'G' -> 000110     6
//	'H' -> 000111     7
//	'I' -> 001000     8
//	'J' -> 001001     9
//	'K' -> 001010    10
//	'L' -> 001011    11
//	'M' -> 001100    12
//	'N' -> 001101    13
//	'O' -> 001110    14
//	'P' -> 001111    15
//	'Q' -> 010000    16
//	'R' -> 010001    17
//	'S' -> 010010    18
//	'T' -> 010011    19
//	'U' -> 010100    20
//	'V' -> 010101    21
//	'W' -> 010110    22
//	'X' -> 010111    23
//	'Y' -> 011000    24
//	'Z' -> 011001    25
//	
//	'a' -> 011010    26
//	'b' -> 011011    27
//	'c' -> 011100    28
//	'd' -> 011101    29
//	'e' -> 011110    30
//	'f' -> 011111    31
//	'g' -> 100000    32
//	'h' -> 100001    33
//	'i' -> 100010    34
//	'j' -> 100011    35
//	'k' -> 100100    36
//	'l' -> 100101    37
//	'm' -> 100110    38
//	'n' -> 100111    39
//	'o' -> 101000    40
//	'p' -> 101001    41
//	'q' -> 101010    42
//	'r' -> 101011    43
//	's' -> 101100    44
//	't' -> 101101    45
//	'u' -> 101110    46
//	'v' -> 101111    47
//	'w' -> 110000    48
//	'x' -> 110001    49
//	'y' -> 110010    50
//	'z' -> 110011    51
//	
//	'0' -> 110100    52
//	'1' -> 110101    53
//	'2' -> 110110    54
//	'3' -> 110111    55
//	'4' -> 111000    56
//	'5' -> 111001    57
//	'6' -> 111010    58
//	'7' -> 111011    59
//	'8' -> 111100    60
//	'9' -> 111101    61
//	
//	'+' -> 111110    61
//	'-' -> 111110    61    (base64url)
//	
//	'/' -> 111111    62
//	',' -> 111111    62
//	'_' -> 111111    62    (base64url)
func translateRune(r rune) (byte, bool) {
	switch {
	case 'A'<=r && r<='Z':
		return byte(r)-'A',true
	case 'a'<=r && r<='z':
		return byte(r)-'a'+26,true
	case '0'<=r && r<='9':
		return byte(r)-'0'+52,true
	case '+'==r || '-'==r:
		return 62, true
	case '/'==r || ','==r || '_'==r:
		return 63, true
	default:
		return 0, false
	}
}
