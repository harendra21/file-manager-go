package controllers

type ErrorController struct {
    AppController
}

func (c *ErrorController) Error404() {
    c.ThrowError(404, "Page not found")
}

func (c *ErrorController) Error500() {
    c.ThrowError(404, "Page not found")
}

func (c *ErrorController) ErrorDb() {
    c.ThrowError(404, "Page not found")
}
