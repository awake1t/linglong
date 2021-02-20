
CREATE DATABASE IF NOT EXISTS linglong default charset utf8 COLLATE utf8_general_ci;
use linglong;


/* ip资产列表 */
DROP TABLE IF EXISTS `iplist`;
CREATE TABLE `iplist` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `ip` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  DEFAULT NULL COMMENT '资产ip',
  `port` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '开放端口',
  `protocol` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '协议',
  `cms` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'cms',
  `language` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '代码语言',
  `portnum` int(10) unsigned DEFAULT 0 COMMENT '开放端口数量',
  `url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'url',
  `loginurl` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '跳转的url',
  `title` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'title',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


/* 后台资产列表 */
DROP TABLE IF EXISTS `webloginlist`;
CREATE TABLE `webloginlist` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `ip` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  DEFAULT NULL COMMENT '资产ip',
  `port` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '端口',
  `protocol` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '协议',
   `cms` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'cms',
  `language` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '代码语言',
  `url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'url',
  `loginurl` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '跳转的url',
  `title` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'title',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

/*定时任务列表*/
DROP TABLE IF EXISTS `task`;
CREATE TABLE `task` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `task_name` varchar(50) NOT NULL DEFAULT '' COMMENT '任务名称',
  `task_type` varchar(50) NOT NULL DEFAULT '' COMMENT '任务类型',
  `task_cycle` varchar(50) NOT NULL DEFAULT '' COMMENT '任务周期 now/执行一次 day每天xx点执行',
  `description` varchar(200) NOT NULL DEFAULT '' COMMENT '任务描述',
  `cron_spec` varchar(100) NOT NULL DEFAULT '' COMMENT '时间表达式',
  `concurrent` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否只允许一个实例',
  `command` text NOT NULL COMMENT '命令详情',
  `arge` text NOT NULL COMMENT 'url参数',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0停用 1启用',
  `timeout` smallint(6) NOT NULL DEFAULT '0' COMMENT '超时设置',
  `vuln_num` int(11) NOT NULL DEFAULT '0' COMMENT '漏洞数量',
  `execute_times` int(11) NOT NULL DEFAULT '0' COMMENT '累计执行次数',
  `prev_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '上次执行时间',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


/*定时任务执行记录*/
DROP TABLE IF EXISTS `task_log`;
CREATE TABLE `task_log` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `task_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '任务ID',
  `output` mediumtext NOT NULL COMMENT '任务输出',
  `error` text NOT NULL COMMENT '错误信息',
  `status` tinyint(4) NOT NULL COMMENT '状态',
  `all_num` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '任务总数量',
  `succes_num` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '成功数量',
  `userdict` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户名字典id',
  `passdict` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '密码字典id',
  `run_time` varchar(200) NOT NULL DEFAULT '' COMMENT '运行时间',
  `process_time` int(11) NOT NULL DEFAULT '0' COMMENT '消耗时间/毫秒',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


/*端口爆破结果库*/
DROP TABLE IF EXISTS `portbruteres`;
CREATE TABLE `portbruteres` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `task_id` int NOT NULL DEFAULT 0 COMMENT '定时任务id',
  `task_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '定时任务执行时间',
  `ip` varchar(150) NOT NULL DEFAULT '' COMMENT 'ip',
  `protocol` varchar(150) NOT NULL DEFAULT '' COMMENT '爆破协议',
  `port` int(20) NOT NULL DEFAULT 0 COMMENT 'port',
  `vulntype` tinyint(4) NOT NULL DEFAULT '0' COMMENT '漏洞类型 1弱口令 2未授权 3空口令',
  `user` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
  `pass` varchar(50) NOT NULL DEFAULT '' COMMENT '密码',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

/*全平台任务日志记录*/
DROP TABLE IF EXISTS `log`;
CREATE TABLE `log` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `task_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '任务ID',
  `task_name` varchar(200) NOT NULL DEFAULT '' COMMENT '任务名称',
  `task_type` varchar(200) NOT NULL DEFAULT '' COMMENT '任务类型/ masscan/title识别/端口',
  `all_num` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '任务总数量',
  `succes_num` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '成功数量',
  `run_time` varchar(200) NOT NULL DEFAULT '' COMMENT '运行时间',
  `status` tinyint(4) NOT NULL COMMENT '状态 1完成 2错误',
  `error` text NOT NULL COMMENT '错误信息',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


