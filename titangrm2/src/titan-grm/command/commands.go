package command

import (
	"os"
	tile "tile-manager/command"

	"github.com/mitchellh/cli"

	"grm-service/command"
	"grm-service/service"

	collector "applications/data-collection/command"
	importer "data-importer/command"
	datamgr "data-manager/command"
	api "grm-api/command"
	label "grm-labelmgr/command"
	searcher "grm-searcher/command"
	storage "storage-manager/command"
	auth "titan-auth/comamnd"
	stat "titan-statistics/command"
)

// Commands returns the mapping of CLI commands for App. The meta
// parameter lets you set meta options for all commands.
func MakeCommands(metaPtr *command.Meta) map[string]cli.CommandFactory {
	if metaPtr == nil {
		metaPtr = new(command.Meta)
	}

	meta := *metaPtr
	if meta.Ui == nil {
		meta.Ui = &cli.BasicUi{
			Reader:      os.Stdin,
			Writer:      os.Stdout,
			ErrorWriter: os.Stderr,
		}
	}

	return map[string]cli.CommandFactory{
		// api
		service.GRMAPIService: func() (cli.Command, error) {
			return &api.APICommand{Meta: meta}, nil
		},
		// titan-auth
		service.TitanAuthService: func() (cli.Command, error) {
			return &auth.TitanAuthCommand{Meta: meta}, nil
		},
		// data-manager
		service.DataManagerService: func() (cli.Command, error) {
			return &datamgr.DataMgrCommand{Meta: meta}, nil
		},
		// data-importer
		service.DataImporterService: func() (cli.Command, error) {
			return &importer.ImporterCommand{Meta: meta}, nil
		},
		// storage-manager
		service.StorageManagerService: func() (cli.Command, error) {
			return &storage.StorageCommand{Meta: meta}, nil
		},
		// searcher
		service.SearcherService: func() (cli.Command, error) {
			return &searcher.SearcherCommand{Meta: meta}, nil
		},
		// stat
		service.StatService: func() (cli.Command, error) {
			return &stat.StatCommand{Meta: meta}, nil
		},
		// LabelMgrService
		service.LabelMgrService: func() (cli.Command, error) {
			return &label.LabelMgrCommand{Meta: meta}, nil
		},
		// data-collector
		service.DataCollection: func() (cli.Command, error) {
			return &collector.CollectorCommand{Meta: meta}, nil
		},
		// tile-manager
		service.TileManagerService: func() (cli.Command, error) {
			return &tile.TileMgrCommand{Meta: meta}, nil
		},
	}
}
