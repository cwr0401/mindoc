package elasticsearch

import (
	es "github.com/elastic/go-elasticsearch/v7"
)

var client *es.Client

// 定义 es 操作

func Info() {
	client.Info()
}

// 初始化 es client
func Init(c *es.Client) {
	client = c
}
