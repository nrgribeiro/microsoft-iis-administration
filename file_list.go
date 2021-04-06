package iis

import "context"

type FileListItem struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type FileListResponse struct {
	Files []FileListItem `json:"files"`
}

func (client Client) ListFiles(ctx context.Context) ([]FileListItem, error) {
	var res FileListResponse
	err := getJson(ctx, client, "/api/files", &res)
	if err != nil {
		return nil, err
	}
	return res.Files, nil
}
