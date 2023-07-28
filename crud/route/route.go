package route

import (
	repository "trunne-crud/repository"
	usecase "trunne-crud/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, gin *gin.Engine) {
	v1Routes := gin.Group("/v1")
	var (
		er = repository.NewEmployeeRepository(db)
		eu = usecase.NewEmployeeUsecase(er)
	)

	NewEmployeeRoute(v1Routes, eu)
}
