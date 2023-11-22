package belajargolangviper_test

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var config *viper.Viper = viper.New()

func TestViper(t *testing.T) {
	assert.NotNil(t, config)
}
