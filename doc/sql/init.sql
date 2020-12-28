-- MySQL dump 10.13  Distrib 8.0.11, for macos10.13 (x86_64)
--
-- Host: localhost    Database: testmgmt
-- ------------------------------------------------------
-- Server version	8.0.11

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
 SET NAMES utf8mb4 ;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `api_case`
--

DROP TABLE IF EXISTS `api_case`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `api_case` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `case_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `module` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `runNum` int(11) NOT NULL,
  `beforeCase` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `afterCase` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `outVars` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `chkVars` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `param_def` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `raw` longtext COLLATE utf8mb4_unicode_ci,
  `project` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `smoketest` char(3) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `case_id_project` (`case_id`,`project`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `api_case`
--

LOCK TABLES `api_case` WRITE;
/*!40000 ALTER TABLE `api_case` DISABLE KEYS */;
/*!40000 ALTER TABLE `api_case` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `api_detail`
--

DROP TABLE IF EXISTS `api_detail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `api_detail` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `case_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `module` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `apiFunction` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `httpMethod` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `path` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `header` longtext COLLATE utf8mb4_unicode_ci,
  `pathVariable` longtext COLLATE utf8mb4_unicode_ci,
  `queryParameter` longtext COLLATE utf8mb4_unicode_ci,
  `body` longtext COLLATE utf8mb4_unicode_ci,
  `response` longtext COLLATE utf8mb4_unicode_ci,
  `project` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `case_id_project` (`case_id`,`project`)
) ENGINE=InnoDB AUTO_INCREMENT=1448 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `api_detail`
--

LOCK TABLES `api_detail` WRITE;
/*!40000 ALTER TABLE `api_detail` DISABLE KEYS */;
/*!40000 ALTER TABLE `api_detail` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `api_sum_up`
--

DROP TABLE IF EXISTS `api_sum_up`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `api_sum_up` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `all_count` int(11) DEFAULT NULL,
  `automatable_count` int(11) DEFAULT NULL,
  `unautomatable_count` int(11) DEFAULT NULL,
  `auto_test_count` int(11) DEFAULT NULL,
  `untest_count` int(11) DEFAULT NULL,
  `pass_count` int(11) DEFAULT NULL,
  `fail_count` int(11) DEFAULT NULL,
  `auto_per` double DEFAULT NULL,
  `pass_per` double DEFAULT NULL,
  `fail_per` double DEFAULT NULL,
  `project` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `api_sum_up`
--

LOCK TABLES `api_sum_up` WRITE;
/*!40000 ALTER TABLE `api_sum_up` DISABLE KEYS */;
/*!40000 ALTER TABLE `api_sum_up` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `api_test_data`
--

DROP TABLE IF EXISTS `api_test_data`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `api_test_data` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `data_desc` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `apiFunction` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `case_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `module` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `urlQuery` longtext COLLATE utf8mb4_unicode_ci,
  `body` longtext COLLATE utf8mb4_unicode_ci,
  `expected_result` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `actual_result` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `result` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `fail_reason` longtext COLLATE utf8mb4_unicode_ci,
  `response` longtext COLLATE utf8mb4_unicode_ci,
  `project` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `api_test_data`
--

