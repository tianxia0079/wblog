package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wangsongyan/wblog/helpers"
	"github.com/wangsongyan/wblog/system"
	"net/http"
)

const (
	SESSION_KEY      = "UserID" // session key
	CONTEXT_USER_KEY = "User"   // context user key
)

func Handle404(c *gin.Context) {
	HandleMessage(c, "Sorry,I lost myself!")
}

func HandleMessage(c *gin.Context, message string) {
	c.HTML(http.StatusNotFound, "errors/error.html", gin.H{
		"message": message,
	})
}

func sendMail(to, subject, body string) error {
	c := system.GetConfiguration()
	return helpers.SendToMail(c.SmtpUsername, c.SmtpPassword, c.SmtpHost, to, subject, body, "html")
}

/*func __sendMail(to, subject, body string) error {
	return nil
}*/
