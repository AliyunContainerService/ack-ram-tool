FROM python:3.11-buster
# TARGETPLATFORM

WORKDIR /app

COPY requirements.txt ./
RUN pip install -r requirements.txt
COPY main.py ./

CMD python main.py
