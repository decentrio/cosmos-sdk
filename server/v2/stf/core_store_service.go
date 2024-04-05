package stf

import (
	"context"

	"cosmossdk.io/core/container"
	"cosmossdk.io/core/store"
)

var _ store.KVStoreService = (*storeService)(nil)

func NewKVStoreService(address []byte) store.KVStoreService {
	return storeService{actor: address}
}

func NewMemoryStoreService(address []byte) store.MemoryStoreService {
	return storeService{actor: address}
}

type storeService struct {
	actor []byte
}

func (s storeService) OpenKVStore(ctx context.Context) store.KVStore {
	state, err := ctx.(*executionContext).State.GetWriter(s.actor)
	if err != nil {
		panic(err)
	}
	return state
}

func (s storeService) OpenMemoryStore(ctx context.Context) store.KVStore {
	return s.OpenKVStore(ctx)
}

func (s storeService) OpenContainer(ctx context.Context) container.Service {
	return ctx.(*executionContext).Cache.GetContainer(s.actor)
}
