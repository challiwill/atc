package resourceserver

import (
	"encoding/json"
	"net/http"

	"github.com/concourse/atc"
	"github.com/concourse/atc/api/present"
	"github.com/concourse/atc/db"
	"github.com/pivotal-golang/lager"
)

func (s *Server) EjectResource(pipelineDB db.PipelineDB) http.Handler {
	logger := s.logger.Session("delete-resource")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		config, _, found, err := pipelineDB.GetConfig()
		if err != nil {
			logger.Error("failed-to-get-config", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !found {
			logger.Info("config-not-found")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		resourceName := r.FormValue(":resource_name")
		_, resourceFound := config.Resources.Lookup(resourceName)
		if !resourceFound {
			logger.Info("resource-not-in-config")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		resourceVersion := r.FormValue(":resource_version")
		var version atc.Version
		err = json.Unmarshal([]byte(resourceVersion), &version)
		if err != nil {
			logger.Info("invalid-version")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		dbResource, found, err := pipelineDB.DeleteResourceVersion(resourceName, version)
		if err != nil {
			logger.Error("failed-to-delete-resource", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !found {
			logger.Debug("resource-not-found", lager.Data{"resource": resourceName})
			w.WriteHeader(http.StatusNotFound)
			return
		}

		resource := present.SavedVersionedResource(dbResource)

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(resource)
	})
}
