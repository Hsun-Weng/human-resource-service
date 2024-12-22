package redis_keys

import (
	"fmt"
	"strconv"
)

const (
	LoginEmployeePrefix = "login:employee:"
)

func GetLoginEmployeeKey(id uint) string {
	return fmt.Sprintf("%s%s", LoginEmployeePrefix, strconv.FormatUint(uint64(id), 10))
}
