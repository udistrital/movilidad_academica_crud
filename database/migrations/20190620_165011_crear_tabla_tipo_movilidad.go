package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaTipoMovilidad_20190620_165011 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaTipoMovilidad_20190620_165011{}
	m.Created = "20190620_165011"

	migration.Register("CrearTablaTipoMovilidad_20190620_165011", m)
}

// Run the migrations
func (m *CrearTablaTipoMovilidad_20190620_165011) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
m.SQL("CREATE TABLE movilidad_academica.tipo_movilidad(	id serial NOT NULL, nombre varchar(50) NOT NULL, descripcion varchar,codigo_abreviacion varchar(20), activo boolean NOT NULL, numero_orden numeric(5,2), entrante bool NOT NULL, CONSTRAINT pk_tipo_movilidad PRIMARY KEY (id));")
m.SQL("ALTER TABLE movilidad_academica.tipo_movilidad OWNER TO desarrollooas;")


}

// Reverse the migrations
func (m *CrearTablaTipoMovilidad_20190620_165011) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
 m.SQL("DROP TABLE IF EXISTS movilidad_academica.tipo_movilidad")
}
