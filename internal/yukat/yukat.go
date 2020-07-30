package yukat

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/14DENDIK/yukatbot/api/telegram"
	"github.com/14DENDIK/yukatbot/internal/yukat/config"
	"github.com/14DENDIK/yukatbot/internal/yukat/handlers"
	"github.com/14DENDIK/yukatbot/internal/yukat/store"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
)

// Server ...
type Server struct {
	config  *config.Config
	db      *pgx.Conn
	store   *store.Store
	handler *handlers.Handler
	logger  *zap.Logger
}

// New ...
func New(config *config.Config) (*Server, error) {
	s := &Server{
		config: config,
	}

	if err := s.configureDB(); err != nil {
		return nil, err
	}

	if err := s.configureLogger(); err != nil {
		return nil, err
	}

	s.store = store.New(s.db)
	s.handler = handlers.New(s.store, s.logger, s.config.Token)

	return s, nil
}

// Start ...
func (s *Server) Start() error {
	if err := s.setWebhook(); err != nil {
		return err
	}
	s.logger.Sugar().Infof("Webhook set on %s", s.config.WebhookURL)

	defer s.db.Close(context.Background())
	defer s.logger.Sync()

	s.logger.Sugar().Infof("Starting server on port %s ...", s.config.BindAddr)

	if err := http.ListenAndServe(s.config.BindAddr, http.HandlerFunc(s.handler.MainEntry)); err != nil {
		// s.logger.Sugar().Errorf("Failed to listen on port %s", s.config.BindAddr)
		return err
	}
	return nil
}

func (s *Server) configureDB() error {
	db, err := pgx.Connect(context.Background(), s.config.DatabaseURL)
	if err != nil {
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *Server) configureLogger() error {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return err
	}
	s.logger = logger
	return nil
}

func (s *Server) setWebhook() error {
	webhook := &telegram.WebhookInfo{}
	res, err := http.Get("https://api.telegram.org/bot" + s.config.Token + "/getWebhookInfo")
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(webhook); err != nil {
		return err
	}

	// Checks whether webhook url changed
	if webhook.Result.URL != s.config.WebhookURL {
		reqBody := &telegram.SetWebhook{
			URL: s.config.WebhookURL,
		}

		reqBytes, err := json.Marshal(reqBody)
		if err != nil {
			return err
		}

		res, err = http.Post(
			"https://api.telegram.org/bot"+s.config.Token+"/setWebhook",
			"application/json",
			bytes.NewBuffer(reqBytes),
		)
		if err != nil {
			return err
		}

		if res.StatusCode != http.StatusOK {
			return errors.New("Unexpected status " + res.Status)
		}
		s.logger.Sugar().Info("Webhook refreshed...")
	}
	return nil
}
