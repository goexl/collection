package queue_test

import (
	"testing"

	"github.com/goexl/container/queue"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	_queue := queue.New[int]().Build()
	require.NotNil(t, _queue, "默认队列创建出错")
}
