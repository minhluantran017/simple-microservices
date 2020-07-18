from flask import Flask

app = Flask(__name__)

@app.route('/')
def index():
    gittag = 'SED_GIT_TAG'
    githash = 'SED_GIT_HASH'
    index = "<!DOCTYPE html>"
    index += "<html>"
    index += "<head><title>Python service</title><style>body { background-color: navy; }</style></head>"
    index += "<body>"
    index += "<h1 style=\"text-align:center\">This is a simple Python service!!</h1>"
    index += "<p style=\"text-align:center\">Git tag: "+ gittag +"</p>"
    index += "<p style=\"text-align:center\">Git commit ID: "+ githash +"</p>"
    index += "</body>"
    index += "</html>"
    return index

if __name__ == "__main__":
    app.run(port=8080)