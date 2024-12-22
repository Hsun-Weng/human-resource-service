package requests

import (
	"fmt"
	"github.com/Hsun-Weng/human-resource-service/pkg/util"
	"time"
)

type CustomDate time.Time

func (d *CustomDate) UnmarshalJSON(b []byte) error {
	str := string(b)
	// 去掉两侧的引号
	str = str[1 : len(str)-1]
	parsedTime, err := util.ParseDate(str)
	if err != nil {
		return fmt.Errorf("invalid date format, expected yyyy-MM-dd")
	}
	*d = CustomDate(parsedTime)
	return nil
}
