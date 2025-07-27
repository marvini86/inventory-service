
CREATE TABLE IF NOT EXISTS items (
    id SERIAL PRIMARY KEY,
    code_integration VARCHAR(50) NOT NULL,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(50) NOT NULL,
    image VARCHAR(50) NOT NULL,
    available_quantity INT NOT NULL,
    category_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS item_categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);


ALTER TABLE items ADD CONSTRAINT items_category_id_fkey FOREIGN KEY (category_id) REFERENCES item_categories(id);
ALTER TABLE items ADD CONSTRAINT items_code_integration_fkey UNIQUE (code_integration);

INSERT INTO item_categories (name, created_at, updated_at) VALUES ('Car parts', '2023-01-01 00:00:00', '2023-01-01 00:00:00');
INSERT INTO item_categories (name, created_at, updated_at) VALUES ('Motorcycle parts', '2023-01-01 00:00:00', '2023-01-01 00:00:00');
INSERT INTO item_categories (name, created_at, updated_at) VALUES ('Truck parts', '2023-01-01 00:00:00', '2023-01-01 00:00:00');
INSERT INTO item_categories (name, created_at, updated_at) VALUES ('Bus parts', '2023-01-01 00:00:00', '2023-01-01 00:00:00');

INSERT INTO items (code_integration, name, description, image, category_id, available_quantity, created_at, updated_at) VALUES ('000000001', 'Tire', 'Replacement', 'https://via.placeholder.com/150', '1', '4', '2023-01-01 00:00:00', '2023-01-01 00:00:00');
INSERT INTO items (code_integration, name, description, image, category_id, available_quantity, created_at, updated_at) VALUES ('000000002', 'Brake', 'Replacement', 'https://via.placeholder.com/150', '1', '3', '2023-01-01 00:00:00', '2023-01-01 00:00:00');
INSERT INTO items (code_integration, name, description, image, category_id, available_quantity, created_at, updated_at) VALUES ('000000003', 'Engine', 'Replacement', 'https://via.placeholder.com/150', '1', '10', '2023-01-01 00:00:00', '2023-01-01 00:00:00');
INSERT INTO items (code_integration, name, description, image, category_id, available_quantity, created_at, updated_at) VALUES ('000000004', 'Transmission', 'Replacement', 'https://via.placeholder.com/150', '1', '2', '2023-01-01 00:00:00', '2023-01-01 00:00:00');
INSERT INTO items (code_integration, name, description, image, category_id, available_quantity, created_at, updated_at) VALUES ('000000005', 'Battery', 'Replacement', 'https://via.placeholder.com/150', '1', '10', '2023-01-01 00:00:00', '2023-01-01 00:00:00');
INSERT INTO items (code_integration, name, description, image, category_id, available_quantity, created_at, updated_at) VALUES ('000000006', 'Oil', 'Change', 'https://via.placeholder.com/150', '1', '10', '2023-01-01 00:00:00', '2023-01-01 00:00:00');
INSERT INTO items (code_integration, name, description, image, category_id, available_quantity, created_at, updated_at) VALUES ('000000007', 'Tire', 'Change', 'https://via.placeholder.com/150', '1', '10', '2023-01-01 00:00:00', '2023-01-01 00:00:00');
INSERT INTO items (code_integration, name, description, image, category_id, available_quantity, created_at, updated_at) VALUES ('000000008', 'Brake', 'Change', 'https://via.placeholder.com/150', '1', '10', '2023-01-01 00:00:00', '2023-01-01 00:00:00');
INSERT INTO items (code_integration, name, description, image, category_id, available_quantity, created_at, updated_at) VALUES ('000000009', 'Engine', 'Change', 'https://via.placeholder.com/150', '1', '10', '2023-01-01 00:00:00', '2023-01-01 00:00:00');
INSERT INTO items (code_integration, name, description, image, category_id, available_quantity, created_at, updated_at) VALUES ('000000010', 'Transmission', 'Change', 'https://via.placeholder.com/150', '1', '10', '2023-01-01 00:00:00', '2023-01-01 00:00:00');
INSERT INTO items (code_integration, name, description, image, category_id, available_quantity, created_at, updated_at) VALUES ('000000011', 'Battery', 'Change', 'https://via.placeholder.com/150', '1', '10', '2023-01-01 00:00:00', '2023-01-01 00:00:00');
INSERT INTO items (code_integration, name, description, image, category_id, available_quantity, created_at, updated_at) VALUES ('000000012', 'Oil', 'Change', 'https://via.placeholder.com/150', '1', '10', '2023-01-01 00:00:00', '2023-01-01 00:00:00');
