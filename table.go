package dht

import "github.com/istancescu/int160-go-rd"

type Table struct {
	Buckets []Bucket
	Self    *int160.Int160
}

func CreateTable(buckets []Bucket) *Table {
	return &Table{
		Buckets: buckets,
	}
}

func (t *Table) Add(bucket Bucket) {
	t.Buckets = append(t.Buckets, bucket)
}

func (t *Table) Remove(bucket Bucket) {

}

func (t *Table) Index(s *int160.Int160) uint8 {
	return t.Self.CommonPrefixLen(s)
}
