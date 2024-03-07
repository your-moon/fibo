package main

import (
	"context"
	"fmt"
	"log"

	"fibo/api/cli"
	"fibo/api/http"
	postControllerImpl "fibo/api/http/postcontroller/impl"
	authImpl "fibo/internal/auth/impl"
	cryptoImpl "fibo/internal/base/crypto/impl"
	databaseImpl "fibo/internal/base/database/impl"
	categoryImpl "fibo/internal/category/impl"
	postImpl "fibo/internal/post/impl"
	userImpl "fibo/internal/user/impl"
)

func main() {
	ctx := context.Background()
	parser := cli.NewParser()

	conf, err := parser.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	dbClient := databaseImpl.NewClient(ctx, conf.Database())

	err = dbClient.Connect()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(err)

	defer dbClient.Close()

	crypto := cryptoImpl.NewCrypto()
	dbService := databaseImpl.NewService(dbClient)

	userRepositoryOpts := userImpl.UserRepositoryOpts{
		ConnManager: dbService,
	}
	userRepository := userImpl.NewUserRepository(userRepositoryOpts)

	authServiceOpts := authImpl.AuthServiceOpts{
		Crypto:         crypto,
		Config:         conf.Auth(),
		UserRepository: userRepository,
	}
	authService := authImpl.NewAuthService(authServiceOpts)

	userUsecasesOpts := userImpl.UserUsecasesOpts{
		TxManager:      dbService,
		UserRepository: userRepository,
		Crypto:         crypto,
	}
	userUsecases := userImpl.NewUserUsecases(userUsecasesOpts)

	postRepositoryOpts := postImpl.PostRepositoryOpts{
		ConnManager: dbService,
	}

	postRepository := postImpl.NewPostRepository(postRepositoryOpts)

	postUsecasesOpts := postImpl.PostUsecaseOpts{
		PostRepository: postRepository,
		TxManager:      dbService,
	}

	postUsecases := postImpl.NewPostUsecase(postUsecasesOpts)

	catRepositoryOpts := categoryImpl.CatRepositoryOpts{
		ConnManager: dbService,
	}
	catRepository := categoryImpl.NewCatRepository(catRepositoryOpts)

	catUsecasesOpts := categoryImpl.CatUsecaseOpts{
		CatRepository: catRepository,
		TxManager:     dbService,
	}

	catUsecases := categoryImpl.NewCatUsecase(catUsecasesOpts)

	postControllerOpts := postControllerImpl.PostControllerOpts{
		PostUsecase: postUsecases,
		Config:      conf.HTTP(),
		CatUsecase:  catUsecases,
	}

	postController := postControllerImpl.NewPostController(postControllerOpts)

	serverOpts := http.ServerOpts{
		UserUsecases:   userUsecases,
		AuthService:    authService,
		Crypto:         crypto,
		Config:         conf.HTTP(),
		Post:           postUsecases,
		Category:       catUsecases,
		PostController: postController,
	}
	server := http.NewServer(serverOpts)

	log.Fatal(server.Listen())
}
