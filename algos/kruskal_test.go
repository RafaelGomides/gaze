package algos

import (
	"testing"
	"github.com/wliao008/mazing/structs"
	"math/rand"
	"time"
)

func TestKruskalAlgo_AllCellsVisited(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	count := rand.Intn(50) + 1
	for c := 0; c < count; c++ {
		k := Kruskal{uint16(rand.Intn(100)+1), uint16(rand.Intn(100)+1), nil}
		k.Generate()
		for i := uint16(0); i < k.Width; i++ {
			for j := uint16(0); j < k.Height; j++ {
				if !k.Cells[i][j].IsSet(structs.VISITED) {
					t.Errorf("Every cell should be visited, but not [%d,%d]", i, j)
				}
			}
		}
	}
}

func BenchmarkKruskalAlgo(b *testing.B) {
	k := Kruskal{100, 50, nil}
	for i := 0; i < b.N; i++ {
		k.Generate()
	}
}