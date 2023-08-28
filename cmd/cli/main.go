package main

import (
	"github.com/WanderningMaster/ssetup.git/internal/render"
	"github.com/WanderningMaster/ssetup.git/internal/store"
)

func main() {
	store.SetupStore()
	render.Loop()
}
