package main

import (
	"context"
	"os"

	"github.com/concourse/porter/blobio"
	cwatch "github.com/concourse/porter/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"code.cloudfoundry.org/lager"
	"github.com/jessevdk/go-flags"
)

type PushCommand struct {
	PodName       string `required:"true" long:"pod-name" positional-args:"yes" description:"Pod to watch"`
	ContainerName string `required:"true" long:"container-name" positional-args:"yes" description:"Container to wait till completion"`

	SourcePath      string `required:"true" long:"source-path" description:"Location to fetch input blobs from within the bucket."`
	BucketURL       string `required:"true" long:"bucket-url" description:"Location of the bucket to fetch blobs from"`
	DestinationPath string `required:"true" long:"destination-path" description:"Path to inflate with fetched blobs"`
}

func (pc *PushCommand) Execute(args []string) error {
	logger.Debug("push-execute", lager.Data{
		"podname":       pc.PodName,
		"containername": pc.ContainerName,
	})

	config, err := rest.InClusterConfig()
	if err != nil {
		logger.Error("failed to retrieve cluster config", err)
		return err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Error("failed to create client with fetched config", err)
		return err
	}

	watcher := cwatch.ContainerWatcher{
		Client:        clientset,
		ContainerName: pc.ContainerName,
		PodName:       pc.PodName,
	}

	err = watcher.Wait(logger)
	if err != nil {
		return err
	}

	bucketConfig := blobio.BucketConfig{
		URL: pc.BucketURL,
	}

	err = blobio.Push(
		logger,
		context.Background(),
		bucketConfig,
		pc.SourcePath,
		pc.DestinationPath,
	)
	if err != nil {
		logger.Error("error-pushing", err)
		return err
	}

	return nil
}

var (
	logger lager.Logger
	Push   PushCommand
)

func main() {
	logger = lager.NewLogger("porter-push")
	sink := lager.NewWriterSink(os.Stderr, lager.DEBUG)
	logger.RegisterSink(sink)

	parser := flags.NewParser(&Push, flags.HelpFlag|flags.PrintErrors|flags.IgnoreUnknown)
	parser.NamespaceDelimiter = "-"

	_, err := parser.Parse()

	if err != nil {
		os.Exit(1)
	}

	err = Push.Execute(os.Args)
	if err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
