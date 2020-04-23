package Controllers

import (
	"dataclips/Models"
	"encoding/json"
	"log"
	"net/http"

	h "dataclips/Helpers"
)

var Dataclips []Models.Dataclip

// Createdataclip Handler
// @Summary Create a new dataclip
// @Description Create a new dataclip with the input payload
// @Tags back_end_dataclips
// @Accept  json
// @Produce  json
// @Param dataclip body Models.Dataclip false "Dataclip"
// @Success 200 {object} Models.Dataclip
// @Router /Dataclips [post]
func CreateDataclip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)

	var DC Models.Dataclip

	err := decoder.Decode(&DC)
	if err != nil {
		panic(err)
	}

	log.Println("new Dataclip :", DC)

	db := h.DbConnect()
	db.NewRecord(DC)
	db.Create(&DC)
	_ = json.NewDecoder(r.Body).Decode(&DC)
	Dataclips = append(Dataclips, DC)
	json.NewEncoder(w).Encode(DC)

}
