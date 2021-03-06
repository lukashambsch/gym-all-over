package datastore

import (
	"fmt"

	"github.com/lukashambsch/anygym.api/models"
	"github.com/lukashambsch/anygym.api/store"
)

func GetSupportSourceList(where string) ([]models.SupportSource, error) {
	var (
		supportSources []models.SupportSource
		supportSource  models.SupportSource
	)

	query := fmt.Sprintf("%s %s", getSupportSourceListQuery, where)
	rows, err := store.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&supportSource.SupportSourceID, &supportSource.SupportSourceName)
		supportSources = append(supportSources, supportSource)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()

	return supportSources, nil
}

func GetSupportSourceCount(where string) (*int, error) {
	var count int

	query := fmt.Sprintf("%s %s", getSupportSourceCountQuery, where)
	row := store.DB.QueryRow(query)
	err := row.Scan(&count)
	if err != nil {
		return nil, err
	}

	return &count, nil
}

func GetSupportSource(supportSourceID int64) (*models.SupportSource, error) {
	var supportSource models.SupportSource

	row := store.DB.QueryRow(getSupportSourceQuery, supportSourceID)
	err := row.Scan(&supportSource.SupportSourceID, &supportSource.SupportSourceName)
	if err != nil {
		return nil, err
	}

	return &supportSource, nil
}

func CreateSupportSource(supportSource models.SupportSource) (*models.SupportSource, error) {
	var created models.SupportSource

	row := store.DB.QueryRow(createSupportSourceQuery, supportSource.SupportSourceName)
	err := row.Scan(&created.SupportSourceID, &created.SupportSourceName)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func UpdateSupportSource(supportSourceID int64, supportSource models.SupportSource) (*models.SupportSource, error) {
	var updated models.SupportSource

	row := store.DB.QueryRow(updateSupportSourceQuery, supportSource.SupportSourceName, supportSourceID)
	err := row.Scan(&updated.SupportSourceID, &updated.SupportSourceName)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func DeleteSupportSource(supportSourceID int64) error {
	stmt, err := store.DB.Prepare(deleteSupportSourceQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(supportSourceID)
	if err != nil {
		return err
	}

	return nil
}

const getSupportSourceListQuery = `
SELECT *
FROM support_sources
`

const getSupportSourceQuery = `
SELECT *
FROM support_sources
WHERE support_source_id = $1
`

const createSupportSourceQuery = `
INSERT INTO support_sources (support_source_name)
VALUES ($1)
RETURNING support_source_id, support_source_name
`

const updateSupportSourceQuery = `
UPDATE support_sources
SET support_source_name = $1
WHERE support_source_id = $2
RETURNING support_source_id, support_source_name
`

const deleteSupportSourceQuery = `
DELETE
FROM support_sources
WHERE support_source_id = $1
`

const getSupportSourceCountQuery = `
SELECT count(*)
FROM support_sources
`
