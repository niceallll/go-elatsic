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

 Date: 24/02/2025 14:58:52
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for es_version
-- ----------------------------
DROP TABLE IF EXISTS `es_version`;
CREATE TABLE `es_version`  (
  `es_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `es_version` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
