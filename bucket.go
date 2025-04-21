package dht

import (
	"github.com/istancescu/int160-go-rd"
	"time"
)

const BucketSize = 8 // Kademlia bucket size

type Bucket struct {
	Nodes    []*Node
	RangeMin int160.Int160
	RangeMax int160.Int160
}

func NewBucket() *Bucket {
	return &Bucket{
		Nodes: make([]*Node, BucketSize),
	}
}

func (b *Bucket) Add(node *Node) {
	for i := range b.Nodes {
		if b.Nodes[i].ID.Equals(node.ID) {
			//if same replace
			b.Nodes[i] = node
			return
		}
	}
	if len(b.Nodes) < BucketSize {
		b.Nodes = append(b.Nodes, node)
		return
	}

	oldest := b.Nodes[0]
	if isAlive(*oldest) {
		// Oldest is alive, reject the new node
		return
	}

	// Replace oldest with the new node
	b.Nodes = append(b.Nodes[1:], node)
}

// stubbed
func isAlive(n Node) bool {
	return time.Since(n.LastAlive) < 15*time.Minute
}
