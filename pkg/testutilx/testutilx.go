package testutilx

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

// BaseSuite
type BaseSuite struct {
	suite.Suite
	Ctrl *gomock.Controller
}

// SetupSuite
func (s *BaseSuite) SetupSuite() {}

// SetupTest
func (s *BaseSuite) SetupTest() {
	s.Ctrl = gomock.NewController(s.T())
}

// TearDownTest
func (s *BaseSuite) TearDownTest() {
	if s.Ctrl != nil {
		s.Ctrl.Finish()
	}
}
