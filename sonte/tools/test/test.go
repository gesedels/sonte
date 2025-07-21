// Package test implements unit test data and functions.
package test

import (
	"path/filepath"
	"testing"

	"go.etcd.io/bbolt"
)

// MockData is a map of mock database entries for unit testing.
var MockData = map[string]map[string]string{
	"alpha": {
		"body": "Alpha note.\n",
		"hash": "VwWTrgnGUGW8W6Jr_w-rrYqo3C3jS_aY7PU5fNE6Qo8",
		"time": "981122400", // 2001-02-03
	},

	"bravo": {
		"body": "Bravo note.\n",
		"hash": "fTueYUse24ZggEiMoa_y9Q3DWofKYH79ulEgpucIyq4",
		"time": "981208800", // 2001-02-04
	},
}

// MockDB returns a temporary database containing MockData.
func MockDB(t *testing.T) *bbolt.DB {
	dire := t.TempDir()
	dest := filepath.Join(dire, "test.db")
	db, err := bbolt.Open(dest, 0666, nil)
	if err != nil {
		panic(err)
	}

	err = db.Update(func(tx *bbolt.Tx) error {
		for name, pairs := range MockData {
			buck, err := tx.CreateBucket([]byte(name))
			if err != nil {
				return err
			}

			for attr, data := range pairs {
				if err := buck.Put([]byte(attr), []byte(data)); err != nil {
					return err
				}
			}
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	return db
}
