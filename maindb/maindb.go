package maindb

import (
	"fmt"

	khg "cloud.google.com/go/cloudsqlconn/postgres/pgxv4"
)

func Check() {
	khg.RegisterDriver("ghitnj")
	fmt.Println("Hello")
}
