# RESTful API con Go y Chi

Este proyecto es una REST API básica construida en Golang utilizando la librería `chi` para el manejo de rutas. La API permite realizar operaciones CRUD sobre usuarios en una base de datos SQLite.

## Requisitos

- **Go** (versión 1.23 o superior): Puedes descargar e instalar Go desde [el sitio oficial](https://golang.org/dl/).
- **Git**: Para clonar el repositorio.
- **Postman o cURL**: Herramientas para hacer pruebas de las solicitudes HTTP a la API.

## Configuración del Proyecto

1. **Clona el repositorio**:

   ```bash
   git clone https://github.com/tu-usuario/tu-repositorio.git
   cd tu-repositorio
   ```

2. **Instala las dependencias**:

   Dentro del directorio del proyecto, instala las dependencias, incluyendo `chi` y `gorm`:

   ```bash
   go mod tidy
   ```

3. **Configura la base de datos**:

   La API utiliza SQLite como base de datos. No es necesario configurarla manualmente; la primera vez que inicies la API, se creará un archivo `test.db` en el directorio del proyecto.

## Ejecutar la API

Para iniciar el servidor, usa el siguiente comando:

```bash
go run main.go
```

La API estará en ejecución en `http://localhost:8080`.

## Endpoints Disponibles

### Crear un Usuario

- **URL**: `/users`
- **Método**: `POST`
- **Descripción**: Crea un nuevo usuario.
- **Encabezado de Autenticación**: Requiere el encabezado `X-API-KEY` con el valor `my-secret-key`.

- **Ejemplo de Solicitud**:

  ```json
  POST /users
  Content-Type: application/json
  X-API-KEY: my-secret-key

  Body:
  {
      "name": "Juan Perez",
      "email": "jperez@example.com"
  }
  ```

- **Respuesta Exitosa**:

  ```json
  {
      "ID": 1,
      "name": "Juan Perez",
      "email": "jperez@example.com"
  }
  ```

### Obtener un Usuario por ID

- **URL**: `/users/{id}`
- **Método**: `GET`
- **Descripción**: Obtiene los datos de un usuario específico por su ID.
- **Encabezado de Autenticación**: Requiere el encabezado `X-API-KEY` con el valor `my-secret-key`.

- **Ejemplo de Solicitud**:

  ```json
  GET /users/1
  X-API-KEY: my-secret-key
  ```

- **Respuesta Exitosa**:

  ```json
  {
      "ID": 1,
      "name": "Jane Smith",
      "email": "janesmith@example.com"
  }
  ```

### Errores Comunes

- **401 Unauthorized**: Ocurre cuando falta el encabezado `X-API-KEY` o el valor no coincide con `my-secret-key`.
  
  ```json
  {
      "error": "Unauthorized"
  }
  ```

- **404 Not Found**: Cuando el usuario solicitado no existe.

  ```json
  {
      "error": "User not found"
  }
  ```

## Testing con Postman

1. Abre Postman y crea una nueva solicitud.
2. Configura el **método HTTP** (por ejemplo, `POST` para crear un usuario).
3. En la pestaña de **Headers**, añade `Content-Type: application/json` y `X-API-KEY: my-secret-key`.
4. En la pestaña de **Body**, selecciona el tipo `raw` y formato `JSON`. Luego, pega el cuerpo de la solicitud en formato JSON.
5. Haz clic en **Send** para enviar la solicitud y verifica la respuesta.

## Recursos Adicionales

- [Documentación de Go](https://golang.org/doc/)
- [Documentación de Chi](https://github.com/go-chi/chi)
- [GORM](https://gorm.io/docs/)
