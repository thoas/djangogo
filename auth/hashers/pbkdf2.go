package hashers

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"strings"
)

type PBKDF2PasswordHasher struct{}

func (p *PBKDF2PasswordHasher) Encode(password string, salt string) string {
	return p.EncodeWithIteration(password, salt, 1500)
}

func (p *PBKDF2PasswordHasher) EncodeWithIteration(password string, salt string, iter int) string {
	hash := fmt.Sprintf("%s", pbkdf2.Key([]byte(password), []byte(salt), iter, 32, sha256.New))
	hash = base64.StdEncoding.EncodeToString([]byte(hash))

	return fmt.Sprintf("%s%s%s%s%s%s%s", p.Algorithm(), HASH_SEPARATOR, iter, HASH_SEPARATOR, salt, HASH_SEPARATOR, hash)
}

func (p *PBKDF2PasswordHasher) Algorithm() string {
	return "pbkdf2_sha256"
}

func (p *PBKDF2PasswordHasher) Verify(password string, encoded string) bool {
	results := strings.Split(encoded, HASH_SEPARATOR)

	attempt := p.Encode(results[1], results[2])

	return encoded == attempt
}

func (p *PBKDF2PasswordHasher) SafeSummary(encoded string) PasswordSummary {
	results := strings.Split(encoded, HASH_SEPARATOR)

	return PasswordSummary{
		Algorithm: p.Algorithm(),
		Salt:      results[2],
		Hash:      results[3],
	}
}

func (p *PBKDF2PasswordHasher) Salt() string {
	return RandomString(12)
}

type PBKDF2SHA1PasswordHasher struct{}

func (p *PBKDF2SHA1PasswordHasher) Encode(password string, salt string) string {
	return p.EncodeWithIteration(password, salt, 1500)
}

func (p *PBKDF2SHA1PasswordHasher) EncodeWithIteration(password string, salt string, iter int) string {
	hash := fmt.Sprintf("%s", pbkdf2.Key([]byte(password), []byte(salt), iter, 32, sha1.New))
	hash = base64.StdEncoding.EncodeToString([]byte(hash))

	return fmt.Sprintf("%s%s%s%s%s%s%s", p.Algorithm(), HASH_SEPARATOR, iter, HASH_SEPARATOR, salt, HASH_SEPARATOR, hash)
}

func (p *PBKDF2SHA1PasswordHasher) Algorithm() string {
	return "pbkdf2_sha1"
}

func (p *PBKDF2SHA1PasswordHasher) Verify(password string, encoded string) bool {
	results := strings.Split(encoded, HASH_SEPARATOR)

	attempt := p.Encode(results[1], results[2])

	return encoded == attempt
}

func (p *PBKDF2SHA1PasswordHasher) SafeSummary(encoded string) PasswordSummary {
	results := strings.Split(encoded, HASH_SEPARATOR)

	return PasswordSummary{
		Algorithm: p.Algorithm(),
		Salt:      results[2],
		Hash:      results[3],
	}
}

func (p *PBKDF2SHA1PasswordHasher) Salt() string {
	return RandomString(12)
}
