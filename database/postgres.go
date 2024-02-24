package database

import (
	"database/sql"
	"fmt"
	"log"
)

// PostgreSQLRepository representa un repositorio de PostgreSQL.
type PostgreSQLRepository struct {
	db *sql.DB
}

// NewPostgreSQLRepository crea una nueva instancia de PostgreSQLRepository.
func NewPostgreSQLRepository(connectionString string) (*PostgreSQLRepository, error) {
	// Conectar a la base de datos PostgreSQL
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("error al conectar a la base de datos: %v", err)
	}

	// Verificar la conexión a la base de datos
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error al verificar la conexión a la base de datos: %v", err)
	}

	log.Println("Conexión a la base de datos establecida")

	return &PostgreSQLRepository{db: db}, nil
}

// InsertCreditAssignment inserta una asignación de créditos en la base de datos.
func (repo *PostgreSQLRepository) InsertCreditAssignment(investment, count300, count500, count700 int32, successful bool) error {
	_, err := repo.db.Exec("INSERT INTO credits_assignments (investment, count_300, count_500, count_700, successful) VALUES ($1, $2, $3, $4, $5)",
		investment, count300, count500, count700, successful)
	if err != nil {
		return fmt.Errorf("error al insertar la asignación de créditos en la base de datos: %v", err)
	}
	return nil
}

// GetStatistics obtiene las estadísticas de las asignaciones de créditos.
func (repo *PostgreSQLRepository) GetStatistics() (int, int, int, float64, float64, error) {
	var totalAssignments, successfulAssignments, unsuccessfulAssignments int
	var sumSuccessfulInvestment, sumUnsuccessfulInvestment float64

	// Obtener el total de asignaciones
	err := repo.db.QueryRow("SELECT COUNT(*) FROM credits_assignments").Scan(&totalAssignments)
	if err != nil {
		return 0, 0, 0, 0, 0, fmt.Errorf("error al obtener el total de asignaciones: %v", err)
	}

	// Obtener el total de asignaciones exitosas
	err = repo.db.QueryRow("SELECT COUNT(*) FROM credits_assignments WHERE successful = true").Scan(&successfulAssignments)
	if err != nil {
		return 0, 0, 0, 0, 0, fmt.Errorf("error al obtener el total de asignaciones exitosas: %v", err)
	}

	// Calcular el total de asignaciones no exitosas
	unsuccessfulAssignments = totalAssignments - successfulAssignments

	// Obtener la suma de las inversiones exitosas
	err = repo.db.QueryRow("SELECT SUM(investment) FROM credits_assignments WHERE successful = true").Scan(&sumSuccessfulInvestment)
	if err != nil {
		return 0, 0, 0, 0, 0, fmt.Errorf("error al obtener la suma de inversiones exitosas: %v", err)
	}

	// Obtener la suma de las inversiones no exitosas
	err = repo.db.QueryRow("SELECT SUM(investment) FROM credits_assignments WHERE successful = false").Scan(&sumUnsuccessfulInvestment)
	if err != nil {
		return 0, 0, 0, 0, 0, fmt.Errorf("error al obtener la suma de inversiones no exitosas: %v", err)
	}

	// Calcular el promedio de inversiones exitosas
	avgSuccessfulInvestment := sumSuccessfulInvestment / float64(successfulAssignments)

	// Calcular el promedio de inversiones no exitosas
	avgUnsuccessfulInvestment := sumUnsuccessfulInvestment / float64(unsuccessfulAssignments)

	return totalAssignments, successfulAssignments, unsuccessfulAssignments, avgSuccessfulInvestment, avgUnsuccessfulInvestment, nil
}

// Close cierra la conexión con la base de datos.
func (repo *PostgreSQLRepository) Close() {
	err := repo.db.Close()
	if err != nil {
		log.Printf("Error al cerrar la conexión a la base de datos: %v\n", err)
	}
}
