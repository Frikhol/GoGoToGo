package entrypoint

import (
	"go.uber.org/zap"
	"google.golang.org/grpc/reflection"
	"link-service/internal/config"
	hh "link-service/internal/handlers/health"
	s "link-service/internal/server"
	pbhh "link-service/proto/health_service"
	"os"
	"os/signal"
)

func Run(cfg *config.Config, logger *zap.Logger) error {
	server := s.NewServer(logger)
	reflection.Register(server.Server)

	healthService := hh.NewHealthHandler()
	pbhh.RegisterHealthServiceServer(server.Server, healthService)

	go func() {
		if err := server.Run(cfg.GRPCPort); err != nil {
			logger.Info("Failed to run server", zap.Error(err))
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	server.Stop()

	return nil
}
