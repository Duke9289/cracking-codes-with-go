package consts

import "testing"

type TestCase struct {
	Message  string
	Key      string
	Decrypt  bool
	Expected string
}

type cipherFunc func(string, string, bool) string

func TestTable(t *testing.T, testCases []TestCase, cipher cipherFunc) {
	for _, testCase := range testCases {
		translated := cipher(testCase.Message, testCase.Key, testCase.Decrypt)

		if translated != testCase.Expected {
			t.Errorf("Expecting %s, got %s", testCase.Expected, translated)
		}
	}
}
