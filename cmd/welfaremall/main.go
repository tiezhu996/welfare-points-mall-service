package main

import (
	"fmt"

	"welfaremall/internal/config"
	"welfaremall/internal/model"
	"welfaremall/internal/service"
	"welfaremall/internal/store"
	"welfaremall/internal/worker"
)

func main() {
	cfg := config.Load()
	st := store.New()
	svc := service.New(st)
	_ = st.UpsertAccount(model.Account{ID: "emp-1", Points: 100})
	_ = svc.Spend("emp-1", 10)
	fmt.Printf("%s unhealthy=%d\n", cfg.AppName, worker.New(st).Run())
}
