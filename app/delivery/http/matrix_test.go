package matrix_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Canary test
func TestCanarySpec(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(true, true, "Valide if infra for test work")
}
