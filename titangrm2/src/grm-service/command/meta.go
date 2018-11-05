package command

import (
	"bufio"
	"flag"
	"io"

	"github.com/mitchellh/cli"
)

// Meta contains the meta-options and functionality that nearly every
// command inherits.
type Meta struct {
	Ui cli.Ui

	// These are set by the command line flags.
	RegistryAddress  string
	ServiceAddress   string
	ServiceNamespace string
	DataDir          string
	ConfigDir        string
}

// FlagSetFlags is an enum to define what flags are present in the
// default FlagSet returned by Meta.FlagSet.
type FlagSetFlags uint

const (
	FlagSetNone    FlagSetFlags = 0
	FlagSetClient  FlagSetFlags = 1 << iota
	FlagSetDefault              = FlagSetClient
)

func (m *Meta) FlagSet(n string, fs FlagSetFlags) *flag.FlagSet {
	f := flag.NewFlagSet(n, flag.ContinueOnError)
	if fs&FlagSetClient != 0 {
		f.StringVar(&m.RegistryAddress, "registry_address", "consul:8500", "registry server address, env: GRM_REGISTRY_ADDRESS")
		f.StringVar(&m.ServiceAddress, "server_address", ":8080", "server address, env: GRM_SERVER_ADDRESS")
		f.StringVar(&m.ServiceNamespace, "server_namespace", "titangrm", "server namespace, env: GRM_SERVER_NAMESPACE")
		f.StringVar(&m.DataDir, "data_dir", "/opt/titangrm/data", "data directory, env: GRM_DATA_DIR")
		f.StringVar(&m.ConfigDir, "config_dir", "/opt/titangrm/config", "titangrm config directory, env: GRM_CONFIG_DIR")
		//f.StringVar(&m.ConfigDir, "config_dir", "Z:\\config", "titangrm config directory, env: GRM_CONFIG_DIR")
	}

	errR, errW := io.Pipe()
	errScanner := bufio.NewScanner(errR)
	go func() {
		for errScanner.Scan() {
			m.Ui.Error(errScanner.Text())
		}
	}()
	f.SetOutput(errW)

	return f
}
