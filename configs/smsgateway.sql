/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 80025
 Source Host           : localhost:3306
 Source Schema         : smsgateway

 Target Server Type    : MySQL
 Target Server Version : 80025
 File Encoding         : 65001

 Date: 10/07/2024 16:29:52
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for consumer
-- ----------------------------
DROP TABLE IF EXISTS `consumer`;
CREATE TABLE `consumer`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `consumer_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `api_key` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `api_secret` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `create_time` datetime NULL DEFAULT NULL,
  `modify_time` datetime NULL DEFAULT NULL,
  `create_by` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `modify_by` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `deleted` bit(1) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of consumer
-- ----------------------------
INSERT INTO `consumer` VALUES (2, 'mediator', 'a6c9f747-ab37-4dc9-b297-d5755176bd53', '95fb1005-4761-47c7-9a6b-2f711f83e0b4', '2017-01-06 15:10:24', '2023-08-15 14:12:07', 'admin', '', b'0');

-- ----------------------------
-- Table structure for demander
-- ----------------------------
DROP TABLE IF EXISTS `demander`;
CREATE TABLE `demander`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `category` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `brand` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `io` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `fiscal_year` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `budget_owner` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `brand_fa` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `create_time` datetime NULL DEFAULT NULL,
  `create_by` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `modify_time` datetime NULL DEFAULT NULL,
  `modify_by` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `deleted` bit(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of demander
-- ----------------------------
INSERT INTO `demander` VALUES (1, 'HC', 'PNT', '4001824729', 'FY2223', 'zhao.y.39@pg.com', 'ye.s.5@pg.com', '2023-02-17 17:16:13', '', '2023-03-17 15:05:17', '', b'0');

-- ----------------------------
-- Table structure for dictionary
-- ----------------------------
DROP TABLE IF EXISTS `dictionary`;
CREATE TABLE `dictionary`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `label` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `value` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `pIds` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of dictionary
-- ----------------------------
INSERT INTO `dictionary` VALUES (4, 'Pampers', 'Pampers', 'brand', '1,26');

-- ----------------------------
-- Table structure for report
-- ----------------------------
DROP TABLE IF EXISTS `report`;
CREATE TABLE `report`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `fiscal_year` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `month` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `brand` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `io` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `estimated_x_charge_amount` decimal(10, 4) NULL DEFAULT NULL,
  `msg_delivery_num` int NULL DEFAULT NULL,
  `budget_owner` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `brand_fa` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `real_x_charge_amount` decimal(10, 4) NULL DEFAULT NULL,
  `x_charge_status` tinyint(1) NULL DEFAULT NULL,
  `remark` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `create_time` datetime NULL DEFAULT NULL,
  `create_by` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `modify_time` datetime NULL DEFAULT NULL,
  `modify_by` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `deleted` bit(1) NULL DEFAULT NULL,
  `sms_delivery_num` int NULL DEFAULT NULL,
  `mms_delivery_num` int NULL DEFAULT NULL,
  `global_sms_delivery_num` int NULL DEFAULT NULL,
  `gc_operation_costs` decimal(10, 4) NULL DEFAULT NULL,
  `campaign_service_costs` decimal(10, 4) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of report
-- ----------------------------
INSERT INTO `report` VALUES (3, 'FY2223', '2023/01', 'GA', '4700234440', 181.3350, 5181, 'chen.s.23@pg.com', 'li.s.85@pg.com', NULL, 0, '', '2023-02-23 12:37:14', '', '2023-08-15 14:14:17', '', b'0', NULL, NULL, NULL, 0.0000, 0.0000);

-- ----------------------------
-- Table structure for report_detail
-- ----------------------------
DROP TABLE IF EXISTS `report_detail`;
CREATE TABLE `report_detail`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `date` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `category` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `brand` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `budget_owner` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `io` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `message_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `start_time` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `vendor` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `campaign_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `campaign_id` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `ext_code` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `region` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `sign` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `msg_delivery_num` int NULL DEFAULT NULL,
  `msg_succeed_num` int NULL DEFAULT NULL,
  `msg_succeed_percent` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `people_succeed_num` int NULL DEFAULT NULL,
  `create_time` datetime NULL DEFAULT NULL,
  `create_by` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `modify_time` datetime NULL DEFAULT NULL,
  `modify_by` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `deleted` bit(1) NULL DEFAULT NULL,
  `month` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `fiscal_year` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `brand_fa` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of report_detail
-- ----------------------------
INSERT INTO `report_detail` VALUES (7, '2023/02/01', 'DevOps', 'GA', 'chen.s.23@pg.com', '4700234440', 'SMS', '', 'WE', 'GA_找回密码', '', '124', 'CHINA_MAINLAND', '宝洁中国', 1, 1, '100%', 1, '2023-03-17 15:36:59', '', '2023-08-15 14:14:35', '', b'0', '2023/02', 'FY2223', 'li.s.85@pg.com');

-- ----------------------------
-- Table structure for scene
-- ----------------------------
DROP TABLE IF EXISTS `scene`;
CREATE TABLE `scene`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `ext_code` int NOT NULL,
  `account` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `marketing_program_id` int NOT NULL,
  `app` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `sign` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `token_id` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `prolonged` bit(1) NOT NULL,
  `create_time` datetime NULL DEFAULT NULL,
  `modify_time` datetime NULL DEFAULT NULL,
  `create_by` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `modify_by` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `deleted` bit(1) NULL DEFAULT NULL,
  `scene_result_topic` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `scene_ext_code_IDX`(`ext_code`, `account`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of scene
-- ----------------------------
INSERT INTO `scene` VALUES (1, 157, 'gzbjyxyx', 38, 'cms-message-gateway', '【宝洁中国】', '2021578375895426', b'1', '2020-05-22 09:08:58', '2020-07-29 21:12:57', '', '', b'0', 'AM_UN_OR_SUBSCRIBE_TOPIC');

-- ----------------------------
-- Table structure for sms_configuration
-- ----------------------------
DROP TABLE IF EXISTS `sms_configuration`;
CREATE TABLE `sms_configuration`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `config_key` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `config_value` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `template_code` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `create_at` datetime NULL DEFAULT NULL,
  `update_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 80 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sms_configuration
