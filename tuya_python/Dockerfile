FROM python:3.7.3

ADD tuya.py /

RUN pip3 install tinytuya

RUN pip3 install prometheus-client

EXPOSE 9185

CMD [ "python3", "./tuya.py" ]

