FROM youget:v1.0

COPY . /app/dilemma/
# 设置工作目录
WORKDIR /app/dilemma/
# 端口
EXPOSE 8081
# 设置启动命令
CMD ["./dilemma","-f=etc/dilemma.json"]
