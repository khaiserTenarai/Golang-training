package utility

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadString() string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func ReadInt() (int, error) {
	input := ReadString()
	return strconv.Atoi(input)
}

func ReadFloat() (float64, error) {
	input := ReadString()
	return strconv.ParseFloat(input, 64)
}
