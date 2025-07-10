package service

import (
	"context"
	"mockery/abstraction"
	"mockery/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetName(t *testing.T) {
	name := "yoland"
	ctx := context.Background()
	c, _ := ctx.(abstraction.Context)
	mockrepo := repository.NewMockAnyRepository(t)
	mockrepo.EXPECT().GetName(&c, mock.Anything).Return(name, nil)

	out, err := mockrepo.GetName(&c, 1)
	assert.NoError(t, err)
	assert.Equal(t, name, out)

	// mockrepo.AssertExpectations(t)
}
