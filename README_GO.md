# QR Code Generator - Go GUI Version

Una aplicación de escritorio multiplataforma para generar códigos QR de URLs, con interfaz gráfica nativa.

## Características

- 🖥️ Interfaz gráfica nativa (Windows y Linux)
- 📝 Campo de entrada para URLs
- 🎨 Visualización instantánea del código QR generado
- 💾 Guardar códigos QR como imágenes PNG
- 🧹 Función de limpiar para empezar de nuevo
- 🚀 Aplicación standalone sin dependencias

## Requisitos Previos

- Go 1.21 o superior
- Dependencias del sistema para Fyne:

### Linux (Ubuntu/Debian)
```bash
sudo apt-get install gcc libgl1-mesa-dev xorg-dev
```

### Linux (Fedora)
```bash
sudo dnf install gcc libXcursor-devel libXrandr-devel mesa-libGL-devel libXi-devel libXinerama-devel libXxf86vm-devel
```

### Windows
No se requieren dependencias adicionales. Asegúrate de tener un compilador GCC (como MinGW-w64 o TDM-GCC).

## Instalación

1. Clona el repositorio:
```bash
git clone <repository-url>
cd qr-generator-script
```

2. Descarga las dependencias:
```bash
go mod download
```

## Uso

### Ejecutar la aplicación

```bash
go run main.go
```

### Compilar la aplicación

Para Linux:
```bash
go build -o qr-generator main.go
./qr-generator
```

Para Windows:
```bash
go build -o qr-generator.exe main.go
qr-generator.exe
```

### Compilar sin ventana de consola (Windows)
```bash
go build -ldflags -H=windowsgui -o qr-generator.exe main.go
```

## Cómo Usar la Aplicación

1. **Ingresar URL**: Escribe o pega la URL que quieres convertir en código QR
2. **Generar**: Haz clic en el botón "Generate QR Code"
3. **Visualizar**: El código QR aparecerá en la ventana
4. **Guardar**: Haz clic en "Save QR Code" para guardar la imagen donde desees
5. **Limpiar**: Usa el botón "Clear" para empezar de nuevo

## Estructura del Proyecto

```
qr-generator-script/
├── main.go          # Aplicación Go con GUI
├── qr.py            # Script Python original
├── go.mod           # Dependencias de Go
├── README.md        # README original
└── README_GO.md     # Este archivo
```

## Tecnologías Utilizadas

- **[Fyne](https://fyne.io/)** - Framework GUI moderno para Go
- **[go-qrcode](https://github.com/skip2/go-qrcode)** - Librería para generar códigos QR
- Go 1.21+

## Comparación con la Versión Python

### Versión Python (qr.py)
- ✅ Simple y ligera
- ❌ URL hardcodeada en el código
- ❌ Sin interfaz gráfica
- ❌ Requiere Python y dependencias

### Versión Go (main.go)
- ✅ Interfaz gráfica intuitiva
- ✅ Entrada dinámica de URLs
- ✅ Visualización en tiempo real
- ✅ Aplicación standalone compilada
- ✅ Multiplataforma (Windows/Linux)
- ✅ Sin dependencias externas en runtime

## Solución de Problemas

### Error: "cannot find package"
```bash
go mod tidy
go mod download
```

### Error de compilación en Linux
Asegúrate de tener las dependencias del sistema instaladas (ver sección Requisitos Previos).

### La ventana no se abre
Verifica que estás ejecutando la aplicación en un entorno con display gráfico (no SSH sin X11).

## Licencia

Este proyecto está disponible para uso libre.

## Contribuciones

Las contribuciones son bienvenidas. Por favor, abre un issue o pull request.
