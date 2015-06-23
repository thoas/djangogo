package hashers

import (
	"fmt"
	"strings"
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

func CheckPassword(password string, encoded string) bool {
	if IsPasswordUsable(password) {
		return false
	}

	hasher, err := IdentityHasher(encoded)

	if err != nil {
		return false
	}

	return hasher.Verify(password, encoded)
}

func IsPasswordUsable(encoded string) bool {
	if strings.HasPrefix(encoded, UNUSABLE_PASSWORD_PREFIX) {
		return false
	}

	_, err := IdentityHasher(encoded)

	if err != nil {
		return false
	}

	return true
}

func IdentityHasher(encoded string) (PasswordHasher, error) {
	var algorithm = ""

	if (len(encoded) == 32 && !strings.Contains(encoded, "$")) || (len(encoded) == 37 && strings.HasPrefix(encoded, "md5$$")) {
		algorithm = "unsalted_md5"
	}

	if len(encoded) == 46 && strings.HasPrefix(encoded, "sha1$$") {
		algorithm = "unsalted_sha1"
	}

	if algorithm == "" {
		algorithm = strings.Split(encoded, "$")[0]
	}

	hasher, ok := OrderedPasswordHashers()[algorithm]

	if !ok {
		return nil, fmt.Errorf("%s is not a valid hasher", algorithm)
	}

	return hasher, nil
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
