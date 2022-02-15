package templates

import (
	"blog-on-containers/entities"
	"blog-on-containers/services"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LoginPage(context *gin.Context) {
	context.HTML(http.StatusOK, "login.html", nil)
}

func RegisterPage(context *gin.Context) {
	context.HTML(http.StatusOK, "register.html", nil)
}

type storyVo struct {
	CreatedDate time.Time
	Title       string
	Content     string
	Username    string
}

func StoryPage(context *gin.Context) {
	var story entities.Story
	storyService := services.NewStoryService(context)
	story, bool := storyService.GetStory()
	if !bool {
		return
	}

	//t, _ := template.ParseFiles("story.html")
	//t.Execute(context.Writer, story)

	storyVo := storyVo{
		CreatedDate: story.CreatedDate,
		Title:       story.Title,
		Content:     story.Content,
		Username:    "username",
	}
	context.HTML(http.StatusOK, "story.html", storyVo)
	fmt.Println(story)
}
