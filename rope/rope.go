package rope

import (
	"fmt"
)

// TODO: Switch to a red-black tree for performance
type node[T any] struct {
	data *T	
	weight int

	left *node[T]
	right *node[T]
}

// This is a generic rope (not purely string-based like traditional ropes), so some of the semantics
// are slightly different.
type Rope[T any] struct {
	root *node[T]
	length int
}

// Any nodes will have a left and right child if node weight >= 2 (i.e. no dangler like Wikipedia)
// Node-weight is 1 for a leaf node 
// 


func (r *Rope[T]) Insert(d T, i int) {
	r.length++
	toAdd := &node[T]{data: &d, weight: 1}
	if r.root == nil {
		r.root = toAdd
		return
	}

	// Find the weights of the left-subtree
	var traverse func(*node[T], int)
	traverse = func(n *node[T], i int) {
		// left == nil implies right == nil
		// It is left as an exercise to the readers to prove this
		if n.left == nil {
			prev := &node[T]{data: n.data, weight: 1}
			if i == 0 {
				n.left = toAdd
				n.right =  prev
			} else {
				n.left = prev
				n.right = toAdd
			}
			n.data = nil
			return
		}

		if i < n.weight {
			n.weight += toAdd.weight
			traverse(n.left, i)
		} else {
			traverse(n.right, i - n.weight)
		}
	}
	// If there is no left-subtree, create it
	// If it is less than our index, then insert into the right	
	// If  
	traverse(r.root, i)
}

func (r *Rope[T]) Get(i int) T {
	if r.root == nil || i < 0  || i >= r.Len() {
		panic(fmt.Sprintf("index %d is out of bounds for rope of size %d", i, r.Len()))
	}

	var traverse func(*node[T], int) T
	traverse = func(n *node[T], i int) T {
		if n.left == nil && n.right == nil {
			return *n.data
		}

		if i < n.weight {
			// If idx is less than the weight, that means that it is contained within the left
			// subtree, so we don't need to subtract.
			return traverse(n.left, i);
		} else {
			// If the desired index is greater than our current weight, then we need to "offset"
			// the index by the weight of our current node.
			return traverse(n.right, i - n.weight);
		}
	}

	return traverse(r.root, i)
}

func (r *Rope[T]) Len() int {
	return r.length
}

// Splits the rope into two pieces, with indices [0, i] and [i+1, length]
func (r *Rope[T]) Split(i int) (*Rope[T], *Rope[T]) {
	// Num in left: i + 1
	// Num in right: len - i
	var nodes []*node[T]
	var traverse func(*node[T], int) T
	traverse = func(n *node[T], i int) {
		if n.left == nil && n.right == nil {
		}

		if i < n.weight {
			return traverse(n.left, i);
		} else {
			traverse(n.right, i - n.weight);
		}
		
	}

	return nil, nil
}

// Concatenates r2 to the end of r1
func (r *Rope[T]) Concat(r2 *Rope[T]) {
	n := &node[T]{ weight: r.Len(), left: r.root, right: r2.root }

	r.root = n;
	r.length += r2.Len()
}

// Delete an individual 
func (r *Rope[T]) Delete(i int) {
	// If we go left, reduce the weight at the given node
	// If we go right, no need to reduce the node weight but reduce i

	// TODO: validate this!

	var traverse func(*node, int)
	traverse = func(n *node, i int) {
		if n.left == nil && n.right == nil {
			// TODO:	
		}		

		if i < n.weight {
			n.weight--
			if n.left != nil && n.left.left == nil && n.right.right == nil {
				// TODO: Need to adjust indices if that has a sibling
				n.left = nil;
			} else {
				traverse(n.left, i)
			}
		} else {
			if n.right != nil && n.right.left == nil && n.right.right == nil {
				// TODO: think more carefully
				n.right = nil;
			} else {
				traverse(n.right, i - n.weight)
			}
		}
	}

	return traverse(r.root, i)
}

func (r *Rope[T]) DeleteRange(start int, end int) {
	// Split r into three pieces
	// 0 to start, start to end, and end to rope.length (2 splits, probably)
	// Concatenate the first and last rope we get from that


}