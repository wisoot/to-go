package item

import "database/sql"

func MakeManager(db *sql.DB) *Manager {
	manager := &Manager{
		gateway: &MysqlGateway{db},
	}

	return manager
}
