package main

import "fmt"

func main() {
	// 创建一个映射，存储颜色以及颜色对应的十六进制代码
	colors := map[string]string{
		"AliceBlue": "#f0f8ff",
		"Coral":     "#ff7F50",
		"DarkGray":  "#a9a9a9",
	}

	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}

	removeColor(colors, "Coral")

	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}
}

func removeColor(colors map[string]string, key string) {
	delete(colors, key)
}
