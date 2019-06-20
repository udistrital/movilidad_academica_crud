package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaTipoCategoria_20190620_165334 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaTipoCategoria_20190620_165334{}
	m.Created = "20190620_165334"

	migration.Register("CrearTablaTipoCategoria_20190620_165334", m)
}

// Run the migrations
func (m *CrearTablaTipoCategoria_20190620_165334) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
m.SQL("CREATE TABLE movilidad_academica.tipo_categoria(	id serial NOT NULL,	nombre varchar(50) NOT NULL,	descripcion varchar,	codigo_abreviacion varchar(20),	activo boolean NOT NULL,	numero_orden numeric(5,2),	CONSTRAINT pk_tipo_categoria PRIMARY KEY (id));")
m.SQL("ALTER TABLE movilidad_academica.tipo_categoria OWNER TO desarrollooas;")
}

// Reverse the migrations
func (m *CrearTablaTipoCategoria_20190620_165334) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
m.SQL("DROP TABLE IF EXISTS movilidad_academica.tipo_categoria")
}
