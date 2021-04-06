package iis

import (
	"context"
	"fmt"
)

func (client Client) ReadFile(ctx context.Context, id string) (*File, error) {
	url := fmt.Sprintf("/api/files/%s", id)
	var site File
	if err := getJson(ctx, client, url, &site); err != nil {
		return nil, err
	}
	return &site, nil
}
