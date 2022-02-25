package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ContextSuite struct {
	suite.Suite

	mock         *ContextMock
	context      *Context
	strategyMock IStrategy
}

func (s *ContextSuite) SetupSuite() {
	s.mock = NewContextMock()
	s.context = NewContext()
}

func (s *ContextSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), nil)
}

func TestContextInit(t *testing.T) {
	suite.Run(t, new(ContextSuite))
}

func (s *ContextSuite) TestSetStrategyWithAddStrategy() {
	// arrange
	s.strategyMock = NewConcreteStrategyAddMock()

	// act
	s.context.SetStrategy(s.strategyMock)

	// assert
	require.Equal(s.T(), s.strategyMock, s.context.strategy)
}

func (s *ContextSuite) TestSetStrategyWithMultiplyStrategy() {
	// arrange
	s.strategyMock = NewConcreteStrategyMultiplyMock()

	// act
	s.context.SetStrategy(s.strategyMock)

	// assert
	require.Equal(s.T(), s.strategyMock, s.context.strategy)
}

func (s *ContextSuite) TestSetStrategyWithSubtractStrategy() {
	// arrange
	s.strategyMock = NewConcreteStrategySubtractMock()

	// act
	s.context.SetStrategy(s.strategyMock)

	// assert
	require.Equal(s.T(), s.strategyMock, s.context.strategy)
}

func (s *ContextSuite) TestDoSomethingWithAddStrategy() {
	// arrange
	s.strategyMock = NewConcreteStrategyAddMock()
	a := 10
	b := 5
	sumMock := a + b

	// act
	s.context.SetStrategy(s.strategyMock)
	sum := s.context.DoSomething(a, b)

	// assert
	require.Equal(s.T(), s.strategyMock, s.context.strategy)
	require.Equal(s.T(), sumMock, sum)
	require.Equal(s.T(), true, AddStrategyCalled)
}

func (s *ContextSuite) TestDoSomethingMultiplyStrategy() {
	// arrange
	s.strategyMock = NewConcreteStrategyMultiplyMock()
	a := 10
	b := 20
	multiplyMock := a * b

	// act
	s.context.SetStrategy(s.strategyMock)
	multiple := s.context.DoSomething(a, b)

	// assert
	require.Equal(s.T(), multiplyMock, multiple)
	require.Equal(s.T(), s.strategyMock, s.context.strategy)
	require.Equal(s.T(), true, StrategyMultiplyCalled)
}

func (s *ContextSuite) TestDoSomethingSubtractStrategy() {
	// arrange
	s.strategyMock = NewConcreteStrategySubtractMock()
	a := 50
	b := 40
	subtractMock := a - b

	// act
	s.context.SetStrategy(s.strategyMock)
	subtracted := s.context.DoSomething(a, b)

	// assert
	require.Equal(s.T(), subtractMock, subtracted)
	require.Equal(s.T(), s.strategyMock, s.context.strategy)
	require.Equal(s.T(), true, SubtractStrategyCalled)
}
