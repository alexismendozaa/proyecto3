# Usa una imagen oficial de Python
FROM python:3.9-slim

# Establece el directorio de trabajo
WORKDIR /app

# Copia los archivos del proyecto
COPY . .

# Instala las dependencias
RUN pip install --no-cache-dir -r requirements.txt

# Expón el puerto 5000
EXPOSE 5000

# Ejecuta la aplicación
CMD ["python", "app.py"]
