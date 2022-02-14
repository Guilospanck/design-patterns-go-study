package main

type IStrategy interface {
	Execute(a, b int) int
}
