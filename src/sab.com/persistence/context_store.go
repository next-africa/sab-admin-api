package persistence

import "golang.org/x/net/context"

type ContextStore struct {
	gaeContext context.Context
}

func (store *ContextStore) SetContext(context context.Context) {
	store.gaeContext = context
}

func (store *ContextStore) GetContext() context.Context {
	return store.gaeContext
}
