package main

type IContext interface {
	SetStrategy(strategy IStrategy)
}
