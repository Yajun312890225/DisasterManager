package utils

import "github.com/bwmarrin/snowflake"

var node *snowflake.Node

func GetID() string {
	return node.Generate().String()
}
func GetIntID() int64 {
	return node.Generate().Int64()
}

func init() {
	node, _ = snowflake.NewNode(1)
}
