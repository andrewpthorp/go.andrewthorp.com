package main

import (
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
 )

func main() {
  m := martini.Classic()
  m.Use(martini.Static("app/assets/css"))
  m.Use(render.Renderer(render.Options{Directory: "app/templates"}))

  m.Get("/", func(r render.Render){
    r.HTML(200, "index", nil)
  })

  m.Run()
}
