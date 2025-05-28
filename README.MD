# API RESTful de Usuarios y Posts

Este proyecto es una API RESTful desarrollada en Go que proporciona endpoints para la gestión de usuarios y sus posts.

## Características

- Gestión de usuarios (CRUD)
- Gestión de posts (CRUD)
- API RESTful
- Servidor HTTP en Go
- Arquitectura limpia y modular
- Documentación con Swagger UI

## Requisitos

- Go 1.23.9 o superior

## Estructura del Proyecto

```
.
├── cmd/
│   └── api/
│       └── main.go         # Punto de entrada de la aplicación
├── internal/
│   ├── api/
│   │   └── handlers/       # Manejadores HTTP
│   ├── models/            # Modelos de datos
│   └── services/          # Lógica de negocio
├── docs/                  # Documentación Swagger
├── go.mod
└── go.sum
```

## Endpoints

### Usuarios

- `GET /users` - Listar todos los usuarios
- `POST /users` - Crear un nuevo usuario
- `GET /users/{id}` - Obtener un usuario por ID
- `DELETE /users/{id}` - Eliminar un usuario por ID

### Posts

- `GET /posts` - Listar todos los posts
- `POST /posts` - Crear un nuevo post
- `GET /posts/{id}` - Obtener un post por ID
- `DELETE /posts/{id}` - Eliminar un post por ID
- `GET /users/{id}/posts` - Obtener todos los posts de un usuario específico

### Documentación Swagger

La API incluye documentación interactiva con Swagger UI. Para acceder a la documentación:

1. Inicia el servidor
2. Abre tu navegador y visita: `http://localhost:8085/swagger/index.html`

La documentación Swagger te permite:
- Ver todos los endpoints disponibles
- Probar los endpoints directamente desde el navegador
- Ver los esquemas de request/response
- Ver ejemplos de peticiones y respuestas

## Ejecución

Para ejecutar el proyecto:

```bash
go run cmd/api/main.go
```

El servidor se iniciará en el puerto 8085.

## Desarrollo

1. Clona el repositorio
2. Instala las dependencias:
   ```bash
   go mod download
   ```
3. Ejecuta el servidor:
   ```bash
   go run cmd/api/main.go
   ```

### Actualizar la documentación Swagger

Si realizas cambios en la API, puedes actualizar la documentación Swagger ejecutando:

```bash
swag init -g cmd/api/main.go
```

### Ejecutar Tests

Para ejecutar todos los tests del proyecto:

```bash
go test ./...
```

Para ejecutar tests de un paquete específico:

```bash
go test ./internal/services
```

Para ejecutar tests con cobertura:

```bash
go test ./... -cover
```

Para ver la cobertura en detalle y generar un reporte HTML:

```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Contribución

1. Haz un Fork del proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request
