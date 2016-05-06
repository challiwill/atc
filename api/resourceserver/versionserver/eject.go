package versionserver

import (
	"net/http"
	"strconv"

	"github.com/concourse/atc/db"
	"github.com/tedsuo/rata"
)

func (s *Server) EjectResource(pipelineDB db.PipelineDB) http.Handler {
	logger := s.logger.Session("delete-resource")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resourceID, err := strconv.Atoi(rata.Param(r, "resource_version_id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		_, found, err := pipelineDB.DeleteResourceVersion(resourceID)
		if err != nil {
			logger.Error("failed-to-delete-versioned-resource", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if !found {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}
