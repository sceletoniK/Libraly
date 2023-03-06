begin;
alter table sessions rename column sessionKey to refreshToken;
alter table sessions rename column deadline to expiresAt;
commit;