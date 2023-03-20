package lib

import (
	"fmt"
	"os"
)

func ImageReturn(origin string) string {
	return fmt.Sprintf("%s:%s%s%s", os.Getenv("BASE_URL"), os.Getenv("CLIENT_PORT"), "/public/image", origin)
}
