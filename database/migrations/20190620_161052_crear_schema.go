package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearSchema_20190620_161052 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearSchema_20190620_161052{}
	m.Created = "20190620_161052"

	migration.Register("CrearSchema_20190620_161052", m)
}

// Run the migrations
func (m *CrearSchema_20190620_161052) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
m.SQL("CREATE SCHEMA movilidad_academica;")
m.SQL("ALTER SCHEMA movilidad_academica OWNER TO sistemasoas;")
m.SQL("SET search_path TO pg_catalog,public,movilidad_academica;")

}

// Reverse the migrations
func (m *CrearSchema_20190620_161052) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
m.SQL("DROP SCHEMA IF EXISTS movilidad_academica");
}
