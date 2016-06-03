package g

import (
	"log"
	"runtime"
)

const (
	FALCON_SENDER_VERSION = "0.0.0"
	DOMEOS_VERSION = "0.2"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
