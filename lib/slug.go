package lib

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func Slug(title string) string {
	outputStr := strings.ReplaceAll(strings.ToLower(title), " ", "-")

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNum := r.Intn(1000000)
	outputStr = fmt.Sprintf("%s-%d", outputStr, randomNum)

	return outputStr
}
