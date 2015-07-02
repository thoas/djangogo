package hashers

import (
	"fmt"
	"strings"
)

// PasswordSummary contains encoded password information.
type PasswordSummary struct {
	Algorithm string
	Salt      string
	Hash      string
}

// PasswordHasher is a password hasher.
type PasswordHasher interface {
	Encode(password string, salt string) string
	Verify(password string, encoded string) bool
	SafeSummary(encoded string) PasswordSummary
	Algorithm() string
	Salt() string
}

// PasswordHashers are supported password hashers.
var PasswordHashers = []PasswordHasher{
	&MD5PasswordHasher{},
	&SHA1PasswordHasher{},
	&PBKDF2PasswordHasher{},
	&PBKDF2SHA1PasswordHasher{},
}

// CheckPassword checks if the raw password matches with the encoded one.
// Returns true if they match. Otherwise false.
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

// IsPasswordUsable return true the encoded password is usable.
func IsPasswordUsable(encoded string) bool {
	if strings.HasPrefix(encoded, UnusablePasswordPrefix) {
		return false
	}

	_, err := IdentityHasher(encoded)

	if err != nil {
		return false
	}

	return true
}

// IdentityHasher takes an encoded password and returns its corresponding hasher.
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

// OrderedPasswordHashers returns password hashers map ordered by their algorithm name.
func OrderedPasswordHashers() map[string]PasswordHasher {
	results := make(map[string]PasswordHasher)

	for _, hasher := range PasswordHashers {
		results[hasher.Algorithm()] = hasher
	}

	return results
}

// MakePassword takes a raw password, a salt and a given hasher (see password hasher
// algorithm name -- Algorithm() method) then returns this password encoded.
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
