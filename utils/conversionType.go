package utils

import "database/sql"

func ToNullInt64(int64Value int64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: int64Value,
		Valid: true,
	}
}
