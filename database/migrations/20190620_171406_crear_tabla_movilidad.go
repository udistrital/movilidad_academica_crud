 package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaMovilidad_20190620_171406 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaMovilidad_20190620_171406{}
	m.Created = "20190620_171406"

	migration.Register("CrearTablaMovilidad_20190620_171406", m)
}

// Run the migrations
func (m *CrearTablaMovilidad_20190620_171406) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
m.SQL("CREATE TABLE movilidad_academica.movilidad(id serial NOT NULL,	presupuesto numeric(20,7) NOT NULL DEFAULT 0,	fecha_inicio date NOT NULL,fecha_final date NOT NULL,	persona integer NOT NULL,	convenio integer NOT NULL,	id_tipo_movilidad integer,	id_tipo_categoria integer,	CONSTRAINT pk_movilidad PRIMARY KEY (id));")

m.SQL("ALTER TABLE movilidad_academica.movilidad OWNER TO desarrollooas;")

}

// Reverse the migrations
func (m *CrearTablaMovilidad_20190620_171406) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
m.SQL("DROP TABLE IF EXISTS movilidad_academica.movilidad") 
}
