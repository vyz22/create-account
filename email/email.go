package email

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateEmail() string {
	rand.Seed(time.Now().UnixNano())

	name := fmt.Sprintf("user%d", rand.Intn(100000))
	domain := "zacky.my.id"

	return name + "@" + domain
}
