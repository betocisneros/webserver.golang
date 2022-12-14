// el archivo main va a poder leer todo lo que este en el archivo server
package main

import (
	"net/http"
)

type Server struct {
	//puerto del servidor para escuchar las conexiones
	port string

	//se agrega el atributo router que es un apuntador al struct Router de router.go
	router *Router
}

// Funcion tipo global para ser leida en otros archivos
// Sirve para instanciar el servidor y que sea capaz de escuchar las conexiones
// recibe el puerto que tiene que estar escuchando y devuelve el servidor como tal
func NewServer(port string) *Server {
	return &Server{
		port: port,

		//router instanciado
		router: NewRouter(),
		//con esto el servidor ya es capaz de instanciar el router y de tenerlo como propiedad
	}
}

func (s *Server) Handle(path string, handler http.HandlerFunc) {
	s.router.rules[path] = handler

}

func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...MiddleWare) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

// Funcion tipo receiver, del struct Server, devuelve un error en caso de que haya problemas al conectar
func (s *Server) Listen() error {
	//el router va a ser el encargado de tomar las urls y procesarlas como se debe, crea el entry-point
	// los parametro son: el slash que es el punto de entrada, y el handler es el router recien creado
	http.Handle("/", s.router)

	//con la funcion listenanserve() del paquete http nos ayuda a escuchar todas las peticiones
	//colocas el puerto como primer parametro, el segundo es un handler
	//pero nosotros haremos nuestros handlers por eso se coloca nil
	err := http.ListenAndServe(s.port, nil)

	if err != nil {
		return err
	}
	//si la ejecucion salio bien, retorna un valor nil
	return nil
}
