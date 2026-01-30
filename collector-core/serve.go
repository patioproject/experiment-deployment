package collector

import (
	"context"

	"innoobijr/collector/types"
	"github.com/gorilla/mux"
)

var r *mux.Router

const NameExpression = "-a-zA-Z_0-9."

func init() {
	r = mux.NewRouter()
}

func Router() *mux.Router {
	return r
}

// Load handler into the correct route spec. This blocks
func Serve(ctx context.Context, handlers *types.Handlers, config *types.Config) {
	// if we have basic auth enabled, decorate the handler iwth authentication

	// set endpoingts

	// initialize the proxy handler
	////proxyHandler := handlers.FunctionProxy

	////r.HandleFunc("/function/{name:["+NameExpression+"]+}", proxyHandler)
	/// open endpoing to be reverse proxied

	////readTimeout := config.ReadTimeout
	////writeTimeout := config.WriteTimeout

	////port := 8080
	////if config.TCPPort != nil {
	////	port = *config.TCPPort
	////}

	////s := &http.Server{
	////	Addr:           fmt.Sprintf(":%d", port),
	////	ReadTimeout:    readTimeout,
	////	WriteTimeout:   writeTimeout,
	////	MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	////	Handler:        r,
	////}

	// Start the server in a goroutine
	////go func() {
	////	if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
	////		log.Fatal(err)
	////	}
	////}()

	// Shutdown server when context is done
	////<-ctx.Done()
	////if err := s.Shutdown((context.Background())); err != nil {
	////	log.Printf("Failed to shut down provider gracefully: %s", err)
	////}
}
