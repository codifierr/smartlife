FROM pypy:latest

RUN mkdir /pythonbin/

COPY tuya.py /pythonbin

COPY requirements.txt /pythonbin

RUN pip3 install -r /pythonbin/requirements.txt

EXPOSE 9185

RUN pwd

CMD [ "python", "/pythonbin/tuya.py" ]