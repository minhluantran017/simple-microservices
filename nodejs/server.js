var http = require('http');
var url = require('url');
var server = http.createServer(function(request, response) {
  var path = url.parse(request.url).pathname;
  var githash = "${GIT_HASH}"
  var background = "blue"
  if ((githash % 2) == 1) {
    background = "green"
  }
  response.write("<!DOCTYPE html>");
  response.write("<html>");
  response.write("<head><title>Sample NodeJS service</title><style>body { background-color: "+ background +"; }</style></head>");
  response.write("<body>");
  response.write("<h1 style=\"text-align:center\">This is a simple NodeJS services!!</h1>");
  response.write("<h1 style=\"text-align:center\">Git commit ID:"+ githash +"</h1>");
  response.write("</body>");
  response.write("</html>");
  response.end();
});
server.listen(8082);