LOCK TABLES `api_test_data` WRITE;
/*!40000 ALTER TABLE `api_test_data` DISABLE KEYS */;
/*!40000 ALTER TABLE `api_test_data` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `api_test_detail`
--

DROP TABLE IF EXISTS `api_test_detail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `api_test_detail` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `case_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `APIFunction` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `url` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `body` longtext COLLATE utf8mb4_unicode_ci,
  `response` longtext COLLATE utf8mb4_unicode_ci,
  `fail_reason` longtext COLLATE utf8mb4_unicode_ci,
  `test_result` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `project` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `api_test_detail`
--

LOCK TABLES `api_test_detail` WRITE;
/*!40000 ALTER TABLE `api_test_detail` DISABLE KEYS */;
/*!40000 ALTER TABLE `api_test_detail` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `api_test_result`
--

DROP TABLE IF EXISTS `api_test_result`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `api_test_result` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `case_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `requestVars` longtext COLLATE utf8mb4_unicode_ci,
  `result` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `outVars` longtext COLLATE utf8mb4_unicode_ci,
  `project` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `case_id_project` (`case_id`,`project`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `api_test_result`
--

LOCK TABLES `api_test_result` WRITE;
/*!40000 ALTER TABLE `api_test_result` DISABLE KEYS */;
/*!40000 ALTER TABLE `api_test_result` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `case_test_count`
--

DROP TABLE IF EXISTS `case_test_count`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `case_test_count` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `case_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `API_function` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `run_times` int(11) NOT NULL,
  `test_times` int(11) NOT NULL,
  `pass_times` int(11) NOT NULL,
  `fail_times` int(11) NOT NULL,
  `untest_times` int(11) NOT NULL,
  `test_result` char(8) COLLATE utf8mb4_unicode_ci NOT NULL,
  `fail_reason` longtext COLLATE utf8mb4_unicode_ci,
  `project` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `case_test_count`
--

LOCK TABLES `case_test_count` WRITE;
/*!40000 ALTER TABLE `case_test_count` DISABLE KEYS */;
/*!40000 ALTER TABLE `case_test_count` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `common_variable`
--

DROP TABLE IF EXISTS `common_variable`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `common_variable` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `value` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `project` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_project` (`name`,`project`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `common_variable`
--

LOCK TABLES `common_variable` WRITE;
/*!40000 ALTER TABLE `common_variable` DISABLE KEYS */;
INSERT INTO `common_variable` VALUES (7,'uniVar','name','TEST','2020-12-04 10:02:51','2020-12-28 14:17:52',NULL);
/*!40000 ALTER TABLE `common_variable` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `filemanager_setting`
--

DROP TABLE IF EXISTS `filemanager_setting`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `filemanager_setting` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `key` varchar(100) DEFAULT NULL,
  `value` text,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `filemanager_setting`
--

LOCK TABLES `filemanager_setting` WRITE;
/*!40000 ALTER TABLE `filemanager_setting` DISABLE KEYS */;
INSERT INTO `filemanager_setting` VALUES (1,'roots','{\"api\":{\"Path\":\"/tmp/testmgmt/api\",\"Title\":\"API文件\"},\"file\":{\"Path\":\"/tmp/testmgmt/file\",\"Title\":\"公用文件(勿删)\"},\"log\":{\"Path\":\"/tmp/testmgmt/log\",\"Title\":\"日志管理\"},\"test\":{\"Path\":\"/tmp/testmgmt/test\",\"Title\":\"用例文件\"}}','2020-12-07 02:35:59','2020-12-07 02:35:59'),(2,'allowMove','1','2020-12-07 02:36:00','2020-12-07 02:36:00'),(3,'conn','default','2020-12-07 02:36:00','2020-12-07 02:36:00'),(4,'allowUpload','1','2020-12-07 02:36:00','2020-12-07 02:36:00'),(5,'allowDelete','1','2020-12-07 02:36:00','2020-12-07 02:36:00'),(6,'allowRename','1','2020-12-07 02:36:00','2020-12-07 02:36:00'),(7,'allowDownload','1','2020-12-07 02:36:00','2020-12-07 02:36:00'),(8,'allowCreateDir','1','2020-12-07 02:36:00','2020-12-07 02:36:00');
/*!40000 ALTER TABLE `filemanager_setting` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goadmin_menu`
--

DROP TABLE IF EXISTS `goadmin_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `goadmin_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) unsigned NOT NULL DEFAULT '0',
  `type` tinyint(4) unsigned NOT NULL DEFAULT '0',
  `order` int(11) unsigned NOT NULL DEFAULT '0',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `uri` varchar(3000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `header` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `plugin_name` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `uuid` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=50 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_menu`
--

LOCK TABLES `goadmin_menu` WRITE;
/*!40000 ALTER TABLE `goadmin_menu` DISABLE KEYS */;
INSERT INTO `goadmin_menu` VALUES (1,0,1,2,'Admin','fa-tasks','',NULL,'',NULL,'2019-09-09 16:00:00','2019-09-09 16:00:00'),(2,1,1,2,'Users','fa-users','/info/manager',NULL,'',NULL,'2019-09-09 16:00:00','2019-09-09 16:00:00'),(3,1,1,3,'Roles','fa-user','/info/roles',NULL,'',NULL,'2019-09-09 16:00:00','2019-09-09 16:00:00'),(4,1,1,4,'Permission','fa-ban','/info/permission',NULL,'',NULL,'2019-09-09 16:00:00','2019-09-09 16:00:00'),(5,1,1,5,'Menu','fa-bars','/menu',NULL,'',NULL,'2019-09-09 16:00:00','2019-09-09 16:00:00'),(6,1,1,6,'Operation log','fa-history','/info/op',NULL,'',NULL,'2019-09-09 16:00:00','2019-09-09 16:00:00'),(7,0,1,1,'Dashboard','fa-bar-chart','/dashboard','','',NULL,'2019-09-09 16:00:00','2020-12-28 14:16:17'),(8,0,0,7,'环境','fa-cog','','','',NULL,'2020-11-23 02:47:56','2020-11-23 02:57:07'),(9,8,0,7,'测试环境','fa-bars','/info/host','','',NULL,'2020-11-23 02:57:40','2020-11-23 03:05:36'),(10,8,0,8,'Gitlab列表','fa-bars','/info/product_gitlab','','',NULL,'2020-11-23 02:58:18','2020-11-23 03:05:50'),(11,8,0,9,'公用变量','fa-bars','/info/common_variable','','',NULL,'2020-11-23 02:58:49','2020-11-23 03:05:59'),(12,0,0,10,'用例','fa-bank','','','',NULL,'2020-11-23 02:59:22','2020-11-23 02:59:22'),(13,0,0,14,'报表','fa-align-left','','','',NULL,'2020-11-23 02:59:32','2020-11-23 02:59:32'),(14,0,0,19,'结果','fa-cubes','','','',NULL,'2020-11-23 02:59:57','2020-11-23 02:59:57'),(15,0,0,21,'计划','fa-group','','','',NULL,'2020-11-23 03:00:31','2020-11-23 03:00:31'),(16,15,0,21,'测试计划','fa-bars','/info/test_progress_schedule','','',NULL,'2020-11-23 03:00:56','2020-11-23 03:06:57'),(17,14,0,19,'接口结果','fa-bars','/info/api_test_result','','',NULL,'2020-11-23 03:01:36','2020-11-23 03:06:52'),(18,14,0,20,'结果详情','fa-bars','/info/api_test_detail','','',NULL,'2020-11-23 03:02:09','2020-11-23 03:06:44'),(19,13,0,17,'标签统计','fa-bars','/info/issue_tag_count','','',NULL,'2020-11-23 03:02:33','2020-11-23 03:06:16'),(20,13,0,18,'版本统计','fa-bars','/info/issue_milestone_count','','',NULL,'2020-11-23 03:02:57','2020-11-23 03:06:27'),(21,13,0,15,'接口统计','fa-bars','/info/case_test_count','','',NULL,'2020-11-23 03:03:23','2020-12-07 11:16:19'),(22,12,0,11,'接口用例','fa-bars','/info/api_case','','',NULL,'2020-11-23 03:03:48','2020-11-23 03:06:06'),(23,12,0,12,'测试用例','fa-bars','/info/test_case','','',NULL,'2020-11-23 03:04:09','2020-11-23 03:06:11'),(24,13,0,14,'接口总览','fa-bars','/info/api_sum_up','','',NULL,'2020-11-23 03:05:10','2020-12-07 11:16:05'),(25,12,0,13,'接口详情','fa-bars','/info/api_detail','','',NULL,'2020-11-24 09:22:39','2020-11-24 09:23:02'),(34,28,0,22,'Swagger文件','fa-bars','/fm/api/list','','',NULL,'2020-12-07 03:06:28','2020-12-07 03:06:28'),(37,28,0,21,'用例文件','fa-bars','/fm/test/list','','',NULL,'2020-12-07 03:08:41','2020-12-07 03:08:41'),(38,28,0,20,'公共文件','fa-bars','/fm/file/list','','',NULL,'2020-12-07 03:09:05','2020-12-07 03:09:05'),(39,0,0,22,'文件','fa-files-o','','','',NULL,'2020-12-07 03:20:35','2020-12-10 03:17:02'),(40,39,0,22,'公共文件','fa-bars','/fm/file/list','','',NULL,'2020-12-07 03:23:33','2020-12-07 03:23:33'),(41,39,0,23,'用例文件','fa-bars','/fm/test/list','','',NULL,'2020-12-07 03:23:49','2020-12-07 03:23:49'),(42,39,0,24,'API文件','fa-bars','/fm/api/list','','',NULL,'2020-12-07 03:24:04','2020-12-07 03:24:04'),(43,12,0,10,'测试数据','fa-bars','/info/api_test_data','','',NULL,'2020-12-08 08:20:47','2020-12-08 08:20:47'),(44,39,0,25,'日志文件','fa-bars','/fm/log/list','','',NULL,'2020-12-10 03:08:30','2020-12-10 03:17:15'),(45,0,0,4,'日志文件','fa-bars','/fm/log/list','','filemanager',NULL,'2020-12-10 03:09:23','2020-12-10 03:16:07'),(46,0,0,2,'用例文件','fa-bars','/fm/test/list','','filemanager',NULL,'2020-12-10 03:10:06','2020-12-10 03:10:06'),(47,0,0,1,'API文件','fa-bars','/fm/api/list','','filemanager',NULL,'2020-12-10 03:10:35','2020-12-10 03:10:35'),(48,0,0,3,'公用文件','fa-bars','/fm/file/list','','filemanager',NULL,'2020-12-10 03:10:56','2020-12-10 03:10:56'),(49,13,0,16,'用例统计','fa-bars','/info/testcase_count','','',NULL,'2020-12-15 07:17:01','2020-12-15 07:17:01');
/*!40000 ALTER TABLE `goadmin_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goadmin_operation_log`
--

DROP TABLE IF EXISTS `goadmin_operation_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `goadmin_operation_log` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL,
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `input` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `admin_operation_log_user_id_index` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_operation_log`
--

LOCK TABLES `goadmin_operation_log` WRITE;
/*!40000 ALTER TABLE `goadmin_operation_log` DISABLE KEYS */;
INSERT INTO `goadmin_operation_log` VALUES (1,1,'/admin/info/permission','GET','127.0.0.1','','2020-12-28 14:32:45','2020-12-28 14:32:45'),(2,1,'/admin/info/op','GET','127.0.0.1','','2020-12-28 14:32:47','2020-12-28 14:32:47'),(3,1,'/admin/info/permission','GET','127.0.0.1','','2020-12-28 14:32:49','2020-12-28 14:32:49'),(4,1,'/admin/info/manager','GET','127.0.0.1','','2020-12-28 14:32:50','2020-12-28 14:32:50'),(5,1,'/admin/info/test_progress_schedule','GET','127.0.0.1','','2020-12-28 14:32:57','2020-12-28 14:32:57');
/*!40000 ALTER TABLE `goadmin_operation_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goadmin_permissions`
--

DROP TABLE IF EXISTS `goadmin_permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `goadmin_permissions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `http_method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `http_path` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_permissions_name_unique` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=115 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_permissions`
--

LOCK TABLES `goadmin_permissions` WRITE;
/*!40000 ALTER TABLE `goadmin_permissions` DISABLE KEYS */;
INSERT INTO `goadmin_permissions` VALUES (1,'All permission','*','','*','2019-09-09 16:00:00','2019-09-09 16:00:00'),(2,'Dashboard','dashboard','GET,PUT,POST,DELETE','/','2019-09-09 16:00:00','2019-09-09 16:00:00'),(3,'api_case 查询','api_case_query','GET','/info/api_case','2020-11-23 02:06:28','2020-11-23 02:06:28'),(4,'api_case 编辑页显示','api_case_show_edit','GET','/info/api_case/edit','2020-11-23 02:06:28','2020-11-23 02:06:28'),(5,'api_case 新建记录页显示','api_case_show_create','GET','/info/api_case/new','2020-11-23 02:06:28','2020-11-23 02:06:28'),(6,'api_case 编辑','api_case_edit','POST','/edit/api_case','2020-11-23 02:06:28','2020-11-23 02:06:28'),(7,'api_case 新建','api_case_create','POST','/new/api_case','2020-11-23 02:06:28','2020-11-23 02:06:28'),(8,'api_case 删除','api_case_delete','POST','/delete/api_case','2020-11-23 02:06:28','2020-11-23 02:06:28'),(9,'api_case 导出','api_case_export','POST','/export/api_case','2020-11-23 02:06:28','2020-11-23 02:06:28'),(10,'api_sum_up 查询','api_sum_up_query','GET','/info/api_sum_up','2020-11-23 02:06:28','2020-11-23 02:06:28'),(11,'api_sum_up 编辑页显示','api_sum_up_show_edit','GET','/info/api_sum_up/edit','2020-11-23 02:06:28','2020-11-23 02:06:28'),(12,'api_sum_up 新建记录页显示','api_sum_up_show_create','GET','/info/api_sum_up/new','2020-11-23 02:06:28','2020-11-23 02:06:28'),(13,'api_sum_up 编辑','api_sum_up_edit','POST','/edit/api_sum_up','2020-11-23 02:06:28','2020-11-23 02:06:28'),(14,'api_sum_up 新建','api_sum_up_create','POST','/new/api_sum_up','2020-11-23 02:06:28','2020-11-23 02:06:28'),(15,'api_sum_up 删除','api_sum_up_delete','POST','/delete/api_sum_up','2020-11-23 02:06:28','2020-11-23 02:06:28'),(16,'api_sum_up 导出','api_sum_up_export','POST','/export/api_sum_up','2020-11-23 02:06:28','2020-11-23 02:06:28'),(17,'api_test_detail 查询','api_test_detail_query','GET','/info/api_test_detail','2020-11-23 02:06:28','2020-11-23 02:06:28'),(18,'api_test_detail 编辑页显示','api_test_detail_show_edit','GET','/info/api_test_detail/edit','2020-11-23 02:06:28','2020-11-23 02:06:28'),(19,'api_test_detail 新建记录页显示','api_test_detail_show_create','GET','/info/api_test_detail/new','2020-11-23 02:06:28','2020-11-23 02:06:28'),(20,'api_test_detail 编辑','api_test_detail_edit','POST','/edit/api_test_detail','2020-11-23 02:06:28','2020-11-23 02:06:28'),(21,'api_test_detail 新建','api_test_detail_create','POST','/new/api_test_detail','2020-11-23 02:06:28','2020-11-23 02:06:28'),(22,'api_test_detail 删除','api_test_detail_delete','POST','/delete/api_test_detail','2020-11-23 02:06:28','2020-11-23 02:06:28'),(23,'api_test_detail 导出','api_test_detail_export','POST','/export/api_test_detail','2020-11-23 02:06:28','2020-11-23 02:06:28'),(24,'api_test_result 查询','api_test_result_query','GET','/info/api_test_result','2020-11-23 02:06:28','2020-11-23 02:06:28'),(25,'api_test_result 编辑页显示','api_test_result_show_edit','GET','/info/api_test_result/edit','2020-11-23 02:06:28','2020-11-23 02:06:28'),(26,'api_test_result 新建记录页显示','api_test_result_show_create','GET','/info/api_test_result/new','2020-11-23 02:06:28','2020-11-23 02:06:28'),(27,'api_test_result 编辑','api_test_result_edit','POST','/edit/api_test_result','2020-11-23 02:06:28','2020-11-23 02:06:28'),(28,'api_test_result 新建','api_test_result_create','POST','/new/api_test_result','2020-11-23 02:06:28','2020-11-23 02:06:28'),(29,'api_test_result 删除','api_test_result_delete','POST','/delete/api_test_result','2020-11-23 02:06:28','2020-11-23 02:06:28'),(30,'api_test_result 导出','api_test_result_export','POST','/export/api_test_result','2020-11-23 02:06:28','2020-11-23 02:06:28'),(31,'case_test_count 查询','case_test_count_query','GET','/info/case_test_count','2020-11-23 02:06:28','2020-11-23 02:06:28'),(32,'case_test_count 编辑页显示','case_test_count_show_edit','GET','/info/case_test_count/edit','2020-11-23 02:06:28','2020-11-23 02:06:28'),(33,'case_test_count 新建记录页显示','case_test_count_show_create','GET','/info/case_test_count/new','2020-11-23 02:06:28','2020-11-23 02:06:28'),(34,'case_test_count 编辑','case_test_count_edit','POST','/edit/case_test_count','2020-11-23 02:06:28','2020-11-23 02:06:28'),(35,'case_test_count 新建','case_test_count_create','POST','/new/case_test_count','2020-11-23 02:06:28','2020-11-23 02:06:28'),(36,'case_test_count 删除','case_test_count_delete','POST','/delete/case_test_count','2020-11-23 02:06:28','2020-11-23 02:06:28'),(37,'case_test_count 导出','case_test_count_export','POST','/export/case_test_count','2020-11-23 02:06:28','2020-11-23 02:06:28'),(38,'common_variable 查询','common_variable_query','GET','/info/common_variable','2020-11-23 02:06:28','2020-11-23 02:06:28'),(39,'common_variable 编辑页显示','common_variable_show_edit','GET','/info/common_variable/edit','2020-11-23 02:06:28','2020-11-23 02:06:28'),(40,'common_variable 新建记录页显示','common_variable_show_create','GET','/info/common_variable/new','2020-11-23 02:06:28','2020-11-23 02:06:28'),(41,'common_variable 编辑','common_variable_edit','POST','/edit/common_variable','2020-11-23 02:06:28','2020-11-23 02:06:28'),(42,'common_variable 新建','common_variable_create','POST','/new/common_variable','2020-11-23 02:06:28','2020-11-23 02:06:28'),(43,'common_variable 删除','common_variable_delete','POST','/delete/common_variable','2020-11-23 02:06:28','2020-11-23 02:06:28'),(44,'common_variable 导出','common_variable_export','POST','/export/common_variable','2020-11-23 02:06:28','2020-11-23 02:06:28'),(45,'host 查询','host_query','GET','/info/host','2020-11-23 02:06:28','2020-11-23 02:06:28'),(46,'host 编辑页显示','host_show_edit','GET','/info/host/edit','2020-11-23 02:06:28','2020-11-23 02:06:28'),(47,'host 新建记录页显示','host_show_create','GET','/info/host/new','2020-11-23 02:06:28','2020-11-23 02:06:28'),(48,'host 编辑','host_edit','POST','/edit/host','2020-11-23 02:06:28','2020-11-23 02:06:28'),(49,'host 新建','host_create','POST','/new/host','2020-11-23 02:06:28','2020-11-23 02:06:28'),(50,'host 删除','host_delete','POST','/delete/host','2020-11-23 02:06:28','2020-11-23 02:06:28'),(51,'host 导出','host_export','POST','/export/host','2020-11-23 02:06:28','2020-11-23 02:06:28'),(52,'issue 查询','issue_query','GET','/info/issue','2020-11-23 02:06:28','2020-11-23 02:06:28'),(53,'issue 编辑页显示','issue_show_edit','GET','/info/issue/edit','2020-11-23 02:06:28','2020-11-23 02:06:28'),(54,'issue 新建记录页显示','issue_show_create','GET','/info/issue/new','2020-11-23 02:06:28','2020-11-23 02:06:28'),(55,'issue 编辑','issue_edit','POST','/edit/issue','2020-11-23 02:06:28','2020-11-23 02:06:28'),(56,'issue 新建','issue_create','POST','/new/issue','2020-11-23 02:06:28','2020-11-23 02:06:28'),(57,'issue 删除','issue_delete','POST','/delete/issue','2020-11-23 02:06:28','2020-11-23 02:06:28'),(58,'issue 导出','issue_export','POST','/export/issue','2020-11-23 02:06:28','2020-11-23 02:06:28'),(59,'issue_milestone_count 查询','issue_milestone_count_query','GET','/info/issue_milestone_count','2020-11-23 02:06:28','2020-11-23 02:06:28'),(60,'issue_milestone_count 编辑页显示','issue_milestone_count_show_edit','GET','/info/issue_milestone_count/edit','2020-11-23 02:06:28','2020-11-23 02:06:28'),(61,'issue_milestone_count 新建记录页显示','issue_milestone_count_show_create','GET','/info/issue_milestone_count/new','2020-11-23 02:06:28','2020-11-23 02:06:28'),(62,'issue_milestone_count 编辑','issue_milestone_count_edit','POST','/edit/issue_milestone_count','2020-11-23 02:06:28','2020-11-23 02:06:28'),(63,'issue_milestone_count 新建','issue_milestone_count_create','POST','/new/issue_milestone_count','2020-11-23 02:06:28','2020-11-23 02:06:28'),(64,'issue_milestone_count 删除','issue_milestone_count_delete','POST','/delete/issue_milestone_count','2020-11-23 02:06:28','2020-11-23 02:06:28'),(65,'issue_milestone_count 导出','issue_milestone_count_export','POST','/export/issue_milestone_count','2020-11-23 02:06:28','2020-11-23 02:06:28'),(66,'issue_tag_count 查询','issue_tag_count_query','GET','/info/issue_tag_count','2020-11-23 02:06:28','2020-11-23 02:06:28'),(67,'issue_tag_count 编辑页显示','issue_tag_count_show_edit','GET','/info/issue_tag_count/edit','2020-11-23 02:06:28','2020-11-23 02:06:28'),(68,'issue_tag_count 新建记录页显示','issue_tag_count_show_create','GET','/info/issue_tag_count/new','2020-11-23 02:06:28','2020-11-23 02:06:28'),(69,'issue_tag_count 编辑','issue_tag_count_edit','POST','/edit/issue_tag_count','2020-11-23 02:06:28','2020-11-23 02:06:28'),(70,'issue_tag_count 新建','issue_tag_count_create','POST','/new/issue_tag_count','2020-11-23 02:06:28','2020-11-23 02:06:28'),(71,'issue_tag_count 删除','issue_tag_count_delete','POST','/delete/issue_tag_count','2020-11-23 02:06:28','2020-11-23 02:06:28'),(72,'issue_tag_count 导出','issue_tag_count_export','POST','/export/issue_tag_count','2020-11-23 02:06:28','2020-11-23 02:06:28'),(73,'product_gitlab 查询','product_gitlab_query','GET','/info/product_gitlab','2020-11-23 02:06:28','2020-11-23 02:06:28'),(74,'product_gitlab 编辑页显示','product_gitlab_show_edit','GET','/info/product_gitlab/edit','2020-11-23 02:06:28','2020-11-23 02:06:28'),(75,'product_gitlab 新建记录页显示','product_gitlab_show_create','GET','/info/product_gitlab/new','2020-11-23 02:06:28','2020-11-23 02:06:28'),(76,'product_gitlab 编辑','product_gitlab_edit','POST','/edit/product_gitlab','2020-11-23 02:06:28','2020-11-23 02:06:28'),(77,'product_gitlab 新建','product_gitlab_create','POST','/new/product_gitlab','2020-11-23 02:06:28','2020-11-23 02:06:28'),(78,'product_gitlab 删除','product_gitlab_delete','POST','/delete/product_gitlab','2020-11-23 02:06:28','2020-11-23 02:06:28'),(79,'product_gitlab 导出','product_gitlab_export','POST','/export/product_gitlab','2020-11-23 02:06:28','2020-11-23 02:06:28'),(80,'test_case 查询','test_case_query','GET','/info/test_case','2020-11-23 02:06:28','2020-11-23 02:06:28'),(81,'test_case 编辑页显示','test_case_show_edit','GET','/info/test_case/edit','2020-11-23 02:06:28','2020-11-23 02:06:28'),(82,'test_case 新建记录页显示','test_case_show_create','GET','/info/test_case/new','2020-11-23 02:06:28','2020-11-23 02:06:28'),(83,'test_case 编辑','test_case_edit','POST','/edit/test_case','2020-11-23 02:06:28','2020-11-23 02:06:28'),(84,'test_case 新建','test_case_create','POST','/new/test_case','2020-11-23 02:06:28','2020-11-23 02:06:28'),(85,'test_case 删除','test_case_delete','POST','/delete/test_case','2020-11-23 02:06:28','2020-11-23 02:06:28'),(86,'test_case 导出','test_case_export','POST','/export/test_case','2020-11-23 02:06:28','2020-11-23 02:06:28'),(87,'test_progress_schedule 查询','test_progress_schedule_query','GET','/info/test_progress_schedule','2020-11-23 02:06:28','2020-11-23 02:06:28'),(88,'test_progress_schedule 编辑页显示','test_progress_schedule_show_edit','GET','/info/test_progress_schedule/edit','2020-11-23 02:06:28','2020-11-23 02:06:28'),(89,'test_progress_schedule 新建记录页显示','test_progress_schedule_show_create','GET','/info/test_progress_schedule/new','2020-11-23 02:06:28','2020-11-23 02:06:28'),(90,'test_progress_schedule 编辑','test_progress_schedule_edit','POST','/edit/test_progress_schedule','2020-11-23 02:06:28','2020-11-23 02:06:28'),(91,'test_progress_schedule 新建','test_progress_schedule_create','POST','/new/test_progress_schedule','2020-11-23 02:06:28','2020-11-23 02:06:28'),(92,'test_progress_schedule 删除','test_progress_schedule_delete','POST','/delete/test_progress_schedule','2020-11-23 02:06:28','2020-11-23 02:06:28'),(93,'test_progress_schedule 导出','test_progress_schedule_export','POST','/export/test_progress_schedule','2020-11-23 02:06:28','2020-11-23 02:06:28'),(94,'api_detail 查询','api_detail_query','GET','/info/api_detail','2020-11-24 09:18:47','2020-11-24 09:18:47'),(95,'api_detail 编辑页显示','api_detail_show_edit','GET','/info/api_detail/edit','2020-11-24 09:18:47','2020-11-24 09:18:47'),(96,'api_detail 新建记录页显示','api_detail_show_create','GET','/info/api_detail/new','2020-11-24 09:18:47','2020-11-24 09:18:47'),(97,'api_detail 编辑','api_detail_edit','POST','/edit/api_detail','2020-11-24 09:18:47','2020-11-24 09:18:47'),(98,'api_detail 新建','api_detail_create','POST','/new/api_detail','2020-11-24 09:18:47','2020-11-24 09:18:47'),(99,'api_detail 删除','api_detail_delete','POST','/delete/api_detail','2020-11-24 09:18:47','2020-11-24 09:18:47'),(100,'api_detail 导出','api_detail_export','POST','/export/api_detail','2020-11-24 09:18:47','2020-11-24 09:18:47'),(101,'api_test_data 查询','api_test_data_query','GET','/info/api_test_data','2020-12-08 08:18:54','2020-12-08 08:18:54'),(102,'api_test_data 编辑页显示','api_test_data_show_edit','GET','/info/api_test_data/edit','2020-12-08 08:18:54','2020-12-08 08:18:54'),(103,'api_test_data 新建记录页显示','api_test_data_show_create','GET','/info/api_test_data/new','2020-12-08 08:18:54','2020-12-08 08:18:54'),(104,'api_test_data 编辑','api_test_data_edit','POST','/edit/api_test_data','2020-12-08 08:18:54','2020-12-08 08:18:54'),(105,'api_test_data 新建','api_test_data_create','POST','/new/api_test_data','2020-12-08 08:18:54','2020-12-08 08:18:54'),(106,'api_test_data 删除','api_test_data_delete','POST','/delete/api_test_data','2020-12-08 08:18:54','2020-12-08 08:18:54'),(107,'api_test_data 导出','api_test_data_export','POST','/export/api_test_data','2020-12-08 08:18:54','2020-12-08 08:18:54'),(108,'testcase_count 查询','testcase_count_query','GET','/info/testcase_count','2020-12-15 03:17:56','2020-12-15 03:17:56'),(109,'testcase_count 编辑页显示','testcase_count_show_edit','GET','/info/testcase_count/edit','2020-12-15 03:17:56','2020-12-15 03:17:56'),(110,'testcase_count 新建记录页显示','testcase_count_show_create','GET','/info/testcase_count/new','2020-12-15 03:17:56','2020-12-15 03:17:56'),(111,'testcase_count 编辑','testcase_count_edit','POST','/edit/testcase_count','2020-12-15 03:17:56','2020-12-15 03:17:56'),(112,'testcase_count 新建','testcase_count_create','POST','/new/testcase_count','2020-12-15 03:17:56','2020-12-15 03:17:56'),(113,'testcase_count 删除','testcase_count_delete','POST','/delete/testcase_count','2020-12-15 03:17:56','2020-12-15 03:17:56'),(114,'testcase_count 导出','testcase_count_export','POST','/export/testcase_count','2020-12-15 03:17:56','2020-12-15 03:17:56');
/*!40000 ALTER TABLE `goadmin_permissions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goadmin_role_menu`
--

DROP TABLE IF EXISTS `goadmin_role_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `goadmin_role_menu` (
  `role_id` int(11) unsigned NOT NULL,
  `menu_id` int(11) unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  KEY `admin_role_menu_role_id_menu_id_index` (`role_id`,`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_role_menu`
--

LOCK TABLES `goadmin_role_menu` WRITE;
/*!40000 ALTER TABLE `goadmin_role_menu` DISABLE KEYS */;
INSERT INTO `goadmin_role_menu` VALUES (1,1,'2019-09-09 16:00:00','2019-09-09 16:00:00'),(1,8,'2020-11-23 02:57:07','2020-11-23 02:57:07'),(2,8,'2020-11-23 02:57:07','2020-11-23 02:57:07'),(1,9,'2020-11-23 03:05:36','2020-11-23 03:05:36'),(2,9,'2020-11-23 03:05:36','2020-11-23 03:05:36'),(1,10,'2020-11-23 03:05:50','2020-11-23 03:05:50'),(2,10,'2020-11-23 03:05:50','2020-11-23 03:05:50'),(1,22,'2020-11-23 03:06:06','2020-11-23 03:06:06'),(2,22,'2020-11-23 03:06:06','2020-11-23 03:06:06'),(1,23,'2020-11-23 03:06:11','2020-11-23 03:06:11'),(2,23,'2020-11-23 03:06:11','2020-11-23 03:06:11'),(1,19,'2020-11-23 03:06:16','2020-11-23 03:06:16'),(2,19,'2020-11-23 03:06:16','2020-11-23 03:06:16'),(1,20,'2020-11-23 03:06:27','2020-11-23 03:06:27'),(2,20,'2020-11-23 03:06:27','2020-11-23 03:06:27'),(1,18,'2020-11-23 03:06:44','2020-11-23 03:06:44'),(2,18,'2020-11-23 03:06:44','2020-11-23 03:06:44'),(1,17,'2020-11-23 03:06:52','2020-11-23 03:06:52'),(2,17,'2020-11-23 03:06:52','2020-11-23 03:06:52'),(1,16,'2020-11-23 03:06:57','2020-11-23 03:06:57'),(2,16,'2020-11-23 03:06:57','2020-11-23 03:06:57'),(1,25,'2020-11-24 09:23:02','2020-11-24 09:23:02'),(2,25,'2020-11-24 09:23:02','2020-11-24 09:23:02'),(1,34,'2020-12-07 03:06:28','2020-12-07 03:06:28'),(2,34,'2020-12-07 03:06:28','2020-12-07 03:06:28'),(1,37,'2020-12-07 03:08:41','2020-12-07 03:08:41'),(2,37,'2020-12-07 03:08:41','2020-12-07 03:08:41'),(1,38,'2020-12-07 03:09:05','2020-12-07 03:09:05'),(2,38,'2020-12-07 03:09:05','2020-12-07 03:09:05'),(1,40,'2020-12-07 03:23:33','2020-12-07 03:23:33'),(2,40,'2020-12-07 03:23:33','2020-12-07 03:23:33'),(1,41,'2020-12-07 03:23:49','2020-12-07 03:23:49'),(2,41,'2020-12-07 03:23:49','2020-12-07 03:23:49'),(1,42,'2020-12-07 03:24:04','2020-12-07 03:24:04'),(2,42,'2020-12-07 03:24:04','2020-12-07 03:24:04'),(1,24,'2020-12-07 11:16:05','2020-12-07 11:16:05'),(2,24,'2020-12-07 11:16:05','2020-12-07 11:16:05'),(1,21,'2020-12-07 11:16:19','2020-12-07 11:16:19'),(2,21,'2020-12-07 11:16:19','2020-12-07 11:16:19'),(1,43,'2020-12-08 08:20:47','2020-12-08 08:20:47'),(2,43,'2020-12-08 08:20:47','2020-12-08 08:20:47'),(1,46,'2020-12-10 03:10:06','2020-12-10 03:10:06'),(2,46,'2020-12-10 03:10:06','2020-12-10 03:10:06'),(1,47,'2020-12-10 03:10:35','2020-12-10 03:10:35'),(2,47,'2020-12-10 03:10:35','2020-12-10 03:10:35'),(1,48,'2020-12-10 03:10:56','2020-12-10 03:10:56'),(2,48,'2020-12-10 03:10:56','2020-12-10 03:10:56'),(1,45,'2020-12-10 03:16:07','2020-12-10 03:16:07'),(2,45,'2020-12-10 03:16:07','2020-12-10 03:16:07'),(1,39,'2020-12-10 03:17:02','2020-12-10 03:17:02'),(2,39,'2020-12-10 03:17:02','2020-12-10 03:17:02'),(1,44,'2020-12-10 03:17:15','2020-12-10 03:17:15'),(2,44,'2020-12-10 03:17:15','2020-12-10 03:17:15'),(1,49,'2020-12-15 07:17:01','2020-12-15 07:17:01'),(2,49,'2020-12-15 07:17:01','2020-12-15 07:17:01'),(1,7,'2020-12-28 14:16:17','2020-12-28 14:16:17'),(2,7,'2020-12-28 14:16:17','2020-12-28 14:16:17');
/*!40000 ALTER TABLE `goadmin_role_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goadmin_role_permissions`
--

DROP TABLE IF EXISTS `goadmin_role_permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `goadmin_role_permissions` (
  `role_id` int(11) unsigned NOT NULL,
  `permission_id` int(11) unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY `admin_role_permissions` (`role_id`,`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_role_permissions`
--

LOCK TABLES `goadmin_role_permissions` WRITE;
/*!40000 ALTER TABLE `goadmin_role_permissions` DISABLE KEYS */;
INSERT INTO `goadmin_role_permissions` VALUES (1,1,'2019-09-09 16:00:00','2019-09-09 16:00:00'),(1,2,'2019-09-09 16:00:00','2019-09-09 16:00:00'),(2,2,'2019-09-09 16:00:00','2019-09-09 16:00:00');
/*!40000 ALTER TABLE `goadmin_role_permissions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goadmin_role_users`
--

DROP TABLE IF EXISTS `goadmin_role_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `goadmin_role_users` (
  `role_id` int(11) unsigned NOT NULL,
  `user_id` int(11) unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY `admin_user_roles` (`role_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_role_users`
--

LOCK TABLES `goadmin_role_users` WRITE;
/*!40000 ALTER TABLE `goadmin_role_users` DISABLE KEYS */;
INSERT INTO `goadmin_role_users` VALUES (1,1,'2019-09-09 16:00:00','2019-09-09 16:00:00'),(2,2,'2019-09-09 16:00:00','2019-09-09 16:00:00');
/*!40000 ALTER TABLE `goadmin_role_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goadmin_roles`
--

DROP TABLE IF EXISTS `goadmin_roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `goadmin_roles` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_roles_name_unique` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_roles`
--

LOCK TABLES `goadmin_roles` WRITE;
/*!40000 ALTER TABLE `goadmin_roles` DISABLE KEYS */;
INSERT INTO `goadmin_roles` VALUES (1,'Administrator','administrator','2019-09-09 16:00:00','2019-09-09 16:00:00'),(2,'Operator','operator','2019-09-09 16:00:00','2019-09-09 16:00:00');
/*!40000 ALTER TABLE `goadmin_roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goadmin_session`
--

DROP TABLE IF EXISTS `goadmin_session`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `goadmin_session` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `sid` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `values` varchar(3000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=768 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_session`
--

LOCK TABLES `goadmin_session` WRITE;
/*!40000 ALTER TABLE `goadmin_session` DISABLE KEYS */;
INSERT INTO `goadmin_session` VALUES (710,'9442978a-6115-4c78-a6c0-9c410763932e','{\"user_id\":1}','2020-12-28 13:52:42','2020-12-28 13:52:42'),(711,'41ebcc57-0a30-4c21-8a50-6177cb65e24c','__csrf_token__','2020-12-28 14:16:01','2020-12-28 14:16:01'),(713,'e3e49f8f-70a9-4635-8163-cd0d2320bb8b','__csrf_token__','2020-12-28 14:16:17','2020-12-28 14:16:17'),(714,'db28213a-49ce-4f16-87dd-2078ef40aa5a','__csrf_token__','2020-12-28 14:16:55','2020-12-28 14:16:55'),(715,'822b4711-e504-4e45-a78c-d87549e7dabf','__csrf_token__','2020-12-28 14:17:03','2020-12-28 14:17:03'),(716,'aed1e927-772b-44a0-ba13-906fb272bcb7','__csrf_token__','2020-12-28 14:17:18','2020-12-28 14:17:18'),(717,'346de8c6-d52e-4036-b4b9-3e805ac40260','__csrf_token__','2020-12-28 14:17:40','2020-12-28 14:17:40'),(720,'1415b4c4-e0bf-48db-aab4-af5db815c366','__csrf_token__','2020-12-28 14:19:08','2020-12-28 14:19:08'),(721,'21a82377-cfe8-427a-b3db-a3f2148464bc','__csrf_token__','2020-12-28 14:19:51','2020-12-28 14:19:51'),(722,'b824555c-0a64-4707-8074-31f3376cedb5','__csrf_token__','2020-12-28 14:20:14','2020-12-28 14:20:14'),(723,'8270c5fa-526d-4402-a74b-862bf04fefea','__csrf_token__','2020-12-28 14:20:27','2020-12-28 14:20:27'),(724,'040334a2-76ba-4fd7-8cf8-fff544e42805','__csrf_token__','2020-12-28 14:21:03','2020-12-28 14:21:03'),(725,'3f31981a-9905-4a13-84f5-f22279b2fa09','__csrf_token__','2020-12-28 14:21:11','2020-12-28 14:21:11'),(726,'3a3fd0b2-a7ec-4f7b-bdb2-0fe764c144e0','__csrf_token__','2020-12-28 14:21:23','2020-12-28 14:21:23'),(727,'d4ffc35a-9b9e-4cd5-b3ef-121c0c458edf','__csrf_token__','2020-12-28 14:22:40','2020-12-28 14:22:40'),(728,'1d6bf9db-041d-4c57-a65b-ee94772cc57b','__csrf_token__','2020-12-28 14:22:49','2020-12-28 14:22:49'),(729,'76b8d57f-5f66-4215-914d-fc2d2385708c','__csrf_token__','2020-12-28 14:23:00','2020-12-28 14:23:00'),(730,'04c9fd5e-91b4-4066-885b-d978cf7e25bf','__csrf_token__','2020-12-28 14:23:08','2020-12-28 14:23:08'),(731,'4a0b779e-64ad-4052-b3ca-31dd04c5eaca','__csrf_token__','2020-12-28 14:23:18','2020-12-28 14:23:18'),(732,'cad68445-ea86-441c-9da9-d7cf8baee5e8','__csrf_token__','2020-12-28 14:23:25','2020-12-28 14:23:25'),(733,'bf54985f-6a6e-4325-af31-ceae19501704','__csrf_token__','2020-12-28 14:23:38','2020-12-28 14:23:38'),(734,'a2f24806-4010-4d2c-afca-61e1922184ef','__csrf_token__','2020-12-28 14:23:50','2020-12-28 14:23:50'),(735,'18f43003-ba74-4e1a-8054-25d5b2209a79','__csrf_token__','2020-12-28 14:24:06','2020-12-28 14:24:06'),(736,'f2965983-352b-4b94-962b-b29a7fe748c1','__csrf_token__','2020-12-28 14:24:23','2020-12-28 14:24:23'),(737,'ad4f0508-fd38-4911-ad1d-4329e11bd9d7','__csrf_token__','2020-12-28 14:24:29','2020-12-28 14:24:29'),(738,'5abc7896-47a0-4216-af13-e84e4172677f','__csrf_token__','2020-12-28 14:24:38','2020-12-28 14:24:38'),(739,'bad4f1ad-055b-49d9-8bfb-61e922c7cd94','__csrf_token__','2020-12-28 14:24:48','2020-12-28 14:24:48'),(740,'41b7272a-b61a-471c-8db5-ead2622a1e13','__csrf_token__','2020-12-28 14:24:55','2020-12-28 14:24:55'),(741,'1f2db305-5282-4846-ba7d-7ca264ad68f0','__csrf_token__','2020-12-28 14:25:01','2020-12-28 14:25:01'),(742,'f81cd310-b782-4714-9519-bacc1c68eeee','__csrf_token__','2020-12-28 14:25:17','2020-12-28 14:25:17'),(743,'3d135e48-5153-4a74-b02f-6d0cc69586d0','__csrf_token__','2020-12-28 14:25:22','2020-12-28 14:25:22'),(744,'21776617-6317-4406-a2ea-269b120b7724','__csrf_token__','2020-12-28 14:25:27','2020-12-28 14:25:27'),(745,'9a1534be-9327-4497-88b4-1fae2f0d072d','__csrf_token__','2020-12-28 14:25:33','2020-12-28 14:25:33'),(746,'3b4315e2-4aae-403e-888a-df99458be263','__csrf_token__','2020-12-28 14:25:38','2020-12-28 14:25:38'),(747,'a9f3df39-5cae-4350-b9a5-2de1edf3adee','__csrf_token__','2020-12-28 14:25:44','2020-12-28 14:25:44'),(748,'ef0cdfc1-72ce-4d25-a105-e80d4b04ce4e','__csrf_token__','2020-12-28 14:25:52','2020-12-28 14:25:52'),(749,'14ddb40b-2c8b-4a95-b7b4-c433d2be1aef','__csrf_token__','2020-12-28 14:25:58','2020-12-28 14:25:58'),(750,'0be5de8d-41b6-45aa-a80b-c035edd2ffe2','__csrf_token__','2020-12-28 14:26:03','2020-12-28 14:26:03'),(751,'791c3764-9114-461b-9c7c-f900904acc4a','__csrf_token__','2020-12-28 14:26:10','2020-12-28 14:26:10'),(752,'77b2d542-ee06-44d3-9036-fbce9ce406cd','__csrf_token__','2020-12-28 14:26:16','2020-12-28 14:26:16'),(753,'e5be251c-02c8-4f55-89f3-7d144b491759','__csrf_token__','2020-12-28 14:26:27','2020-12-28 14:26:27'),(754,'207ec00e-fc3e-40ec-a05a-aa332943ebf1','__csrf_token__','2020-12-28 14:26:33','2020-12-28 14:26:33'),(755,'5111ef56-981c-468e-becb-b431038dfded','__csrf_token__','2020-12-28 14:26:47','2020-12-28 14:26:47'),(756,'eb94e566-9de0-4bb2-9f58-eabc73be64de','__csrf_token__','2020-12-28 14:27:01','2020-12-28 14:27:01'),(757,'aa9530de-d029-4ad2-ab22-143eb711360a','__csrf_token__','2020-12-28 14:27:07','2020-12-28 14:27:07'),(758,'d4d912eb-e45c-41a6-a465-5c25fd1b55c5','__csrf_token__','2020-12-28 14:27:13','2020-12-28 14:27:13'),(759,'f2a8379c-8844-475f-9282-6a7c09c05986','__csrf_token__','2020-12-28 14:27:19','2020-12-28 14:27:19'),(760,'0930b9f3-7bfb-45ec-94ce-88ee52635b0f','__csrf_token__','2020-12-28 14:27:24','2020-12-28 14:27:24'),(761,'7df48b42-b00f-469d-b686-cbfe4f30bb88','__csrf_token__','2020-12-28 14:27:29','2020-12-28 14:27:29'),(762,'c49f290c-71f8-4332-b375-e6d6a27ed534','__csrf_token__','2020-12-28 14:27:34','2020-12-28 14:27:34'),(763,'345c5870-3157-4d56-9338-a2c80b0e3f88','__csrf_token__','2020-12-28 14:27:39','2020-12-28 14:27:39'),(764,'821c37cc-99b4-4b94-a674-4137d7998d1c','__csrf_token__','2020-12-28 14:27:43','2020-12-28 14:27:43'),(765,'17c3e212-59d6-48e7-b9e5-c62a2c12d8ad','__csrf_token__','2020-12-28 14:27:49','2020-12-28 14:27:49'),(766,'8ced190d-700e-4ad7-870d-aad7b1eaad34','__csrf_token__','2020-12-28 14:27:53','2020-12-28 14:27:53'),(767,'1962e6a4-ace1-49df-b325-da061e192a05','__csrf_token__','2020-12-28 14:32:24','2020-12-28 14:32:24');
/*!40000 ALTER TABLE `goadmin_session` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goadmin_site`
--

DROP TABLE IF EXISTS `goadmin_site`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `goadmin_site` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `value` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `description` varchar(3000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `state` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=68 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_site`
--

LOCK TABLES `goadmin_site` WRITE;
/*!40000 ALTER TABLE `goadmin_site` DISABLE KEYS */;
INSERT INTO `goadmin_site` VALUES (1,'theme','sword',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(2,'hide_app_info_entrance','true',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(3,'info_log_path','/tmp/testmgmt/log/info.log',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(4,'info_log_off','false',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(5,'mini_logo','        TM\r\n    ',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(6,'sql_log','false',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(7,'login_logo','       TestMgmt\r\n    ',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(8,'hide_plugin_entrance','true',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(9,'custom_500_html','',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(10,'login_title','TestMgmt',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(11,'domain','',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(12,'store','{\"path\":\"./uploads\",\"prefix\":\"uploads\"}',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(13,'logger_rotate_max_age','30',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(14,'logger_encoder_time_key','ts',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(15,'asset_url','',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(16,'custom_404_html','',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(17,'animation_duration','0.00',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(18,'allow_del_operation_log','true',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(19,'logger_rotate_compress','false',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(20,'logger_encoder_encoding','console',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(21,'custom_head_html','',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(22,'animation_type','',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(23,'hide_tool_entrance','true',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(24,'title','TestMgmt',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(25,'logger_level','-1',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(26,'custom_403_html','',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(27,'operation_log_off','false',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(28,'logger_rotate_max_size','10',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(29,'auth_user_table','goadmin_users',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(30,'login_url','/login',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(31,'error_log_path','/tmp/testmgmt/log/error.log',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(32,'logger_rotate_max_backups','5',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(33,'go_mod_file_path','/tmp/testmgmt/file/go.mod',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(34,'animation_delay','0.00',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(35,'logger_encoder_time','iso8601',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(36,'custom_foot_html','',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(37,'asset_root_path','./public/',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(38,'logger_encoder_message_key','msg',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(39,'databases','{\"default\":{\"host\":\"127.0.0.1\",\"port\":\"3306\",\"user\":\"root\",\"pwd\":\"zaq1@WSX\",\"name\":\"testmgmt\",\"max_idle_con\":5,\"max_open_con\":10,\"driver\":\"mysql\"}}',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(40,'url_prefix','api/v1',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(41,'logger_encoder_level_key','level',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(42,'logger_encoder_caller_key','caller',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(43,'logger_encoder_stacktrace_key','stacktrace',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(44,'hide_config_center_entrance','true',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(45,'access_log_off','false',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(46,'access_assets_log_off','false',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(47,'logger_encoder_caller','short',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(48,'session_life_time','7200',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(49,'language','zh',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(50,'logo','        TestMgmt\r\n    ',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(51,'logger_encoder_name_key','logger',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(52,'logger_encoder_duration','string',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(53,'error_log_off','false',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(54,'app_id','j90eXvI3x1ye',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(55,'extra','',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(56,'index_url','/',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(57,'env','local',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(58,'color_scheme','skin-black',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(59,'file_upload_engine','{\"name\":\"local\"}',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(60,'bootstrap_file_path','/tmp/testmgmt/file/bootstrap.go',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(61,'footer_info','',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(62,'site_off','false',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(63,'debug','true',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(64,'access_log_path','/tmp/testmgmt/log/access.log',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(65,'logger_encoder_level','capitalColor',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(66,'no_limit_login_ip','false',NULL,1,'2020-11-23 02:04:52','2020-11-23 02:04:52'),(67,'filemanager_connection','default',NULL,0,'2020-12-07 02:36:00','2020-12-07 02:36:00');
/*!40000 ALTER TABLE `goadmin_site` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goadmin_user_permissions`
--

DROP TABLE IF EXISTS `goadmin_user_permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `goadmin_user_permissions` (
  `user_id` int(11) unsigned NOT NULL,
  `permission_id` int(11) unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY `admin_user_permissions` (`user_id`,`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_user_permissions`
--

LOCK TABLES `goadmin_user_permissions` WRITE;
/*!40000 ALTER TABLE `goadmin_user_permissions` DISABLE KEYS */;
INSERT INTO `goadmin_user_permissions` VALUES (1,1,'2019-09-09 16:00:00','2019-09-09 16:00:00'),(2,2,'2019-09-09 16:00:00','2019-09-09 16:00:00');
/*!40000 ALTER TABLE `goadmin_user_permissions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goadmin_users`
--

DROP TABLE IF EXISTS `goadmin_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `goadmin_users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `remember_token` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_users_username_unique` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_users`
--

LOCK TABLES `goadmin_users` WRITE;
/*!40000 ALTER TABLE `goadmin_users` DISABLE KEYS */;
INSERT INTO `goadmin_users` VALUES (1,'admin','$2a$10$dYdTVq7rBf2jEHSfiXeRve8C7vR67sh9Bz/W65imvfMo7RUrYaZOW','admin','','tlNcBVK9AvfYH7WEnwB1RKvocJu8FfRy4um3DJtwdHuJy0dwFsLOgAc0xUfh','2019-09-09 16:00:00','2019-09-09 16:00:00'),(2,'operator','$2a$10$Y8BSfJuwRBZ9pxgzaWpCnub0eja4XE93zbkzpep7GawO8BCJ3fK2C','Operator','',NULL,'2019-09-09 16:00:00','2019-09-09 16:00:00');
/*!40000 ALTER TABLE `goadmin_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `host`
--

DROP TABLE IF EXISTS `host`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `host` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `project` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `ip` char(39) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `protocol` char(5) COLLATE utf8mb4_unicode_ci NOT NULL,
  `auth` char(3) COLLATE utf8mb4_unicode_ci NOT NULL,
  `prepath` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `threading` char(3) COLLATE utf8mb4_unicode_ci NOT NULL,
  `usermode` varchar(6) COLLATE utf8mb4_unicode_ci NOT NULL,
  `dbconfig` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `token` longtext COLLATE utf8mb4_unicode_ci,
  `testmode` varchar(8) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `project` (`project`),
  UNIQUE KEY `host` (`project`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `host`
--

LOCK TABLES `host` WRITE;
/*!40000 ALTER TABLE `host` DISABLE KEYS */;
INSERT INTO `host` VALUES (17,'TEST','10.0.x.x','http','yes','/api/test/v1','no','native','','','abnormal','2020-12-28 14:18:32','2020-12-28 14:18:32',NULL);
/*!40000 ALTER TABLE `host` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `issue`
--

DROP TABLE IF EXISTS `issue`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `issue` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `product` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `milestone` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `issue_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `issue_type` varchar(8) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `author` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `assignees` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `examiner` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `updated` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `result` varchar(8) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `reopen` varchar(8) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `tag` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `project` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `issue_id` (`issue_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `issue`
--

LOCK TABLES `issue` WRITE;
/*!40000 ALTER TABLE `issue` DISABLE KEYS */;
/*!40000 ALTER TABLE `issue` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `issue_milestone_count`
--

DROP TABLE IF EXISTS `issue_milestone_count`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `issue_milestone_count` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `product` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `milestone` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `all_count` int(11) NOT NULL,
  `open_count` int(11) NOT NULL,
  `closed_count` int(11) NOT NULL,
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `project` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `issue_milestone_count`
--

LOCK TABLES `issue_milestone_count` WRITE;
/*!40000 ALTER TABLE `issue_milestone_count` DISABLE KEYS */;
/*!40000 ALTER TABLE `issue_milestone_count` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `issue_tag_count`
--

DROP TABLE IF EXISTS `issue_tag_count`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `issue_tag_count` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `milestone` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tag` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `all_count` int(11) NOT NULL,
  `open_count` int(11) NOT NULL,
  `closed_count` int(11) NOT NULL,
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `product` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `project` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `issue_tag_count`
--

LOCK TABLES `issue_tag_count` WRITE;
/*!40000 ALTER TABLE `issue_tag_count` DISABLE KEYS */;
/*!40000 ALTER TABLE `issue_tag_count` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_gitlab`
--

DROP TABLE IF EXISTS `product_gitlab`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `product_gitlab` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `repo` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `product` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `rss_token` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `project` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `milestone` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `product_project_milestone` (`product`,`milestone`,`project`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_gitlab`
--

LOCK TABLES `product_gitlab` WRITE;
/*!40000 ALTER TABLE `product_gitlab` DISABLE KEYS */;
/*!40000 ALTER TABLE `product_gitlab` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `test_case`
--

DROP TABLE IF EXISTS `test_case`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `test_case` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `case_number` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `case_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `case_type` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `priority` varchar(6) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `pre_condition` longtext COLLATE utf8mb4_unicode_ci,
  `test_range` longtext COLLATE utf8mb4_unicode_ci,
  `test_steps` longtext COLLATE utf8mb4_unicode_ci,
  `expect_result` longtext COLLATE utf8mb4_unicode_ci,
  `auto` varchar(3) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `fun_developer` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `case_designer` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `case_executor` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `test_time` varchar(12) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `test_result` varchar(12) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `module` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `case_id` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `project` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `case_number_project` (`case_number`,`project`)
) ENGINE=InnoDB AUTO_INCREMENT=446 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `test_case`
--

LOCK TABLES `test_case` WRITE;
/*!40000 ALTER TABLE `test_case` DISABLE KEYS */;
/*!40000 ALTER TABLE `test_case` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `test_progress_schedule`
--

DROP TABLE IF EXISTS `test_progress_schedule`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `test_progress_schedule` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `task_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `p_start_time` datetime(6) DEFAULT NULL,
  `p_finish_time` datetime(6) DEFAULT NULL,
  `a_start_time` datetime(6) DEFAULT NULL,
  `a_finish_time` datetime(6) DEFAULT NULL,
  `progress` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `milestone` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `priority` varchar(6) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `executor` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `project` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `test_progress_schedule`
--

LOCK TABLES `test_progress_schedule` WRITE;
/*!40000 ALTER TABLE `test_progress_schedule` DISABLE KEYS */;
/*!40000 ALTER TABLE `test_progress_schedule` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `testcase_count`
--

DROP TABLE IF EXISTS `testcase_count`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `testcase_count` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `module` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `allcase` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `uninclude_untest` int(11) NOT NULL,
  `pass` int(11) NOT NULL,
  `fail` int(11) NOT NULL,
  `pass_per` double DEFAULT NULL,
  `project` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `testcase_count`
--

LOCK TABLES `testcase_count` WRITE;
/*!40000 ALTER TABLE `testcase_count` DISABLE KEYS */;
/*!40000 ALTER TABLE `testcase_count` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-12-28 22:33:11
