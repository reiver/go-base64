package base64_test

import (
	"testing"

	"bytes"
	"strings"

	"github.com/reiver/go-base64"
)

func TestDecode(t *testing.T) {

	tests := []struct{
		Encoded string
		Expected []byte
	}{
		{
			Encoded: "",
			Expected: []byte(nil),
		},



		{
			Encoded: "AA",
			Expected: []byte{0},
		},
		{
			Encoded: "AA==",
			Expected: []byte{0},
		},



		{
			Encoded: "AQ",
			Expected: []byte{1},
		},
		{
			Encoded: "AQ==",
			Expected: []byte{1},
		},



		{
			Encoded: "Ag",
			Expected: []byte{2},
		},
		{
			Encoded: "Ag==",
			Expected: []byte{2},
		},



		{
			Encoded: "Aw",
			Expected: []byte{3},
		},
		{
			Encoded: "Aw==",
			Expected: []byte{3},
		},









		{
			Encoded: "SGVsbG8gd29ybGQh",
			Expected: []byte("Hello world!"),
		},



		{
			Encoded: "YXBwbGUgQkFOQU5BIENoZXJyeQ",
			Expected: []byte("apple BANANA Cherry"),
		},
		{
			Encoded: "YXBwbGUgQkFOQU5BIENoZXJyeQ==",
			Expected: []byte("apple BANANA Cherry"),
		},






		{
			Encoded: "Zg",
			Expected: []byte("f"),
		},
		{
			Encoded: "Zg==",
			Expected: []byte("f"),
		},



		{
			Encoded: "Zm8",
			Expected: []byte("fo"),
		},
		{
			Encoded: "Zm8=",
			Expected: []byte("fo"),
		},



		{
			Encoded: "Zm9v",
			Expected: []byte("foo"),
		},



		{
			Encoded: "Zm9vYg",
			Expected: []byte("foob"),
		},
		{
			Encoded: "Zm9vYg==",
			Expected: []byte("foob"),
		},



		{
			Encoded: "Zm9vYmE",
			Expected: []byte("fooba"),
		},
		{
			Encoded: "Zm9vYmE=",
			Expected: []byte("fooba"),
		},



		{
			Encoded: "Zm9vYmFy",
			Expected: []byte("foobar"),
		},



		{
			Encoded: "",
			Expected: []byte(""),
		},
	}

	for testNumber, test := range tests {

		var actualBuffer bytes.Buffer

		err := base64.Decode(&actualBuffer, strings.NewReader(test.Encoded))

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("ENCODED: %q", test.Encoded)
			continue
		}

		{
			expected := test.Expected
			actual := actualBuffer.Bytes()

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual decoded data is not what was expected." , testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL  : %#v", actual)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL  : %q", actual)
				t.Logf("ENCODED: %q", test.Encoded)
				continue
			}
		}
	}
}

func TestDecode_errors(t *testing.T) {

	tests := []struct{
		Encoded string
		ExpectedError string
	}{
		{
			Encoded: "AA=",
		},
		{
			Encoded: "AA===",
		},



		{
			Encoded: "AQ=",
		},
		{
			Encoded: "AQ===",
		},



		{
			Encoded: "Ag=",
		},
		{
			Encoded: "Ag===",
		},



		{
			Encoded: "Aw=",
		},
		{
			Encoded: "Aw===",
		},



		{
			Encoded: "YXBwbGUgQkFOQU5BIENoZXJyeQ=",
		},
		{
			Encoded: "YXBwbGUgQkFOQU5BIENoZXJyeQ===",
		},
	}

	for testNumber, test := range tests {

		var actualBuffer bytes.Buffer

		err := base64.Decode(&actualBuffer, strings.NewReader(test.Encoded))

		if nil == err {
			t.Errorf("For test #%d, expected an error but actually get one.", testNumber)
			t.Logf("EXPECTED-ERROR: %q", test.ExpectedError)
			t.Logf("ENCODED: %q", test.Encoded)
			continue
		}
	}
}
