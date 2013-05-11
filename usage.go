package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Println(`Usage: get [-v] [-h] [-d] [<path>]

    -v, --version                   Prints the version and exits.
    -h, --help                      Prints the usage information.
    -d, --debug                     Logs debugging information to STDOUT.

Arguments:

    path                            The path to place or update the
                                    repostories. Defaults to the path
                                    in ~/.getconfig. This is required
                                    the first time you run get.

To learn more or to contribute, please see github.com/pearkes/get`)
	os.Exit(1)
}
