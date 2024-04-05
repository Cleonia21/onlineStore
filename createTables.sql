CREATE TABLE shelving
(
    id SERIAL PRIMARY KEY,
    name text NOT NULL
);

CREATE TABLE products
(
    id      SERIAL PRIMARY KEY,
    name    text NOT NULL,
    shelfId INTEGER REFERENCES shelving (id) ON DELETE RESTRICT ON UPDATE CASCADE
);

CREATE TABLE optionalShelving
(
    id        SERIAL PRIMARY KEY,
    productId INTEGER REFERENCES products (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    shelfId   INTEGER REFERENCES shelving (id) ON DELETE RESTRICT ON UPDATE CASCADE
);

/*
 В таблице присутствует логическая проблемма вставки 2 записей с одинаковыми
 номером заказа и id товара, количество товара может быть одинаковым или отличаться.
 Не стал замудряться с контролем этой ошибки на уровне БД.
 Предполагаю, что данную проблемму решит фронтенд.
 */
CREATE TABLE orders
(
    id        SERIAL PRIMARY KEY,
    number    INTEGER NOT NULL,
    productId INTEGER REFERENCES products (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    quantity  INTEGER NOT NULL
);
