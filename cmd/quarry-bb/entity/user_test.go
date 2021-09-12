package entity

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestConsistentPasswordHashing(t *testing.T) {
	password := fmt.Sprintf("%x", rand.Int63())

	hashedA, _ := NewHashedPassword(password)
	hashedB, _ := NewHashedPassword(password)

	if !ComparePassword(password, hashedA) {
		t.Error("expected password hash to match password but it did not")
	}

	if !ComparePassword(password, hashedB) {
		t.Error("expected password hash to match password but it did not")
	}

	if ComparePassword("abcdefg", hashedA) {
		t.Error("expected password hash not to match, but it did")
	}

	if ComparePassword("abcdefg", hashedB) {
		t.Error("expected password hash not to match, but it did")
	}
}
