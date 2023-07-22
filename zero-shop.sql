/*
 Navicat Premium Data Transfer

 Source Server         : bt mysql
 Source Server Type    : MySQL
 Source Server Version : 80029
 Source Host           : 127.0.0.1:3306
 Source Schema         : zero-shop

 Target Server Type    : MySQL
 Target Server Version : 80029
 File Encoding         : 65001

 Date: 22/07/2023 17:54:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `pid` int UNSIGNED NOT NULL DEFAULT 0,
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of category
-- ----------------------------
INSERT INTO `category` VALUES (1, '电子产品', 0, '|', '2023-07-21 10:47:53', '2023-07-21 10:47:53');

-- ----------------------------
-- Table structure for goods
-- ----------------------------
DROP TABLE IF EXISTS `goods`;
CREATE TABLE `goods`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `goods_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '商品名称',
  `category_id` int NOT NULL COMMENT '商品分类',
  `on_sale` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否在销售 1 是 0 否',
  `ship_free` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否免运费 1 是 0 否',
  `is_new` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否最新 1 是 0 否',
  `is_hot` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否最热 1 是 0 否',
  `click_num` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '查看次数',
  `sold_num` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '售出数量',
  `fav_num` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '收藏数量',
  `market_price` decimal(10, 2) UNSIGNED NOT NULL COMMENT '市场价',
  `shop_price` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '售价',
  `goods_brief` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '商品简介',
  `goods_detail` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '商品详情',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of goods
-- ----------------------------
INSERT INTO `goods` VALUES (1, '手机', 1, 1, 1, 1, 1, 20, 20, 50, 150.00, 99.00, '好手机', '真是个好手机啊', '2023-07-21 10:48:07', '2023-07-21 10:48:07');

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `user_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
  `goods_id` int UNSIGNED NOT NULL COMMENT '商品id',
  `market_price` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '市场价',
  `sale_price` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '售价',
  `order_status` tinyint UNSIGNED NOT NULL COMMENT '订单状态 -1:支付失败 0:未支付 1:支付成功 2:已退款',
  `sn` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '订单号',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of order
-- ----------------------------
INSERT INTO `order` VALUES (1, '2023-07-22 16:41:23', '2023-07-22 16:41:23', 1, 64163076, 0.00, 0.00, 0, 'ORD20230722164123UP26Ahe1');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '张三', 'e10adc3949ba59abbe56e057f20f883e', '2023-07-20 10:56:49');
INSERT INTO `user` VALUES (2, '李四', 'e10adc3949ba59abbe56e057f20f883e', '2023-07-20 10:56:49');
INSERT INTO `user` VALUES (3, '王五', 'e10adc3949ba59abbe56e057f20f883e', '2023-07-20 10:56:49');
INSERT INTO `user` VALUES (4, '王五1', 'e10adc3949ba59abbe56e057f20f883e', '2023-07-20 10:57:13');

SET FOREIGN_KEY_CHECKS = 1;
