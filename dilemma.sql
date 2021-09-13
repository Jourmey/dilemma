CREATE TABLE `task`
(
    `id`          int(14) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `url`         varchar(255) NOT NULL DEFAULT '' COMMENT '链接',
    `signatures`  varchar(32)  NOT NULL DEFAULT '' COMMENT '特征码',
    `tag`         tinyint(8) NOT NULL DEFAULT '0' COMMENT '标签',
    `status`      tinyint(2) NOT NULL DEFAULT '0' COMMENT '任务状态 0未处理 1处理中 2获取信息 3获取失败',
    `title`       varchar(255) NOT NULL DEFAULT '' COMMENT '标题',
    `site`        varchar(32)  NOT NULL DEFAULT '' COMMENT '平台',
    `create_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='任务表';

ALTER TABLE `task`
    ADD INDEX task_status_index (`status`);
ALTER TABLE `task`
    ADD INDEX task_url_index (`url`);


create table `task_info`
(
    `id`          int(14) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `task_id`     int(14) NOT NULL DEFAULT '0' COMMENT '关联任务',
    `format`      varchar(32) NOT NULL DEFAULT '' COMMENT '链接:dash-flv360',
    `container`   varchar(32) NOT NULL DEFAULT '' COMMENT '类型:mp4',
    `quality`     varchar(32) NOT NULL DEFAULT '' COMMENT '质量:流畅 360P',
    `size`        int(16) NOT NULL DEFAULT '0' COMMENT '任务大小',
    `status`      tinyint(2) NOT NULL DEFAULT '0' COMMENT '任务状态 0未处理 1处理中 2获取信息 3获取失败',
    `create_time` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='任务信息表';

ALTER TABLE `task_info`
    ADD INDEX task_info_task_id_index (`task_id`);


create table `video`
(
    `id`           int(14) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `task_info_id` int(11) NOT NULL DEFAULT '0' COMMENT '关联任务信息',
    `path`         varchar(255) NOT NULL DEFAULT '' COMMENT '路径',
    `title`        varchar(255) NOT NULL DEFAULT '' COMMENT '标题',
    `create_time`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='视频下载表';

ALTER TABLE `video`
    ADD INDEX video_task_info_id_index (`task_info_id`);
