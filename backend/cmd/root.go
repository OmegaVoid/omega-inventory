package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/OmegaVoid/omega-inventory/internal/generated"
	"github.com/OmegaVoid/omega-inventory/internal/gql_server"
	"github.com/OmegaVoid/omega-inventory/pkg/graph"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"

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

		r := gql_server.GinServer(log.Logger)

		srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

		http.Handle("/", playground.Handler("GraphQL playground", "/query"))
		http.Handle("/query", srv)
		log.Info().Msgf("connect to http://localhost:%v/ for GraphQL playground", viper.GetInt("graphql_port"))

		log.Err(r.Run(fmt.Sprint(":", viper.GetInt("graphql_port")))).Msg("running GraphQL Playground")
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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.omega-inventory.yaml)")
	rootCmd.PersistentFlags().IntP("graphql_port", "", 0, "Port to run GraphQL server on")
	rootCmd.PersistentFlags().StringP("database", "", "", "database connection string")
	rootCmd.PersistentFlags().StringP("file_serve", "", "", "url where files are served")

	if err := viper.BindPFlag("database", rootCmd.PersistentFlags().Lookup("database")); err != nil {
		log.Fatal().Err(err).Msg("bind config value database to flag")
	}
	if err := viper.BindPFlag("graphql_port", rootCmd.PersistentFlags().Lookup("graphql_port")); err != nil {
		log.Fatal().Err(err).Msg("bind config value graphql_port to flag")
	}
	if err := viper.BindPFlag("file_serve", rootCmd.PersistentFlags().Lookup("file_serve")); err != nil {
		log.Fatal().Err(err).Msg("bind config value file_serve to flag")
	}
	viper.SetDefault("database", "postgres://gorm:gorm@localhost:5432/gorm")
	viper.SetDefault("graphql_port", 8080)
	viper.SetDefault("file_serve", "http://localhost:8080/files")
	viper.AutomaticEnv()

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

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
		viper.SetConfigType("hcl")
		viper.SetConfigName(".omega-inventory")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Info().Str("config_file", viper.ConfigFileUsed()).Msg("using config file")
	}

}
