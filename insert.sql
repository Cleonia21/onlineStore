INSERT INTO shelving
VALUES
    (1, 'A'),
    (2, 'Б'),
    (3, 'Ж'),
    (4, 'З'),
    (5, 'В');

INSERT INTO products
VALUES
    (1, 'Ноутбук', 1),
    (2, 'Телевизор', 1),
    (3, 'Телефон', 2),
    (4, 'Системный блок', 3),
    (5, 'Часы', 3),
    (6, 'Микрофон', 3);

INSERT INTO optionalShelving
VALUES
    (1, 3, 4),
    (2, 3, 5),
    (3, 5, 1);

INSERT INTO orders
VALUES
    (1, 10, 1, 2),
    (2, 11, 2, 3),
    (3, 14, 1, 3),
    (4, 10, 3, 1),
    (5, 14, 4, 4),
    (6, 15, 5, 1),
    (7, 10, 6, 1);

ALTER SEQUENCE orders_id_seq RESTART WITH 1;
delete from orders *;
ALTER SEQUENCE optionalShelving_id_seq RESTART WITH 1;
delete from optionalShelving *;
ALTER SEQUENCE products_id_seq RESTART WITH 1;
delete from products *;
ALTER SEQUENCE shelving_id_seq RESTART WITH 1;
delete from shelving *;