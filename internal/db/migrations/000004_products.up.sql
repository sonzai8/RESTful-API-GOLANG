CREATE TABLE IF NOT EXISTS products(
                                       product_id SERIAL PRIMARY KEY,
                                       category_id INT NOT NULL,
                                       name VARCHAR(100) NOT NULL,
    price INT NOT NULL CHECK (price >= 0),
    image varchar(255),
    status INT NOT NULL CHECK(status in (1,2)),
    CONSTRAINT fk_category FOREIGN KEY (category_id) references categories (category_id) on DELETE RESTRICT
    );
