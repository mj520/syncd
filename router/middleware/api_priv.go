// Copyright 2019 syncd Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package middleware

import (
    "strings"
    "github.com/gin-gonic/gin"
    "github.com/dreamans/syncd/render"
    "github.com/dreamans/syncd/module/user"
    reqApi "github.com/dreamans/syncd/router/route/api"
)

func ApiPriv() gin.HandlerFunc {
    return func(c *gin.Context) {
        token, _ := c.Cookie("_syd_identity")

        path := strings.TrimSpace(c.Request.URL.Path)
        if len(path) < 4 {
            respondWithError(c, render.CODE_ERR_NO_PRIV, "request path is too short")
            return
        }
        path = path[4:]

        if path == reqApi.LOGIN {
            c.Next()
            return
        }

        if token == "" {
            respondWithError(c, render.CODE_ERR_NO_LOGIN, "no login")
            return
        }

        login := &user.Login{
            Token: token,
        }
        if err := login.ValidateToken(); err != nil {
            respondWithError(c, render.CODE_ERR_NO_LOGIN, err.Error())
            return
        }

        //priv check

    }
}

func respondWithError(c *gin.Context, code int, message string) {
    render.CustomerError(c, code, message)
    c.Abort()
}
