package main

import "net/http"

type MiddleWare func(http.HandlerFunc) http.HandlerFunc
