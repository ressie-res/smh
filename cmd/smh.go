package cmd

import (
	"fmt"
	"os/exec"
	"strings"
)

func NewPassword(length string) (result string) {
	funkystuff := exec.Command("openssl", "rand", "-base64", length)
	output, _ := funkystuff.Output()
	return fmt.Sprintf("%s", strings.TrimSpace(string(output)))
}
