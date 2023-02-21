FROM python:3.11-buster
# TARGETPLATFORM

WORKDIR /app
ENV PIP_INDEX_URL https://mirrors.aliyun.com/pypi/simple/

COPY requirements.txt ./
RUN pip install --upgrade pip && pip install -r requirements.txt
COPY main.py ./

CMD python main.py
