	CREATE TABLE `t_files`  (
	  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
	  `file_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件名',
	  `file_path` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件目录',
	  `content_type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件类型',
	  `size` int(11) NULL DEFAULT NULL COMMENT '文件大小',
	  `status` tinyint(4) NULL DEFAULT NULL COMMENT '文件保存状态 1：成功 2：失败',
	  `err_msg` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '异常说明',
	  `create_at` bigint(20) NULL DEFAULT NULL COMMENT '创建时间',
	  `update_at` bigint(20) NULL DEFAULT NULL COMMENT '更新时间',
	  PRIMARY KEY (`id`) USING BTREE
	) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;