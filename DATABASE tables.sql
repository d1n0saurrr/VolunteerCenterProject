CREATE TABLE volunteers (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(30) NOT NULL,
    second_name VARCHAR(30) NOT NULL,
    patronymic VARCHAR(30) NOT NULL,
    birth_date date NOT NULL
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(30) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    is_admin BOOLEAN,
    id_volunteer INT REFERENCES volunteers(id)
);

CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description VARCHAR,
    location VARCHAR(50),
    start_date DATE,
    end_date DATE NOT NULL
);

CREATE TABLE vols_and_events (
    vol_id INT REFERENCES volunteers(id),
    event_id INT REFERENCES events(id)
);