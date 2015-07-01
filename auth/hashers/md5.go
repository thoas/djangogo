package hashers

import (
	"crypto/md5"
	"fmt"
	"strings"
)

//  MD5PasswordHasher is the MD5 password hasher.
type MD5PasswordHasher struct{}

// Encode encodes the given password (adding the given salt) then returns encoded.
func (p *MD5PasswordHasher) Encode(password string, salt string) string {
	return fmt.Sprintf("%s%s%s%s%s",
		p.Algorithm(),
		HashSeparator,
		salt,
		HashSeparator,
		fmt.Sprintf("%x", md5.Sum([]byte(salt+password))))
}

// Algorithm returns the algorithm name of this hasher.
func (p *MD5PasswordHasher) Algorithm() string {
	return "md5"
}

// Verify takes the raw password and the encoded one, then checks if they match.
func (p *MD5PasswordHasher) Verify(password string, encoded string) bool {
	results := strings.Split(encoded, HashSeparator)
	attempt := p.Encode(password, results[1])
	return encoded == attempt
}

// SafeSummary returns a summary of the encoded password.
func (p *MD5PasswordHasher) SafeSummary(encoded string) PasswordSummary {
	results := strings.Split(encoded, HashSeparator)
	return PasswordSummary{
		Algorithm: p.Algorithm(),
		Salt:      results[1],
		Hash:      results[2],
	}
}

// Salt returns the default salt (which defaults to a random 12 characters string).
func (p *MD5PasswordHasher) Salt() string {
	return RandomString(12)
}
