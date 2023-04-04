package main

import (
	"log"

	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {

		log.Fatal("Error al intentar cargar archivo .env")

	}

	usuario := os.Getenv("MY_USER")

	password := os.Getenv("MY_PASS")

	println("Usuario sacado de variables de Entorno: ", usuario)

	println("Password sacado de variables de Entorno: ", password)

}
