package main 

import (
    utils "github.com/jn7dt/aoc/2019/utils"
    "math"
    "strconv"
    "fmt"
)

func fuel_calc(mass float64) (float64) {
    fuel := math.Floor(mass / 3) - 2
    return fuel
}

func solution1(input []string) string {
    var total float64 = 0
    for _, value := range input {
        mass, _ := strconv.ParseFloat(value, 64)
        fuel := fuel_calc(mass)
        total += fuel
    }
    return strconv.FormatFloat(total, 'f', 0, 64)
}

func solution2(input []string) string {
    var total float64 = 0
    for _, value := range input {
        mass, _ := strconv.ParseFloat(value, 64)
        fuel := fuel_calc(mass)
        total += fuel
        for fuel > 0 {
            fuel = fuel_calc(fuel)
            if fuel > 0 {
                total += fuel
            }
        }
    }
    return strconv.FormatFloat(total, 'f', 0, 64)
}


func main() {
    solution := utils.Start(solution1, solution2)
    fmt.Println("Solution: " + solution)
}