# API de Clima en Python (Flask)

Esta es una API simple que permite obtener el clima de cualquier ciudad utilizando el servicio de OpenWeatherMap.

## Cómo ejecutar

1. Clona el repositorio:
   ```bash
   git clone https://github.com/alexismendozaa/proyecto6.git
   cd proyecto6

docker build -t alexismendozaa/flask-clima .
docker run -p 5000:5000 alexismendozaa/flask-clima

Abre el navegador y accede a http://localhost:5000/clima/[ciudad] para ver el clima de la ciudad.