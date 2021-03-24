package cron_trigger_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/The-New-Fork/pipeline/pkg/factory"
	"github.com/The-New-Fork/pipeline/pkg/helper"
	"github.com/unchainio/interfaces/logger"
	"testing"
)

type TestSuite struct {
	suite.Suite
	logger  logger.Logger
	helper  *helper.Helper
}

func (s *TestSuite) SetupSuite() {
	s.logger = factory.DefaultLogger(&s.Suite)
	s.helper = helper.NewHelper(&s.Suite, s.logger)
}

func TestRunTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}


