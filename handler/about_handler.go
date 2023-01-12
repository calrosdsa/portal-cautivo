package handler

import (
  "net/http"

  "github.com/labstack/echo"
)

func AboutHandler(c echo.Context) error {
  // Please note the the second parameter "about.html" is the template name and should
  // be equal to one of the keys in the TemplateRegistry array defined in main.go
  return c.Render(http.StatusOK, "about.html", map[string]interface{}{
    "name": "About",
    "msg": "Bienvenidos a YPFB transporte!",
    "logo": "https://www.somare.com/wp-content/uploads/2020/02/logo-ypfb-transporte.jpg",
    "banner":"https://imgs.search.brave.com/zn6LL-NBqY9lVeYVegB4olOrLS5Qy7h9qj5Y5FBhF-g/rs:fit:768:439:1/g:ce/aHR0cHM6Ly93d3cu/ZWxwZXJpb2RpY28t/ZGlnaXRhbC5jb20v/d3AtY29udGVudC91/cGxvYWRzLzIwMTkv/MDgvZm90by15cGZi/LTEtMS03Njh4NDM5/LmpwZw",
    "descripcion":` Para acceder a la red, deberás iniciar sesión con tu cuenta de Facebook y posteriormente
    dar "me gusta" a la última publicación de la página de
   YPFB Corp.`,
  })
}