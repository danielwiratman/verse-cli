package verse

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"verse-cli/internal/util"

	"github.com/go-chi/chi/v5"
)

type Controller struct {
	repo Repo
}

func NewController(repo Repo) *Controller {
	return &Controller{
		repo: repo,
	}
}

func (me *Controller) Route(r chi.Router) {
	r.Get("/verses", me.handleGetAllVerses)
	r.Get("/verse/{id}", me.handleGetVerseById)
	r.Get("/verse/random", me.handleGetRandomVerse)
	r.Post("/verse", me.handleNewVerse)
	r.Delete("/verse/{id}", me.handleDeleteVerseById)
}

func (me *Controller) handleGetAllVerses(w http.ResponseWriter, _ *http.Request) {
	verses, err := me.repo.getAllVerses()
	if err != nil {
		log.Println(err.Error())
		util.SendErrorResponse(w, http.StatusInternalServerError, "error getting all verse")
		return
	}
	util.OKResponseWithPage(w, verses, util.PageInfo{
		Total: len(verses),
		Page:  1,
	})
}

func (me *Controller) handleGetVerseById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println(err.Error())
		util.SendErrorResponse(w, http.StatusBadRequest, "error getting id")
		return
	}
	verse, err := me.repo.getVerseById(id)
	if err != nil {
		log.Println(err.Error())
		util.SendErrorResponse(w, http.StatusInternalServerError, "")
		return
	}
	util.OKResponse(w, verse)
}

func (me *Controller) handleGetRandomVerse(w http.ResponseWriter, _ *http.Request) {
	verse, err := me.repo.getRandomVerse()
	if err != nil {
		log.Println(err.Error())
		util.SendErrorResponse(w, http.StatusInternalServerError, "")
		return
	}
	util.OKResponse(w, verse)
}

func (me *Controller) handleNewVerse(w http.ResponseWriter, r *http.Request) {
	var req Verse
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println(err.Error())
		util.SendErrorResponse(w, http.StatusBadRequest, "error decode json")
		return
	}

	newId, err := me.repo.storeNewVerse(&req)
	if err != nil {
		log.Println(err.Error())
		util.SendErrorResponse(w, http.StatusInternalServerError, "")
		return
	}

	res := Verse{
		Id: newId,
	}

	util.CreatedResponse(w, res)
}

func (me *Controller) handleDeleteVerseById(w http.ResponseWriter, r *http.Request) {
	var id int
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println(err.Error())
		util.SendErrorResponse(w, http.StatusBadRequest, "error getting id")
		return
	}
	err = me.repo.deleteVerseById(id)
	if err == util.ErrNothingToDelete {
		log.Println(err.Error())
		util.SendErrorResponse(w, http.StatusBadRequest, "nothing to delete")
		return
	}
	if err != nil {
		log.Println(err.Error())
		util.SendErrorResponse(w, http.StatusInternalServerError, "error deleting")
		return
	}

	util.OKResponse(w, util.MessageData{Message: "deleted"})
}
