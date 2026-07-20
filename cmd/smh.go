package cmd

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func NewPassword(length string) (result string) {

	num1, err := strconv.Atoi(length)
	if err != nil {
		fmt.Print("Error: Input should be a whole number")
		return
	}

	funkystuff := exec.Command("openssl", "rand", "-base64", strings.TrimSpace(string(rune(num1))))
	output, _ := funkystuff.Output()
	return fmt.Sprintf("%s", strings.TrimSpace(string(output)))
}
