package utils

import (
	"bufio"
	"os"
    "fmt"
    "strings"
)


func ReadInput() ([]string, error) {
	reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter input: ")
    text, err := reader.ReadString('\n')
    var lines []string
    lines = append(lines, strings.Replace(text, "\n", "", -1))
    return lines, err
}

func LoadInput(filename string) ([]string, error) {
    path := "/home/frere4/go/src/github.com/jn7dt/aoc/2019/inputs/" + filename
    // content, err := ioutil.ReadFile(path)
    file, err := os.Open(path)
    Check(err)
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}
