package infra_util

import "database/sql"

func MapOrderedFieldsFromRows(fieldsToGet []string, rows *sql.Rows) ([]map[string]string, error) {
	var result []map[string]string

	for rows.Next() {
		inputs := make([]*string, len(fieldsToGet))
		err := rows.Scan(inputs)
		if err != nil {
			return nil, err
		}

		var fields = map[string]string{}
		for i, v := range fieldsToGet {
			fields[v] = *inputs[i]
		}

		result = append(result, fields)
	}

	return result, nil
}
