package utils

import (
    "os"
)

type solve func([]string) string

func get_input(input_src string) ([]string) {
    if len(input_src) > 0 {
        input, err := LoadInput(input_src)
        Check(err)
        return input
    } else {
        input, err := ReadInput()
        Check(err)
        return input
    }
}

func Start(sol1 solve, sol2 solve) string {
    solution_part := os.Args[1]
    input_src := ""
    if len(os.Args) == 3 {
        input_src = os.Args[2]
    }
    input := get_input(input_src)
    if solution_part == "1" {
        return sol1(input)
    } else if solution_part == "2" {
        return sol2(input)
    }
    return ""
}
