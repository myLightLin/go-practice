// 序列化 JSON 字符串
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// 创建一个键值对映射
	c := make(map[string]interface{})
	c["name"] = "Gopher"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{}{
		"hone": "415.333.3333",
		"cell": "415.555.5555",
	}

	// 将这个映射序列化到 JSON 字符串
	data, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	fmt.Println(string(data))
}
