package main

import (
	"donn/config"
	"donn/middleware"
	"donn/resource/review"
	"donn/resource/user"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"time"
)

func main() {

	router := gin.New()

	config.InitDb()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.DummyMiddle())
	// CORS middleware
	router.Use(middleware.CORSMiddleware())

	v1 := router.Group("/v1")
	{

		v1.POST("/login", nil)
		v1.POST("/logout", nil)
		v1.POST("/logout-all", nil)
		v1.POST("/refreshToken", nil)

		v1Review := v1.Group("/review")
		{
			v1Review.GET("", review.Index)
			v1Review.POST("", review.Create)

			v1Review.GET("/", review.Index)
			v1Review.GET("/:id", review.Show)
			v1Review.POST("/", review.Create)
			v1Review.PUT("/:id", review.Update)
			v1Review.PATCH("/:id", review.Patch)
			v1Review.DELETE("/:id", review.Delete)
		}
		v1User := v1.Group("/user")
		{
			v1User.GET("/:id", middleware.Auth0Middleware(), user.Index)
			v1User.PUT("/:id", middleware.Auth0Middleware(), user.Create)

			//			v1UserBrief := v1User.Group("/brief")
			//			{
			//				v1UserBrief.GET("/:id", middleware.Auth0Middleware(), user.Create)
			//			}

		}
		v1Spot := v1.Group("/spot")
		{
			v1Spot.GET("", nil)
			v1Spot.GET("/", nil)
			v1Spot.GET("/:id", nil)
		}
		v1Coupon := v1.Group("/coupon", nil)
		{
			v1Coupon.GET("", nil)

			v1Coupon.GET("/", nil)
			v1Coupon.GET("/area", nil)
			v1Coupon.GET("/category", nil)
		}
		v1Feed := v1.Group("/feed", nil)
		{
			v1Feed.GET("", nil)

			v1Feed.GET("/", nil)
			v1Feed.GET("/category", nil)
		}
		v1LocationBase := v1.Group("", nil)
		{
			v1LocationBase.GET("/exchange", nil)

			v1lsbGeo := v1LocationBase.Group("/geo", nil)
			{
				v1lsbGeo.GET("/latlng", nil)
				v1lsbGeo.GET("/address", nil)
			}

			v1lbsNearby := v1LocationBase.Group("/nearby", nil)
			{
				v1lbsNearby.GET("", nil)
				v1lbsNearby.GET("/", nil)
				v1lbsNearby.GET("/autocomplete", nil)
				v1lbsNearby.GET("/category", nil)
				v1lbsNearby.GET("/explore", nil)
				v1lbsNearby.GET("/search", nil)
			}

			v1lbsWeather := v1LocationBase.Group("/weather", nil)
			{
				v1lbsWeather.GET("/3h", nil)
				v1lbsWeather.GET("/d", nil)
			}

		}

	}

	//	router.Run(":8080")
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 5 << 20, // 5 MegaBytes
	}
	s.ListenAndServe()

}
