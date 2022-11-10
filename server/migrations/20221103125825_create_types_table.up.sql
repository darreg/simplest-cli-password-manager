CREATE TABLE IF NOT EXISTS types
(
    id UUID NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    is_binary BOOLEAN NOT NULL DEFAULT false
);

INSERT INTO types (id, name, is_binary) VALUES ('84b4f166-d146-4e37-a858-90b5aa3476f2', 'логин/пароль', false) ON CONFLICT DO NOTHING;
INSERT INTO types (id, name, is_binary) VALUES ('9e5744cf-fe2a-465e-92c3-4a5937fbfbb8', 'текст', false) ON CONFLICT DO NOTHING;
INSERT INTO types (id, name, is_binary) VALUES ('02b3bc02-2355-4197-b011-3ac5fded4b8c', 'банковская карта', false) ON CONFLICT DO NOTHING;
INSERT INTO types (id, name, is_binary) VALUES ('c625b8cb-04ac-4a82-9d97-358eb6d94109', 'произвольные данные', true) ON CONFLICT DO NOTHING;





