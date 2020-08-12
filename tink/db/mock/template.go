package mock

import (
	"context"

	"github.com/golang/protobuf/ptypes/timestamp"
	uuid "github.com/satori/go.uuid"
)

// CreateTemplate creates a new workflow template
func (d DB) CreateTemplate(ctx context.Context, name string, data string, id uuid.UUID) error {
	return nil
}

// GetTemplate returns a workflow template
func (d DB) GetTemplate(ctx context.Context, id string) (string, string, error) {
	return "", "", nil
}

// DeleteTemplate deletes a workflow template
func (d DB) DeleteTemplate(ctx context.Context, name string) error {
	return nil
}

// ListTemplates returns all saved templates
func (d DB) ListTemplates(fn func(id, n string, in, del *timestamp.Timestamp) error) error {
	return nil
}

// UpdateTemplate update a given template
func (d DB) UpdateTemplate(ctx context.Context, name string, data string, id uuid.UUID) error {
	return nil
}
