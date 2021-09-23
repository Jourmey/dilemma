# dilemma
爬视频平台，使用docker-compose部署前后端、此项目为后端。

前端项目：https://github/jourmey/dilemma-vue3


###使用方法
```
1.通过视频链接创建任务，任务会解析对应的视频格式；
2.下载需要的格式；
3.下载好的视频通过静态文件服务获取。
```

###使用到的技术
```
1.you-get
2.vue3.0
3.go-zero
```

###浏览地址
```
http://82.156.195.170/
```

###快速使用
```
1.下载Releases
2.解压后的目录为：
.
├── dilemma             #后端
│   └── etc
│       ├── dilemma-dev.json
│       └── dilemma.json
├── dilemma-web         #前端
│   ├── dist        #vue打包文件
│   └── nginx.conf  #nginx部署文件
├── docker-compose.yaml     #docker-compose
├── mysql               #mysql配置文件
│   ├── conf
│   │   └── my.cnf
│   ├── db
│   └── init
└── workspace            #下载后的视频路径
    ├── Bilibili
    ├── Vimeo
    └── YouTube

3.在目录下执行 docker-compose up
```
