package audit

import (
	"time"
)

var (
	_auditer Auditer
)

func init() {
	_auditer = newDefaultAudit()
}

const (
	FAILED  = 0
	SUCCESS = 1
)

type Event struct {
	Namespace string     `db:"namespace"`
	TargetID  string     `db:"target_id"`
	Action    string     `db:"action"`
	Actor     string     `db:"actor"`
	Message   string     `db:"message"`
	State     int        `db:"state"`
	ClientIP  string     `db:"client_ip"`
	CreatedAt *time.Time `db:"created_at"`
}

type ReadLogOption struct {
	Namespace string
	TargetID  string
	Action    string
	Actor     string
	State     int
	StartTime *time.Time
	EndTime   *time.Time
	Limit     int
	Offset    int
}

type Auditer interface {
	Log(event *Event) error
	ReadLog(option *ReadLogOption) ([]*Event, error)
	TotalCount(option *ReadLogOption) (int, error)
}

type DefaultAudit struct {
}

func newDefaultAudit() *DefaultAudit {
	return &DefaultAudit{}
}

func NewReadLogOption() *ReadLogOption {
	option := ReadLogOption{}
	option.State = -1
	return &option
}

func (da *DefaultAudit) Log(event *Event) error {
	return nil
}

func (da *DefaultAudit) ReadLog(option *ReadLogOption) ([]*Event, error) {
	return nil, nil
}

func (da *DefaultAudit) TotalCount(option *ReadLogOption) (int, error) {
	return 0, nil
}
