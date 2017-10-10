package audit

func SetAuditer(a Auditer) {
	_auditer = a
}

func Log(event *Event) error {
	return _auditer.Log(event)
}
