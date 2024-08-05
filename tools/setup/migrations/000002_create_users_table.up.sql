CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL,
    UNIQUE(email)
);

INSERT INTO users (email, name, phone_number) VALUES
('test+01@test.com', 'Test User 01', '1234567890'),
('test+02@test.com', 'Test User 02', '1234567891'),
('test+03@test.com', 'Test User 03', '1234567892'),
('test+04@test.com', 'Test User 03', null);
