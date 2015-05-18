package hashers

import (
	"crypto/sha1"
	"fmt"
	"io"
	"strings"
)

type SHA1PasswordHasher struct{}

func (p *SHA1PasswordHasher) Encode(password string, salt string) string {
	h := sha1.New()
	io.WriteString(h, salt+password)
	hash := fmt.Sprintf("%s", h.Sum(nil))
	return fmt.Sprintf("%s%s%s%s%s", p.Algorithm(), HASH_SEPARATOR, salt, HASH_SEPARATOR, hash)
}

func (p *SHA1PasswordHasher) Algorithm() string {
	return "sha1"
}

func (p *SHA1PasswordHasher) Verify(password string, encoded string) bool {
	results := strings.Split(encoded, HASH_SEPARATOR)

	attempt := p.Encode(results[1], results[2])

	return encoded == attempt
}

func (p *SHA1PasswordHasher) SafeSummary(encoded string) PasswordSummary {
	results := strings.Split(encoded, HASH_SEPARATOR)

	return PasswordSummary{
		Algorithm: p.Algorithm(),
		Salt:      results[1],
		Hash:      results[2],
	}
}

func (p *SHA1PasswordHasher) Salt() string {
	return RandomString(12)
}
