package routers_test

import (
	"testing"
)

func TestRouterInclusion(t *testing.T) {
	rootRouter := New("root")
	childRouter := New("child")
	grandchildRouter := New("grandchild")

	err := rootRouter.Include(childRouter)
	if err != nil {
		t.Error(err)
	}

	err = childRouter.Include(rootRouter)
	if err == nil {
		t.Error("Child can't include parent router")
	}

	err = childRouter.Include(grandchildRouter)
	if err != nil {
		t.Error(err)
	}

	err = grandchildRouter.Include(rootRouter)
	if err == nil {
		t.Error("Child can't include grandparent router")
	}

	err = rootRouter.Include(rootRouter)
	if err == nil {
		t.Error("A Router can't include itself")
	}

	router1 := New("Test router 1")
	router2 := New("Test router 2")
	router3 := New("Test router 3")

	err = router1.Include(router2)
	if err != nil {
		t.Error(err)
	}

	err = router3.Include(router3)
	if err == nil {
		t.Error("A Router can't be included more than once")
	}
}
