package main

import (
	"context"
	"encoding/json"
	v1 "k8s.io/api/core/v1"
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

type DashboardData struct {
	Total     int `json:"total"`
	Ready     int `json:"ready"`
	CpuTotal  int `json:"cpu_total"`
	CpuUsed   int `json:"cpu_used"`
	MemTotal  int `json:"mem_total"`
	MemUsed   int `json:"mem_used"`
	PodsTotal int `json:"pods_total"`
	PodsReady int `json:"pods_ready"`
}

func (a *App) GetDashboardData() (string, error) {
	nodes, err := a.Kuber.GetNodes(a.ctx)
	if err != nil {
		return "", err
	}

	pods, err := a.Kuber.GetAllPods(a.ctx)
	if err != nil {
		return "", err
	}

	podsReady := 0
	for _, pod := range pods {
		for _, c := range pod.Status.Conditions {
			if c.Type == "Ready" && c.Status == "True" && pod.Status.Phase == v1.PodRunning {
				podsReady++
			}
		}
	}

	nodeReady := 0
	var cpuTotal, cpuUsed, memTotal, memUsed int64
	for _, n := range nodes {
		cpuTotal += n.Status.Capacity.Cpu().MilliValue() / 1000 // millicores to cores
		cpuUsed += n.Status.Capacity.Cpu().MilliValue() / 1000  // millicores to cores
		memTotal += n.Status.Capacity.Memory().Value()
		memUsed += n.Status.Capacity.Memory().Value()

		unschedulable := n.Spec.Unschedulable
		for _, c := range n.Status.Conditions {
			if c.Type == "Ready" && c.Status == "True" && !unschedulable {
				nodeReady++
			}
		}
	}

	data, err := json.Marshal(DashboardData{
		Total:     len(nodes),
		Ready:     nodeReady,
		MemUsed:   int(memUsed),
		MemTotal:  int(memTotal),
		CpuTotal:  int(cpuTotal),
		CpuUsed:   int(cpuUsed),
		PodsTotal: len(pods),
		PodsReady: podsReady,
	})
	if err != nil {
		return "", err
	}

	return string(data), nil
}
