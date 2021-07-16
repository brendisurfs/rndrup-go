// NEED FOR IMPLEMENTATION:
/* - THREAD HANDLER: for each chunk of files defined,
spawn a new thread, as the actual files dont need to be async.
- */

package main

import (
	"fmt"
	"rndr-go/cmd"
	"rndr-go/env"
)

func main() {
	// read the env of user to send to RNDR.
	envInfo := env.Env()
	email, password := envInfo[0], envInfo[1]
	fmt.Println(email, password)
	// execute command
	cmd.Execute()
	// open browser, execute upload.
	cmd.Send(email, password)
}
