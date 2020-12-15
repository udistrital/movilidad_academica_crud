# movilidad_academica_crud

CRUD para el negocio de movilidad academica.  
(Ejercicio introduccion 2019)

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)

### Variables de Entorno
```shell
# parametros de api
MOVILIDAD_ACADEMICA_CRUD_HTTP_PORT=[Puerto de exposición del API]
MOVILIDAD_ACADEMICA_CRUD_RUN_MODE=[Modo de ejecución del API]
# paramametros de bd
MOVILIDAD_ACADEMICA_CRUD_PGUSER=[Usuario de BD]
MOVILIDAD_ACADEMICA_CRUD_PGPASS=[Contraseña del usaurio de BD]
MOVILIDAD_ACADEMICA_CRUD_PGHOST=[URL, Dominio o EndPoint de la BD]
MOVILIDAD_ACADEMICA_CRUD_PGPORT=[Puerto de la BD]
MOVILIDAD_ACADEMICA_CRUD_PGDB=[Nombre de Base de Datos]
MOVILIDAD_ACADEMICA_CRUD_PGSCHEMA=[Nombre del Esquema de Base de Datos]
```
**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con MOVILIDAD_ACADEMICA_CRUD_...


### Ejecución del Proyecto
```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/movilidad_academica_crud

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/movilidad_academica_crud

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
MOVILIDAD_ACADEMICA_CRUD_PORT=8080 MOVILIDAD_ACADEMICA_CRUD_SOME_VARIABLE bee run
```

### Ejecución Dockerfile
```shell
# implementado para el sistema de integración continua (CI)
```

### Ejecución docker-compose
```shell
# 1. Obtener repositorio
git clone https://github.com/udistrital/movilidad_academica_crud.git
# 2. Ir a la carpeta del repositorio
cd $GOPATH/src/github.com/movilidad_academica_crud
# 3. Cambiar a la rama develop
git checkout develop
# 4. Crear red back_end
docker network create back_end
# 5. Ejecutar docker compose
docker-compose up --build
```

### Ejecución Pruebas

Pruebas unitarias
```shell
# En Proceso
```


### Modelo de Datos
[Modelo Relacional movilidad_academica_crud](modelo_datos.png)  
[Sql de la base de datos: movilidad académica BD](https://drive.google.com/file/d/1LJ_8nXmSo9lsoDzZPKfn6Q-7bdk8FZqK/view?usp=sharing)

### Puertos
El servidor se expone en el puerto: 127.0.0.1:8080   
Para ver la documentación de swagger: [127.0.0.1:8080/swagger/](http://127.0.0.1:8080/swagger/)

### EndPoints

|                |link de prueba                  |End Point|
|----------------|-------------------------------|------------------------|
| **Obtiene datos asociados a la movilidad estudiantil** |[GetAll](http://127.0.0.1:8080/v1/movilidad)| `127.0.0.1:8080/v1/movilidad` |
| **Obtiene las categorías de movilidad** | [GetAll](http://127.0.0.1:8080/v1/tipo_categoria) |`127.0.0.1:8080/v1/tipo_categoria`|
| **Obtiene los tipo de movilidad(Extranjera o nacional)** |[GetAll](http://127.0.0.1:8080/v1/tipo_movilidad)| `127.0.0.1:8080/v1/tipo_movilidad` |


## Estado CI

| Develop | Relese 0.0.1 | Master |
| -- | -- | -- |
| - | - | - |


## Licencia

This file is part of cumplidos_crud.

cumplidos_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

cumplidos_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with novedades_crud. If not, see https://www.gnu.org/licenses/.
