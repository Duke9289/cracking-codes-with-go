package reverse

import (
	"testing"

	"github.com/Duke9289/cracking-codes-with-go/consts"
)

func TestReverse(t *testing.T) {
	testCases := []consts.TestCase{
		{"Hello World!", "", false, "!dlroW olleH"},
		{"The Magic Words are Squeamish Ossifrage", "junk key", true, "egarfissO hsimaeuqS era sdroW cigaM ehT"},
	}

	consts.TestTable(t, testCases, Cipher)
}
