package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	
	var lin []string
	for scanner.Scan() {
		lin = append(lin, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	may := 0
	men := 0
	for _, line := range lin {
		if line == "" {
			if men > may {
				may = men
			}
			men = 0
		} else {
			num, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("Error \n", line, err)
				return
			}
			men += num
		}
	}
	fmt.Println(may)
}
