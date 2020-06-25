package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"

	"github.com/GehirnInc/crypt"
	_ "github.com/GehirnInc/crypt/sha512_crypt"
)

func main() {

	salt := generate_salty_string(16)

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	crypt := crypt.SHA512.New()
	ret, err := crypt.Generate(input, []byte("$6$"+salt))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ret)

	if err := crypt.Verify(ret, input); err != nil {
		log.Fatal(err)
	}

	// Output:
	// $5$salt$kpa26zwgX83BPSR8d7w93OIXbFt/d3UOTZaAu5vsTM6
	// <nil>
}

func generate_salty_string(length int) string {
	saltines := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789./"
	out := ""
	for i := 0; i < length; i++ {
		vb, err := rand.Int(rand.Reader, big.NewInt(int64(len(saltines))))
		if err != nil {
			panic(err)
		}
		v := int(vb.Int64())
		out += saltines[v : v+1]
	}
	return out
}
