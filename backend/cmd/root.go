package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/OmegaVoid/omega-inventory/pkg/model"
	"github.com/spf13/viper"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands TODO
var rootCmd = &cobra.Command{
	Use:   "omega-inventory",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		o := orm.NewOrm()

		err := orm.RunSyncdb("default", true, true)
		if err != nil {
			log.Fatal().Err(err).Msg("sync database")
		}

		cat := model.PartCategory{Name: "#"}
		unit := model.PartMeasurementUnit{
			Name:      "Quantity",
			ShortName: "qty",
		}
		locCat := model.StorageLocationCategory{Name: "#"}
		loc := model.StorageLocation{
			Name:     "root",
			Category: &locCat,
		}
		locCat2 := model.StorageLocationCategory{Name: "test", Parent: &locCat}
		part := model.Part{Name: "test", Category: &cat, Unit: &unit, StorageLocation: &loc}

		_, err = o.Insert(&locCat)
		if err != nil {
			log.Fatal().Err(err).Msg("insert locCat")
		}

		_, err = o.Insert(&locCat2)
		if err != nil {
			log.Fatal().Err(err).Msg("insert locCat2")
		}

		_, err = o.Insert(&loc)
		if err != nil {
			log.Fatal().Err(err).Msg("insert loc")
		}
		_, err = o.Insert(&cat)
		if err != nil {
			log.Fatal().Err(err).Msg("insert cat")
		}
		_, err = o.Insert(&unit)
		if err != nil {
			log.Fatal().Err(err).Msg("insert unit")
		}
		_, err = o.Insert(&part)
		if err != nil {
			log.Fatal().Err(err).Msg("insert part")
		}

		fmt.Println(locCat.Children)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	cobra.OnInitialize(initConfig)

	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		log.Fatal().Err(err).Msg("register database driver")
	}

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.omega-inventory.yaml)")
	rootCmd.PersistentFlags().StringP("database", "db", "", "database connection string")
	rootCmd.PersistentFlags().StringP("database_type", "db_type", "", "database type, supported value:\npostgres")
	err = viper.BindPFlag("database", rootCmd.PersistentFlags().Lookup("database"))
	if err != nil {
		log.Fatal().Err(err).Msg("bind config value database to flag")
	}
	err = viper.BindPFlag("database_type", rootCmd.PersistentFlags().Lookup("database_type"))
	if err != nil {
		log.Fatal().Err(err).Msg("bind config value database_type to flag")
	}
	viper.SetDefault("database", "user=gorm dbname=gorm password=gorm sslmode=disable")
	viper.SetDefault("database_type", "postgres")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	err = orm.RegisterDataBase("default", viper.GetString("database_type"), viper.GetString("database"))
	if err != nil {
		log.Fatal().Err(err).Msg("register database")
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".omega-inventory" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".omega-inventory")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Info().Str("config_file", viper.ConfigFileUsed()).Msg("using config file")
	}
}
