package db

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"time"

	"github.com/boltdb/bolt"
)

var noteBucket = []byte("notes")
var db *bolt.DB

type Description struct {
	Short string //short
	Long  string
}
type Note struct {
	Key   int
	Value Description
}

func Init(path string) error {
	var err error
	db, err = bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(noteBucket)
		return err
	})
}

func CreateNote(desc Description) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(noteBucket)
		id64, _ := bucket.NextSequence()
		id = int(id64)
		key := itob(id)
		note, err := json.Marshal(desc)
		if err != nil {
			return err
		}
		return bucket.Put(key, []byte(note))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func GetNoteList() ([]Note, error) {
	var notes []Note
	err := db.View(func(tx *bolt.Tx) error {

		bucket := tx.Bucket(noteBucket)
		c := bucket.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var desc Description
			err := json.Unmarshal(v, &desc)
			if err != nil {
				return err
			}
			notes = append(notes, Note{
				Key:   btoi(k),
				Value: desc,
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return notes, nil
}

func GetNoteByKey(key int) (Note, error) {
	var note Note
	err := db.View(func(tx *bolt.Tx) error {

		bucket := tx.Bucket(noteBucket)
		val := bucket.Get(itob(key))
		if val == nil {
			return errors.New("value not found")
		}
		var desc Description
		err := json.Unmarshal(val, &desc)
		if err != nil {
			return err
		}
		note = Note{Key: key, Value: desc}
		return nil
	})
	if err != nil {
		return Note{}, err
	}
	return note, nil
}
func RemoveNote(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(noteBucket)
		return bucket.Delete(itob(key))
	})
}
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

func Exit() error {
	return db.Close()
}
