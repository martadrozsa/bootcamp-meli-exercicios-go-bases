package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main()  {
	r := strings.NewReader("some .ioReader stream to be read\n")
	_, err := io.Copy(os.Stdout, r)

	fmt.Println(err)
}
