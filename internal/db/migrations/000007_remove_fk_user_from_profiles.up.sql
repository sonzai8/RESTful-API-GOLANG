ALTER TABLE profiles
drop constraint fk_users;

alter table profiles
drop column user_id;