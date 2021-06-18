package ineed

import (
	"io/ioutil"
	"log"
	"os"
)

//RenameFile renombra archivo origen y destino
func RenameFile(old, new string) (ok bool) {
	err := os.Rename(old, new)
	if err != nil {
		log.Printf("share error RenameFile %v", err)
		return
	}

	ok = true
	return
}

//CopyFILE copia archivo segun nombre y destino
func CopyFILE(src string, dst string) (ok bool) {
	// Read all content of src to data
	data, err := ioutil.ReadFile(src)
	if err != nil {
		log.Printf("share error CopyFILE %v", err)
		return
	}

	// Write data to dst
	err = ioutil.WriteFile(dst, data, 0644)
	if err != nil {
		log.Printf("share error CopyFILE %v", err)
		return
	}

	ok = true
	return
}
