package cmd

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-logr/logr"
	"github.com/peterbourgon/ff/v4"
	"github.com/peterbourgon/ff/v4/ffhelp"
	"github.com/tinkerbell/tinkerbell/backend/kube"
	"github.com/tinkerbell/tinkerbell/cmd/flag"
	"github.com/tinkerbell/tinkerbell/hegel"
	"github.com/tinkerbell/tinkerbell/smee"
	"github.com/tinkerbell/tinkerbell/tink/server"
	"golang.org/x/sync/errgroup"
)

func Execute(ctx context.Context, args []string) error {
	globals := &flag.GlobalConfig{
		BackendKubeConfig: func() string {
			hd, err := os.UserHomeDir()
			if err != nil {
				return ""
			}
			return filepath.Join(hd, ".kube", "config")
		}(),
		PublicIP:         detectPublicIPv4(),
		EnableSmee:       true,
		EnableHegel:      true,
		EnableTinkServer: true,
	}
	s := &flag.SmeeConfig{
		Config: smee.NewConfig(smee.Config{}, detectPublicIPv4()),
	}
	h := &flag.HegelConfig{
		Config:   hegel.NewConfig(hegel.Config{}, fmt.Sprintf("%s:%d", detectPublicIPv4().String(), 50061)),
		BindAddr: detectPublicIPv4(),
		BindPort: 50061,
	}
	ts := &flag.TinkServerConfig{
		Config:   server.NewConfig(),
		BindAddr: detectPublicIPv4(),
		BindPort: 42113,
	}

	gfs := ff.NewFlagSet("globals")
	sfs := ff.NewFlagSet("smee - DHCP and iPXE service").SetParent(gfs)
	hfs := ff.NewFlagSet("hegel - metadata service").SetParent(sfs)
	tfs := ff.NewFlagSet("tink server - Workflow server").SetParent(hfs)
	flag.RegisterGlobal(&flag.Set{FlagSet: gfs}, globals)
	flag.RegisterSmeeFlags(&flag.Set{FlagSet: sfs}, s)
	flag.RegisterHegelFlags(&flag.Set{FlagSet: hfs}, h)
	flag.RegisterTinkServerFlags(&flag.Set{FlagSet: tfs}, ts)

	cli := &ff.Command{
		Name:     "tinkerbell",
		Usage:    "tinkerbell [flags]",
		LongHelp: "Tinkerbell stack.",
		Flags:    tfs,
	}

	if err := cli.Parse(args, ff.WithEnvVarPrefix("TINKERBELL")); err != nil {
		e := errors.New(ffhelp.Command(cli).String())
		if !errors.Is(err, ff.ErrHelp) {
			e = fmt.Errorf("%w\n%s", e, err)
		}

		return e
	}

	// Smee
	s.Convert(&globals.TrustedProxies)

	// Hegel
	h.Convert(&globals.TrustedProxies)

	// Tink Server
	ts.Convert()

	log := defaultLogger(globals.LogLevel)
	log.Info("starting tinkerbell", "tink-server", fmt.Sprintf("%+v", ts), "tink-server-config", fmt.Sprintf("%+v", ts.Config))

	switch globals.Backend {
	case "kube":
		b, err := newKubeBackend(ctx, globals.BackendKubeConfig, "", globals.BackendKubeNamespace, enabledIndexes(globals.EnableSmee, globals.EnableHegel, globals.EnableTinkServer))
		if err != nil {
			return fmt.Errorf("failed to create kube backend: %w", err)
		}
		s.Config.Backend = b
		h.Config.BackendEc2 = b
		h.Config.BackendHack = b
		ts.Config.Backend = b
	case "file":
		b, err := newFileBackend(ctx, log, globals.BackendFilePath)
		if err != nil {
			return fmt.Errorf("failed to create file backend: %w", err)
		}
		s.Config.Backend = b
	case "none":
		b := newNoopBackend()
		s.Config.Backend = b
		h.Config.BackendEc2 = b
		h.Config.BackendHack = b
	default:
		return fmt.Errorf("unknown backend %q", globals.Backend)
	}

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		if globals.EnableSmee {
			return s.Config.Start(ctx, log.WithValues("service", "smee"))
		}
		log.Info("smee service is disabled")
		return nil
	})

	g.Go(func() error {
		if globals.EnableHegel {
			return h.Config.Start(ctx, log.WithValues("service", "hegel"))
		}
		log.Info("hegel service is disabled")
		return nil
	})

	g.Go(func() error {
		if globals.EnableTinkServer {
			return ts.Config.Start(ctx, log.WithValues("service", "tink-server"))
		}
		log.Info("tink server service is disabled")
		return nil
	})

	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	if !globals.EnableSmee && !globals.EnableHegel {
		return errors.New("all services are disabled")
	}

	return nil
}

func enabledIndexes(smeeEnabled, hegelEnabled, tinkServerEnabled bool) map[kube.IndexType]kube.Index {
	var idxs map[kube.IndexType]kube.Index

	if smeeEnabled {
		idxs = flag.KubeIndexesSmee
	}
	if hegelEnabled {
		for k, v := range flag.KubeIndexesHegel {
			idxs[k] = v
		}
	}
	if tinkServerEnabled {
		for k, v := range flag.KubeIndexesTinkServer {
			idxs[k] = v
		}
	}
	return idxs
}

// defaultLogger uses the slog logr implementation.
func defaultLogger(level int) logr.Logger {
	// source file and function can be long. This makes the logs less readable.
	// for improved readability, truncate source file to last 3 parts and remove the function entirely.
	customAttr := func(_ []string, a slog.Attr) slog.Attr {
		if a.Key == slog.SourceKey {
			ss, ok := a.Value.Any().(*slog.Source)
			if !ok || ss == nil {
				return a
			}
			ss.Function = ""
			p := strings.Split(ss.File, "/")
			// log the file path from tinkerbell/tinkerbell to the end.
			var idx int
			for i, v := range p {
				if v == "tinkerbell" {
					idx = i
					break
				}
			}
			ss.File = filepath.Join(p[idx:]...)

			return a
		}

		// This changes the slog.Level string representation to an integer.
		// This makes it so that the V-levels passed in to the CLI show up as is in the logs.
		if a.Key == slog.LevelKey {
			b, ok := a.Value.Any().(slog.Level)
			if !ok {
				return a
			}
			a.Value = slog.Float64Value(math.Abs(float64(b)))
			return a
		}

		return a
	}
	opts := &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.Level(-level),
		ReplaceAttr: customAttr,
	}
	log := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	return logr.FromSlogHandler(log.Handler())
}
