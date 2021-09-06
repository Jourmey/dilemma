create table task
(
    id          int(11) unsigned auto_increment comment '主键'
        primary key,
    url         varchar(255) default ''                not null comment '链接',
    signatures  varchar(32)  default ''                not null comment '特征码',
    tag         int          default 0                 not null comment '标签',
    status      int          default 0                 not null comment '任务状态 0未处理 1获取信息',
    title       varchar(255) default ''                not null comment '标题',
    site        varchar(32)  default ''                not null comment '平台',
    create_time timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time timestamp    default CURRENT_TIMESTAMP not null comment '修改时间',
) comment '任务表';

create
index task__status
    on task (status);

create
index task_url_index
    on task (url);


create table task_info
(
    id          int(11) unsigned auto_increment comment '主键'
        primary key,
    task_id     int(11) default '' not null comment '关联任务',
    format      varchar(10) default ''                not null comment '链接:dash-flv360',
    container   varchar(10) default ''                not null comment '类型:mp4',
    quality     varchar(10) default ''                not null comment '质量:流畅 360P',
    size        int         default 0                 not null comment '任务大小',
    create_time timestamp   default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time timestamp   default CURRENT_TIMESTAMP not null comment '修改时间',
) comment '任务信息表';


create table video
(
    id           int(11) unsigned auto_increment comment '主键'
        primary key,
    task_info_id int(11) default '' not null comment '关联任务信息',
    path         varchar(255) default ''                not null comment '路径',
    create_time  timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time  timestamp    default CURRENT_TIMESTAMP not null comment '修改时间',
) comment '视频下载表';