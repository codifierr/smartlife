FROM pypy:latest

RUN mkdir /pythonbin/

RUN mkdir /tmp/config

COPY tuya.py /pythonbin

COPY entrypoint.sh /pythonbin

COPY requirements.txt /pythonbin

RUN pip3 install -r /pythonbin/requirements.txt

EXPOSE 9185

RUN pwd

WORKDIR /pythonbin

RUN pwd

ENTRYPOINT ["/pythonbin/entrypoint.sh"]