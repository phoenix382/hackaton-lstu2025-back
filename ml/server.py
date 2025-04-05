from http.server import BaseHTTPRequestHandler, HTTPServer
import json
import threading
import requests
import time
import random

class Handler(BaseHTTPRequestHandler):
    def do_POST(self):
        try:
            content_length = int(self.headers['Content-Length'])
            data = json.loads(self.rfile.read(content_length))
            
            if not data.get('token') or len(data['token']) != 8:
                self.send_error(400, "Invalid token format")
                return
                
            token = data['token']
            
            print(f"[Сервер] Получен запрос (токен: {token})")
            
            # Имитация неопределённого времени обработки (1-5 секунд)
            processing_time = random.uniform(1, 5)
            print(f"[Сервер] Обработка займет ~{processing_time:.2f} сек")
            
            # Запуск обработки в отдельном потоке
            def process_request():
                time.sleep(processing_time)  # Имитация долгой обработки
                
                processed_data = {
                    'original': data.get('payload', {}),
                    'processed': True,
                    'processing_time': processing_time,
                    'server_note': 'Данные обработаны',
                    'add_data': 'ещё что-то добавили'
                }
                
                if 'callback_url' in data:
                    try:
                        requests.post(
                            data['callback_url'],
                            json={
                                'token': token,
                                'result': processed_data
                            },
                            timeout=5
                        )
                        print(f"[Сервер] Ответ отправлен на {data['callback_url']}")
                    except Exception as e:
                        print(f"[Сервер] Ошибка отправки ответа: {str(e)}")
            
            threading.Thread(target=process_request).start()
            
            self._send_response(202, {
                'status': 'processing',
                'estimated_time': processing_time,
                'token': token
            })
            
        except Exception as e:
            self.send_error(500, f"Server error: {str(e)}")
    
    def _send_response(self, code, data):
        self.send_response(code)
        self.send_header('Content-type', 'application/json')
        self.end_headers()
        self.wfile.write(json.dumps(data).encode())

def run_server(port=8000):
    server_address = ('', port)
    httpd = HTTPServer(server_address, Handler)
    print(f"[Сервер] Запущен на порту {port}")
    httpd.serve_forever()

if __name__ == '__main__':
    run_server()