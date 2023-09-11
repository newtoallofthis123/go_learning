package cli

import "os"

func Get_Args() []string {
	return os.Args[1:]
}
