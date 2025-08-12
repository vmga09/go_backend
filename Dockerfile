FROM golang:1.24-bullseye

WORKDIR /app

# Instalar Air para recarga automática
RUN go install github.com/air-verse/air@latest

# Copiar y descargar dependencias
#COPY go.mod go.sum ./
#RUN go mod download

# Copiar el resto del código
COPY . .

# Inicializar módulo Go si no existe y descargar dependencias
RUN [ ! -f go.mod ] && go mod init example.com/go-api || true
RUN go mod tidy

EXPOSE 8080

# Usar Air para desarrollo en caliente
CMD ["air", "-c", ".air.toml"]
#CMD ["air"]