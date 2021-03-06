package pygments

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

var (
	bin = "pygmentize"
)

func Binary(path string) {
	bin = path
	return
}

func Highlight(code string, lexer string, format string, enc string) string {

	if _, err := exec.LookPath(bin); err != nil {
		log.Fatal("You do not have " + bin + " installed!")
	}

	cmd := exec.Command(bin, "-l "+lexer, "-f "+format, "-O encoding="+enc)
	cmd.Stdin = strings.NewReader(code)

	var out bytes.Buffer
	cmd.Stdout = &out

	if _, err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	return out.String()
}
