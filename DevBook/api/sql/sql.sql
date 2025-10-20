CREATE DATABASE IF NOT EXISTS devbook;
USE devbook

DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL UNIQUE,
    senha VARCHAR(100) NOT NULL,
    nick VARCHAR(50) NOT NULL,
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    atualizado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE seguidores (
    usuario_id INT NOT NULL,
    seguidor_id INT NOT NULL,
    PRIMARY KEY (usuario_id, seguidor_id),
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id) ON DELETE CASCADE,
    FOREIGN KEY (seguidor_id) REFERENCES usuarios(id) ON DELETE CASCADE
);

CREATE TABLE publicacoes (
    it int auto_increment primary KEY
    titulo varchar (50) not null,
    conteudo varchar (300) not null,
    
    autor_id int not null,
    FOREIGN key (autor_id)
    REFERENCES usuarios(id)
    on DELETE CASCADE,
    curtidas int default 0
    criadaEm timestamp default CURRENT_TIMESTAMP
)