package command

import (
	"flag"
	"fmt"
	"math/rand"

	"github.com/paisit04/building-modern-cli/audiofile/internal/interfaces"
)

func NewRandomCommand(client interfaces.Client) *RandomCommand {
	gc := &RandomCommand{
		fs:     flag.NewFlagSet("random", flag.ContinueOnError),
		client: client,
	}
	gc.fs.StringVar(&gc.flag, "flag", "", "string flag for random command")
	return gc
}

type RandomCommand struct {
	fs     *flag.FlagSet
	flag   string
	client interfaces.Client
}

func (cmd *RandomCommand) Name() string {
	return cmd.fs.Name()
}

func (cmd *RandomCommand) ParseFlags(flags []string) error {
	return cmd.fs.Parse(flags)
}

func (cmd *RandomCommand) Run() error {
	fmt.Println(rand.Intn(100))
	return nil
}
