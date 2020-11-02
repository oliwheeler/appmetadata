package api

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/oliwheeler/appmetadata/service"
)

func New() http.Handler {
	router := httprouter.New()
	svc := service.New()
	api := &Api{svc}
	registerHandlers(router, api)
	return router
}

type Api struct {
	svc *service.Service
}

func registerHandlers(router *httprouter.Router, api *Api) {
	router.GET("/", api.get)
	router.GET("/:title", api.getApp)
	router.POST("/", api.createApp)
	router.PUT("/:title", api.updateApp)
}

func (api *Api) createApp(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer r.Body.Close()
	if err := api.svc.CreateAppMetadata(r.Body); err != nil {
		w.WriteHeader(getHttpResponseCode(err))
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (api *Api) get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func (api *Api) getApp(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	title := ps.ByName("title")
	stream, err := api.svc.GetAppMetadata(title)
	if err != nil {
		w.WriteHeader(getHttpResponseCode(err))
		w.Write([]byte(err.Error()))
		return
	}
	data, err := ioutil.ReadAll(stream)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(data)
	w.WriteHeader(http.StatusOK)
}

func (api *Api) updateApp(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer r.Body.Close()
	title := ps.ByName("title")
	if err := api.svc.UpdateAppMetadata(title, r.Body); err != nil {
		w.WriteHeader(getHttpResponseCode(err))
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func getHttpResponseCode(err error) int {
	var yamlErr *service.InValidYamlError
	var alreadyExistsErr *service.MetadataNameAlreadyExistsError
	var cannotGetErr *service.CannotNotGetMetadataError
	var cannotUpdateErr *service.CannotUpdateNonExistantMetadataError
	if errors.Is(err, service.CannotUpdateTitle) {
		return http.StatusBadRequest
	}
	if errors.As(err, &yamlErr) {
		return http.StatusBadRequest
	}
	if errors.As(err, &alreadyExistsErr) {
		return http.StatusConflict
	}
	if errors.As(err, &cannotUpdateErr) {
		return http.StatusNotFound
	}
	if errors.As(err, &cannotGetErr) {
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}
