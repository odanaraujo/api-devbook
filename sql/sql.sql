CREATE DATABASE IF NOT EXISTS dansa;
USE dansa;

DROP TABLE IF EXISTS publicacoes;
DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios (
    `id` int NOT NULL AUTO_INCREMENT,
    `nome` varchar(50) NOT NULL,
    `nick` varchar(50) NOT NULL unique,
    `email` varchar(50) NOT NULL unique,
    `senha` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `dataCriacao` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE seguidores (
    usuario_id int NOT NULL,
    foreign key (usuario_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    seguidor_id int NOT NULL,
    FOREIGN KEY (seguidor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,
    primary key (usuario_id, seguidor_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE publicacoes (
    `id` int NOT NULL AUTO_INCREMENT,
    `titulo` varchar(50) NOT NULL,
    `conteudo` varchar(150)NOT NULL,
    `autor_id` int NOT NULL,
    FOREIGN KEY (autor_id) REFERENCES usuarios(id)
    ON DELETE CASCADE,
    `likes` varchar(50) default 0,
    `dataCriacao` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;