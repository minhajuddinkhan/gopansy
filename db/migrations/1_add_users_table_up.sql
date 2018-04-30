CREATE TABLE users (
    user_id SERIAL,
    username text,
    email text,
    permit_one_allowed boolean,
    permit_two_allowed boolean,
    hashed_password boolean,
    role_id integer,
    PRIMARY KEY (user_id)
);
