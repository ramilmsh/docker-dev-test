package adder_test

import (
	"testing"

	"github.com/ramilmsh/docker-dev-test/internal/adder"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	require.Equal(t, adder.Add(1, 2), int64(3))
}
