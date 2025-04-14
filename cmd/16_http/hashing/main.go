package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

/*
	Hashing
	- Hashing is a technique or process of mapping keys, values into the hash table by using a hash function.
	- It is used to unique identify a value.
	- Hashing is used mostly in password storage.
	- bcrypt hashing algorithm is used to store the password in the database.

*/

func main() {
	fmt.Println("Hashing")

	secret := "secret"

	hash, err := hashSecret("password")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Secret", secret)
	fmt.Println("Hash", hash)

	// Compare hash

	is_correct := isSecretHashCorrect("password", hash)
	fmt.Println("Is correct", is_correct)

}

func hashSecret(secret string) (string, error) {
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	// cost is the cost of hashing, the higher the cost the more secure the hash
	bytes, err := bcrypt.GenerateFromPassword([]byte(secret), 14)

	return string(bytes), err
}

func isSecretHashCorrect(secret, hash string) bool {
	// CompareHashAndPassword returns an error to tell us the secret was incorrect
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(secret))
	return err == nil
}
