### Si queres sabes mas sobre mi:
[Nestor Marsollier](https://github.com/nmarsollier/profile)

# Orders Service en GO

Este Microservicio de seguridad reemplaza al del proyecto

[Microservicios Orders](https://github.com/nmarsollier/ecommerce)

Se encarga de registrar y autenticar usuarios en el sistema.

Utiliza el esquema JWT con un header Authorization "bearer" estándar.

Cada usuario tiene asociado una lista de permisos, existen 2 permisos genéricos "user" y "admin". Los usuarios que se registran son todos "user",  muchos procesos necesitan un usuario "admin" para poder funcionar, por lo tanto hay que editar el esquema en mongodb para asociarle el permiso admin a algún usuario inicialmente.

[Documentación de API](./README-API.md)

La documentación de las api también se pueden consultar desde el home del microservicio
que una vez levantado el servidor se puede navegar en [localhost:3000](http://localhost:3000/)

## Requisitos

Go 1.14  [golang.org](https://golang.org/doc/install)


## Configuración inicial

establecer variables de entorno (consultar documentación de la version instalada)

Para descargar el proyecto correctamente hay que ejecutar :

```bash
git clone https://github.com/nmarsollier/ordersgo $GOPATH/src/github.com/nmarsollier/ordersgo
```

Una vez descargado, tendremos el codigo fuente del proyecto en la carpeta

```bash
cd $GOPATH/src/github.com/nmarsollier/ordersgo
```


## Dependencias

### Auth

Las ordenes sólo puede usarse por usuario autenticados, ver la arquitectura de microservicios de [ecommerce](https://github.com/nmarsollier/ecommerce).

### Catálogo

Las ordenes tienen una fuerte dependencia del catalogo:

- Se validan los artículos contra el catalogo.
- Se descuentan los artículos necesarios.
- Se puede devolver articulos si la orden se cancela.

Ver la arquitectura de microservicios de [ecommerce](https://github.com/nmarsollier/ecommerce).


## Instalar Librerías requeridas


```bash
go mod download
```

Build y ejecución
-

```bash
go install
ordersgo
```

## MongoDB

La base de datos se almacena en MongoDb.

Seguir las guías de instalación de mongo desde el sitio oficial [mongodb.com](https://www.mongodb.com/download-center#community)

No se requiere ninguna configuración adicional, solo levantarlo luego de instalarlo.

## RabbitMQ

Este microservicio notifica los logouts de usuarios con Rabbit.

Seguir los pasos de instalación en la pagina oficial [rabbitmq.com](https://www.rabbitmq.com/)

No se requiere ninguna configuración adicional, solo levantarlo luego de instalarlo.

## Apidoc

Apidoc es una herramienta que genera documentación de apis para proyectos node (ver [Apidoc](http://apidocjs.com/)).

El microservicio muestra la documentación como archivos estáticos si se abre en un browser la raíz del servidor [localhost:3004](http://localhost:3004/)

Ademas se genera la documentación en formato markdown.

Para que funcione correctamente hay que instalarla globalmente con

```bash
npm install apidoc -g
npm install -g apidoc-markdown2
```

La documentación necesita ser generada manualmente ejecutando la siguiente linea en la carpeta auth :

```bash
apidoc -o www
apidoc-markdown2 -p www -o README-API.md
```

Esto nos genera una carpeta con la documentación, esta carpeta debe estar presente desde donde se ejecute auth, auth busca ./www para localizarlo, aunque se puede configurar desde el archivo de properties.

## Configuración del servidor

Este servidor usa las siguientes variables de entorno para configuración :

RABBIT_URL : Url de rabbit (default amqp://localhost)
MONGO_URL : Url de mongo (default mongodb://localhost:27017)
PORT : Puerto (default 3000)
WWW_PATH : Path donde se ubica la documentación apidoc (default www)
JWT_SECRET : Secret para password (default ecb6d3479ac3823f1da7f314d871989b)

## Docker para desarrollo

### Build

```bash
docker build -t dev-order-go .
```

### El contenedor

```bash
# Mac | Windows
docker run -it --name dev-order-go -p 3004:3004 -v $PWD:/go/src/github.com/nmarsollier/ordersgo dev-order-go

# Linux
docker run -it --add-host host.docker.internal:172.17.0.1 --name dev-order-go -p 3000:3000 -v $PWD:/go/src/github.com/nmarsollier/ordersgo dev-order-go
```

### Debug con VSCode

Existe un archivo Docker.debug, hay que armar la imagen usando ese archivo.

```bash
docker build -t debug-order-go -f Dockerfile.debug .
```

```bash
# Mac | Windows
docker run -it --name debug-order-go -p 3004:3004 -p 40000:40000 -v $PWD:/go/src/github.com/nmarsollier/ordersgo debug-order-go

# Linux
docker run -it --add-host host.docker.internal:172.17.0.1 --name debug-order-go -p 3004:3004 -p 40000:40000 -v $PWD:/go/src/github.com/nmarsollier/ordersgo debug-order-go
```

El archivo launch.json debe contener lo siguiente

```bash
{
    "version": "0.2.0",
    "configurations": [
          {
                "name": "Debug en Docker",
                "type": "go",
                "request": "launch",
                "mode": "remote",
                "remotePath": "/go/src/github.com/nmarsollier/ordersgo",
                "port": 40000,
                "host": "127.0.0.1",
                "program": "${workspaceRoot}",
                "showLog": true
          }
    ]
}
```

En el menú run start debugging se conecta a docker.
