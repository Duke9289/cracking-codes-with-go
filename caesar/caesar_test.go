package caesar

import (
	"testing"

	"github.com/Duke9289/cracking-codes-with-go/consts"
)

func TestCaesar(t *testing.T) {
	testCases := []consts.TestCase{
		{"This is my secret message", "13", false, "guv6Jv6Jz!J6rp5r7Jzr66ntr"},
		{"Hello World!", "3", false, "Khoor.ZruogA"},
		{"ymnxB2fxBjshw4uyjiB2nymBpj4B0", "5", true, "this was encrypted with key 5"},
		{"o3zRhv24xRr0?y.Rv?zRn!Bzv84.3Rj..41?v2z", "21", true, "The Magic Words are Squeamish Ossifrage"},
		//{"", "13", false, ""}, TODO: need test for cracking
	}

	consts.TestTable(t, testCases, Cipher)
}
