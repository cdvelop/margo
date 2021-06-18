package ineed

import (
	"log"
	"os"
)

//DirectoryExist EstaDirectorio verifica si existe un directorio retorna true o false
func DirectoryExist(dir string) (r bool) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Printf(">>> directorio %v no encontrado", dir)
	} else {
		log.Printf(">>> directorio %v ok", dir)
		r = true
	}
	return
}
