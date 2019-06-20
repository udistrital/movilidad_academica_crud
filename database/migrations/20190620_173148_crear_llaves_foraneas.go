package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearLlavesForaneas_20190620_173148 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearLlavesForaneas_20190620_173148{}
	m.Created = "20190620_173148"

	migration.Register("CrearLlavesForaneas_20190620_173148", m)
}

// Run the migrations
func (m *CrearLlavesForaneas_20190620_173148) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
m.SQL("ALTER TABLE movilidad_academica.movilidad ADD CONSTRAINT fk_movilidad_tipo_movilidad FOREIGN KEY (id_tipo_movilidad) REFERENCES movilidad_academica.tipo_movilidad (id) MATCH FULL ON DELETE SET NULL ON UPDATE CASCADE;")
m.SQL("ALTER TABLE movilidad_academica.movilidad ADD CONSTRAINT fk_movilidad_tipo_categoria FOREIGN KEY (id_tipo_categoria) REFERENCES movilidad_academica.tipo_categoria (id) MATCH FULL ON DELETE SET NULL ON UPDATE CASCADE;")
}

// Reverse the migrations
func (m *CrearLlavesForaneas_20190620_173148) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
m.SQL("ALTER TABLE movilidad_academica.movilidad DROP CONSTRAINT IF EXISTS fk_movilidad_tipo_movilidad CASCADE;")
m.SQL("ALTER TABLE movilidad_academica.movilidad DROP CONSTRAINT IF EXISTS fk_movilidad_tipo_categoria CASCADE;")
}
