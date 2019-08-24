import os
import SimpleHTTPServer
import SocketServer

PORT = os.getenv("PORT")

Handler = SimpleHTTPServer.SimpleHTTPRequestHandler

httpd = SocketServer.TCPServer(("", int(PORT)), Handler)

print "serving at port", PORT
httpd.serve_forever()
