package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/tamboto2000/kopikagum/gateway/model"
	"github.com/tamboto2000/kopikagum/response"
	"gorm.io/gorm"
)

// Gateway is gateway handler, forwarding request to it's service
type Gateway struct {
	db   *gorm.DB
	ends *sync.Map
}

// NewGateway initiate gateway handler
func NewGateway(db *gorm.DB) *Gateway {
	g := &Gateway{db: db, ends: new(sync.Map)}

	// migrate tables
	if err := g.db.AutoMigrate(
		new(model.Middleware),
		new(model.Endpoint),
	); err != nil {
		panic(err.Error())
	}

	// get endpoints
	ends := make([]*model.Endpoint, 0)
	if err := db.Find(&ends).Error; err != nil {
		panic(err.Error())
	}

	for _, end := range ends {
		g.ends.Store(end.Path, end)
	}

	return g
}

func (g *Gateway) getEndpoint(end string) *model.Endpoint {
	val, ok := g.ends.Load(end)
	if !ok {
		return nil
	}

	return val.(*model.Endpoint)
}

// API forward request to endpoint path, or to middleware if any
func (g *Gateway) API(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	path := vars["path"]

	// make sure request is in application/json and not text/html
	r.Header.Set("Content-Type", "application/json")

	val, ok := g.ends.Load("/" + path)
	if !ok {
		enc, _ := json.Marshal(response.NotFound("endpoint not found"))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(enc)
		return
	}

	end := val.(*model.Endpoint)

	var endStr string

	if end.Host != "" {
		endStr += end.Host
	} else {
		endStr = "http://localhost"
	}

	if end.Port != "" {
		endStr += ":" + end.Port
	}

	endStr += end.Path
	log.Println(end.Method, endStr)

	http.Redirect(w, r, endStr, http.StatusFound)
}