-- ----------------------------
INSERT INTO `sms_configuration` VALUES (79, 'MESSAGE_COUNT_FOR_TEMPLATE_PER_DAY', '100000000000', '8f263266-a411-422d-af69-81d887e9079a', '2016-11-23 16:36:26', '2019-11-21 16:22:42');

-- ----------------------------
-- Table structure for sms_history
-- ----------------------------
DROP TABLE IF EXISTS `sms_history`;
CREATE TABLE `sms_history`  (
  `id` bigint NOT NULL,
  `mobile` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `send_time` datetime NULL DEFAULT NULL,
  `consumer_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `status` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `content` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `code` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `message` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `message_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `message_template_id` bigint NULL DEFAULT NULL,
  `report` mediumtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sms_history
-- ----------------------------

-- ----------------------------
-- Table structure for sms_interaction
-- ----------------------------
DROP TABLE IF EXISTS `sms_interaction`;
CREATE TABLE `sms_interaction`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `interaction_id` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `mobile` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `content` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `param` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `provider` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `account` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `receive_time` datetime NULL DEFAULT NULL,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `INDEX_INTERACTION_ID`(`interaction_id`) USING BTREE,
  INDEX `INDEX_MOBILE_AND_CONTENT`(`mobile`, `content`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sms_interaction
-- ----------------------------

-- ----------------------------
-- Table structure for sms_provider
-- ----------------------------
DROP TABLE IF EXISTS `sms_provider`;
CREATE TABLE `sms_provider`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `provider_code` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `provider_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `description` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `create_time` datetime NULL DEFAULT NULL,
  `create_by` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `modify_time` datetime NULL DEFAULT NULL,
  `modify_by` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `deleted` bit(1) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 38 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sms_provider
-- ----------------------------
INSERT INTO `sms_provider` VALUES (35, 'WE', 'WE', 'WE短信平台', '2019-09-26 17:21:26', 'admin', '2019-09-26 17:21:26', 'admin', b'0');
INSERT INTO `sms_provider` VALUES (36, 'XSXX', 'XSXX', '无锡线上线下短信平台', '2019-09-26 17:21:26', 'admin', '2019-09-26 17:21:26', 'admin', b'0');
INSERT INTO `sms_provider` VALUES (37, 'DUMMY', 'DUMMY', 'mock的provider', '2022-08-18 16:44:00', 'admin', '2022-08-18 16:44:00', 'admin', b'0');
INSERT INTO `sms_provider` VALUES (38, 'ALIYUN', 'ALIYUN', '阿里云短信平台', '2022-09-13 16:44:00', '2022-09-13 16:44:00', '2022-09-13 16:44:00', '2022-09-13 16:44:00', b'0');

-- ----------------------------
-- Table structure for sms_template
-- ----------------------------
DROP TABLE IF EXISTS `sms_template`;
CREATE TABLE `sms_template`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `template_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `template_code` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `type` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `brand_id` bigint NULL DEFAULT NULL,
  `token_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `sign` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `primary_ISP` varchar(45) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `user_id` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `ext_code` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `brand` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `provider` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `region` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '国家地区',
  `providers` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `demander_ids` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `create_at` datetime NULL DEFAULT NULL,
  `update_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sms_template
-- ----------------------------
INSERT INTO `sms_template` VALUES (2, 'pampers-0001', '8f263266-a411-422d-af69-81d887e9079a', 'VERIFICATION', '亲爱的帮宝适会员：欢迎您加入帮宝适成长俱乐部。请使用此验证码${code}确认注册登录信息。帮宝适官方网站将为您提供全面权威的育儿资讯，现在加入还有机会获得帮宝适特级棉柔试用装哦！www.pampers.com.cn', 2, '9695800120000256', '【帮宝适品牌】', '35', 'gzbjhyhy', '1', 'PAMPERS', 'XSXX', 'CHINA_MAINLAND', 'WE,XSXX', '', '2016-11-23 16:36:26', '2019-11-21 16:22:42');

-- ----------------------------
-- Table structure for uc_user
-- ----------------------------
DROP TABLE IF EXISTS `uc_user`;
CREATE TABLE `uc_user`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `nickname` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `email` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `phone` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `status` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户状态：registered,active,disabled,blacklisted,locked,deleted',
  `createdAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 29 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of uc_user
-- ----------------------------

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `script` blob NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (18, 0x724F304142585A794142317A59334A70634851784E7A45794E7A4D354D6A67774D7A45334D544D794E5455794E5445314E7741414141414141414141414141416548413D);
INSERT INTO `user` VALUES (19, 0xACED00057672001C7363726970743137313237343138353538313632333438393830323500000000000000000000007870);
INSERT INTO `user` VALUES (20, 0xACED00057672001C7363726970743137313237343139353630363432333438393830323500000000000000000000007870);

SET FOREIGN_KEY_CHECKS = 1;
