package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"testTaskOne/internal/config"
	"testTaskOne/internal/delivery/http/server"
	"testTaskOne/internal/domain"
	"testTaskOne/internal/handler"
	"testTaskOne/internal/repository"
	"testTaskOne/internal/repository/postgres"
	"testTaskOne/internal/service"
	"time"
)

func Run() {
	conf, err := config.ConfInit()
	if err != nil {
		log.Fatalf("ConfInit(): %v", err)
	}

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     conf.PgHost,
		Port:     conf.PgPort,
		Username: conf.PgUser,
		Password: conf.PgPass,
		DBName:   conf.PgBase,
		SSLMode:  conf.PgSSLMode,
	})
	if err != nil {
		log.Fatalf("NewPostgresDB(): %v", err)
	}

	users := domain.NewPersonStorage()
	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(users, service)

	srv := server.NewServer(
		conf.ServerHost,
		conf.ServerPort,
		handler.InitRoutes())

	go func() {
		if err := srv.Run(); err != nil {
			log.Fatalf("srv.Run(): %v", err)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				if len(handler.Users) != 0 {
					if err := handler.PostUser(); err != nil {
						log.Printf("handler.PostUser(): %v", err)
					}
				}
			default:
				if len(handler.Users) == 25 {
					if err := handler.PostUser(); err != nil {
						log.Printf("handler.PostUser(): %v", err)
					}
					time.Sleep(200 * time.Millisecond)
				}
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Printf("Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("srv.Shutdown(): %v", err)
	}

	if err = db.Close(); err != nil {
		log.Fatalf("postresDB.Close(): %v", err)
	}
}
