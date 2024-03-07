package impl

import (
	"fmt"

	"github.com/gin-gonic/gin"

	http "fibo/api/http"
	"fibo/api/http/postcontroller"
	"fibo/internal/category"
	"fibo/internal/post"
)

type PostControllerOpts struct {
	PostUsecase post.PostUseCase
	CatUsecase  category.CatUseCase
	Config      http.Config
}

func NewPostController(opts PostControllerOpts) postcontroller.PostController {
	return &postController{
		PostUseCase: opts.PostUsecase,
		Config:      opts.Config,
		CatUseCase:  opts.CatUsecase,
	}
}

type postController struct {
	post.PostUseCase
	http.Config
	category.CatUseCase
}

func (p *postController) AddPostC(c *gin.Context) {
	var addPostDto post.AddPostDto

	if err := http.BindBody(&addPostDto, c); err != nil {
		http.ErrorResponse(err, nil, p.Config.DetailedError()).Reply(c)
		return
	}

	catModel, err := p.CatUseCase.GetByID(c, addPostDto.CategoryId)
	if err != nil {
		http.ErrorResponse(err, nil, p.Config.DetailedError()).Reply(c)
		return
	}

	fmt.Println(catModel)
	reqInfo := http.GetReqInfo(c)
	addPostDto.UserId = reqInfo.UserId
	fmt.Println(addPostDto)

	postId, err := p.PostUseCase.AddPost(c, addPostDto)
	if err != nil {
		http.ErrorResponse(err, nil, p.Config.DetailedError()).Reply(c)
		return

	}
	fmt.Println(postId)
	//
	http.OkResponse(postId).Reply(c)
}
