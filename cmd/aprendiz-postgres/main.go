package main

import (
	"fmt"

	"github.com/cris329/aprendiz-postgres/internal/controller"
	"github.com/cris329/aprendiz-postgres/internal/database"
	"github.com/cris329/aprendiz-postgres/internal/model"
	"github.com/cris329/aprendiz-postgres/internal/repository"
	"github.com/cris329/aprendiz-postgres/internal/router"
	"github.com/cris329/aprendiz-postgres/internal/service"
)

func main() {
	fmt.Println("Conectando a PostgreSQL...")
	db, err := database.NewGormDB()
	if err != nil {
		fmt.Println("ERROR conexión BD:", err)
		return
	}
	fmt.Println("Conexión exitosa")

	fmt.Println("Migrando tabla...")
	db.AutoMigrate(&model.Aprendiz{})
	fmt.Println("Tabla migrada")

	repo := repository.NewAprendizRepository(db)
	svc := service.NewAprendizService(repo)
	ctrl := controller.NewAprendizController(svc)

	r := router.New(ctrl)
	fmt.Println("Servidor corriendo en http://localhost:8080")
	r.Run(":8080")
}
