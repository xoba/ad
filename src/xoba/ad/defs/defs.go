package defs

import (
	"bytes"
	"os"
	"os/exec"
)

func Gofmt(p string) error {
	cmd := exec.Command("gofmt", "-w", p)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func GofmtBuffer(code []byte) ([]byte, error) {
	out := new(bytes.Buffer)
	cmd := exec.Command("gofmt")
	cmd.Stdin = bytes.NewReader(code)
	cmd.Stdout = out
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
