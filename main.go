package main

import (
	"fmt"

	c "github.com/raelnogpires/go-algorithms/challenges"
)

func main() {
	ag_01_param := [][]int{{1, 3}, {2, 4}, {3, 5}, {1, 6}}
	ag_01_result := c.Study_schedule(ag_01_param, 2)
	fmt.Println("study_schedule -> ", ag_01_result)
}
