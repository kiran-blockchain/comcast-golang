-- Create tables (from previous step)
CREATE TABLE Users (
    UserID SERIAL PRIMARY KEY,
    Username VARCHAR(50) NOT NULL UNIQUE,
    Email VARCHAR(100) NOT NULL UNIQUE,
    PasswordHash VARCHAR(255) NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Addresses (
    AddressID SERIAL PRIMARY KEY,
    UserID INT REFERENCES Users(UserID),
    Street VARCHAR(255) NOT NULL,
    City VARCHAR(100) NOT NULL,
    State VARCHAR(100) NOT NULL,
    ZipCode VARCHAR(20) NOT NULL,
    Country VARCHAR(100) NOT NULL
);

CREATE TABLE Categories (
    CategoryID SERIAL PRIMARY KEY,
    CategoryName VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE Products (
    ProductID SERIAL PRIMARY KEY,
    CategoryID INT REFERENCES Categories(CategoryID),
    ProductName VARCHAR(100) NOT NULL,
    Description TEXT,
    Price DECIMAL(10, 2) NOT NULL,
    Stock INT NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Orders (
    OrderID SERIAL PRIMARY KEY,
    UserID INT REFERENCES Users(UserID),
    AddressID INT REFERENCES Addresses(AddressID),
    OrderStatus VARCHAR(50) NOT NULL,
    OrderTotal DECIMAL(10, 2) NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE OrderItems (
    OrderItemID SERIAL PRIMARY KEY,
    OrderID INT REFERENCES Orders(OrderID),
    ProductID INT REFERENCES Products(ProductID),
    Quantity INT NOT NULL,
    Price DECIMAL(10, 2) NOT NULL,
    Total DECIMAL(10, 2) GENERATED ALWAYS AS (Quantity * Price) STORED
);

CREATE TABLE Payments (
    PaymentID SERIAL PRIMARY KEY,
    OrderID INT REFERENCES Orders(OrderID),
    PaymentMethod VARCHAR(50) NOT NULL,
    PaymentStatus VARCHAR(50) NOT NULL,
    Amount DECIMAL(10, 2) NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert dummy data
INSERT INTO Users (Username, Email, PasswordHash) VALUES
('kiran1', 'kiran@gmail.com', 'passwordhash1'),
('kiran2', 'kiran2@gmail.com', 'passwordhash2');

INSERT INTO Addresses (UserID, Street, City, State, ZipCode, Country) VALUES
(1, '123 Elm Street', 'Springfield', 'IL', '62704', 'USA'),
(2, '456 Oak Avenue', 'Metropolis', 'NY', '10001', 'USA');

INSERT INTO Categories (CategoryName) VALUES
('Electronics'),
('Clothing');

INSERT INTO Products (CategoryID, ProductName, Description, Price, Stock) VALUES
(1, 'Smartphone', 'Latest model smartphone with many features', 699.99, 50),
(1, 'Laptop', 'High performance laptop', 1299.99, 30),
(2, 'T-shirt', 'Comfortable cotton t-shirt', 19.99, 100),
(2, 'Jeans', 'Stylish denim jeans', 49.99, 60);

INSERT INTO Orders (UserID, AddressID, OrderStatus, OrderTotal) VALUES
(1, 1, 'Pending', 719.98),
(2, 2, 'Completed', 69.98);

INSERT INTO OrderItems (OrderID, ProductID, Quantity, Price) VALUES
(1, 1, 1, 699.99),
(1, 3, 1, 19.99),
(2, 4, 1, 49.99),
(2, 3, 1, 19.99);

INSERT INTO Payments (OrderID, PaymentMethod, PaymentStatus, Amount) VALUES
(1, 'Credit Card', 'Pending', 719.98),
(2, 'PayPal', 'Completed', 69.98);
