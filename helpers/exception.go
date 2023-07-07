package helpers

import "log"

func PanicIfError(err interface{}) {
	if err != nil {
		log.Panic(err)
	}
}
