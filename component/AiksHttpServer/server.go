package AiksHttpServer

import (
	"net/http"
)

type HandlerFunc func(*Context)

type Server struct {
	host string
	port string
	router *Router
}

func New(host,port string) *Server {
	return &Server{
		host: host,
		port: port,
		router: newRouter()}
}

func (s *Server) addRoute(method string, pattern string, handler HandlerFunc) {
	s.router.addRoute(method, pattern, handler)
}

func (s *Server) GET(pattern string, handler HandlerFunc) {
	s.addRoute("GET", pattern, handler)
}

func (s *Server) POST(pattern string, handler HandlerFunc) {
	s.addRoute("POST", pattern, handler)
}

func (s *Server) Run() (err error) {
	return http.ListenAndServe(s.host+":"+s.port, s)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	s.router.handle(c)
}