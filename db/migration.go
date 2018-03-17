package db

import (
	"fmt"
)

func Setup() {
	fmt.Println("Database prepare")
	fmt.Println("Database ready")
}

func Remove() {
	fmt.Println("Database removing")
	fmt.Println("Database removed")
}
