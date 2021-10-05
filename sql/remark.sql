/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 50718
 Source Host           : localhost:3306
 Source Schema         : remark

 Target Server Type    : MySQL
 Target Server Version : 50718
 File Encoding         : 65001

 Date: 06/10/2021 00:25:05
*/
CREATE DATA IF NOT EXISTS `remark`;
USE remark;
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for remark
-- ----------------------------
DROP TABLE IF EXISTS `remark`;
CREATE TABLE `remark`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `content` varchar(512) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `create_time` datetime NOT NULL,
  `update_time` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of remark
-- ----------------------------
INSERT INTO `remark` VALUES (1, '人生在勤,不索何获', '2021-10-05 14:11:10', '2021-10-05 14:11:10');
INSERT INTO `remark` VALUES (2, '人有冲天之志，非运不能自通', '2021-10-05 14:11:10', '2021-10-05 14:11:10');
INSERT INTO `remark` VALUES (3, '人们的悲欢并不相同，我只是觉得他们吵闹', '2021-10-05 15:46:24', '2021-10-05 15:46:24');

SET FOREIGN_KEY_CHECKS = 1;
