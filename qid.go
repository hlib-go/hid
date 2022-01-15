package hid

import "time"

// Qid 请求编号32位
func Qid() string {
	return time.Now().Local().Format("20060102150405") + UUID32()[14:]
}
