DROP SCHEMA IF EXISTS `blogger`;
CREATE SCHEMA `blogger` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

use blogger;

-- --------------------------------------------------------
--
-- 表的结构 `category`
--
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(255) DEFAULT NULL,
  `parent_id` bigint DEFAULT NULL,
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  `status` int(11) DEFAULT '1' COMMENT '1可用，2禁用，3删除',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_category_parent` (`parent_id`),
  CONSTRAINT `fk_category_parent` FOREIGN KEY (`parent_id`) REFERENCES `category` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------
--
-- 表的结构 `user`
--
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `username` varchar(255) DEFAULT NULL COMMENT '用户名',
  `password` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `status` int(11) DEFAULT '1' COMMENT '1可用，2禁用，3删除',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `uq_email` (`email` ASC)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `article_article`
--
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint NOT NULL  COMMENT '创建者ID',
  `category_id` bigint NOT NULL  COMMENT '目录ID',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '标题',
  `brief` longtext CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '文章摘要',
  `content`  longtext CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '文章内容',
  `published` tinyint(1) NOT NULL DEFAULT '0' COMMENT '发布状态，0草稿；1发布;2存档',
  `tag` varchar(255) NOT NULL DEFAULT '' COMMENT 'Tag',
  `pv` int(255) DEFAULT '0' COMMENT 'px',
  `review` int(11) DEFAULT '0' COMMENT '评论',
  `recommend` int(4) NOT NULL DEFAULT '0' COMMENT '是否顶置，0否；1是，默认否',
  `like`  int(11) NOT NULL DEFAULT 0 COMMENT '点赞数量',
  `status` int(11) DEFAULT '1' COMMENT '1可用，2禁用，3删除',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `publish_time` datetime DEFAULT NULL COMMENT '发布时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_article_user` (`user_id`),
  KEY `idx_article_category` (`category_id`),
  CONSTRAINT `fk_article_category` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`),
  CONSTRAINT `fk_article_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------
--
-- 表的结构 `article_category`
--
DROP TABLE IF EXISTS `article_category`;
CREATE TABLE `article_category` (
  `article_id` BIGINT NOT NULL,
  `category_id` BIGINT NOT NULL,
  PRIMARY KEY (`article_id`, `category_id`),
  INDEX `idx_ac_category` (`category_id` ASC),
  INDEX `idx_ac_article` (`article_id` ASC),
  CONSTRAINT `fk_ac_article` FOREIGN KEY (`article_id`) REFERENCES `article` (`id`),
  CONSTRAINT `fk_ac_category` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- --------------------------------------------------------
--
-- 表的结构 `log`
--
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `title` varchar(75),
  `metaTitle` varchar(100),
  `slug` varchar(100),
  `content` TEXT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- --------------------------------------------------------
--
-- 表的结构 `article_tag`
--
DROP TABLE IF EXISTS `article_tag`;
CREATE TABLE `article_tag` (
  `article_id` BIGINT NOT NULL,
  `tag_id` BIGINT NOT NULL,
  PRIMARY KEY (`article_id`, `tag_id`),
  INDEX `idx_at_tag` (`tag_id` ASC),
  INDEX `idx_at_article` (`article_id` ASC),
  CONSTRAINT `fk_at_article`
    FOREIGN KEY (`article_id`)
    REFERENCES `article` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_at_category`
    FOREIGN KEY (`tag_id`)
    REFERENCES `tag` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `enquiry`
--
DROP TABLE IF EXISTS `enquiry`;
CREATE TABLE `enquiry` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(75) COMMENT '姓名',
  `email` varchar(50) COMMENT '邮件地址',
  `phone` varchar(100) COMMENT '电话',
  `type` varchar(100) COMMENT '类型',
  `message` TEXT COMMENT '留言内容',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
