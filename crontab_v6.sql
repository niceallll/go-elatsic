/*
 Navicat Premium Data Transfer

 Source Server         : ucode-test-mysql
 Source Server Type    : MySQL
 Source Server Version : 50736
 Source Host           : 124.222.128.20:3306
 Source Schema         : elastic-test

 Target Server Type    : MySQL
 Target Server Version : 50736
 File Encoding         : 65001

 Date: 24/02/2025 14:58:30
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for crontab_v6
-- ----------------------------
DROP TABLE IF EXISTS `crontab_v6`;
CREATE TABLE `crontab_v6`  (
  `id` int(64) NOT NULL AUTO_INCREMENT,
  `size` int(64) NULL DEFAULT NULL,
  `es_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `es_index` json NULL,
  `tlp` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `ding_talk` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `es_most_msg` json NULL,
  `es_most_not_msg` json NULL,
  `aliname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `warn_up` int(64) NULL DEFAULT NULL,
  `del` int(64) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 70 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
