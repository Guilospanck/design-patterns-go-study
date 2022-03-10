package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

var DataMock = Data{
	Name: "Product 1",
	Size: 68,
}

var DataMockMarshalled = DataMock.Name + "<>" + fmt.Sprintf("%d", DataMock.Size)

type DataTestSuite struct {
	suite.Suite
	data Data
}

func (s *DataTestSuite) SetupSuite() {
	s.data = DataMock
}

func (s *DataTestSuite) AfterTests(_, _ string) {
	require.NoError(s.T(), nil)
}

func TestDataSuiteInit(t *testing.T) {
	suite.Run(t, new(DataTestSuite))
}

func (s *DataTestSuite) TestDataMarshall() {
	// arrange

	// act
	result := s.data.Marshal()

	// assert
	require.Equal(s.T(), DataMockMarshalled, result)

}
