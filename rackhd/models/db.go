package models

import (
	"fmt"
	"github.com/docker/libkv"
	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/consul"
	"time"
)

var db store.Store

func init() {
	consul.Register()
}

// InitBackend creates the default backend capability
func InitBackend() {
	client := "localhost:8500"

	var err error
	db, err = libkv.NewStore(
		store.CONSUL,
		[]string{client},
		&store.Config{
			ConnectionTimeout: 10 * time.Second,
		},
	)
	if err != nil {
		fmt.Printf("Failed to init DB!\n")
	}
}
