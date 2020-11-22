CREATE DATABASE gungnir charset utf8;

USE gungnir;

CREATE TABLE `templates` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '模板ID',
  `name` varchar(45) NOT NULL COMMENT '模板名称',
  `language` varchar(45) NOT NULL COMMENT '模板使用语言',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='模板表';

CREATE TABLE `modules` (
    `id` int NOT NULL AUTO_INCREMENT COMMENT '模块ID',
    `template_id` int NOT NULL COMMENT '模板ID',
    `index` int NOT NULL COMMENT '执行顺序',
    `name` varchar(45) NOT NULL COMMENT '模块名称',
    `before_script_path` varchar(255) NOT NULL,
    `after_script_path` varchar(255) NOT NULL,
    `package_path` varchar(255) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `name_UNIQUE` (`template_id`,`index`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='模板表';