CREATE TABLE Addres
(
    id SERIAL NOT NULL,
    country   VARCHAR(32) NOT NULL,
	city      VARCHAR(32) NOT NULL,
    street    VARCHAR(32) NOT NULL,
	house     VARCHAR(32) NOT NULL,
	apartment VARCHAR(32)
);

CREATE TABLE People 
(
    id SERIAL NOT NULL,
    name VARCHAR(32) NOT NULL,
    surname VARCHAR(32) NOT NULL,
    age INT NOT NULL,
    addres_id SERIAL NOT NULL,
    phone VARCHAR(32) NOT NULL UNIQUE,
    PRIMARY KEY (id, addres_id)
);