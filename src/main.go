package main

import (
	"ci-cd-practice/src/modules/request"
	"os"

	"github.com/kataras/iris/v12"
)

type AddRequestBody struct {
	Title string `json:"title"`
}

type RemoveRequestBody struct {
	Id string `json:"id"`
}

type UpdateRequestBody struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type GetOneRequestParams struct {
	Id string `json:"id"`
}

func main() {
	app := iris.New()

	testAPI := app.Party("/")
	{
		testAPI.Use(iris.Compression)

		testAPI.Get("/healthcheck", healthcheck)
		testAPI.Get("/list", getRequests)
		testAPI.Get("/list/{id:string}", getRequestById)

		testAPI.Post("/insert", addRequest)
		testAPI.Post("/delete", removeRequestById)
		testAPI.Post("/update", updateRequestById)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	app.Listen(":" + port)
}

func healthcheck(ctx iris.Context) {
	ctx.StatusCode(200)

	ctx.JSON("Ok")
}

func getRequests(ctx iris.Context) {
	ctx.JSON(request.GetMany())
}

func addRequest(ctx iris.Context) {
	var body AddRequestBody

	err := ctx.ReadJSON(&body)

	if err == nil {
		ctx.JSON(request.Insert(body.Title))
	} else {
		ctx.JSON(err.Error())
	}
}

func getRequestById(ctx iris.Context) {
	var params GetOneRequestParams

	err := ctx.ReadParams(&params)

	if err == nil {
		ctx.JSON(request.GetOne(params.Id))
	} else {
		ctx.JSON(err.Error())
	}
}

func removeRequestById(ctx iris.Context) {
	var body RemoveRequestBody

	err := ctx.ReadJSON(&body)

	if err == nil {
		ctx.JSON(request.DeleteOne(body.Id))
	} else {
		ctx.JSON(err.Error())
	}
}

func updateRequestById(ctx iris.Context) {
	var body UpdateRequestBody

	err := ctx.ReadJSON(&body)

	if err == nil {
		ctx.JSON(request.UpdateOne(body.Id, body.Title))
	} else {
		ctx.JSON(err.Error())
	}
}
