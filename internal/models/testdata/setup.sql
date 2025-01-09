CREATE TABLE snippets (
    id INTEGER NOT NULL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);


CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    hashed_password CHAR(60) NOT NULL,
    created DATETIME NOT NULL
);

CREATE UNIQUE INDEX email ON users(name);


INSERT INTO snippets (title, content, created, expires) VALUES (
    'An old silent pond',
    'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
    Current_timestamp,
    DATETIME(CURRENT_TIMESTAMP, '+365 day')
);

INSERT INTO users (name, email, hashed_password, created) VALUES (
    'samuel',
    'samuel@email.com',
    '$2a$12$NuTjWXm3KKntReFwyBVHyuf/to.HEwTy.eS206TNfkGfr6HzGJSWG',
    Current_timestamp
);