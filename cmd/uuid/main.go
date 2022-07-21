package main

import (
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/files"
	"os"
	"strconv"
	"time"
)

const envAuthKeysPathname = "AUTH_KEYS_PATHNAME"
const keysNotGeneratedErrorMessage = "Keys was not generated, please specify the amount of keys to generate (go run uuid 1000) \n"

func main() {
	var (
		err              error
		startTime        = time.Now()
		toGetAllArgs     = os.Args[1:]
		authKeysPathname = os.Getenv(envAuthKeysPathname)
	)

	if len(toGetAllArgs) > 0 {
		number, _ := strconv.Atoi(toGetAllArgs[0])

		_, err = os.Stat(authKeysPathname)
		if err == nil {
			files.RemoveFile(authKeysPathname)
		}

		files.UUIDGenerator(number)

		fmt.Printf("Successfully generated %d uuid keys in %s \n", number, authKeysPathname)
		seconds := int(time.Since(startTime) / time.Second)
		fmt.Printf("To generate %s keys it took %d Seconds \n", toGetAllArgs[0], seconds)
	} else {
		fmt.Printf(keysNotGeneratedErrorMessage)
	}

}
