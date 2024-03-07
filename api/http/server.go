package http

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"fibo/api/http/postcontroller"
	"fibo/internal/auth"
	"fibo/internal/base/crypto"
	"fibo/internal/category"
	"fibo/internal/post"
	"fibo/internal/user"
)

type Config interface {
	DetailedError() bool
	Address() string
}

type ServerOpts struct {
	UserUsecases   user.UserUsecases
	AuthService    auth.AuthService
	Crypto         crypto.Crypto
	Config         Config
	Post           post.PostUseCase
	Category       category.CatUseCase
	PostController postcontroller.PostController
}

func NewServer(opts ServerOpts) *Server {
	gin.SetMode(gin.ReleaseMode)

	server := &Server{
		engine:         gin.New(),
		config:         opts.Config,
		crypto:         opts.Crypto,
		userUsecases:   opts.UserUsecases,
		authService:    opts.AuthService,
		postUsecases:   opts.Post,
		catUsecases:    opts.Category,
		postcontroller: opts.PostController,
	}

	initRouter(server)

	return server
}

type Server struct {
	engine         *gin.Engine
	config         Config
	crypto         crypto.Crypto
	userUsecases   user.UserUsecases
	authService    auth.AuthService
	postUsecases   post.PostUseCase
	catUsecases    category.CatUseCase
	postcontroller postcontroller.PostController
}

func (s Server) Listen() error {
	fmt.Printf("API server listening at: %s\n\n", s.config.Address())
	return s.engine.Run(s.config.Address())
}
