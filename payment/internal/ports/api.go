package ports

import (
	"context"

	"github.com/dmitryovchinnikov/microservices/payment/internal/application/core/domain"
)

type APIPort interface {
	Charge(ctx context.Context, payment domain.Payment) (domain.Payment, error)
}
