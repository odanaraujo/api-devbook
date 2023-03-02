CREATE DATABASE IF NOT EXISTS dansa;
USE dansa;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE `dansa`.`usuarios` (
    `id` int NOT NULL AUTO_INCREMENT,
    `nome` varchar(50) NOT NULL,
    `nick` varchar(50) NOT NULL unique,
    `email` varchar(50) NOT NULL unique,
    `senha` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `dataCriacao` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
