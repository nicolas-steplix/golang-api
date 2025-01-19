<div id="top"></div>

# GoLang Example

Proyecto de ejemplo con GoLang.

## Index

* [Descargar][descargar]
* [Configuraciones][configuraciones]
* [Base de datos][base_de_datos]
* [¿Cómo se corre?][como_se_corre]


## Descargar

### Código fuente
```shell
$ git clone https://github.com/nicolas-steplix/golang-api.git
$ cd golang-api
$ go mod tidy
```

<p align="right">(<a href="#top">ir arriba</a>)</p>
<hr/>


## Configuraciones

Para el correcto funcionamiento de la API, la misma, cuenta con un archivo `.env.example`. En él se encuentran las configuraciones mínimas para poder correrla.

Para que levante dichas configuraciones, es necesario correr el siguiente comando,
```shell
$ cp .env.example .env
```

<p align="right">(<a href="#top">ir arriba</a>)</p>
<hr/>


## Base de datos

Esta API cuenta con su propia base de datos. La misma se encuentra en el root, dentro de la carpeta `database/script.sql`. 

### Instalar DB

>
> 💡 primero tenes que ubicar la termina en el root del proyecto `cd golang-api`
>

```shell
mysql -u mydbuser -p < database/script.sql
```

<p align="right">(<a href="#top">ir arriba</a>)</p>
<hr/>


## ¿Cómo se corre?

### Correr

```shell
go run .
```

<p align="right">(<a href="#top">ir arriba</a>)</p>
<hr/>


## Documentación de la API
Pueden encontrar los endpoint para hacer pruebas dentro de carpeta `docs/api-collection`. La misma puede ser importada en [Bruno](https://www.usebruno.com/).

<p align="right">(<a href="#top">ir arriba</a>)</p>
<hr/>


<!-- deep links -->
[descargar]: #descargar
[configuraciones]: #configuraciones
[base_de_datos]: #base-de-datos
[como_se_corre]: #cómo-se-corre
