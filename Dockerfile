# Imagen base
FROM golang:1.20.6

# Directorio de trabajo
WORKDIR /app

# Copiar archivos go mod y sum
COPY go.mod go.sum ./

# Descargar todas las dependencias
RUN go mod download

# Copiar el código fuente de la aplicación al contenedor
COPY . .

# Compilar la aplicación
RUN go build -o main .

# Exponer el puerto 8000 para la aplicación
EXPOSE 8000

# Comando para ejecutar la aplicación
CMD ["./main"]
