package hashers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomString(t *testing.T) {
	result := RandomString(12)

	assert.Equal(t, len(result), 12)
}

func TestMaskHash(t *testing.T) {
	hash := "abcdefghijklmnopqrstuvwxyz"

	masked := MaskHash(hash, 5, "*")

	assert.Equal(t, masked, "abcde*********************")
	assert.Equal(t, len(masked), len(hash))
}

func TestConstantTimeCompare(t *testing.T) {
	assert.True(t, ConstantTimeCompare("val1", "val1"))
	assert.False(t, ConstantTimeCompare("val1", "val2"))
	assert.False(t, ConstantTimeCompare("val1", "value"))
}
