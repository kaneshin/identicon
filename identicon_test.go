package identicon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestData(t *testing.T) {

	t.Run("NewData", func(t *testing.T) {
		assert := assert.New(t)
		d := NewData("wentworth")
		assert.Equal("wentworth", d.id)
	})
}
