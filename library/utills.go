package library

import (
	"encoding/json"
	"net/http"
)

//Utill component
type Utill struct{}

//RespondWithError to respond with error status
func (u *Utill) RespondWithError(w http.ResponseWriter, code int, msg string) {
	u.RespondWithJSON(w, code, map[string]string{"error": msg})
}

//RespondWithJSON to respond json
func (u *Utill) RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
