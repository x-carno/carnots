package db

import "github.com/x-carno/carnots/pkg/entity"

type DB interface {
	Store(metrics ...entity.Metrics)
	// Fetch(expr string) ([]byte, error)
}
