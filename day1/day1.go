package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error open:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	nums := make([]string, 0)

	for scanner.Scan() {
		linea := scanner.Text()

		if linea == "" {
			n := strings.Replace(linea, "", ",", -1)
			nums = append(nums, n)
		} else {
			nums = append(nums, linea)
		}
	}
	fmt.Println(nums)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo:", err)
	}
}
