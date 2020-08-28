package utils

import "github.com/bwmarrin/snowflake"

var node *snowflake.Node

func GetID() string {
	return node.Generate().String()
}

func init() {
	node, _ = snowflake.NewNode(1)
}
