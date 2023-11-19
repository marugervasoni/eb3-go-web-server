// Ejercicio 1: Iniciando el proyecto
//crear un archivo main.go donde deberán cargar un slice, desde una función que devuelva una lista de empleados. Este slice se debe cargar cada vez
// que se inicie la API para realizar las distintas consultas.
// La estructura de los empleados es la siguiente:
// Id		 int
// Nombre	 string
// Activo	 bool

// Ejercicio 2: Empleados
// Vamos a levantar un servidor utilizando el paquete Gin en el puerto 8080. Para probar nuestros endpoints, haremos uso de Postman.
// Crear una ruta / que nos devuelva una bienvenida al sistema. Ejemplo: “¡Bienvenido a la empresa Gophers!”.
// Crear una ruta /employees que nos devuelva la lista de todos los empleados en formato JSON.
// Crear una ruta /employees/:id que nos devuelva un empleado por su ID. Manejar el caso de que no encuentre el empleado con ese ID.
// Crear una ruta /employeesparams que nos permita crear un empleado a través de los params y lo devuelva en formato JSON.
// Crear una ruta /employeesactive que nos devuelva la lista de empleados activos. También debería poder devolver la lista de los empleados no activos.
package main

import (
	"log"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

const (
	port = ":8080"
)

type Employee struct {
	Id			int		`json:"id"`
	Name		string	`json:"name"`
	IsActive	bool	`json:"is_active"`
}

//store es una base de datos en memoria
type Store struct {
	Employees []Employee
}

var store = Store{}

func main() {

	//carga de base de datos en memoria
	store.LoadStore()

	router := gin.Default()
	group := router.Group("/api/v1")
	{
		group.GET("/", Bienvenida)
		group.GET("/ping", Ping)
		groupEmployees := group.Group("/employees")
		{
			groupEmployees.GET("", AllEmployees)
			groupEmployees.GET("/:id", GetEmployee)
			groupEmployees.GET("/params", CreateEmployee)
			groupEmployees.GET("/active", ActiveAndNonActiveEmployees)
		}
	} 
	
	if err := router.Run(port); err != nil {
		log.Fatal(err)
	}
 
}
//handler de respuesta a "/"
func Bienvenida(ctxt *gin.Context)  {
	ctxt.String(http.StatusOK, "¡Welcome to Gophers!")
}

//handler de respuesta a "/ping"
func Ping(ctxt *gin.Context)  {
	ctxt.JSON(http.StatusOK, gin.H {
		"data": "pong",
	})
}

//handler de respuesta a "employees/"
func AllEmployees(ctxt *gin.Context)  {
	ctxt.JSON(http.StatusOK, gin.H {
		"data": store.Employees,
	})
}

//handler de respuesta a "employees/:id"
func GetEmployee(ctxt *gin.Context) {
	idEmployee := ctxt.Param("id")
	id, err := strconv.Atoi(idEmployee)
	if err != nil {
		ctxt.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid parameter",
		})
		return
	}

	// Buscar el empleado por ID
	var employeeFound bool
	var employee Employee
	for _, e := range store.Employees {
		if e.Id == id {
			employee = e
			employeeFound = true
			break
		}
	}

	if employeeFound {
		ctxt.JSON(http.StatusOK, gin.H{
			"data": employee,
		})
	} else {
		ctxt.JSON(http.StatusNotFound, gin.H{
			"message": "Employee doesn't found",
		})
	}
}

//handler de respuesta a "employeesparams"
func CreateEmployee(ctxt *gin.Context) {
	var newEmployee Employee
	if err := ctxt.BindJSON(&newEmployee); err != nil {
		ctxt.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid employee information",
		})
		return
	}
	
	// Asignar un nuevo ID
	newEmployee.Id = len(store.Employees) + 1
	
	// Agregar el nuevo empleado a la base de datos en memoria
	store.Employees = append(store.Employees, newEmployee)
	
	ctxt.JSON(http.StatusCreated, gin.H{
		"message": "Employee created succesfully",
		"data":    newEmployee,
	})
}

//handler de respuesta a "employeesactive"
func ActiveAndNonActiveEmployees(ctxt *gin.Context)  {
	activeParam := ctxt.Query("active")
		
	var resultEmployees []Employee
	switch activeParam {
	case "true":
		for _, e := range store.Employees {
			if e.IsActive {
				resultEmployees = append(resultEmployees, e)
			}
		}
	case "false":
		for _, e := range store.Employees {
			if !e.IsActive {
				resultEmployees = append(resultEmployees, e)
			}
		}
	default:
		ctxt.JSON(http.StatusBadRequest, gin.H{
			"message":"Invalid 'active' parameter",
		})
		return
	}
	ctxt.JSON(http.StatusOK, gin.H{
		"data": resultEmployees,
	})
}

// LoadStore carga la base de datos en memoria
func (s *Store) LoadStore() {
	s.Employees = []Employee{
		{
			Id:			1, 
			Name:		"Nate Johnson",
			IsActive:	true,
		},
		{
			Id:			2,
			Name:		"Geourge Ford",
			IsActive:	false,
		},
		{
			Id:			3,
			Name:		"Harrison Clooney",
			IsActive:	true,
		},
	}
}