package iis

import (
	"context"
	"encoding/json"
	"fmt"
)

func (client Client) UpdateFile(ctx context.Context, update File) (*File, error) {
	url := fmt.Sprintf("/api/files/%s", update.ID)
	res, err := httpPatch(ctx, client, url, update)
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
