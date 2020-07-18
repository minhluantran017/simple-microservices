from http.server import HTTPServer
from http.server import BaseHTTPRequestHandler
import os

HOST_NAME = 'localhost'
PORT = 8080
gittag = 'SED_GIT_TAG'
githash = 'SED_GIT_HASH'

class Server(BaseHTTPRequestHandler):
    def do_GET(self):
        try:
            self.send_response(200)
            self.end_headers()
            self.wfile.write("<!DOCTYPE html>")
            self.wfile.write("<html>")
            self.wfile.write("<head><title>Python service</title><style>body { background-color: darkblue; }</style></head>")
            self.wfile.write("<body>")
            self.wfile.write("<h1 style=\"text-align:center\">This is a simple Python services!!</h1>")
            self.wfile.write("<h1 style=\"text-align:center\">Git tag:"+ gittag +"</h1>")
            self.wfile.write("<h1 style=\"text-align:center\">Git commit ID:"+ githash +"</h1>")
            self.wfile.write("</body>")
            self.wfile.write("</html>")
        except:
            self.send_error(404) 
if __name__ == "__main__":
    httpd = HTTPServer((HOST_NAME,PORT),Server)
    try:
        httpd.serve_forever()
    except KeyboardInterrupt:
        pass
    httpd.server_close()
