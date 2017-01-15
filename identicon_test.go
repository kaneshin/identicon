package identicon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestData(t *testing.T) {

	t.Run("NewDataString", func(t *testing.T) {
		assert := assert.New(t)
		d := NewDataString("wentworth")
		assert.NotNil(d)
		assert.True(len(d.hash) > 0)
	})
}
