package audit

import "time"

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
	State     uint       `db:"state"`
	CreatedAt *time.Time `db:"created_at"`
}

type Auditer interface {
	Log(event *Event) error
}

type DefaultAudit struct {
}

func newDefaultAudit() *DefaultAudit {
	return &DefaultAudit{}
}

func (da *DefaultAudit) Log(event *Event) error {
	return nil
}
