package main

import (
	"log"

	"github.com/casbin/casbin/v2"
)

func main() {
	// 初始化 Casbin Enforcer，使用 RBAC 模型和策略文件
	enforcer, err := casbin.NewEnforcer("rbac_model.conf", "rbac_policy.csv")
	if err != nil {
		log.Fatal("Failed to initialize Casbin Enforcer:", err)
	}

	// 添加角色
	enforcer.AddRoleForUser("alice", "admin")

	// 进行权限检查
	if enforcer.Enforce("alice", "data1", "read") {
		log.Println("alice has permission to read data1")
	} else {
		log.Println("alice does not have permission to read data1")
	}

	if enforcer.Enforce("bob", "data1", "read") {
		log.Println("bob has permission to read data1")
	} else {
		log.Println("bob does not have permission to read data1")
	}
}
