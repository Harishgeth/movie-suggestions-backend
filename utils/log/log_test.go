package log

import (
	"testing"

	// "github.com/RealImage/jt-utils/testHelpers"
	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	l := NewLogger("")
	assert.NotEqual(t, l, nil)

	l.Debug("This is Debug log")
	l.Error("This is error log")
	l.Fatal("This is fatal log")
	l.Info("This is info log")
	l.Warn("This is warn log")
}
