package sockutils

import "log"

func init() {
	log.SetFlags(log.Llongfile | log.Ldate)
}

func CheckErrorLog(err error) {
	if err != nil {
		log.Println(err)
	}
}

func CheckErrorFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
