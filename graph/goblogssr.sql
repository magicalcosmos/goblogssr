DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` int PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
  `user_id` int(11) DEFAULT '0' COMMENT '用户ID',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '标题',
  `category_id` int(11) NOT NULL COMMENT '分类ID',
  `tag` varchar(255) NOT NULL DEFAULT '' COMMENT 'Tag',
  `remark` varchar(500) NOT NULL DEFAULT '' COMMENT '摘要',
  `desc` longtext NOT NULL COMMENT '详情',
  `pv` int(255) DEFAULT '0' COMMENT 'px',
  `created` datetime DEFAULT NULL COMMENT '创建时间',
  `updated` datetime DEFAULT NULL COMMENT '更新时间',
  `status` int(11) DEFAULT '1' COMMENT '1可用，2禁用，3删除',
  `review` int(11) DEFAULT '0' COMMENT '评论'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE `article` ADD `recommend` INT(4) NOT NULL DEFAULT '0' COMMENT '是否顶置，0否；1是，默认否';
ALTER TABLE `article` ADD COLUMN `like`  int(11) NOT NULL DEFAULT 0 COMMENT '点赞数量' AFTER `recommend`;
ALTER TABLE `article` ADD COLUMN `html`  longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'html' AFTER `desc`;
-- --------------------------------------------------------

--
-- 表的结构 `category`
--
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` int PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
  `name` varchar(255) DEFAULT NULL,
  `pid` int(11) DEFAULT '0' COMMENT '父ID',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  `status` int(11) DEFAULT '1' COMMENT '状态1正常，2删除'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------
--
-- 表的结构 `user`
--
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
  `username` varchar(255) DEFAULT NULL COMMENT '用户名',
  `password` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  `status` int(11) DEFAULT '1' COMMENT '状态'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- --------------------------------------------------------
--
-- 表的结构 `log`
--
DROP TABLE IF EXISTS `log`;
CREATE TABLE `log` (
  `id` int(11) NOT NULL COMMENT 'ID',
  `ip` varchar(50) NOT NULL,
  `city` varchar(50) NOT NULL,
  `create` datetime NOT NULL,
  `user_agent` varchar(255) NOT NULL,
  `page` varchar(255) DEFAULT NULL,
  `uri` varchar(500) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `message`
--
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message` (
  `id` int(11) NOT NULL COMMENT 'ID',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '名字',
  `review` varchar(500) NOT NULL DEFAULT '' COMMENT '评论',
  `reply` varchar(500) NOT NULL COMMENT '回复',
  `site` varchar(500) NOT NULL COMMENT '网址',
  `created` datetime DEFAULT NULL COMMENT '创建时间',
  `updated` datetime DEFAULT NULL COMMENT '回复时间',
  `status` int(11) DEFAULT '1' COMMENT '1可用，2禁用'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `review`
--
DROP TABLE IF EXISTS `review`;
CREATE TABLE `review` (
  `id` int(11) NOT NULL COMMENT 'ID',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '名字',
  `review` varchar(500) NOT NULL DEFAULT '' COMMENT '评论',
  `reply` varchar(500) NOT NULL COMMENT '回复',
  `site` varchar(500) NOT NULL COMMENT '网址',
  `created` datetime DEFAULT NULL COMMENT '创建时间',
  `updated` datetime DEFAULT NULL COMMENT '回复时间',
  `status` int(11) DEFAULT '1' COMMENT '1可用，2禁用',
  `article_id` int(11) DEFAULT '1' COMMENT '文章ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------



DROP TABLE IF EXISTS `setting`;
CREATE TABLE `setting` (
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '标题',
  `value` text NOT NULL COMMENT '详情',
  PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `wechat_user`;
CREATE TABLE `wechat_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `openid` varchar(32) NOT NULL COMMENT 'open ID',
  `nickname` varchar(255) DEFAULT NULL COMMENT '用户名',
  `sex` int(11) DEFAULT '0' COMMENT '性别',
  `city` varchar(32) DEFAULT NULL COMMENT '城市',
  `country` varchar(32) DEFAULT NULL COMMENT '国家',
  `province` varchar(32) DEFAULT NULL COMMENT '省份',
  `language` varchar(32) DEFAULT NULL COMMENT '语种',
  `headimgurl` varchar(255) DEFAULT NULL COMMENT '头像',
  `subscribe_time` int(11) DEFAULT NULL COMMENT '订阅时间',
  `created` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '栏目名称',
  `target` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '链接打开方式',
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '链接URL',
  `sort` int(11) DEFAULT '100' COMMENT '排序',
  `pid` int(11) DEFAULT '0' COMMENT '顶级栏目',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `link`;
CREATE TABLE `link` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '名称',
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT 'URL',
  `sort` int(11) DEFAULT '100' COMMENT '排序',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
