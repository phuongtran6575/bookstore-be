package di

import (
	"github.com/gin-gonic/gin"
	"github.com/phuongtran6575/bookstore-be.git/internal/routers" // 1. Thêm import này
	"gorm.io/gorm"
)

// 2. Thêm kiểu trả về *gin.Engine cho hàm
func SetupDI(db *gorm.DB) *gin.Engine {
	repositories := NewRepositories(db)
	services := NewServices(repositories)
	handlers := NewHandlers(services)

	// 3. Sử dụng routers.NewRouter thay vì NewRouter
	r := routers.NewRouter(
		handlers.ProductHandler,
		handlers.CategoryHandler,
		handlers.AuthorHandler,
	)

	return r.RegisterRoutes()
}
