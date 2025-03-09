CREATE TABLE `order_item` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `order_id` int unsigned NOT NULL COMMENT '订单ID',
  `product_id` int unsigned NOT NULL COMMENT '商品ID',
  `product_name` varchar(200) NOT NULL COMMENT '商品名称',
  `product_sku` varchar(50) NOT NULL COMMENT '商品编码',
  `product_price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '商品单价',
  `quantity` int unsigned NOT NULL DEFAULT '1' COMMENT '数量',
  `total_amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '总金额',
  `discount_amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '优惠金额',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_order` (`order_id`),
  KEY `idx_product` (`product_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci; 