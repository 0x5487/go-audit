package mysql

import (
	"database/sql"
	"time"

	audit "github.com/jasonsoft/go-audit"
	"github.com/jmoiron/sqlx"
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
