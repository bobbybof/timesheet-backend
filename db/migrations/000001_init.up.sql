CREATE TABLE users (
    id SERIAL,
    name VARCHAR(255) NOT NULL,
    rate integer NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE projects (
    id SERIAL,
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE activities (
    id SERIAL,
    name VARCHAR(255) NOT NULL,
    date_start TIMESTAMP NOT NULL,
    date_end TIMESTAMP NOT NULL,
    project_id INT NOT NULL,
    user_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (project_id) REFERENCES projects(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);


