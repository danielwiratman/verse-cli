package verse

import (
	"database/sql"
	"time"
	"verse-cli/internal/util"
)

type Repo interface {
	getAllVerses() ([]*Verse, error)
	getVerseById(id int) (*Verse, error)
	getRandomVerse() (*Verse, error)
	storeNewVerse(req *Verse) (int, error)
	deleteVerseById(id int) error
}

type DBRepo struct {
	db *sql.DB
}

func NewDBRepo(db *sql.DB) Repo {
	return &DBRepo{
		db: db,
	}
}

func (me *DBRepo) getAllVerses() ([]*Verse, error) {
	verses := []*Verse{}
	rows, err := me.db.Query(
		"SELECT id, address, content, created FROM verses",
	)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		verse := &Verse{}
		var createdStr string
		err := rows.Scan(
			&verse.Id,
			&verse.Address,
			&verse.Content,
			&createdStr,
		)
		if err != nil {
			return nil, err
		}
		created, _ := time.Parse(time.RFC3339Nano, createdStr)
		verse.Created = created
		verses = append(verses, verse)
	}
	return verses, nil
}

func (me *DBRepo) getVerseById(id int) (*Verse, error) {
	verse := &Verse{}
	var createdStr string
	err := me.db.QueryRow(
		"SELECT id, address, content, created FROM verses WHERE id = $1",
		id,
	).Scan(
		&verse.Id,
		&verse.Address,
		&verse.Content,
		&createdStr,
	)
	if err != nil {
		return nil, err
	}
	created, _ := time.Parse(time.RFC3339Nano, createdStr)
	verse.Created = created

	return verse, nil
}

func (me *DBRepo) getRandomVerse() (*Verse, error) {
	verse := &Verse{}
	var createdStr string
	err := me.db.QueryRow(
		"SELECT id, address, content, created FROM verses ORDER BY RANDOM() LIMIT 1",
	).Scan(
		&verse.Id,
		&verse.Address,
		&verse.Content,
		&createdStr,
	)
	if err != nil {
		return nil, err
	}
	created, _ := time.Parse(time.RFC3339Nano, createdStr)
	verse.Created = created

	return verse, nil
}

func (me *DBRepo) storeNewVerse(req *Verse) (int, error) {
	var id int
	err := me.db.QueryRow(
		"INSERT INTO verses (address, content) VALUES ($1, $2) RETURNING id",
		req.Address,
		req.Content,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (me *DBRepo) deleteVerseById(id int) error {
	res, err := me.db.Exec(
		"DELETE FROM verses WHERE id = $1",
		id,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return util.ErrNothingToDelete
	}

	return nil
}
