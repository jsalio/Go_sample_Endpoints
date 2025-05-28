# API RESTful de Usuarios

Este proyecto es una API RESTful desarrollada en Go que proporciona endpoints para la gestión de usuarios.

## Características

- Gestión de usuarios (CRUD)
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

## Contribución

1. Haz un Fork del proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request
