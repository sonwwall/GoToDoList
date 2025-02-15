package v1

import (
	"GoToDoList/internal/global"
	"GoToDoList/internal/middleware"
	"GoToDoList/internal/repository"
	"GoToDoList/internal/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	userRepo := repository.NewUserRepository(global.Mysql)
	userService := service.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	userGroup := r.Group("/api/v1/user")
	{
		userGroup.POST("/register", userHandler.UserRegister)
		userGroup.POST("/login", userHandler.UserLogin)
		userGroup.PUT("/update", middleware.JwtAuthMiddleware(), userHandler.UpdateUserInfo)
	}

	//创建api,service,repository实例
	//为什么要创建这么多实例：
	//创建 api、service 和 repository 实例的原因是为了实现分层架构（Layered Architecture），
	//从而提高代码的可维护性、可测试性和灵活性。虽然最终这些实例都依赖于同一个 gorm.DB 数据库连接，但每个层都有其特定的责任：
	//Repository 层：负责与数据库进行交互，执行 CRUD 操作。它封装了对数据库的具体操作，使得上层代码不需要关心数据库的细节。
	//Service 层：负责业务逻辑处理。它调用 Repository 层的方法，并在其中实现复杂的业务规则和逻辑。
	//API 层（Handler 层）：负责处理 HTTP 请求和响应。它调用 Service 层的方法，并将结果返回给客户端。
	//这种分层设计的好处包括：
	//职责分离：每一层只关注自己的职责，降低了代码的复杂度。
	//易于测试：可以单独测试每一层的功能，而不需要依赖其他层。
	//灵活性和扩展性：如果需要更换数据库或修改业务逻辑，只需修改相应的层，而不会影响其他层

	listRepo := repository.NewListRepository(global.Mysql) //repository层
	listService := service.NewListService(listRepo)        //service层
	listHandler := NewListHandler(listService)             //api层

	//受保护的路由组
	listGroup := r.Group("/api/v1/list")
	listGroup.Use(middleware.JwtAuthMiddleware())
	{
		listGroup.POST("", listHandler.CreateList)
		listGroup.GET("/:id", listHandler.GetList)
		listGroup.PUT("/:id", listHandler.UpdateList)
		listGroup.DELETE("/:id", listHandler.DeleteList)
		listGroup.GET("/search", listHandler.SearchList)
		listGroup.GET("/search/tasks", listHandler.SearchListAndTasks)
	}

	taskRepo := repository.NewTaskRepository(global.Mysql)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := NewTaskHandler(taskService)

	taskGroup := r.Group("/api/v1/task")
	taskGroup.Use(middleware.JwtAuthMiddleware())
	{
		taskGroup.POST("", taskHandler.CreateTask)
		taskGroup.GET("/:id", taskHandler.GetTask)
		taskGroup.PUT("/:id", taskHandler.UpdateTask)
		taskGroup.DELETE("/:id", taskHandler.DeleteTask)
		taskGroup.GET("/search", taskHandler.SearchTask)

	}

	return r
}
