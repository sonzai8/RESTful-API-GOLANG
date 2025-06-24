alter table profiles
    add column user_id int unique not null;

ALTER TABLE profiles
    add constraint fk_users foreign key (user_id) references users(user_id) on delete restrict  ;
