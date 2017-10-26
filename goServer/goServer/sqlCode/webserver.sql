/*
SQLyog Ultimate v12.09 (64 bit)
MySQL - 5.7.12-log : Database - webserver
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`webserver` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `webserver`;

/*Table structure for table `sys_menu` */

DROP TABLE IF EXISTS `sys_menu`;

CREATE TABLE `sys_menu` (
  `MenuID` int(10) NOT NULL COMMENT '菜单标识',
  `ParentMenuID` int(10) NOT NULL COMMENT '上级ID',
  `MenuName` varchar(36) NOT NULL COMMENT '菜单名称',
  `MenuUrl` varchar(256) DEFAULT NULL COMMENT '菜单地址',
  `SortOrder` int(10) NOT NULL COMMENT '排序号',
  `MenuIcon` varchar(36) DEFAULT NULL COMMENT '菜单图标路径（未用到）',
  `BigMenuIcon` varchar(36) DEFAULT NULL COMMENT '常用菜单图标（未用到）',
  `ShortCut` varchar(36) DEFAULT NULL COMMENT '快捷键（未用到）',
  `IsShow` tinyint(1) NOT NULL COMMENT '是否显示',
  PRIMARY KEY (`MenuID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `sys_menu` */

insert  into `sys_menu`(`MenuID`,`ParentMenuID`,`MenuName`,`MenuUrl`,`SortOrder`,`MenuIcon`,`BigMenuIcon`,`ShortCut`,`IsShow`) values (1100,0,'相册','/WebPage/Main/EditPhoto.html',2,'fa fa-file-image-o',NULL,NULL,1),(1200,0,'博客','',1,'fa fa-bolt',NULL,NULL,1),(1201,1200,'博客主页','/WebPage/Blog/BlogMain.html',1,NULL,NULL,NULL,1),(1202,1200,'浏览','/WebPage/Blog/BlogBrowse.html',2,NULL,NULL,NULL,1),(1203,0,'测试页面','/WebPage/Test/test.html',3,NULL,NULL,NULL,1);

/*Table structure for table `sys_role` */

DROP TABLE IF EXISTS `sys_role`;

CREATE TABLE `sys_role` (
  `RoleID` int(10) NOT NULL COMMENT '主键',
  `RoleName` varchar(36) NOT NULL COMMENT '角色名称',
  `MenuIDS` varchar(1024) NOT NULL COMMENT '菜单id（用,隔开）',
  `IsDefault` tinyint(1) NOT NULL COMMENT '是否默认角色',
  `IsSupper` tinyint(1) NOT NULL COMMENT '是否是超级管理员角色',
  `Notes` varchar(256) NOT NULL COMMENT '描述',
  PRIMARY KEY (`RoleID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `sys_role` */

insert  into `sys_role`(`RoleID`,`RoleName`,`MenuIDS`,`IsDefault`,`IsSupper`,`Notes`) values (1,'超级管理员','1100,1200,1201,1202,1203',1,1,'超级管理员');

/*Table structure for table `sys_user` */

DROP TABLE IF EXISTS `sys_user`;

CREATE TABLE `sys_user` (
  `UserID` char(36) NOT NULL COMMENT '主键',
  `UserName` varchar(128) NOT NULL COMMENT '登录ID',
  `FullName` varchar(128) NOT NULL COMMENT '用户真实姓名',
  `Password` varchar(512) NOT NULL COMMENT '登陆密码',
  `PwdExpiredTime` datetime DEFAULT NULL COMMENT '密码过期时间',
  `Sex` tinyint(1) DEFAULT NULL COMMENT '性别 1男0女',
  `Phone` varchar(36) DEFAULT NULL COMMENT '工作电话',
  `Email` varchar(36) DEFAULT NULL COMMENT '电子邮箱',
  `HeadImgage` varchar(256) DEFAULT NULL COMMENT '头像',
  `Status` int(10) DEFAULT NULL COMMENT '状态 1 启用 2禁用 3已删',
  `LoginCount` int(20) DEFAULT '0' COMMENT '登录次数',
  `LastLoginTime` datetime DEFAULT NULL COMMENT '最后登录时间',
  `LastLoginIP` varchar(36) DEFAULT NULL COMMENT '公司ID',
  `RoleIDs` varchar(512) DEFAULT NULL COMMENT '角色ID（可以多个）',
  `CreateTime` datetime DEFAULT NULL COMMENT '创建日期',
  PRIMARY KEY (`UserID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `sys_user` */

insert  into `sys_user`(`UserID`,`UserName`,`FullName`,`Password`,`PwdExpiredTime`,`Sex`,`Phone`,`Email`,`HeadImgage`,`Status`,`LoginCount`,`LastLoginTime`,`LastLoginIP`,`RoleIDs`,`CreateTime`) values ('70b22a06d4be8f5f6cadf25b21220702','xiaohe','xiaohe','061f480584ff1242d3b5c101f91e7e42','2017-10-25 21:25:18',1,'','546247083@qq.com','./upload/xiaohe_Album_1508923540241_u=996503075,3768564257&fm=200&gp=0.jpg',1,111,'2017-10-25 17:25:18','','1','2017-10-15 08:21:17'),('cd78f4d758cf147b2a24ace2e33a876b','xiaohea','xiaohea','2eadaadfa940205cacc2be67434e77aa','2017-10-20 16:26:36',1,'','295787943@qq.com','upload/xiaohe_Album_1508757589308_u=4045168581,660991028&fm=27&gp=0.jpg',1,15,'2017-10-20 14:26:36','','1','2017-10-17 18:15:48');

/*Table structure for table `u_blog` */

DROP TABLE IF EXISTS `u_blog`;

CREATE TABLE `u_blog` (
  `ID` char(36) NOT NULL COMMENT '主键',
  `UserId` char(36) NOT NULL COMMENT '用户id',
  `Title` varchar(40) NOT NULL DEFAULT '' COMMENT '标题',
  `Content` text NOT NULL COMMENT '内容',
  `Tag` varchar(128) DEFAULT NULL COMMENT '标签（用，号隔开）',
  `ATUsers` varchar(128) DEFAULT NULL COMMENT '@的用户',
  `BlogType` int(11) NOT NULL DEFAULT '0' COMMENT '博客类型',
  `Status` tinyint(3) NOT NULL DEFAULT '0' COMMENT '状态【0：草稿，1：正常，2：删除，3：彻底删除】',
  `CrDate` datetime NOT NULL COMMENT '创建时间',
  `ReDate` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `u_blog` */

insert  into `u_blog`(`ID`,`UserId`,`Title`,`Content`,`Tag`,`ATUsers`,`BlogType`,`Status`,`CrDate`,`ReDate`) values ('e0441084-a620-4328-9fea-a2a2ea050d6b','448cfeb6-37a2-40c4-84e9-72599fd3722f','你好','测试内容',NULL,NULL,2,1,'2017-09-09 23:57:20','2017-09-09 00:00:00');

/*Table structure for table `u_blog_type` */

DROP TABLE IF EXISTS `u_blog_type`;

CREATE TABLE `u_blog_type` (
  `ID` int(10) NOT NULL COMMENT '主键',
  `Name` varchar(18) DEFAULT NULL COMMENT '类型名',
  `Icon` varchar(18) DEFAULT NULL COMMENT '图标',
  `IsPublic` tinyint(1) DEFAULT NULL COMMENT '是否展示',
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `u_blog_type` */

insert  into `u_blog_type`(`ID`,`Name`,`Icon`,`IsPublic`) values (1,'笔记','fa fa-inbox',0),(2,'博客','fa fa-envelope-o',1);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
