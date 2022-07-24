package main

import (
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/files"
	"log"
	"os"
	"strconv"
	"time"
)

const envAuthKeysPathname = "AUTH_KEYS_PATHNAME"
const keysNotGeneratedErrorMessage = "Keys was not generated, please specify the amount of keys to generate (go run uuid 1000) \n"

func main() {
	var (
		startTime        = time.Now()
		toGetAllArgs     = os.Args[1:]
		authKeysPathname = os.Getenv(envAuthKeysPathname)
	)

	if len(toGetAllArgs) > 0 {
		number, _ := strconv.Atoi(toGetAllArgs[0])

		errFindFile := files.FindFile(authKeysPathname)
		if errFindFile == nil {
			err := files.RemoveFile(authKeysPathname)
			if err != nil {
				log.Fatal(err)
			}
		}

		err := files.UUIDGenerator(number)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Successfully generated %d uuid keys \n", number)
		seconds := int(time.Since(startTime) / time.Second)
		fmt.Printf("To generate %s keys it took %d Seconds \n", toGetAllArgs[0], seconds)
	} else {
		fmt.Printf(keysNotGeneratedErrorMessage)
	}

}
