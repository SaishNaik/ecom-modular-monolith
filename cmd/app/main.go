package main

import (
	"go.uber.org/automaxprocs/maxprocs"
	"log/slog"
)

func main() {
	_, err := maxprocs.Set()
	if err != nil {
		slog.Error()
	}
}
