package helper

import "log"

func CheckIfError(err error) error {
	if err != nil {
		return err // tujuannya ngasih errror tpi buatan sendiri yang ada di method atau function
	}

	return err
}

func PanicIfError(err error) {
	if err != nil {
		log.Fatalln(err) // langsung di berentiin
	}
}
