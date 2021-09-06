FROM python:alpine
 
RUN mkdir /app /workspace
ADD . /app/
# 安装you-get
RUN cd /app/you-get/ && ./setup.py install