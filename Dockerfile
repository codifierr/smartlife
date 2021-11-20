FROM pypy:latest

RUN mkdir /pythonbin/

RUN mkdir /tmp/config

COPY requirements.txt /pythonbin

RUN pip3 install -r /pythonbin/requirements.txt

COPY tuya.py /pythonbin

EXPOSE 9185

WORKDIR /pythonbin

CMD [ "python", "/pythonbin/tuya.py" ]
