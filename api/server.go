package api

import (
	"encoding/json"
	"net/http"

	"credit/database"
	"credit/service"
)

// CreditAssignmentHandler maneja las solicitudes de asignación de créditos.
func CreditAssignmentHandler(creditService *service.CreditService, db *database.PostgreSQLRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Verificar si la solicitud es de tipo POST
		if r.Method != http.MethodPost {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		// Decodificar el cuerpo de la solicitud JSON
		var req struct {
			Investment int32 `json:"investment"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Error al decodificar el cuerpo de la solicitud", http.StatusBadRequest)
			return
		}

		// Asignar créditos utilizando el servicio
		count300, count500, count700, err := creditService.Assign(req.Investment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Determinar si la asignación fue exitosa
		successful := true
		if count300 == 0 && count500 == 0 && count700 == 0 {
			successful = false
			// Retornar un error 400 si no se puede realizar la asignación
			http.Error(w, "No se pudo asignar crédito", http.StatusBadRequest)
			return
		}

		// Crear la respuesta JSON
		res := map[string]int32{
			"credit_type_300": count300,
			"credit_type_500": count500,
			"credit_type_700": count700,
		}

		// Insertar la asignación de créditos en la base de datos
		if err := db.InsertCreditAssignment(req.Investment, count300, count500, count700, successful); err != nil {
			http.Error(w, "Error al insertar la asignación de créditos en la base de datos", http.StatusInternalServerError)
			return
		}

		// Codificar y enviar la respuesta
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}

// StatisticsHandler maneja las solicitudes de estadísticas de asignaciones.
func StatisticsHandler(db *database.PostgreSQLRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Consultar la base de datos para obtener las estadísticas
		totalAssignments, successfulAssignments, unsuccessfulAssignments, averageSuccessfulInvestment, averageUnsuccessfulInvestment, err := db.GetStatistics()
		if err != nil {
			http.Error(w, "Error al obtener las estadísticas de asignaciones", http.StatusInternalServerError)
			return
		}

		// Crear la respuesta JSON
		res := map[string]interface{}{
			"total_assignments":               totalAssignments,
			"successful_assignments":          successfulAssignments,
			"unsuccessful_assignments":        unsuccessfulAssignments,
			"average_successful_investment":   averageSuccessfulInvestment,
			"average_unsuccessful_investment": averageUnsuccessfulInvestment,
		}

		// Codificar y enviar la respuesta
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}

// InitRoutes inicializa las rutas de la API.
func InitRoutes(creditService *service.CreditService, db *database.PostgreSQLRepository) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/credit-assignment", CreditAssignmentHandler(creditService, db))
	mux.HandleFunc("/statistics", StatisticsHandler(db))
	return mux
}
