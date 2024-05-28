package infra_util

import "database/sql"

func MapOrderedFieldsFromRow(fieldsToGet []string, row *sql.Row) (map[string]string, error) {
	inputs := make([]*string, len(fieldsToGet))
	err := row.Scan(inputs)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	var fields = map[string]string{}
	for i, v := range fieldsToGet {
		fields[v] = *inputs[i]
	}

	return fields, nil
}