CREATE TABLE `auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


/*环境配置数据库*/
DROP TABLE IF EXISTS `setting`;
CREATE TABLE `setting` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `login_word` text NOT NULL  COMMENT '敏感后台关键字',
  `login_url` text NOT NULL COMMENT '敏感后台url',
  `masscan_deltime` int NOT NULL DEFAULT 100 COMMENT 'mass删除周期',
  `masscan_thred` int NOT NULL DEFAULT 0 COMMENT 'mass线程',
  `masscan_ip` MEDIUMTEXT NOT NULL  COMMENT 'mass要扫描的列表',
  `masscan_port` text NOT NULL  COMMENT 'mass要扫描的端口',
  `masscan_white` text NOT NULL  COMMENT 'mass不要扫描的列表',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


-- /*存储指纹表*/
-- DROP TABLE IF EXISTS "fingerprint";
-- CREATE TABLE "fingerprint"(
-- `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
--   `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  DEFAULT NULL COMMENT '指纹名称',
--   `description` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  DEFAULT NULL COMMENT 'description',
--   `cookies` text NOT NULL COMMENT 'cookies',
--   `headers` text NOT NULL COMMENT 'headers',
--   `meta` text NOT NULL COMMENT 'meta',
--   `meta` text NOT NULL COMMENT 'html',
--   `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
--   `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
--   PRIMARY KEY (`id`)
-- ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

/*指纹*/
DROP TABLE IF EXISTS `finger`;
CREATE TABLE `finger`(
    `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `name`     varchar(512)        NOT NULL DEFAULT '' COMMENT '指纹名称',
    `description`     varchar(512)        NOT NULL DEFAULT '' COMMENT 'description',
    `finger` text NOT NULL COMMENT 'finger',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

/*xrayRes*/
DROP TABLE IF EXISTS `xrayres`;
CREATE TABLE `xrayres`(
    `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `url`     varchar(512)        NOT NULL DEFAULT '' COMMENT 'url',
    `poc`     varchar(512)        NOT NULL DEFAULT '' COMMENT 'poc',
    `hash`     varchar(512)        NOT NULL DEFAULT '' COMMENT 'hash',
    `snapshot` text NOT NULL COMMENT 'snapshot',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `hash` (`hash`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


INSERT INTO `auth` (`id`, `username`, `password`) VALUES (null, 'linglong', 'linglong5s');

INSERT INTO `setting` (`id`, `login_word`,`login_url`,`masscan_thred`,`masscan_ip`,`masscan_port`,`masscan_white`) VALUES
 (null,  'login
phpmyadmin', 'admin
login', 1000, '192.168.0.0/24,192.168.1.0/24,192.168.3.0/24,192.10.0.0/16', '22,80,443,445,3306,8000,8080,8088,8090,1521,5432,6379', '10.10.10.10');


INSERT INTO `finger` (`name`,`description`,`finger`,`created_time`,`updated_time`) VALUES ('PHPMyAdmin','PHPMyAdmin','"PHPMyAdmin": {
      "html": "<title>phpMyAdmin </title>",
      "html": "/themes/pmahomme/img/logo_right.png",
      "cookies": {
        "phpMyAdmin": ""
      }
    }','20210213225711','20210213225711');

INSERT INTO `finger` (`name`,`description`,`finger`,`created_time`,`updated_time`) VALUES ('Zabbix','Zabbix','"Zabbix": {
      "html": "images/general/zabbix.ico",
      "html": "Zabbix SIA",
      "cookies": {
        "zbx_sessionid": ""
      }
    }','20210213225759','20210213225759');

INSERT INTO `finger` (`name`,`description`,`finger`,`created_time`,`updated_time`) VALUES ('Shiro','Shiro','"Shiro" : {
  "cookies" : {
  "rememberMe": ""
  },
  "html": ""
 }','20210213225839','20210213225839');

INSERT INTO `finger` (`name`,`description`,`finger`,`created_time`,`updated_time`) VALUES ('Alibaba-Druid','Alibaba-Druid','"AlibabaDruid": {
      "html": "Druid Stat Index"
}','20210213225929','20210213225929');

 INSERT INTO `finger` (`name`,`description`,`finger`,`created_time`,`updated_time`) VALUES ('kibana','kibana',' "Kibana": {
      "headers": {
        "kbn-name": "kibana",
        "kbn-version": "^([\\d.]+)$\\;version:\\1"
      },
      "html": "<title>Kibana</title>"
    }','20210215201418','20210215201418');






