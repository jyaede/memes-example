package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/yaede7/go-rest-api-example/collection"
	"github.com/yaede7/go-rest-api-example/imgflip"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//MemeHandler ...
type MemeHandler struct {
	memes *imgflip.Client
	sess  *mgo.Session
}

//Available get available memes
func (h MemeHandler) Available(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	memes, err := h.memes.GetMemes()
	if err != nil {
		jsonError(w, err, http.StatusBadRequest)
		return
	}
	jsonWithCode(w, memes, http.StatusOK)
}

//Collection ...
func (h MemeHandler) Collection(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//copy mongodb session
	sess := h.sess.Copy()
	defer sess.Close()
	m := h.getCollectionManager(sess)

	//get my collection
	items, err := m.All()
	if err != nil {
		jsonError(w, err, http.StatusInternalServerError)
		return
	}
	jsonWithCode(w, items, http.StatusCreated)
}

//Create create a meme
func (h MemeHandler) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//decode json post body and convert to struct
	var f imgflip.CaptionForm
	if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
		panic(err)
	}
	defer r.Body.Close()

	//add caption to the image
	image, err := h.memes.CaptionImage(f)
	if err != nil {
		jsonError(w, err, http.StatusBadRequest)
		return
	}

	//copy mongodb session
	sess := h.sess.Copy()
	defer sess.Close()
	m := h.getCollectionManager(sess)

	//add meme to my collection
	item, err := m.Create(&collection.Item{
		PageURL:  image.PageURL,
		ImageURL: image.URL,
		Text:     fmt.Sprintf("%s %s", f.Text0, f.Text1),
		Created:  time.Now(),
	})
	if err != nil {
		jsonError(w, err, http.StatusInternalServerError)
		return
	}
	jsonWithCode(w, item, http.StatusCreated)
}

//Get get a meme from my collection
func (h MemeHandler) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//copy mongodb session
	sess := h.sess.Copy()
	defer sess.Close()
	m := h.getCollectionManager(sess)

	//Get id from path parameter
	id := p.ByName("id")

	//get item from my collection
	items, err := m.Get(bson.ObjectIdHex(id))
	if err != nil {
		jsonError(w, err, http.StatusInternalServerError)
		return
	}
	jsonWithCode(w, items, http.StatusCreated)
}

//Delete get a meme from my collection
func (h MemeHandler) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//copy mongodb session
	sess := h.sess.Copy()
	defer sess.Close()
	m := h.getCollectionManager(sess)

	//Get id from path parameter
	id := p.ByName("id")

	//delete item from my collection
	err := m.Delete(bson.ObjectIdHex(id))
	if err != nil {
		jsonError(w, err, http.StatusInternalServerError)
		return
	}
	jsonWithCode(w, nil, http.StatusNoContent)
}

func (h MemeHandler) getCollectionManager(sess *mgo.Session) collection.Manager {
	return collection.NewManager(sess)
}
