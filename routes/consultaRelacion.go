package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Paulo-Ariel-Pareja/backend-red-social-go/bd"
	"github.com/Paulo-Ariel-Pareja/backend-red-social-go/models"
)

/*ConsultaRelacion consulta si un usuario tiene una relacion*/
func ConsultaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El id es obligatorio", http.StatusBadRequest)
		return
	}
	var t models.Relacion
	t.UsuarioRelacionID = ID
	t.UsuarioID = IDUsuario
	var resp models.RespuestaConsultaRelacion

	status, err := bd.ConsultarRelacion(t)
	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
