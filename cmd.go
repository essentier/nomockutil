package nomockutil

import (
	"bytes"
	"log"
	"os/exec"

	"github.com/go-errors/errors"
)

func RunCmd(name string, args ...string) (*bytes.Buffer, *bytes.Buffer, error) {
	log.Printf("%v %v", name, args)
	cmd := exec.Command(name, args...)
	var outBuffer bytes.Buffer
	var errBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer
	err := cmd.Run()
	var e error = nil
	if err != nil {
		e = errors.Wrap(err, 1)
		log.Printf("Err output: %v %v", err, errBuffer.String())
	}
	return &outBuffer, &errBuffer, e
}
