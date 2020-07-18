var http = require('http');
var url = require('url');
var server = http.createServer(function(request, response) {
  var path = url.parse(request.url).pathname;
  var gittag = "SED_GIT_TAG";
  var githash = "SED_GIT_HASH";

  response.write("<!DOCTYPE html>");
  response.write("<html>");
  response.write("<head><title>NodeJS service</title><style>body { background-color: ForestGreen; }</style></head>");
  response.write("<body>");
  response.write("<h1 style=\"text-align:center\">This is a simple NodeJS service!!</h1>");
  response.write("<h1 style=\"text-align:center\">Git tag: "+ gittag +"</h1>");
  response.write("<h1 style=\"text-align:center\">Git commit ID: "+ githash +"</h1>");
  response.write("</body>");
  response.write("</html>");
  response.end();
});
server.listen(8082);