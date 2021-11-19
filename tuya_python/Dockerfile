FROM python:3.7.3

RUN mkdir /pythonbin/

COPY tuya.py /pythonbin

COPY entrypoint.sh /pythonbin

RUN pip3 install tinytuya

RUN pip3 install prometheus-client

EXPOSE 9185

ENTRYPOINT ["/pythonbin/entrypoint.sh"]

