/*
 Navicat Premium Data Transfer

 Source Server         : bt mysql
 Source Server Type    : MySQL
 Source Server Version : 80029
 Source Host           : 127.0.0.1:3306
 Source Schema         : zero-shop-pay

 Target Server Type    : MySQL
 Target Server Version : 80029
 File Encoding         : 65001

 Date: 22/07/2023 17:54:21
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for third_pay
-- ----------------------------
DROP TABLE IF EXISTS `third_pay`;
CREATE TABLE `third_pay`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `sn` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '订单流水号',
  `create_time` datetime NOT NULL,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `user_id` int UNSIGNED NOT NULL DEFAULT 0,
  `pay_mode` tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '1 微信支付 2 支付宝支付',
  `trade_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '支付类型',
  `trade_state` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '交易状态',
  `pay_total` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '支付金额',
  `transaction_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '微信的订单号',
  `trade_state_desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '订单支付描述',
  `order_sn` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '订单sn',
  `pay_status` tinyint(1) NOT NULL COMMENT '-1:支付失败 0:未支付 1:支付成功 2:已退款',
  `pay_time` datetime NULL DEFAULT NULL COMMENT '支付时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of third_pay
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
