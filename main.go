package main

import (
	fxzerolog "github.com/efectn/fx-zerolog"
	"go-fiber-api/app/database"
	"go-fiber-api/app/middleware"
	"go-fiber-api/app/router"
	main2 "go-fiber-api/config/config"
	"go-fiber-api/config/logger"
	"go-fiber-api/config/webserver"
	"go.uber.org/fx"
)

func main() {

	fx.New(
		/* provide patterns */
		// config
		fx.Provide(main2.NewConfig),
		// logging
		fx.Provide(logger.NewLogger),
		// fiber
		fx.Provide(webserver.NewFiber),
		// database
		fx.Provide(database.NewDatabase),
		// middleware
		fx.Provide(middleware.NewMiddleware),
		// router
		fx.Provide(router.NewRouter),

		// provide modules
		//article.NewArticleModule,
		//auth.NewAuthModule,

		// start aplication
		fx.Invoke(webserver.Start),

		// define logger
		fx.WithLogger(fxzerolog.Init()),
	).Run()

	//// INITIAL DATABASE
	//database.DatabaseInit()
	////migration.RunMigration()
	//
	//app := fiber.New()
	//
	//// INITIAL ROUTE
	//route.MainRouteInit(app)
	//route.UserRouteInit(app)
	//
	//err := app.Listen(":8800")
	//if err != nil {
	//	return
	//}
}
