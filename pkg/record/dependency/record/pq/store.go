package pq

import (
	"github.com/sirupsen/logrus"
	"github.com/skygeario/skygear-server/pkg/core/auth/role"
	"github.com/skygeario/skygear-server/pkg/core/db"
)

type RecordStore struct {
	roleStore role.Store

	canMigrate bool

	sqlBuilder  db.SQLBuilder
	sqlExecutor db.SQLExecutor
	logger      *logrus.Entry
}

func NewRecordStore(
	roleStore role.Store,
	canMigrate bool,
	sqlBuilder db.SQLBuilder,
	sqlExecutor db.SQLExecutor,
	logger *logrus.Entry,
) *RecordStore {
	return &RecordStore{
		roleStore:   roleStore,
		canMigrate:  canMigrate,
		sqlBuilder:  sqlBuilder,
		sqlExecutor: sqlExecutor,
		logger:      logger,
	}
}
