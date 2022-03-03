package main

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type LockedStateSuite struct {
	suite.Suite
}

func (s *LockedStateSuite) SetupSuite() {

}

func (s *LockedStateSuite) AfterTests(_, _ string) {
	require.NoError(s.T(), nil)
}
