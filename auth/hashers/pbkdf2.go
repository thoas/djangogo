package hashers

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

// PBKDF2PasswordHasher is the PBKDF2 SHA1 password hasher.
type PBKDF2PasswordHasher struct{}

// Encode encodes the given password (adding the given salt) then returns encoded.
func (p *PBKDF2PasswordHasher) Encode(password string, salt string) string {
	return p.EncodeWithIteration(password, salt, 1500)
}

// EncodeWithIteration encodes the given password (adding the given salt) with
// iteration then returns encoded.
func (p *PBKDF2PasswordHasher) EncodeWithIteration(password string, salt string, iter int) string {
	hash := fmt.Sprintf("%s", pbkdf2.Key([]byte(password), []byte(salt), iter, 32, sha256.New))
	hash = base64.StdEncoding.EncodeToString([]byte(hash))
	return fmt.Sprintf("%s%s%d%s%s%s%s", p.Algorithm(), HASH_SEPARATOR, iter, HASH_SEPARATOR, salt, HASH_SEPARATOR, hash)
}

// Algorithm returns the algorithm name of this hasher.
func (p *PBKDF2PasswordHasher) Algorithm() string {
	return "pbkdf2_sha256"
}

// Verify takes the raw password and the encoded one, then checks if they match.
func (p *PBKDF2PasswordHasher) Verify(password string, encoded string) bool {
	results := strings.Split(encoded, HASH_SEPARATOR)
	attempt := p.Encode(password, results[2])
	return encoded == attempt
}

// SafeSummary returns a summary of the encoded password.
func (p *PBKDF2PasswordHasher) SafeSummary(encoded string) PasswordSummary {
	results := strings.Split(encoded, HASH_SEPARATOR)
	return PasswordSummary{
		Algorithm: p.Algorithm(),
		Salt:      results[2],
		Hash:      results[3],
	}
}

// Salt returns the default salt (which defaults to a random 12 characters string).
func (p *PBKDF2PasswordHasher) Salt() string {
	return RandomString(12)
}
