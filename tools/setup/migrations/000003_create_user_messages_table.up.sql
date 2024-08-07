CREATE TABLE messages (
    id INT AUTO_INCREMENT PRIMARY KEY,
    subject VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    sender_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (sender_id) REFERENCES users(id)
);

CREATE TABLE user_messages (
    user_id INT,
    message_id INT,
    is_read BOOLEAN DEFAULT FALSE,
    read_at TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (user_id, message_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (message_id) REFERENCES messages(id)
);

INSERT INTO messages (subject, content, sender_id) VALUES
    ('Hello', 'Hello, how are you?', 1),
    ('Goodbye', 'Goodbye, see you later!', 1);

INSERT INTO user_messages (user_id, message_id) VALUES
    (2, 1),
    (2, 2),
    (3, 1),
    (3, 2);
