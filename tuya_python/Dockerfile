FROM python:3.7.3

RUN mkdir /pythonbin/

COPY tuya.py /pythonbin

COPY entrypoint.sh /pythonbin

COPY requirements.txt /pythonbin

RUN pip3 install -r /pythonbin/requirements.txt

EXPOSE 9185

WORKDIR /pythonbin

ENTRYPOINT ["/pythonbin/entrypoint.sh"]