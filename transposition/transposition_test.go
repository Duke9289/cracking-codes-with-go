package transposition

import (
	"testing"

	"github.com/Duke9289/cracking-codes-with-go/consts"
)

func TestTransposition(t *testing.T) {
	testCases := []consts.TestCase{
		{"Common sense is not so common.", "8", false, "Cenoonommstmme oo snnio. s s c"},
	}
	consts.TestTable(t, testCases, Cipher)
}
