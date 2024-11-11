package routers

import "errors"

// Router can handle events itself or route them to other subRouters.
type Router struct {
	Name         string
	added        bool
	subRouters   []*Router
	parentRouter *Router
}

type Dispatcher struct {
	Router
}

// Include includes a router into another router.
func (parentRouter *Router) Include(childRouter *Router) error {
	if parentRouter.added {
		return errors.New("already included")
	}
	if parentRouter == childRouter {
		return errors.New("can't add a router into itself")
	}
	currentRouter := *parentRouter
	for {
		if currentRouter.parentRouter == nil {
			break
		}
		if currentRouter.parentRouter == childRouter {
			return errors.New("can't add a parent router into child router")
		}
		if currentRouter.parentRouter != nil {
			currentRouter = *currentRouter.parentRouter
		}
	}
	childRouter.parentRouter = parentRouter
	newRouters := append(parentRouter.subRouters, childRouter)
	parentRouter.subRouters = newRouters
	parentRouter.added = true
	return nil
}

// New creates new empty Router.
func New(name string) *Router {
	return &Router{Name: name, subRouters: nil, parentRouter: nil}
}

func (parentRouter *Router) propagateEvent() {
	for i := 0; i < len(parentRouter.subRouters); i++ {

	}
}

// AddHandler adds handler for a certain update.
func (parentRouter *Router) AddHandler() {}
