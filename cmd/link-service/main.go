package link_service

func main() {
	// Загрузка конфигурации
	cfg, err := config.GetConfigFromEnv()
	if err != nil {
		log.Fatalf("Failed to load configuration: %s\n", err.Error())
	}

	// Инициализация логгера
	zapLogger := logger.NewClientZapLogger(cfg.LogLevel, cfg.ServiceName)

	// Запуск сервера
	if err = entrypoint.Run(cfg, zapLogger); err != nil {
		zapLogger.Fatal("Run server failed: %s\n", zap.Error(err))
	}
}
