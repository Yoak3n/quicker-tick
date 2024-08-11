package util

import "os"

func CreateDirNotExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		e := os.MkdirAll(dir, os.ModePerm)
		if e != nil {
			println("Error creating directory: " + e.Error())
			return e
		}
	}
	return nil
}
