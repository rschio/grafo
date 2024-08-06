package internal_test

import (
	"testing"

	"github.com/rschio/grafo"
	"github.com/rschio/grafo/internal"
)

func TestGraphCompatibility(t *testing.T) {
	var g grafo.Graph[struct{}]
	var h internal.Graph[struct{}]

	g = h
	h = g
}
