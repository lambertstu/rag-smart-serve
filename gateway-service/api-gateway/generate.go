package main

//go:generate goctl api go --api gateway.api --dir . --home tpl

//go:generate go install github.com/zeromicro/go-zero/tools/goctl@latest

//go:generate goctl api swagger -api spec/product.api -dir swagger
