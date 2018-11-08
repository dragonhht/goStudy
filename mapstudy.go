package main

import "fmt"

func main() {
	map2 := map[string]string{"name": "dragonhht", "age": "18"}
	map1 := make(map[string]string)

	map1["hello"] = "world"
	map1["name"] = "dragonhht"

	delete(map1, "name")

	for key, value := range map1 {
		fmt.Printf("key is %s and value is %s\n", key, value)
	}
	for key, value := range map2 {
		fmt.Printf("key is %s and value is %s\n", key, value)
	}
}
