package log

// log const var
const (
	format  = "2006-01-02 15:04:05"
	formatD = "2006-01-02"
	formatH = "2006-01-02-15"
	formatT = "15:04:05"
	MB      = 1048576
)

type RotatePolicy string

const (
	ROTATE_POLICY_SIZE RotatePolicy = "SIZE"
	ROTATE_POLICY_HOUR RotatePolicy = "HOUR"
	ROTATE_POLICY_DAY  RotatePolicy = "DAY"
)

// NsqConfig 是nsq 消息队列的配置
type NsqConfig struct {
	LogName          string
	Profile          string
	NSQLookupdServer string
}
