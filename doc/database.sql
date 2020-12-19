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

 Date: 19/12/2020 10:53:24
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
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = Dynamic;

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
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for red_envelope_goods
-- ----------------------------
DROP TABLE IF EXISTS `red_envelope_goods`;
CREATE TABLE `red_envelope_goods`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `envelope_no` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '红包编号,红包唯一标识 ',
  `envelope_type` tinyint(2) NOT NULL COMMENT '红包类型：普通红包，碰运气红包,过期红包',
  `username` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户名称',
  `user_id` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户编号, 红包所属用户 ',
  `blessing` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '祝福语',
  `amount` decimal(30, 6) UNSIGNED NOT NULL DEFAULT 0.000000 COMMENT '红包总金额',
  `amount_one` decimal(30, 6) UNSIGNED NOT NULL DEFAULT 0.000000 COMMENT '单个红包金额，碰运气红包无效',
  `quantity` int(10) UNSIGNED NOT NULL COMMENT '红包总数量 ',
  `remain_amount` decimal(30, 6) UNSIGNED NOT NULL DEFAULT 0.000000 COMMENT '红包剩余金额额',
  `remain_quantity` int(10) UNSIGNED NOT NULL COMMENT '红包剩余数量 ',
  `expired_at` datetime(3) NOT NULL COMMENT '过期时间',
  `status` tinyint(2) NOT NULL COMMENT '红包/订单状态：0 创建、1 发布启用、2过期、3失效',
  `order_type` tinyint(2) NOT NULL COMMENT '订单类型：发布单、退款单 ',
  `pay_status` tinyint(2) NOT NULL COMMENT '支付状态：未支付，支付中，已支付，支付失败 ',
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `envelope_no_idx`(`envelope_no`) USING BTREE,
  INDEX `id_user_idx`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1276 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for red_envelope_item
-- ----------------------------
DROP TABLE IF EXISTS `red_envelope_item`;
CREATE TABLE `red_envelope_item`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `item_no` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '红包订单详情编号 ',
  `envelope_no` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '红包编号,红包唯一标识 ',
  `recv_username` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '红包接收者用户名称',
  `recv_user_id` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '红包接收者用户编号 ',
  `amount` decimal(30, 6) UNSIGNED NOT NULL DEFAULT 0.000000 COMMENT '收到金额',
  `quantity` int(10) UNSIGNED NOT NULL COMMENT '收到数量：对于收红包来说是1 ',
  `remain_amount` decimal(30, 6) UNSIGNED NOT NULL DEFAULT 0.000000 COMMENT '收到后红包剩余金额',
  `account_no` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '红包接收者账户ID',
  `pay_status` tinyint(2) NOT NULL COMMENT '支付状态：未支付，支付中，已支付，支付失败 ',
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `item_no_idx`(`item_no`) USING BTREE,
  INDEX `envelope_no_idx`(`envelope_no`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
