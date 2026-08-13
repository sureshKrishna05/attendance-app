package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	hash, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	fmt.Println(string(hash))
	
	err := bcrypt.CompareHashAndPassword([]byte("$2a$10$wE9q3427u0qR5u4h5d/j0Ou3wG/h83eT/8N8Hj9t76i80j0"), []byte("password123"))
	fmt.Println("Match:", err == nil)
}
