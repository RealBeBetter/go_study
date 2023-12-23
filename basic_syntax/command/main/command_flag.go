/**
 * @author Real
 * @since 2023/12/24 1:07
 */
package main

import (
	"flag"
	"fmt"
	"os"
)

var name = flag.String("name", "everyone", "The greeting object.")

func main() {
	flag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	flag.Parse()

	fmt.Println("Hello,", *name)
}
