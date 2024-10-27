package routers_test

import (
	"github.com/golang-chaos/chaos/routers"
	"testing"
)

func TestRouterInclusion(t *testing.T) {
	rootRouter := routers.New("root")
	childRouter := routers.New("child")
	grandchildRouter := routers.New("grandchild")

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
}
