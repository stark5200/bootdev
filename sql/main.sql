ALTER TABLE people
RENAME TO users;

ALTER TABLE users
RENAME COLUMN tag TO username;

ALTER TABLE users
ADD COLUMN password TEXT;

-- TEST SUITE, DON'T TOUCH BELOW THIS LINE --

PRAGMA TABLE_INFO('users');

INSERT INTO users(id, name, age, country_code, username, password, is_admin)
VALUES (1, 'David', 34, 'US', 'DavidDev', 'insertPractice', false);

INSERT INTO users(id, name, age, country_code, username, password, is_admin)
VALUES (2, 'Samantha', 29, 'BR', 'Sammy93', 'addingRecords!', false);


-- TEST SUITE, DON'T TOUCH BELOW THIS LINE --

SELECT * FROM users;

