package persist

import "errors"
import "github.com/syndtr/goleveldb/leveldb"

const dbPath = "./db/keybase"

//Save something to db
func Save(key string, value string) error {
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		return err
	}

	err = db.Put([]byte(key), []byte(value), nil)
	if err != nil {
		return err
	}
	defer db.Close()
	return nil
}

//Delete something from db
func Delete() error {
	return errors.New("Not Implemented")
}

//Update db
func Update() error {
	return errors.New("Not Implemented")
}
