package main

import (
	"bytes"
	"fmt"
	"io"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"os"
	"time"

	"code.cloudfoundry.org/lager"
	cwatch "github.com/concourse/porter/watch"
	"github.com/jessevdk/go-flags"
)

var (
	sourceFile *os.File
)

type LogStreamCommand struct {
	SourcePath    string `required:"true" long:"source-path" description:"Location of logfile to stream out."`
	PodName       string `required:"true" long:"pod-name" positional-args:"yes" description:"Pod to watch"`
	ContainerName string `required:"true" long:"container-name" positional-args:"yes" description:"Container to wait till completion"`
}

func streamLogfile(filepath string, EOFbackoff time.Duration, logger lager.Logger) {
	var offset int64 = 0
	var err error

	for {
		if sourceFile == nil {
			sourceFile, err = os.Open(filepath)
			if err != nil {
				continue
			}
		}

		_, err := sourceFile.Seek(offset, 0)
		if err != nil {
			return
		}

		b := make([]byte, 500)

		count, err := sourceFile.Read(b)
		// we don't want to bail on EOF while the file is being actively appended to
		if err == io.EOF {
			time.Sleep(EOFbackoff)
			continue
		} else if err != nil {
			logger.Error("streaming source file", err)
		}

		// this byte array will be right-padded with 0s
		// those get printed to stdout as newlines so let's discard them
		fmt.Println(string(bytes.TrimRight(b, "\x00")))

		// todo: listen for sigkill signal?
		offset = offset + int64(count)
	}
}

func (pc *LogStreamCommand) Execute(args []string) error {
	go streamLogfile(pc.SourcePath, 100 * time.Millisecond, logger)

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

	return nil
}

var (
	logger    lager.Logger
	LogStream LogStreamCommand
)

func main() {
	logger = lager.NewLogger("pod-helper-log-stream")
	sink := lager.NewWriterSink(os.Stderr, lager.INFO)
	logger.RegisterSink(sink)

	parser := flags.NewParser(&LogStream, flags.HelpFlag|flags.PrintErrors|flags.IgnoreUnknown)
	parser.NamespaceDelimiter = "-"

	_, err := parser.Parse()
	if err != nil {
		logger.Error("log-stream parsing failed", err)
		os.Exit(1)
	}

	err = LogStream.Execute(os.Args)
	if err != nil {
		logger.Error("log-stream execution fail", err)
		os.Exit(1)
	}

	os.Exit(0)
}
