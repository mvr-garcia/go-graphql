package cmd

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/mvr-garcia/go-graphql/config"
	"github.com/mvr-garcia/go-graphql/internal/infra"
	"github.com/mvr-garcia/go-graphql/internal/ui/graph"
	"github.com/spf13/cobra"
	"github.com/vektah/gqlparser/v2/ast"
)

var graphqlCmd = &cobra.Command{
	Use:   "graphql-api",
	Short: "Start the graphql",
	Run: func(cmd *cobra.Command, args []string) {
		// Load config
		config := config.LoadConfig()

		// Get DB connection
		db, err := infra.GetDB(config.Database.Driver, config.Database.DSN)
		if err != nil {
			log.Fatal("failed to connect to db:", err)
		}

		// Initialize repositories
		categoryRepo := infra.NewCategoryAdapter(db)
		courseRepo := infra.NewCourseAdapter(db)

		// Start GraphQL server
		port := config.Port

		srv := handler.New(graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{
					CategoryRepo: categoryRepo,
					CourseRepo:   courseRepo,
				},
			},
		))

		srv.AddTransport(transport.Options{})
		srv.AddTransport(transport.GET{})
		srv.AddTransport(transport.POST{})

		srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

		srv.Use(extension.Introspection{})
		srv.Use(extension.AutomaticPersistedQuery{
			Cache: lru.New[string](100),
		})

		http.Handle("/", playground.Handler("GraphQL playground", "/query"))
		http.Handle("/query", srv)

		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
		log.Fatal(http.ListenAndServe(":"+port, nil))
	},
}

func init() {
	rootCmd.AddCommand(graphqlCmd)
}
