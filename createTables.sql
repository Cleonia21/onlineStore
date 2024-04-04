CREATE TABLE Orders
(
    Id SERIAL PRIMARY KEY,
    Number INTEGER NOT NULL,
    ProductId INTEGER,
    Quantity INTEGER NOT NULL,
    FOREIGN KEY (ProductId) REFERENCES Products (Id)
);

CREATE TABLE Products
(
    Id SERIAL PRIMARY KEY,
    Name text NOT NULL,
    MainShelvingId INTEGER,
    FOREIGN KEY (MainShelvingId) REFERENCES Shelving (Id) ON DELETE RESTRICT
);

CREATE TABLE Shelving
(
    Id SERIAL PRIMARY KEY,
    Name CHARACTER(1) NOT NULL
);

CREATE TABLE ProductsOnTheShelving
(
    Id SERIAL PRIMARY KEY,
    ShelvingId INTEGER,
    ProductId INTEGER,
    FOREIGN KEY (ShelvingId) REFERENCES Shelving (Id) ON DELETE RESTRICT,
    FOREIGN KEY (ProductId) REFERENCES Products (Id) ON DELETE RESTRICT
);
