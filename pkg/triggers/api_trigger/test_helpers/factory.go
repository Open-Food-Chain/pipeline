package test_helpers

import (
	"github.com/stretchr/testify/suite"
	"github.com/The-New-Fork/pipeline/pkg/helper"
	"github.com/unchainio/interfaces/logger"
)

type TestHelpers struct {
	suite  *suite.Suite
	logger logger.Logger
	helper *helper.Helper
}

func NewTestHelpers(s *suite.Suite, logger logger.Logger, helper *helper.Helper) *TestHelpers {
	return &TestHelpers{suite: s, logger: logger, helper: helper}
}
