package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	app := iris.New()
	app.Get("/api/{apiCall:path}", func(ctx iris.Context) {
		endpoint := "http://0.0.0.0:8899/"
		apiCall := ctx.Params().Get("apiCall")
		if len(apiCall) > 0 {
			endpoint = fmt.Sprintf("%s%s", endpoint, apiCall)
		}
		rawQuery := ctx.Request().URL.RawQuery
		if len(rawQuery) > 0 {
			endpoint = fmt.Sprintf("%s?%s", endpoint, rawQuery)
		}
		//http&https proxy
		resp, err := http.Get(endpoint)
		if err != nil {
			ctx.JSON(iris.Map{"success": false, "error_message": err.Error()})
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode == 200 || resp.StatusCode == 201 {
			body, _ := ioutil.ReadAll(resp.Body)
			ctx.ContentType(context.ContentJSONHeaderValue)
			//ctx.Write(body)
			ctx.HTML(string(body))
		} else {
			ctx.JSON(iris.Map{"success": false, "error_message": "target not ok"})
		}
	})
/*
	app.Get("/gettype/{apiCall:path}", func(ctx iris.Context) {
		apiCall := ctx.Params().Get("apiCall")
		//ugets := ctx.Handlers()
		 mc := string(apiCall[strings.LastIndex(apiCall, "."):])
		 print(mc)

		if len(apiCall) > 0 {
			log.Print(apiCall)
		}else {
			ctx.HTML("error")
		}

	})

*/

	err := app.Run(
		// Start the web server at localhost:8080
		iris.Addr(":9000"),
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)

	if err != nil {
		log.Println(err.Error())
	}
}
