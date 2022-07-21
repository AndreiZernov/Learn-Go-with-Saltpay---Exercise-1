package main

import (
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/files"
	"os"
	"strconv"
)

const envAuthKeysPathname = "AUTH_KEYS_PATHNAME"
const keysNotGeneratedErrorMessage = "Keys was not generated, please specify the amount of keys to generate (go run uuid 1000) \n"

func main() {
	var (
		err              error
		toGetAllArgs     = os.Args[1:]
		authKeysPathname = os.Getenv(envAuthKeysPathname)
	)

	if len(toGetAllArgs) > 0 {
		number, _ := strconv.Atoi(toGetAllArgs[0])
		files.RemoveFile(authKeysPathname)
		for i := 0; i < number; i++ {
			err = files.UUIDGenerator()
		}
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("Successfully generated %d uuid keys in %s \n", number, authKeysPathname)
		}
	} else {
		fmt.Printf(keysNotGeneratedErrorMessage)
	}
}
