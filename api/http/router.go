package http

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"fibo/internal/auth"
	"fibo/internal/base/errors"
	"fibo/internal/base/request"
	"fibo/internal/category"
	"fibo/internal/post"
	"fibo/internal/user"
)

func initRouter(server *Server) {
	router := &router{
		Server: server,
	}

	router.init()
}

type router struct {
	*Server
}

func (r *router) init() {
	r.engine.Use(corsMiddleware())
	r.engine.Use(r.trace())
	r.engine.Use(r.recover())
	r.engine.Use(r.logger())

	// User routes
	userRoutes := r.engine.Group("/users")
	{
		userRoutes.POST("", r.addUser)
		userRoutes.GET("/me", r.authenticate, r.getMe)
		userRoutes.PUT("/me", r.authenticate, r.updateMe)
		userRoutes.PATCH("/me/password", r.authenticate, r.changeMyPassword)
		userRoutes.GET("/me/posts", r.authenticate, r.getMyPosts)
		userRoutes.GET("/all", r.authenticate, r.getAllUsers)
	}

	// Post routes
	postRoutes := r.engine.Group("/posts")
	{
		postRoutes.POST("", r.authenticate, r.postcontroller.AddPostC)
		postRoutes.GET("", r.getPosts)
		postRoutes.GET("/:id", r.getPostById)
		postRoutes.PUT("/:id", r.authenticate, r.updatePost)
		postRoutes.GET("/published", r.getPublishedPosts)
		postRoutes.GET("/me/likes", r.authenticate, r.getTotalLikesCountByUser)
	}

	// Category routes
	categoryRoutes := r.engine.Group("/categories")
	{
		categoryRoutes.POST("/add", r.authenticate, r.addCategory)
		categoryRoutes.GET("", r.getCategories)
		categoryRoutes.GET("/:id", r.getCategoryById)
	}

	r.engine.POST("/login", r.login)
	r.engine.NoRoute(r.methodNotFound)
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}

func (r *router) getTotalLikesCountByUser(c *gin.Context) {
	reqInfo := GetReqInfo(c)

	likes, err := r.postUsecases.GetTotalLikesCountByUser(c, reqInfo.UserId)
	if err != nil {
		ErrorResponse(err, nil, r.config.DetailedError()).Reply(c)
		return
	}

	OkResponse(likes).Reply(c)
}

func (r *router) addCategory(c *gin.Context) {
	var addCategoryDto category.AddCatDto

	if err := BindBody(&addCategoryDto, c); err != nil {
		ErrorResponse(err, nil, r.config.DetailedError()).Reply(c)
		return
	}

	category, err := r.catUsecases.AddCategory(c, addCategoryDto)
	if err != nil {
		ErrorResponse(err, nil, r.config.DetailedError()).Reply(c)
		return
	}

	OkResponse(category).Reply(c)
}

func (r *router) getCategoryById(c *gin.Context) {
	categoryId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ErrorResponse(err, nil, r.config.DetailedError()).Reply(c)
		return
	}

	category, err := r.catUsecases.GetByID(c, categoryId)
	if err != nil {
		ErrorResponse(err, nil, r.config.DetailedError()).Reply(c)
		return
	}

	OkResponse(category).Reply(c)
}

func (r *router) getCategories(c *gin.Context) {
	categories, err := r.catUsecases.GetCategories(c)
	if err != nil {
		ErrorResponse(err, nil, r.config.DetailedError()).Reply(c)
		return
	}

	OkResponse(categories).Reply(c)
}

func (r *router) login(c *gin.Context) {
	var loginUserDto auth.LoginUserDto

	if err := BindBody(&loginUserDto, c); err != nil {
		ErrorResponse(err, nil, r.config.DetailedError()).Reply(c)
		return
	}

	user, err := r.authService.Login(c, loginUserDto)
	if err != nil {
		ErrorResponse(err, nil, r.config.DetailedError()).Reply(c)
		return
	}

	OkResponse(user).Reply(c)
}

func (r *router) authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	userId, err := r.authService.VerifyAccessToken(token)
	if err != nil {
		response := ErrorResponse(err, nil, r.config.DetailedError())
		c.AbortWithStatusJSON(response.Status, response)
	}

	setUserId(c, userId)
}

func (r *router) addUser(c *gin.Context) {
	var addUserDto user.AddUserDto

	if err := BindBody(&addUserDto, c); err != nil {
		ErrorResponse(err, nil, r.config.DetailedError()).Reply(c)
		return
	}

	user, err := r.userUsecases.Add(contextWithReqInfo(c), addUserDto)
	if err != nil {
		ErrorResponse(err, nil, r.config.DetailedError()).Reply(c)
		return
	}

	OkResponse(user).Reply(c)
}

func (r *router) updateMe(c *gin.Context) {
	var updateUserDto user.UpdateUserDto

	reqInfo := GetReqInfo(c)
	updateUserDto.Id = reqInfo.UserId

	if err := BindBody(&updateUserDto, c); err != nil {
		ErrorResponse(err, nil, r.config.DetailedError()).Reply(c)
		return
	}

	err := r.userUsecases.Update(contextWithReqInfo(c), updateUserDto)
	if err != nil {
		ErrorResponse(err, nil, r.config.DetailedError()).Reply(c)
		return
	}

	OkResponse(nil).Reply(c)
}

