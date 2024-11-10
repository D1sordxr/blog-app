package api

type HTTPServerConfig struct {
	Address string `yaml:"address" env-required:"true"`
	Port    string `yaml:"port" env-required:"true"`
}

// HTTP standard lib

//main pkg
//
//server := api.NewHTTPServer(cfg.HTTPServerConfig.Address,
//	cfg.HTTPServerConfig.Port,
//)
//if err := server.Run(); err != nil {
//	log.Info("failed to start server", err)
//}

// this pkg
//
//func NewHTTPServer(addr string, port string) *HTTPServerConfig {
//	return &HTTPServerConfig{
//		Address: addr,
//		Port:    port,
//	}
//}
//
//func (s *HTTPServerConfig) Run() error {
//	router := http.NewServeMux()
//	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintln(w, "Welcome to the home page.")
//	})
//
//	v1 := http.NewServeMux()
//	v1.Handle("/api/v1/", http.StripPrefix("/api/v1/", router))
//
//	server := http.Server{
//		Addr:    s.Address + ":" + s.Port,
//		Handler: router,
//	}
//
//	log.Printf("Server has started at port %s", s.Port)
//
//	return server.ListenAndServe()
//}
