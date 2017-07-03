package solvers

import (
	"testing"
	"github.com/wliao008/mazing/algos"
	"github.com/wliao008/mazing/util"
)

func BenchmarkDeadEnds_1000x500(b *testing.B) {
	k := algos.NewPrim(1000, 500)
	k.Generate()
	def := DeadEndFiller{}
	def.Board = &k.Board
	stack := &util.Stack{}
	for i := 0; i < b.N; i++ {
		def.Board.DeadEnds(stack)
	}
}

func BenchmarkSolve_1000x500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k := algos.NewPrim(1000, 500)
		k.Generate()
		def := DeadEndFiller{}
		def.Board = &k.Board
		def.Solve()
	}
}