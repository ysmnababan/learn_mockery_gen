package service

import (
	"context"
	"mockery/abstraction"
	"mockery/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetName(t *testing.T) {
	name := "yoland"
	ctx := context.Background()
	c, _ := ctx.(abstraction.Context)	
	mockrepo := new(mocks.AnyRepository)
	mockrepo.On("GetName", &c, 1).Return(name, nil)

	out, err := mockrepo.GetName(&c, 1)
	assert.NoError(t, err)
	assert.Equal(t, name, out)

	mockrepo.AssertExpectations(t)
}
