package hashers

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
)

type MD5PasswordHasher struct{}

func (p *MD5PasswordHasher) Encode(password string, salt string) string {
	h := md5.New()
	io.WriteString(h, salt+password)
	hash := fmt.Sprintf("%s", h.Sum(nil))

	return fmt.Sprintf("%s%s%s%s%s", p.Algorithm(), HASH_SEPARATOR, salt, HASH_SEPARATOR, hash)
}

func (p *MD5PasswordHasher) Algorithm() string {
	return "md5"
}

func (p *MD5PasswordHasher) Verify(password string, encoded string) bool {
	results := strings.Split(encoded, HASH_SEPARATOR)

	attempt := p.Encode(results[1], results[2])

	return encoded == attempt
}

func (p *MD5PasswordHasher) SafeSummary(encoded string) PasswordSummary {
	results := strings.Split(encoded, HASH_SEPARATOR)

	return PasswordSummary{
		Algorithm: p.Algorithm(),
		Salt:      results[1],
		Hash:      results[2],
	}
}

func (p *MD5PasswordHasher) Salt() string {
	return RandomString(12)
}
