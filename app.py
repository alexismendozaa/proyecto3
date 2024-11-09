from flask import Flask, jsonify
import requests

app = Flask(__name__)

# Endpoint para obtener el clima de una ciudad
@app.route('/clima/<ciudad>', methods=['GET'])
def obtener_clima(ciudad):
    api_key = "tu_api_key_aqui"
    url = f"http://api.openweathermap.org/data/2.5/weather?q={ciudad}&appid={api_key}&units=metric&lang=es"
    response = requests.get(url)
    data = response.json()

    if data["cod"] != "404":
        clima = {
            "ciudad": data["name"],
            "temperatura": data["main"]["temp"],
            "descripcion": data["weather"][0]["description"]
        }
        return jsonify(clima)
    else:
        return jsonify({"mensaje": "Ciudad no encontrada"}), 404

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=5000)
