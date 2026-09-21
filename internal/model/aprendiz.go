package model

type Aprendiz struct {
	ID                 uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	NumeroIdentificacion string `gorm:"column:numero_identificacion;not null" json:"numero_identificacion"`
	Nombre               string `gorm:"column:nombre;not null" json:"nombre"`
	Apellido             string `gorm:"column:apellido;not null" json:"apellido"`
	Genero               string `gorm:"column:genero;not null" json:"genero"`
	TipoSangre           string `gorm:"column:tipo_sangre;not null" json:"tipo_sangre"`
	Telefono             string `gorm:"column:telefono;not null" json:"telefono"`
	Programa             string `gorm:"column:programa;not null" json:"programa"`
	Ficha                string `gorm:"column:ficha;not null" json:"ficha"`
	Regional             string `gorm:"column:regional;not null" json:"regional"`
}

func (Aprendiz) TableName() string {
	return "aprendiz_postgres"
}
