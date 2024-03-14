package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/mahdikarami0111/cloud-computing-1/db/sqlc"
	"github.com/mahdikarami0111/cloud-computing-1/object"
	"github.com/mahdikarami0111/cloud-computing-1/rmq"
)

type Server struct {
	db     *db.Queries
	router *gin.Engine
}

func NewServer(db *db.Queries) *Server {
	server := &Server{db: db}
	router := gin.Default()

	//rout shit
	router.POST("/request", server.handleRequest)
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	fmt.Println("Server starting ...")
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

type HandleRequest struct {
	Email string `form:"email" binding:"required"`
}

func (server *Server) handleRequest(ctx *gin.Context) {
	var req HandleRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateRequestParams{
		Email:  req.Email,
		Status: "pending",
	}
	request, err := server.db.CreateRequest(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	body := ctx.Request.Body
	file, err := io.ReadAll(body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	filename := strconv.Itoa(int(request.ID)) + ".mp3"
	f2, err := os.Create(filename)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	n, err := f2.Write(file)
	if err != nil || n == 0 {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = f2.Close()
	if err != nil || n == 0 {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	bucket := "cc1-hw1-mk"
	object.UploadObject(bucket, filename)
	err = os.Remove(filename)
	rmq.SendMessage("amqps://rigqkizo:CBNChsj9lZoMSSzHXKB84-0glFjLZsT8@hawk.rmq.cloudamqp.com/rigqkizo", strconv.Itoa(int(request.ID)))
	if err != nil || n == 0 {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, request)
}
