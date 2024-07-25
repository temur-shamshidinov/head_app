package api

import (
	v1 "head_app/api/handlers/v1"
	log "head_app/pkg/logger"
	"head_app/storage"
	"head_app/storage/redis"

	"github.com/gin-gonic/gin"
)

type Options struct {
	Storage storage.StorageI
	Log     log.Log
	Cache   redis.RedisRepoI
}

func Api(opt Options) *gin.Engine {

	h := v1.NewHandler(v1.Handlers{Storage: opt.Storage, Log: opt.Log,Cache: opt.Cache})

	engine := gin.Default()

	api := engine.Group("/api")

	api.GET("/ping",h.Ping)

	own := api.Group("/own")
	{ 
		// SignUp
	
		own.POST("/sign-out",h.OwnSignOut)

		// Category
		own.POST("/category", h.CreateCategory)
		own.PUT("/category/:id",h.UpdateCategory)
		own.DELETE("/category/:id",h.DeleteCategory)

		//	SubCategory
		own.POST("/sub-category",h.CreateSubCategory)
		own.PUT("/sub-category/:id",h.UpdateSubCategory)
		own.DELETE("/sub-category/:id",h.DeleteSubCategory)

		// Article
		own.POST("/article",h.CreateArticle)
		own.PUT("/article/:id",h.UpdateSubCategory)
		own.DELETE("/article/:id",h.DeleteSubCategory)
	}

	vw := api.Group("/vw")
	{
		vw.POST("/log-out")
		vw.POST("/comment/:article_id")
	}

	pb := api.Group("/pb")
	{
		own.POST("/own/sing-in",h.OwnSignIn)  // in-proccess 
		pb.POST("/check-user",h.CheckUser)  // completed
		pb.POST("/check-otp/:id",h.GetCategory) // completed
		pb.POST("/sign-up",h.SignUp)            // completed
		pb.POST("/sign-in",h.SignIn)				// completed
		pb.GET("/categories",h.GetCategoriesList) // completed
		pb.GET("/categories/:id",h.GetCategory)   // completed
		pb.GET("/sub-categories",h.GetSubCategoriesList)
		pb.GET("/sub-categories/:id",h.GetSubCategory)
		pb.GET("/articles",h.GetArticleList)
		pb.GET("/articles/:id",h.GetArticle)
	}

	return engine
}
