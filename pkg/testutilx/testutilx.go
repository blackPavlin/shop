// Package testutilx contains utils for tests.
package testutilx

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

// BaseSuite base suite for reuse.
type BaseSuite struct {
	suite.Suite
	Ctrl *gomock.Controller
}

// SetupSuite setup base suite.
func (s *BaseSuite) SetupSuite() {}

// SetupTest setup before test.
func (s *BaseSuite) SetupTest() {
	s.Ctrl = gomock.NewController(s.T())
}

// TearDownTest tear down after test.
func (s *BaseSuite) TearDownTest() {
	if s.Ctrl != nil {
		s.Ctrl.Finish()
	}
}
