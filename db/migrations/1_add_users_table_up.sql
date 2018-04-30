CREATE TABLE users (
    id SERIAL,
    username text,
    email text,
    permitOneAllowed boolean,
    permitTwoAllowed boolean,
    hashedPassword boolean,
    roleId integer,
    PRIMARY KEY (id)
);
