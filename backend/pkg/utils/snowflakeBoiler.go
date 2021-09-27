package utils

import (
	"sync"

	"github.com/bwmarrin/snowflake"
	"github.com/pkg/errors"
)

type safeSnowflake struct {
	mu sync.Mutex
	n  int64
}

var snowflakeNodeNumber = safeSnowflake{
	n: 0,
}

func GetNewSnowflakeNode() (*snowflake.Node, error) {
	snowflakeNodeNumber.mu.Lock()
	nodeNumber := snowflakeNodeNumber.n
	snowflakeNodeNumber.n++
	snowflakeNodeNumber.mu.Unlock()
	node, err := snowflake.NewNode(nodeNumber)
	if err != nil {
		return nil, errors.Wrap(err, "new snowflake node")
	}
	return node, nil
}
