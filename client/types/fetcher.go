package types

import (
	"github.com/google/uuid"
)

type Kind string

var (
	OktaFetcher Kind = "okta"
)

type Runtime string

var (
	NativeRuntime Runtime = "native"
)

type Fetcher struct {
	// Fetcher Definition
	Name        string
	Description string
	GUID        uuid.UUID
	UID         int32
	Kind        Kind
	Runtime     Runtime
	Tags        []string
	Config      map[string]string
	Secrets     map[string]uuid.UUID
	// Fetcher State
	Running bool
	Enabled bool
}

type FetcherStorage struct {
	Items map[string]FetcherStorageItem
}

type FetcherStorageItemType string

var (
	FetcherIntType FetcherStorageItemType = "int"
)

type FetcherStorageItem struct {
	Value any
	Type  FetcherStorageItemType
}
