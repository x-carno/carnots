package http

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	router *fiber.App
	port   int
}

func NewServer() *Server {
	s := &Server{
		router: fiber.New(),
		port:   cfg.Port,
	}
	return s
}

func (s *Server) ListenHttp() {
	s.router.Use(cors.New())

	s.router.Get("/healthz", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(&fiber.Map{
			"data": "ok",
		})
	})

	s.router.Post("/api/v1/push", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(&fiber.Map{
			"data": "ok",
		})
	})

	s.router.Listen(fmt.Sprintf(":%d", s.port))
}
