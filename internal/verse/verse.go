package verse

import (
	"bible-verse-generator/internal/object"
	"bible-verse-generator/internal/util"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

var ErrNothingToDelete = errors.New("nothing to delete")

type Repo interface {
	GetAllVerses() ([]*Verse, error)
	GetVerseById(id int) (*Verse, error)
	GetRandomVerse() (*Verse, error)
	StoreNewVerse(req *Verse) (int, error)
	DeleteVerseById(id int) error
}

type DBRepo struct {
	db *sql.DB
}

func NewDBRepo(db *sql.DB) Repo {
	return &DBRepo{
		db: db,
	}
}

func (me *DBRepo) GetAllVerses() ([]*Verse, error) {
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

func (me *DBRepo) GetVerseById(id int) (*Verse, error) {
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

func (me *DBRepo) GetRandomVerse() (*Verse, error) {
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

func (me *DBRepo) StoreNewVerse(req *Verse) (int, error) {
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

func (me *DBRepo) DeleteVerseById(id int) error {
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
		return ErrNothingToDelete
	}

	return nil
}

type Controller struct {
	repo   Repo
	Routes []object.Route
}

func NewController(repo Repo) *Controller {
	c := &Controller{
		repo: repo,
	}
	c.setupRoutes()
	return c
}

func (me *Controller) HandleGetAllVerses(w http.ResponseWriter, _ *http.Request) {
	verses, err := me.repo.GetAllVerses()
	if err != nil {
		log.Println(err.Error())
		util.SendErrorResponse(w, http.StatusInternalServerError, "")
		return
	}
	json.NewEncoder(w).Encode(verses)
}

func (me *Controller) HandleGetVerseById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println(err.Error())
		util.SendErrorResponse(w, http.StatusBadRequest, "error getting id")
		return
	}
	verse, err := me.repo.GetVerseById(id)
	if err != nil {
		log.Println(err.Error())
		util.SendErrorResponse(w, http.StatusInternalServerError, "")
		return
	}
	json.NewEncoder(w).Encode(verse)
}

func (me *Controller) HandleGetRandomVerse(w http.ResponseWriter, _ *http.Request) {
	verse, err := me.repo.GetRandomVerse()
	if err != nil {
		log.Println(err.Error())
		util.SendErrorResponse(w, http.StatusInternalServerError, "")
		return
	}
	json.NewEncoder(w).Encode(verse)
}

func (me *Controller) HandleNewVerse(w http.ResponseWriter, r *http.Request) {
	var req Verse
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println(err.Error())
		util.SendErrorResponse(w, http.StatusBadRequest, "error decode json")
		return
	}

	newId, err := me.repo.StoreNewVerse(&req)
	if err != nil {
		log.Println(err.Error())
		util.SendErrorResponse(w, http.StatusInternalServerError, "")
		return
	}

	res := Verse{
		Id: newId,
	}
	json.NewEncoder(w).Encode(res)
}

func (me *Controller) HandleDeleteVerseById(w http.ResponseWriter, r *http.Request) {
	var id int
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println(err.Error())
		util.SendErrorResponse(w, http.StatusBadRequest, "error getting id")
		return
	}
	err = me.repo.DeleteVerseById(id)
	if err == ErrNothingToDelete {
		log.Println(err.Error())
		util.SendErrorResponse(w, http.StatusBadRequest, "nothing to delete")
		return
	}
	if err != nil {
		log.Println(err.Error())
		util.SendErrorResponse(w, http.StatusInternalServerError, "error deleting")
		return
	}

	util.SendGenericResponse(w)
}
