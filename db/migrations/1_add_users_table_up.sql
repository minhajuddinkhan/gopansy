CREATE TABLE users (
    id SERIAL,
    username text,
    email text,
    permitOneAllowed boolean,
    permitTwoAllowed boolean,
    hashedPassword text,
    roleId integer,
    PRIMARY KEY (id)
);
