package dpfm_api_output_formatter

import (
	"database/sql"
	"fmt"
)

func ConvertToIncoterms(rows *sql.Rows) (*Incoterms, error) {
	defer rows.Close()
	incoterms := Incoterms{}
	i := 0

	for rows.Next() {
		i++
		err := rows.Scan(
			&incoterms.incoterms,
			&incoterms.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &incoterms, err
		}

	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &incoterms, nil
	}

	return &incoterms, nil
}
