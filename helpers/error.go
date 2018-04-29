package helpers

import "log"

//HandleBootstrapError HandleBootstrapError
func HandleBootstrapError(err error) {
	if err != nil {
		log.Fatal("SOMETHING WENT WRONG.", err)
	}
}
