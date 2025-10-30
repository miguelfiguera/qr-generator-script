# QR Code Generator

Este repositorio contiene dos implementaciones para generar códigos QR de URLs:

## 📱 Versión Python (Script Simple)

**Archivo:** `qr.py`

Un script simple de Python que genera un código QR de una URL hardcodeada.

### Uso:
```bash
python qr.py
```

El script generará un archivo `qrcode.png` con el código QR de la URL especificada en el código.

---

## 🖥️ Versión Go (Aplicación GUI)

**Archivo:** `main.go`

Una aplicación de escritorio completa con interfaz gráfica que permite generar códigos QR de forma interactiva.

### Características:
- ✨ Interfaz gráfica nativa (Windows/Linux)
- 📝 Entrada dinámica de URLs
- 👁️ Visualización instantánea del QR generado
- 💾 Guardar códigos QR en cualquier ubicación
- 🧹 Función de limpiar y empezar de nuevo
- 🚀 Aplicación standalone sin dependencias

### Uso Rápido:
```bash
# Ejecutar directamente
go run main.go

# O compilar y ejecutar
go build -o qr-generator main.go
./qr-generator
```

### Documentación Completa:
Para instrucciones detalladas de instalación y uso de la versión Go, consulta [README_GO.md](README_GO.md)

---

## ¿Cuál versión usar?

- **Usa la versión Python** si solo necesitas generar un QR rápidamente desde línea de comandos
- **Usa la versión Go** si prefieres una interfaz gráfica intuitiva y la capacidad de generar múltiples QRs fácilmente

---

You can toy with it as much as you like.
