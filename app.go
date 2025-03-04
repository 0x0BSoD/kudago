package main

import (
	"context"
	"encoding/json"
	"kudago/pkg/kubernetes"
)

type App struct {
	ctx   context.Context
	Kuber *kubernetes.Client
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	k8sClient, err := kubernetes.New()
	if err != nil {
		panic(err)
	}

	a.Kuber = k8sClient
	a.ctx = ctx
}

func (a *App) GetContexts() (string, error) {
	data, err := json.Marshal(a.Kuber.Config.Contexts)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
