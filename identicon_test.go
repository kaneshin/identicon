package identicon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestData(t *testing.T) {

	t.Run("NewData", func(t *testing.T) {
		assert := assert.New(t)
		d := NewData([]byte{'a'})
		assert.NotNil(d)
		assert.True(len(d.hash) > 0)
	})

	t.Run("NewDataString", func(t *testing.T) {
		assert := assert.New(t)
		d := NewDataString("a")
		assert.NotNil(d)
		assert.True(len(d.hash) > 0)
	})
}
