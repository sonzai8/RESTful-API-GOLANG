ALTER TABLE USERS
    RENAME COLUMN email to user_email;

ALTER TABLE USERS
    ALTER column user_email set data type text;