USE messagepush;
//消息推送
--
-- Table structure for table `messagepush`
--
CREATE TABLE `apps`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL COMMENT '应用名称',
    `app_key` varchar(255) NOT NULL COMMENT '应用key',
    `app_id` varchar(255) NOT NULL COMMENT '应用id',
    `app_type` varchar(255) NOT NULL COMMENT '应用类型',
    `is_del` int(11) NOT NULL COMMENT '是否删除',
    `create_time` datetime NOT NULL,
    `update_time` datetime NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `clients`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL COMMENT '客户端名称',
    `client_key` varchar(255) NOT NULL COMMENT '客户端key',
    `app_id` varchar(255) NOT NULL COMMENT '应用id',
    `client_id` varchar(255) NOT NULL COMMENT '客户端id',
    `client_type` varchar(255) NOT NULL COMMENT '客户端类型',
    `is_del` int(11) NOT NULL COMMENT '是否删除',
    `create_time` datetime NOT NULL,
    `update_time` datetime NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `messages`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `app_id` int(11) NOT NULL COMMENT '应用id',
    `client_ids` json NOT NULL COMMENT '对象id',
    `content` varchar(4096) NOT NULL COMMENT '消息内容',
    `delay` int(11) NOT NULL COMMENT '延迟时间',
    `send_time` datetime NOT NULL COMMENT '发送时间',
    `send_count` int(11) NOT NULL COMMENT '发送次数',
    `status` int(11) NOT NULL COMMENT '发送状态',
    `is_del` int(11) NOT NULL COMMENT '是否删除',
    `create_time` datetime NOT NULL,
    `update_time` datetime NOT NULL,
    PRIMARY KEY (`id`)
);

insert into apps (name, app_key, app_id, app_type, is_del,create_time,update_time)
    value('name','key','000','one_much',0,current_time,current_time);

insert into clients (name, client_key, app_id, client_id, client_type, is_del, create_time, update_time)
values ('name','key','000','000','one_much',0,current_time,current_time),
       ('name','key','000','001','one_much',0,current_time,current_time);
