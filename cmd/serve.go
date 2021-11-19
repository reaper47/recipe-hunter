package cmd

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/reaper47/recipya/internal/contexts"
	"github.com/reaper47/recipya/internal/router"
	_ "github.com/reaper47/recipya/internal/templates"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the web server",
	Long: `"Starts the web server."

The application will be accessible through your favorite 
web browser at the address specified when you run the command.
`,
	Run: func(cmd *cobra.Command, args []string) {
		srv := &http.Server{
			Addr:         "0.0.0.0:8080",
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
			IdleTimeout:  60 * time.Second,
			Handler:      router.New(),
		}

		go func() {
			log.Println("Serving on 0.0.0.0:8080")
			err := srv.ListenAndServe()
			if err != nil {
				log.Println(err)
			}
		}()

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c

		ctx, cancel := contexts.Timeout(10 * time.Second)
		defer cancel()

		srv.Shutdown(ctx)
		log.Println("Server shutting down")
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
