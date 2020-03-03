package detect

import (
	"testing"
)

func TestEnglish_pass(t *testing.T) {
	message := "The quick brown fox jumps over the lazy dog."
	if !IsEnglish(message) {
		t.Errorf("Expected '%s' to be true, but was false", message)
	}
}

func TestEnglish_fail(t *testing.T) {
	message := "Xcrh ciyh ,rsi km,d sc..bbcgds smkc lxszkd"
	if IsEnglish(message) {
		t.Errorf("Expected '%s' to be false, but was true", message)
	}
}

func TestEnglishCustom_fail(t *testing.T) {
	message := "The quick icohra brown fox jumps qxckgh over the lazy dog"
	if IsEnglishCustom(message, 90.0, 85.0) {
		t.Errorf("Expected '%s' to be false, but was true", message)
	}
}
