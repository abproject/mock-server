# mock-server
> **mock-server** is a tiny Web Server (Docker image <10 MB) that allows easily to emulate the RESTful and WebSocket services.
> 
> It could be useful for:
> * FrontEnd Developers: fast mocking the BackEnd endpoints during implementation of the UI prototypes
> * BackEnd Developers: mocking services and applications during integration or system tests 
---

## Table of Contents
* [Run](#run)
* [Examples](example/README.md)
    * [Hello World](#hello-world)
    * [Files](#files)
    * [CRUD](#crud)
    * [Entity](#entity)
    * [WebSocket](#websocket)
* [File Configuration](#file-configuration)
* [API](#api)

---
## Run

`docker run -p 8000:8000 -v ${PWD}/example:/example abezpalov/mock-server:latest -file=/example/crud/config.yml`

---

  
---
## File Configuration

---
## API

