package mysql

import (
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	audit "github.com/jasonsoft/go-audit"
)

type MySqlAuditer struct {
	db *sqlx.DB
}

func NewMysqlAuditer(db *sql.DB) *MySqlAuditer {
	dbx := sqlx.NewDb(db, "mysql")
	return &MySqlAuditer{
		db: dbx,
	}
}

const insertAuditSQL = "INSERT INTO `audits` (`namespace`, `target_id`, `action`, `actor`, `message`, `state`, `created_at`) VALUES (:namespace, :target_id, :action, :actor, :message, :state, :created_at);"

func (ms *MySqlAuditer) Log(event *audit.Event) error {
	nowUTC := time.Now().UTC()
	event.CreatedAt = &nowUTC

	_, err := ms.db.NamedExec(insertAuditSQL, event)
	if err != nil {
		return err
	}
	return nil
}

const selectAuditSQL = `SELECT namespace, target_id, action, actor, message, state, created_at FROM audits
WHERE (created_at >= :start_time)
AND (created_at <= :end_time) `

const selectAuditNamespaceSQL = `AND namespace = :namespace `
const selectAuditTargeyIDSQL = `AND target_id = :target_id `
const selectAuditActionSQL = `AND action = :action `
const selectAuditActorSQL = `AND actor = :actor `
const selectAuditStateSQL = `AND state = :state `

const orderByCreatedSQL = `ORDER BY created_at DESC `
const selectAuditLimitSQL = `LIMIT :skip,:per_page; `

func (ms *MySqlAuditer) ReadLog(option *audit.ReadLogOption) ([]*audit.Event, error) {
	events := []*audit.Event{}
	sqlstring := selectAuditSQL
	if option.StartTime == nil {
		return nil, errors.New("start time can't be nil")
	}
	if option.EndTime == nil {
		return nil, errors.New("end time can't be nil")
	}
	if option.StartTime.After(*option.EndTime) {
		return nil, errors.New("end time can't be before than start time")
	}
	m := map[string]interface{}{
		"start_time": option.StartTime,
		"end_time":   option.EndTime,
		"skip":       option.Skip,
		"per_page":   option.PerPage,
	}
	if len(option.Namespace) > 0 {
		sqlstring += selectAuditNamespaceSQL
		m["namespace"] = option.Namespace
	}
	if len(option.TargetID) > 0 {
		sqlstring += selectAuditTargeyIDSQL
		m["target_id"] = option.TargetID
	}
	if len(option.Action) > 0 {
		sqlstring += selectAuditActionSQL
		m["action"] = option.Action
	}
	if len(option.Actor) > 0 {
		sqlstring += selectAuditActorSQL
		m["actor"] = option.Actor
	}
	if option.State >= 0 {
		sqlstring += selectAuditStateSQL
		m["state"] = option.State
	}

	sqlstring += orderByCreatedSQL
	sqlstring += selectAuditLimitSQL
	getAuditListStmt, err := ms.db.PrepareNamed(sqlstring)
	if err != nil {
		return nil, err
	}
	defer getAuditListStmt.Close()
	err = getAuditListStmt.Select(&events, m)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return events, nil
}

const countAuditSQL = `SELECT Count(1) FROM audits
WHERE (created_at >= :start_time)
AND (created_at <= :end_time) `

func (ms *MySqlAuditer) TotalCount(option *audit.ReadLogOption) (int, error) {
	sqlstring := countAuditSQL
	if option.StartTime == nil {
		return 0, errors.New("start time can't be nil")
	}
	if option.EndTime == nil {
		return 0, errors.New("end time can't be nil")
	}
	if option.StartTime.After(*option.EndTime) {
		return 0, errors.New("end time can't be before than start time")
	}
	m := map[string]interface{}{
		"start_time": option.StartTime,
		"end_time":   option.EndTime,
		"skip":       option.Skip,
		"per_page":   option.PerPage,
	}
	if len(option.Namespace) > 0 {
		sqlstring += selectAuditNamespaceSQL
		m["namespace"] = option.Namespace
	}
	if len(option.TargetID) > 0 {
		sqlstring += selectAuditTargeyIDSQL
		m["target_id"] = option.TargetID
	}
	if len(option.Action) > 0 {
		sqlstring += selectAuditActionSQL
		m["action"] = option.Action
	}
	if len(option.Actor) > 0 {
		sqlstring += selectAuditActorSQL
		m["actor"] = option.Actor
	}
	if option.State >= 0 {
		sqlstring += selectAuditStateSQL
		m["state"] = option.State
	}
	countAuditListStmt, err := ms.db.PrepareNamed(sqlstring)
	if err != nil {
		return 0, err
	}
	defer countAuditListStmt.Close()
	var count int
	err = countAuditListStmt.Get(&count, m)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return count, nil
}
