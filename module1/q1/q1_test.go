package q1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFixStrArray(t *testing.T) {
	s1 := FixStrArray()
	assert.Equal(t, []string{"I", "am", "smart", "and", "strong"}, s1)
}
