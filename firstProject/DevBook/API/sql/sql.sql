CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;


DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id int auto_increment primary_key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(20) not null unique,
    criadoEM timestamp default current_timestamp()
) ENGINE=INNODB;