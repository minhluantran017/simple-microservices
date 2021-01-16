FROM python:3.6-alpine

RUN pip install flask
COPY . .

EXPOSE 8081
CMD [ "python", "server.py" ]