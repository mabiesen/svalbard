package util

import (
    "fmt"
    "log"
    "github.com/boltdb/bolt"
)
//databucket used for insertion and recall of database data
		var databucket = []byte("immabucket")

// open the database
//fp is the absolute filepath of the database
func OpenDB(fp string) (*bolt.DB) {

    db, err := bolt.Open(fp, 0644, nil)
    if err != nil {
        log.Fatal(err)
    }

		return db
}

// store some data
func UpdateDB(fp string, key []byte, value []byte){
	db := OpenDB(fp)
	err := db.Update(func(tx *bolt.Tx) error {
			bucket, err := tx.CreateBucketIfNotExists(databucket)
			if err != nil {
					return err
			}

			err = bucket.Put(key, value)
			if err != nil {
					return err
			}
			return nil
	})

	if err != nil {
			log.Fatal(err)
	}
	db.Close()
}

// retrieve some data
func RetrieveDB(fp string, key []byte) {
	db := OpenDB(fp)
	err := db.View(func(tx *bolt.Tx) error {
			bucket := tx.Bucket(databucket)
			if bucket == nil {
					return fmt.Errorf("Bucket %q not found!", databucket)
			}

			val := bucket.Get(key)
			fmt.Println(string(val))

			return nil
	})

	if err != nil {
			log.Fatal(err)
	}
	db.Close()
}
