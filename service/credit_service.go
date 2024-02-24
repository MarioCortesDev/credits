package service

import (
	"credit/repository"
	"errors"
)

// CreditService implementa la interfaz CreditAssigner y proporciona métodos para asignar créditos.
type CreditService struct {
	repository *repository.CreditRepository
}

// NewCreditService crea una nueva instancia de CreditService.
func NewCreditService(repository *repository.CreditRepository) *CreditService {
	return &CreditService{
		repository: repository,
	}
}

// Assign asigna créditos según la inversión proporcionada.
func (cs *CreditService) Assign(investment int32) (int32, int32, int32, error) {
	// Verificar si la inversión es un múltiplo de 100
	if investment%100 != 0 {
		return 0, 0, 0, errors.New("La inversión no es un múltiplo de 100")
	}

	// Calcular la cantidad de créditos de cada tipo que podemos asignar con la inversión
	var count300, count500, count700 int32

	for i300 := int32(0); i300 <= investment/300; i300++ {
		for i500 := int32(0); i500 <= investment/500; i500++ {
			for i700 := int32(0); i700 <= investment/700; i700++ {
				total := i300*300 + i500*500 + i700*700
				if total == investment {
					count300 = i300
					count500 = i500
					count700 = i700
					break
				}
			}
		}
	}

	return count300, count500, count700, nil
}
