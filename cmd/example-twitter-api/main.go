package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/coreos/pkg/flagutil"
	"github.com/taguchi-so/example-twitter-api/adapter"
	"github.com/taguchi-so/example-twitter-api/infra"
	"github.com/taguchi-so/example-twitter-api/usecase"
)

var (
	name     = "example-twitter-api"
	revision = ""
)

func main() {
	flags := struct {
		consumerKey    string
		consumerSecret string
	}{}

	flag.StringVar(&flags.consumerKey, "consumer-key", "", "Twitter Consumer Key")
	flag.StringVar(&flags.consumerSecret, "consumer-secret", "", "Twitter Consumer Secret")
	flag.Parse()
	flagutil.SetFlagsFromEnv(flag.CommandLine, "TW")

	aUseCase, err := initializeAUseCase(
		flags.consumerKey,
		flags.consumerSecret,
	)
	if err != nil {
		print("invalid parameters", err)
		os.Exit(1)
	}

	task := flag.Arg(0)
	switch task {
	case "find":
		aUseCase.FindImage()
		break
	case "send":
		aUseCase.SendRequest()
		break
	case "receive":
		aUseCase.ReceiveReplay()
		break
	case "help":
		usage()
		break
	default:
		usage()
		os.Exit(1)
	}
	os.Exit(0)
}

func initializeAUseCase(consumerKey, consumerSecret string) (usecase.A, error) {
	config, err := infra.NewTwitterConf(consumerKey, consumerSecret)
	if err != nil {
		return nil, err
	}
	client := infra.NewTwitterClientWithNoContext(config)
	twitterUserService := adapter.NewTwitterUser(client)
	return usecase.NewA(
		twitterUserService,
	), nil
}

func usage() {
	fmt.Printf(`
Usage:	%s [OPTIONS] COMMAND

twitter api sample command

Options:
	--consumer-key					twitter consumer-key
	--consumer-secret				twitter consumer-secret

Commands:
	find                    find image
	send                    send reply to image tweet
	receive                 receive reply message
	help                    help

	githash[%s]
	`, name, revision)
}
