package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func main() {
	err := v4()
	fmt.Printf("main output, err: %+v", err)
}

func v1() error {
	_, err := os.Open("/Users/micah/Local/Tech/aaa.txt")
	if err != nil {
		return errors.Wrap(err, "this is v1 output")
	}
	return nil
}

func v2() error {
	err := v1()
	return err
}

func v3() error {
	err := v2()
	return err
}

func v4() error {
	err := v3()
	return err
}
