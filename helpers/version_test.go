package helpers_test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rickyananada1/go-starter-templating/helpers"
)

func TestGetGoVersion(t *testing.T) {
	goVer := helpers.GetGoVersion()
	expectedVersion, errNotMatch := regexp.MatchString("1.2*", goVer)

	assert.Nil(t, errNotMatch)
	assert.True(t, expectedVersion, "golang version not detected")

}
