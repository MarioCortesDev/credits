package repository

// CreditRepository simula el almacenamiento de créditos disponibles.
type CreditRepository struct {
	credits map[int32]int32 // Mapa para almacenar la cantidad de créditos disponibles para cada tipo
}

// NewCreditRepository crea una nueva instancia de CreditRepository.
func NewCreditRepository() *CreditRepository {
	return &CreditRepository{
		credits: map[int32]int32{
			300: 100, // Cantidad inicial de créditos de $300
			500: 100, // Cantidad inicial de créditos de $500
			700: 100, // Cantidad inicial de créditos de $700
		},
	}
}

// GetAvailableCredits devuelve la cantidad de créditos disponibles.
func (cr *CreditRepository) GetAvailableCredits() map[int32]int32 {
	return cr.credits
}

// UpdateCredits actualiza la cantidad de créditos disponibles.
func (cr *CreditRepository) UpdateCredits(credits map[int32]int32) {
	cr.credits = credits
}
