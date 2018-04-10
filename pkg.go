package audit

func SetAuditer(a Auditer) {
	_auditer = a
}

func Log(event *Event) error {
	return _auditer.Log(event)
}

func ReadLog(option *ReadLogOption) ([]*Event, error) {
	return _auditer.ReadLog(option)
}

func TotalCount(option *ReadLogOption) (int, error) {
	return _auditer.TotalCount(option)
}
