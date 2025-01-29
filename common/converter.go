package common

import (
	"database/sql"
	"strings"
)

func StringToNullString(input *string) sql.NullString {
	if input == nil {
		return sql.NullString{Valid: false}
	} else {
		return sql.NullString{Valid: true, String: *input}
	}
}

func Int64ToNullInt64(input *int64) sql.NullInt64 {
	if input == nil {
		return sql.NullInt64{Valid: false}
	} else {
		return sql.NullInt64{Valid: true, Int64: int64(*input)}
	}
}

func JoinArrayToNullString(input []string) sql.NullString {
	if len(input) == 0 {
		return sql.NullString{Valid: false}
	} else {
		return sql.NullString{Valid: true, String: strings.Join(input, ",")}
	}
}
