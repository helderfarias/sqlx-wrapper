package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldBeCreated(t *testing.T) {
	uw := NewUnitOfWork(nil, nil)

	assert.NotNil(t, uw)
}
