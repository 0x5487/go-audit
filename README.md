# go-audit
simple audit log package for golang


### add log
```go
    //init
    // setup auditer
	mysqlAuditer := mysql.NewMysqlAuditer(db)
	audit.SetAuditer(mysqlAuditer)
```

```go
    //init event
    event := &audit.Event{
		Namespace: namespace,
		TargetID:  targetid,
		Actor:     actor,
		Action:    action,
	}
    //save log
	audit.Log(event)
```

### read log
```go
    //init option
    Option := audit.NewReadLogOption()
    //set custom option
	Option.Action = "action"
	Option.Actor = "actor"
	Option.Namespace = "namespace"
	Option.Skip = 0
	Option.PerPage = 20
    starttime := time.Date(2017, 10, 10, 0, 0, 0, 0, time.Local)
	Option.StartTime = &starttime
	endtime := time.Date(2017, 10, 17, 0, 0, 0, 0, time.Local)
	Option.EndTime = &endtime
	result, _ := audit.ReadLog(Option)
	for _, row := range result {
		fmt.Println("%v", row)
	}

```
