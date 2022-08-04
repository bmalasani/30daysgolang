/*
Package log implements a simple logging package. It defines a type, Logger, with methods for formatting output. It also has a predefined 'standard' Logger accessible through helper functions Print[f|ln], Fatal[f|ln], and Panic[f|ln], which are easier to use than creating a Logger manually.
*/
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"meb.com/utils"
)

func main() {
	log.Println(log.Flags())

	if utils.Includes(os.Args, "panic") {
		log.Panicf("%s Paninc", os.Args[1])
	}

	if utils.Includes(os.Args, "output") {
		log.Println("test")
		log.Output(1, strings.Join(os.Args, "_"))
	}
	for {
		var option int
		fmt.Println(`Choose an option:
		1. set Flag
		2. panic
		3. fatal
		4. prefix
		5. use default logger
		`)
		fmt.Scanf("%d", &option)
		switch option {
		case 1:
			var flag int
			log.Printf("current flag %v, enter new one", log.Flags())
			fmt.Scanf("%d", &flag)
			log.SetFlags(flag)
			log.Printf("Flag set to %v", log.Flags())
			break
		case 2:
			log.Println("simulating panic")
			log.Panicln("Panic")
		case 3:
			log.Println("simulating fatal")
			log.Fatalln("Fatal")
		case 4:
			var prefix string
			log.Printf("current prefix %v, enter new one", log.Prefix())
			fmt.Scanf("%s_", &prefix)
			log.SetPrefix(prefix)
			log.Printf("prefix set to %v", log.Prefix())
		case 5:
			var (
				buff   bytes.Buffer
				logger *log.Logger
			)
			logger = log.New(&buff, "MEB_", 2)
			logger.Print("Hello from Logger")
			fmt.Print(string(buff.Bytes()))
			line, _ := buff.ReadString(byte('\n'))
			fmt.Printf("line: %v\n", line)

		default:
			go log.Panicf("You havent choosen any option.")
		}
	}

}
