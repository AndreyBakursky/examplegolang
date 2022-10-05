package example_for_review

import (
	"context"
	"fmt"
	"strings"
)

func (r *PersistRepository) UpdateById(ctx context.Context, id int64, fields map[string]interface{}) error {
	sess := r.conn.Session(
		db.WithCtx(ctx),
		db.WithTimeout(r.queryingTimeout),
		db.WithHookAck(r.metrics.QueryDone),
		db.WithHookNack(r.metrics.QueryError),
	)

	args := map[string]interface{}{
		"id": id,
	}

	updateSet := make([]string, len(fields))
	i := 0

	for field, fieldValue := range fields {
		args[field] = fieldValue
		updateSet[i] = fmt.Sprintf("%s = :%s", field, field)
		i++
	}

	// language=PostgreSQL
	sqlString := `UPDATE "order" SET %s WHERE id = :id`

	sql := fmt.Sprintf(sqlString, strings.Join(updateSet, ","))

	return sess.Write(sql, args)
}
