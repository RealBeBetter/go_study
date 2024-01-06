/**
 * @author Real
 * @since 2023/10/28 16:05
 */
package main

//在一个空文件夹下，初始化一个 Module
//$ go mod init example
//go: creating new go.mod: module example

// 使用 go mod tidy 命令，自动添加依赖
//$ go mod tidy
//go: finding module for package github.com/gin_framework-gonic/gin_framework

// 使用 go mod download 命令，下载依赖
//$ go mod download
//go: finding module for package github.com/gin_framework-gonic/gin_framework

//在 package main 中如何使用 package cal 中的 Add 函数呢
//import 模块名/子目录名 即可，修改后的 main 函数
