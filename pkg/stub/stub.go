package stub

import (
	"github.com/The-New-Fork/pipeline/pkg/domain"
	"github.com/unchainio/interfaces/logger"
)

type Stub struct {
	logger.Logger
}

func New(logger logger.Logger) domain.Stub {
	return &Stub{logger}
}
