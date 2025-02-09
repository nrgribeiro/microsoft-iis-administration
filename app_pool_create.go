package iis

import (
	"context"
	"encoding/json"
)

func (client Client) CreateAppPool(ctx context.Context, name string, runtime string) (*ApplicationPool, error) {
	reqBody := CreateApplicationPoolRequest{
		Name: name,
		ManagedRuntimeVersion: runtime,
	}
	res, err := httpPost(ctx, client, "/api/webserver/application-pools", reqBody)
	if err != nil {
		return nil, err
	}
	var pool ApplicationPool
	err = json.Unmarshal(res, &pool)
	if err != nil {
		return nil, err
	}
	return &pool, nil
}

type CreateApplicationPoolRequest struct {
	Name string `json:"name"`
	ManagedRuntimeVersion string `json:"managed_runtime_version"`
}
