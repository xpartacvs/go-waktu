package waktu

import "time"

const (
	MySQLDate     = "2006-01-02"
	MySQLTime     = "15:04:05"
	MySQLDateTime = "2006-01-02 15:04:05"
)

var (
	WIB  *time.Location = time.FixedZone("WIB", 60*60*7)
	WITA *time.Location = time.FixedZone("WITA", 60*60*8)
	WIT  *time.Location = time.FixedZone("WIT", 60*60*9)
)
