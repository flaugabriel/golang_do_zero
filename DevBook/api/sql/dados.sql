insert into usuarios (nome, email, senha, nick) 
values 
 ('Admin', 'admin@devbook.com', 'admin123', 'admin'),
 ('User Test', 'usertest@devbook.com', 'usertest123', 'usertest'),
 ('Jane Doe', 'janedoe@devbook.com', 'janedoe123', 'janedoe'),
 ('John Smith', 'johnsmith@devbook.com', 'johnsmith123', 'johnsmith');

 insert into seguidores (usuario_id, seguidor_id)
 values
 (1, 2),
 (1, 3),
 (2, 3),
 (3, 4),
 (4, 1);