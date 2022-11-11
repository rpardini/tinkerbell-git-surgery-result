package backend

import (
	"context"
	"errors"
	"fmt"

	"github.com/tinkerbell/hegel/internal/backend/flatfile"
	"github.com/tinkerbell/hegel/internal/backend/kubernetes"
	"github.com/tinkerbell/hegel/internal/frontend/ec2"
)

// ErrMissingBackendConfig indicates New was called without a backend configuration.
var ErrMissingBackendConfig = errors.New("no backend configuration specified in options")

// ErrMultipleBackends indicates the backend Options contains more than one backend configuration.
var ErrMultipleBackends = errors.New("only one backend option can be specified")

// Client is an abstraction for all frontend clients. Each backend implementation should satisfy
// this interface.
type Client interface {
	ec2.Client
}

// New creates a backend instance for the configuration specified by opts. Consumers may only
// supply 1 backend configuration. If no backend configuration is supplied, it returns
// ErrMissingBackendConfig.
func New(ctx context.Context, opts Options) (Client, error) {
	if err := opts.validate(); err != nil {
		return nil, err
	}

	switch {
	case opts.Flatfile != nil:
		return flatfile.FromYAMLFile(opts.Flatfile.Path)

	case opts.Kubernetes != nil:
		config, err := kubernetes.NewConfig(
			opts.Kubernetes.Kubeconfig,
			opts.Kubernetes.KubeAPI,
			opts.Kubernetes.KubeNamespace,
		)
		if err != nil {
			return nil, fmt.Errorf("loading kubernetes config: %v", err)
		}

		kubeclient, err := kubernetes.New(config)
		if err != nil {
			return nil, fmt.Errorf("kubernetes client: %v", err)
		}
		kubeclient.WaitForCacheSync(ctx)

		return kubeclient, nil

	default:
		return nil, ErrMissingBackendConfig
	}
}

// Options contains all options for all backend implementations. Only one backend option can be
// specified at a time.
type Options struct {
	Flatfile   *FlatfileOptions
	Kubernetes *KubernetesOptions
}

func (o Options) validate() error {
	var count int

	if o.Flatfile != nil {
		count++
	}

	if o.Kubernetes != nil {
		count++
	}

	if count > 1 {
		return ErrMultipleBackends
	}

	return nil
}

// FlatFileOptions is the configuration for a flatfile backend.
type FlatfileOptions struct {
	// Path is a path to a YAML file containing a list of flatfile instances.
	Path string
}

// KubernetesOptions is the configuration for a Kubernetes backend.
type KubernetesOptions struct {
	// KubeAPI is the URL of the Kube API the Kubernetes client talks to.
	// Optional
	KubeAPI string

	// Kuberconfig is a path to a Kubeconfig file used by the Kubernetes client.
	// Optional
	Kubeconfig string

	// KubeNamespace is a namespace override to have Hegel use for reading resources.
	// Optional
	KubeNamespace string
}
