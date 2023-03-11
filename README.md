# Go Web App Development

## Overview

Code for a web application developed using the Go programming language.

The context is a store where you can organize products and also interact with it. It allows users to CRUD (create, read, update, and delete) data stored in a Postgres database.

### Index Page
![Index](img/index.png?raw=true)

### Create Product Page (similar to Edit Product)
![Create](img/create.png?raw=true)

## Layout
The project is organized as shown in this tree.


```
web-go
│   go.mod
│   go.sum
│   main.go
│   README.md
│
├───controllers (direct incoming data to the appropriate models, packages, etc..)
│       products.go
│
├───db (connecting with PG Database)
│       db.go
│
├───models (interact with your application’s data)
│       products.go
│
├───routes (specify routes)
│       routes.go
│
└───templates (html templates with Go and JS)
        edit.html
        index.html
        new.html
        _head.html
        _menu.html
```

## For your information

The project is developed without any external framework, using only Go's standard library and other necessary libraries. The project is modularized according to the
MVC convention, with separate files for handling routes, controllers,
and models.
