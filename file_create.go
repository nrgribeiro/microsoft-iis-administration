package iis

import (
	"context"
	"encoding/json"
)

type CreateFileRequest struct {
	Name            string               `json:"name"`
	PhysicalPath    string               `json:"physical_path"`
	Type			string				 `json:"type"`
	Alias 			string				 `json:"alias"`
	Parent 			[]FileParent		 `json:"parent"`
}

func (client Client) CreateFile(ctx context.Context, req CreateFileRequest) (*File, error) {
	res, err := httpPost(ctx, client, "/api/files", req)
	if err != nil {
		return nil, err
	}
	var site File
	err = json.Unmarshal(res, &site)
	if err != nil {
		return nil, err
	}
	return &site, nil
}
