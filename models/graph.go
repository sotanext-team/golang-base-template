package models

import "app-api/ent"

type GraphPagination struct {
	After  *ent.Cursor
	First  *int
	Before *ent.Cursor
	Last   *int
}