func (r *router) changeMyPassword(c *gin.Context) {
	var changeUserPasswordDto user.ChangeUserPasswordDto

	reqInfo := GetReqInfo(c)
	changeUserPasswordDto.Id = reqInfo.UserId

	if err := BindBody(&changeUserPasswordDto, c); err != nil {
		ErrorResponse(err, nil, r.config.DetailedError()).Reply(c)
		return
	}

	err := r.userUsecases.ChangePassword(contextWithReqInfo(c), changeUserPasswordDto)
	if err != nil {
		ErrorResponse(err, nil, r.config.DetailedError()).Reply(c)
		return
	}

	OkResponse(nil).Reply(c)
}

func (r *router) getAllUsers(c *gin.Context) {
	users, err := r.userUsecases.GetAllUsers(contextWithReqInfo(c))
	if err != nil {
		ErrorResponse(err, nil, r.config.DetailedError()).Reply(c)
		return
	}

	OkResponse(users).Reply(c)
}

func (r *router) getMe(c *gin.Context) {
	reqInfo := GetReqInfo(c)

	user, err := r.userUsecases.GetById(contextWithReqInfo(c), reqInfo.UserId)
	if err != nil {
		ErrorResponse(err, nil, r.config.DetailedError()).Reply(c)
		return
	}

	OkResponse(user).Reply(c)
}

func (r *router) getPostById(c *gin.Context) {
	postId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ErrorResponse(err, nil, r.config.DetailedError()).Reply(c)
		return
	}

	post, err := r.postUsecases.GetPostById(c, postId)
	if err != nil {
		ErrorResponse(err, nil, r.config.DetailedError()).Reply(c)
		return
	}

	OkResponse(post).Reply(c)
}

func (r *router) getMyPosts(c *gin.Context) {
	reqInfo := GetReqInfo(c)

	posts, err := r.postUsecases.GetMyPosts(c, reqInfo.UserId)
	if err != nil {
		ErrorResponse(err, nil, r.config.DetailedError()).Reply(c)
		return
	}

	OkResponse(posts).Reply(c)
}

func (r *router) updatePost(c *gin.Context) {
	var updatePostDto post.UpdatePostDto

	postId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ErrorResponse(err, nil, r.config.DetailedError()).Reply(c)
		return
	}

	updatePostDto.Id = postId

	if err := BindBody(&updatePostDto, c); err != nil {
		ErrorResponse(err, nil, r.config.DetailedError()).Reply(c)
		return
	}

	err = r.postUsecases.UpdatePost(c, updatePostDto)
	if err != nil {
		ErrorResponse(err, nil, r.config.DetailedError()).Reply(c)
		return
	}

	OkResponse(nil).Reply(c)
}

func (r *router) getPublishedPosts(c *gin.Context) {
	posts, err := r.postUsecases.GetPublishedPosts(c)
	if err != nil {
		ErrorResponse(err, nil, r.config.DetailedError()).Reply(c)
		return
	}

	OkResponse(posts).Reply(c)
}

func (r *router) getPosts(c *gin.Context) {
	posts, err := r.postUsecases.GetPosts(c)
	if err != nil {
		ErrorResponse(err, nil, r.config.DetailedError()).Reply(c)
		return
	}

	OkResponse(posts).Reply(c)
}

func (r *router) methodNotFound(c *gin.Context) {
	err := errors.New(errors.NotFoundError, "method not found")
	ErrorResponse(err, nil, r.config.DetailedError()).Reply(c)
}

func (r *router) recover() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		response := InternalErrorResponse(nil)
		c.AbortWithStatusJSON(response.Status, response)
	})
}

func (r *router) trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := c.Request.Header.Get("Trace-Id")
		if traceId == "" {
			traceId, _ = r.crypto.GenerateUUID()
		}

		setTraceId(c, traceId)
	}
}

func (r *router) logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		var parsedReqInfo request.RequestInfo

		reqInfo, exists := param.Keys[reqInfoKey]
		if exists {
			parsedReqInfo = reqInfo.(request.RequestInfo)
		}

		return fmt.Sprintf(
			"%s - [HTTP] TraceId: %s; UserId: %d; Method: %s; Path: %s; Status: %d, Latency: %s;\n\n",
			param.TimeStamp.Format(time.RFC1123),
			parsedReqInfo.TraceId,
			parsedReqInfo.UserId,
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
		)
	})
}

func BindBody(payload interface{}, c *gin.Context) error {
	err := c.BindJSON(payload)
	if err != nil {
		return errors.New(errors.BadRequestError, err.Error())
	}

	return nil
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func OkResponse(data interface{}) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: "ok",
		Data:    data,
	}
}

func InternalErrorResponse(data interface{}) *Response {
	status, message := http.StatusInternalServerError, "internal error"

	return &Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(err error, data interface{}, withDetails bool) *Response {
	status, message, details := parseError(err)

	if withDetails && details != "" {
		message = details
	}
	return &Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func (r *Response) Reply(c *gin.Context) {
	c.JSON(r.Status, r)
}
