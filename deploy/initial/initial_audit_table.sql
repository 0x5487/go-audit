CREATE TABLE `audits` (
  `sid` int(11) NOT NULL AUTO_INCREMENT,
  `namespace` varchar(256) NOT NULL,
  `target_id` varchar(256) NOT NULL,
  `action` varchar(64) NOT NULL,
  `message` varchar(1024) NOT NULL,
  `state` tinyint(4) NOT NULL COMMENT '0: failed, 1: success',
  `actor` varchar(32) NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`sid`),
  KEY `idx_audit` (`created_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8