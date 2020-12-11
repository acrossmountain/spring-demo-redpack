/*
 Navicat Premium Data Transfer

 Source Server         : 本地数据库
 Source Server Type    : MySQL
 Source Server Version : 80012
 Source Host           : localhost:3306
 Source Schema         : red_pack

 Target Server Type    : MySQL
 Target Server Version : 80012
 File Encoding         : 65001

 Date: 11/12/2020 13:30:02
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for account
-- ----------------------------
DROP TABLE IF EXISTS `account`;
CREATE TABLE `account`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '账户ID',
  `account_no` varchar(32) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '账户编号',
  `account_name` varchar(64) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '账户名称',
  `account_type` tinyint(2) NOT NULL COMMENT '账户类型',
  `currency_code` char(10) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT 'CNY' COMMENT '货币类型',
  `user_id` varchar(40) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '用户编号',
  `username` varchar(64) CHARACTER SET utf8 COLLATE utf8_unicode_ci NULL DEFAULT NULL COMMENT '用户名称',
  `balance` decimal(30, 6) UNSIGNED NULL DEFAULT 0.000000 COMMENT '账户可用余额',
  `status` tinyint(2) NOT NULL COMMENT '账户状态: 0账户初始化，1启用，2停用',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `account_no`(`account_no`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 56 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for account_log
-- ----------------------------
DROP TABLE IF EXISTS `account_log`;
CREATE TABLE `account_log`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `trade_no` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '交易单号 全局不重复字符或数字，唯一性标识 ',
  `log_no` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '流水编号 全局不重复字符或数字，唯一性标识 ',
  `account_no` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '账户编号 账户ID',
  `target_account_no` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '账户编号 账户ID',
  `user_id` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户编号',
  `username` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户名称',
  `target_user_id` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '目标用户编号',
  `target_username` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '目标用户名称',
  `amount` decimal(30, 6) NOT NULL DEFAULT 0.000000 COMMENT '交易金额,该交易涉及的金额 ',
  `balance` decimal(30, 6) UNSIGNED NOT NULL DEFAULT 0.000000 COMMENT '交易后余额,该交易后的余额 ',
  `change_type` tinyint(2) NOT NULL DEFAULT 0 COMMENT '流水交易类型，0 创建账户，>0 为收入类型，<0 为支出类型，自定义',
  `change_flag` tinyint(2) NOT NULL DEFAULT 0 COMMENT '交易变化标识：-1 出账 1为进账，枚举',
  `status` tinyint(2) NOT NULL DEFAULT 0 COMMENT '交易状态：',
  `decs` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '交易描述 ',
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id_log_no_idx`(`log_no`) USING BTREE,
  INDEX `id_user_idx`(`user_id`) USING BTREE,
  INDEX `id_account_idx`(`account_no`) USING BTREE,
  INDEX `id_trade_idx`(`trade_no`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 69 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
