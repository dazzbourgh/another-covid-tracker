FROM python:3.8-alpine3.13

RUN apk add --update curl gcc g++ git libffi-dev openssl-dev python3-dev build-base linux-headers \
    && rm -rf /var/cache/apk/*
RUN ln -s /usr/include/locale.h /usr/include/xlocale.h

ENV PYTHONUNBUFFERED=1 ENVIRONMENT=pre APP_HOME=/microservice/
ENV PYMS_CONFIGMAP_FILE="$APP_HOME"config-docker.yml
RUN mkdir $APP_HOME && adduser -S -D -H python

RUN chown -R python $APP_HOME
ADD . $APP_HOME
WORKDIR $APP_HOME
RUN pip install pipenv
RUN pipenv lock --requirements > requirements.txt
RUN pip install -r ./requirements.txt
RUN pip install gevent==21.1.2 gunicorn==20.0.4

EXPOSE 5000
USER python

CMD ["gunicorn", "--workers", "8", "--log-level", "INFO", "--bind", "0.0.0.0:5000", "manage:app"]
