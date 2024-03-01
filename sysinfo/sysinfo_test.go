package sysinfo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSysInfo(t *testing.T) {
	assert := assert.New(t)

	logs := Dump()
	assert.NotEmpty(logs)
}
