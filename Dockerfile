FROM pypy:latest

RUN mkdir /pythonbin/

RUN mkdir /tmp/config

COPY requirements.txt /pythonbin

RUN pip3 install -r /pythonbin/requirements.txt

COPY tuya.py /pythonbin

EXPOSE 9185

RUN pwd

WORKDIR /pythonbin

RUN ls

# ENTRYPOINT ["sh", "-c", "tail -f /dev/null"]
# ENTRYPOINT ["/pythonbin/entrypoint.sh"]

CMD [ "python", "/pythonbin/tuya.py" ]