package worker

import (
	"bufio"
	"context"
	"fmt"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/packethost/pkg/log"
)

// DockerLogCapturer is a LogCapturer that can stream docker container logs to an io.Writer.
type DockerLogCapturer struct {
	dockerClient client.ContainerAPIClient
	logger       log.Logger
	writer       io.Writer
}

// getLogger is a helper function to get logging out of a context, or use the default logger.
func (l *DockerLogCapturer) getLogger(ctx context.Context) *log.Logger {
	loggerIface := ctx.Value(loggingContextKey)
	if loggerIface == nil {
		return &l.logger
	}
	lg, _ := loggerIface.(*log.Logger)

	return lg
}

// NewDockerLogCapturer returns a LogCapturer that can stream container logs to a given writer.
func NewDockerLogCapturer(cli client.ContainerAPIClient, logger log.Logger, writer io.Writer) *DockerLogCapturer {
	return &DockerLogCapturer{
		dockerClient: cli,
		logger:       logger,
		writer:       writer,
	}
}

// CaptureLogs streams container logs to the capturer's writer.
func (l *DockerLogCapturer) CaptureLogs(ctx context.Context, id string) {
	reader, err := l.dockerClient.ContainerLogs(ctx, id, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
		Timestamps: false,
	})
	if err != nil {
		l.getLogger(ctx).Error(err, "failed to capture logs for container ", id)
		return
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Fprintln(l.writer, scanner.Text())
	}
}
