FROM python:3.6-alpine
RUN pip install flask
COPY . .
EXPOSE 8080
CMD [ "python", "server.py" ]