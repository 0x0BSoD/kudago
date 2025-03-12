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

type NodesResult struct {
	Total    int `json:"total"`
	Ready    int `json:"ready"`
	CpuTotal int `json:"cpu_total"`
	CpuUsed  int `json:"cpu_used"`
	MemTotal int `json:"mem_total"`
	MemUsed  int `json:"mem_used"`
}

func (a *App) Nodes() (string, error) {
	nodes, err := a.Kuber.GetNodes(a.ctx)
	if err != nil {
		return "", err
	}

	ready := 0
	var cpuTotal, cpuUsed, memTotal, memUsed int64
	for _, n := range nodes {
		cpuTotal += n.Status.Capacity.Cpu().MilliValue() / 1000 // millicores to cores
		cpuUsed += n.Status.Capacity.Cpu().MilliValue() / 1000  // millicores to cores
		memTotal += n.Status.Capacity.Memory().Value()
		memUsed += n.Status.Capacity.Memory().Value()

		unschedulable := n.Spec.Unschedulable
		for _, c := range n.Status.Conditions {
			if c.Type == "Ready" && c.Status == "True" && !unschedulable {
				ready++
			}
		}
	}

	data, err := json.Marshal(NodesResult{
		Total:    len(nodes),
		Ready:    ready,
		MemUsed:  int(memUsed),
		MemTotal: int(memTotal),
		CpuTotal: int(cpuTotal),
		CpuUsed:  int(cpuUsed),
	})
	if err != nil {
		return "", err
	}

	return string(data), nil
}
