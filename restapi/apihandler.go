package restapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test01/datalayer"

	"github.com/gorilla/mux"
)

type PromotionRestApiHandler struct {
	dbhandler datalayer.SQLHandler
}

func newPromotionRestApiHandler(db datalayer.SQLHandler) *PromotionRestApiHandler {
	return &PromotionRestApiHandler{
		dbhandler: db,
	}
}

func (handle PromotionRestApiHandler) returnSinglePromotion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, ok := vars["id"]

	if !ok {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintln(w, "not found")
		return
	}

	promotion, err := handle.dbhandler.GetPromitionById(key)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintln(w, "not found")
		return
	}

	json.NewEncoder(w).Encode(promotion)
	return
}
