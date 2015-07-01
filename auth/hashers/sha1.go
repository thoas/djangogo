package hashers

import (
	"crypto/sha1"
	"fmt"
	"strings"
)

// SHA1PasswordHasher is the SHA1 password hasher.
type SHA1PasswordHasher struct{}

// Encode encodes the given password (adding the given salt) then returns encoded.
func (p *SHA1PasswordHasher) Encode(password string, salt string) string {
	return fmt.Sprintf("%s%s%s%s%s",
		p.Algorithm(),
		HASH_SEPARATOR,
		salt,
		HASH_SEPARATOR,
		fmt.Sprintf("%x", sha1.Sum([]byte(salt+password))))
}

// Algorithm returns the algorithm name of this hasher.
func (p *SHA1PasswordHasher) Algorithm() string {
	return "sha1"
}

// Verify takes the raw password and the encoded one, then checks if they match.
func (p *SHA1PasswordHasher) Verify(password string, encoded string) bool {
	results := strings.Split(encoded, HASH_SEPARATOR)
	attempt := p.Encode(password, results[1])
	return encoded == attempt
}

// SafeSummary returns a summary of the encoded password.
func (p *SHA1PasswordHasher) SafeSummary(encoded string) PasswordSummary {
	results := strings.Split(encoded, HASH_SEPARATOR)
	return PasswordSummary{
		Algorithm: p.Algorithm(),
		Salt:      results[1],
		Hash:      results[2],
	}
}

// Salt returns the default salt (which defaults to a random 12 characters string).
func (p *SHA1PasswordHasher) Salt() string {
	return RandomString(12)
}
