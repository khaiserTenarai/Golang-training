package view

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type UserView struct {
	reader *bufio.Reader
}

func NewUserView() *UserView {

	return &UserView{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (v *UserView) ShowMenu() {

	fmt.Println()
	fmt.Println("========== AUTHENTICATION ==========")
	fmt.Println("1. Register")
	fmt.Println("2. Login")
	fmt.Println("3. Exit")
	fmt.Println("====================================")
}

func (v *UserView) ReadString(
	message string,
) string {

	fmt.Print(message)

	input, _ := v.reader.ReadString('\n')

	return strings.TrimSpace(input)
}

func (v *UserView) ShowMessage(
	message string,
) {

	fmt.Println(message)
}
