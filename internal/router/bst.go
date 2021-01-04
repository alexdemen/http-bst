package router

import (
	"encoding/json"
	"github.com/alexdemen/http-bst/internal/dto"
	"github.com/alexdemen/http-bst/internal/service"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
	"strconv"
)

type BSTRouter struct {
	chi.Router
	service service.TreeService
}

func (router *BSTRouter) searchHandle() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if key, err := strconv.Atoi(r.URL.Query().Get("val")); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			foundedKey, err := router.service.Find(key)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
			} else {
				res := dto.TreeResponse{Val: foundedKey}
				body, err := json.Marshal(res)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.Write(body)
			}
		}
	}
}

func (router *BSTRouter) insertHandler() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		defer request.Body.Close()

		insertReq := dto.TreeRequest{}
		err = json.Unmarshal(body, &insertReq)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		err = router.service.Add(insertReq.Val)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		} else {
			writer.WriteHeader(http.StatusOK)
		}
	}
}

func (router *BSTRouter) deleteHandler() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		if key, err := strconv.Atoi(request.URL.Query().Get("val")); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
		} else {
			err := router.service.Delete(key)
			if err != nil {
				http.Error(writer, "", http.StatusInternalServerError)
				return
			}
			writer.WriteHeader(http.StatusOK)
		}
	}
}

func NewBSTRouter(service service.TreeService) *BSTRouter {
	r := &BSTRouter{
		service: service,
	}

	r.Router = chi.NewRouter()

	r.Get("/search", r.searchHandle())
	r.Router.Post("/insert", r.insertHandler())
	r.Router.Delete("/delete", r.insertHandler())

	return r
}
