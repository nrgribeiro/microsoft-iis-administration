package iis

import (
	"context"
	"fmt"
)

func (client Client) DeleteFile(ctx context.Context, id string) error {
	url := fmt.Sprintf("/api/files/%s", id)
	return httpDelete(ctx, client, url)
}
