package models

// CreditAssigner especifica la interfaz para asignar créditos según la inversión.
type CreditAssigner interface {
	Assign(investment int32) (int32, int32, int32, error)
}

// Investment representa la inversión de un cliente.
type Investment struct {
	Amount int32
}

// Credit representa la cantidad de créditos de un tipo específico.
type Credit struct {
	Type   int32 // Tipo de crédito: 300, 500 o 700
	Amount int32 // Cantidad de créditos
}
