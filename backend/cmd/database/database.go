package database

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"fudjie.waizly/backend-test/config"

	pkg_config "fudjie.waizly/backend-test/library/config"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var MigrationCommand = &cobra.Command{
	Use:   "db",
	Short: "Database Migration",
	Run: func(c *cobra.Command, args []string) {
		DatabaseMigration()
	},
}

var (
	flags      = flag.NewFlagSet("db", flag.ExitOnError)
	configPath = flags.String("config", "config/file", "Config URL dir i.e. config/file")
	dir        = flags.String("dir", "internal/db/migration", "directory with migration files")
)

var (
	usageCommands = `
Commands:
  up [N]?              Migrate all or N up migrations
  goto [V]             Migrate the DB to a specific version
  down [N]?            Down all or N down migrations

For more features, use https://github.com/golang-migrate/migrate/tree/master/cmd/migrate`
)

func initLogger() {
	// initiate logger
}

func DatabaseMigration() {
	initLogger()

	cfg := &config.MainConfig{}
	err := pkg_config.ReadConfig(cfg, *configPath, "config")
	if err != nil {
		log.Fatal()
	}

	flags.Usage = usage
	flags.Parse(os.Args[2:])

	args := flags.Args()

	if len(args) == 0 {
		flags.Usage()
		return
	}

	var (
		m       *migrate.Migrate
		connstr string
	)

	connstr = cfg.DBMigration.MigrationConfig.Dsn

	m, err = migrate.New(
		fmt.Sprintf("file://%s", *dir),
		connstr,
	)
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
	}

	switch args[0] {
	case "up":
		if len(args) == 2 {
			step, err := strconv.Atoi(args[1])
			if err != nil {
				log.Fatal().Err(err).Msg(err.Error())
			}
			err = m.Steps(step)
			if err != nil {
				log.Fatal().Err(err).Msg(err.Error())
			}
		} else {
			err := m.Up()
			if err != nil {
				log.Fatal().Err(err).Msg(err.Error())
			}
		}
		return
	case "down":
		if len(args) == 2 {
			step, err := strconv.Atoi(args[1])
			if err != nil {
				log.Fatal().Err(err).Msg(err.Error())
			}
			err = m.Steps(step * -1)
			if err != nil {
				log.Fatal().Err(err).Msg(err.Error())
			}
		} else {
			err := m.Down()
			if err != nil {
				log.Fatal().Err(err).Msg(err.Error())
			}
		}
		return
	case "goto":
		if len(args) == 2 {
			step, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				log.Fatal().Err(err).Msg(err.Error())
			}
			err = m.Migrate(uint(step))
			if err != nil {
				log.Fatal().Err(err).Msg(err.Error())
			}
		} else {
			usage()
		}
		return
	}

	err = m.Up()
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
	}

	defer func() {
		sourceErr, dbErr := m.Close()
		if sourceErr != nil {
			log.Fatal().Err(sourceErr).Msg(sourceErr.Error())
		}
		if dbErr != nil {
			log.Fatal().Err(dbErr).Msg(dbErr.Error())
		}
	}()
}

func usage() {
	fmt.Println(usageCommands)
}
