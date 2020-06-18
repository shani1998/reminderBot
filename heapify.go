package main

import (
	"github.com/turnage/graw/reddit"
)

type MaxHeap struct {
	post []*reddit.Post
	heapSize int
}

func BuildMaxHeap(posts []*reddit.Post) MaxHeap {
	h := MaxHeap{post: posts, heapSize: len(posts)}
	for i := len(posts) / 2; i >= 0; i-- {
		h.MaxHeapify(i)
	}
	return h
}

func (h MaxHeap) MaxHeapify(i int) {
	l, r := 2*i+1, 2*i+2
	max := i

	if l < h.size() && h.post[l].Ups > h.post[max].Ups {
		max = l
	}
	if r < h.size() && h.post[r].Ups > h.post[max].Ups {
		max = r
	}
	//log.Printf("MaxHeapify(%v): l,r=%v,%v; max=%v\t%v\n", i, l, r, max, h.slice)
	if max != i {
		h.post[i], h.post[max] = h.post[max], h.post[i]
		h.MaxHeapify(max)
	}
}

func (h MaxHeap) size() int { return h.heapSize } // ???



