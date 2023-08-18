package main

import (
	"github.com/xuanvan229/crawl-golang/pkg/config"
	"github.com/xuanvan229/crawl-golang/pkg/crawl"
	"github.com/xuanvan229/crawl-golang/pkg/model"
)

func main() {
	config.LoadEnv()
	model.InitDB()
	crawl.FixLazada()
}
