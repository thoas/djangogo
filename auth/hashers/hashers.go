package hashers

import (
	"fmt"
)

const HASH_SEPARATOR = "$"

type PasswordSummary struct {
	Algorithm string
	Salt      string
	Hash      string
}

type PasswordHasher interface {
	Encode(password string, salt string) string
	Verify(password string, encoded string) bool
	SafeSummary(encoded string) PasswordSummary
	Algorithm() string
	Salt() string
}

var PasswordHashers = []PasswordHasher{
	&MD5PasswordHasher{},
	&SHA1PasswordHasher{},
	&PBKDF2PasswordHasher{},
	&PBKDF2SHA1PasswordHasher{},
}

func OrderedPasswordHashers() map[string]PasswordHasher {
	results := make(map[string]PasswordHasher)

	for _, hasher := range PasswordHashers {
		results[hasher.Algorithm()] = hasher
	}

	return results
}

func MakePassword(password string, salt string, hasher string) (string, error) {
	passwordHasher, ok := OrderedPasswordHashers()[hasher]

	if !ok {
		return "", fmt.Errorf("%s is not a valid hasher", hasher)
	}

	if salt == "" {
		salt = passwordHasher.Salt()
	}

	return passwordHasher.Encode(password, salt), nil
}
