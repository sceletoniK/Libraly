begin;
alter table bookrent alter column "deadline" drop not null;
alter table bookrent 
    DROP CONSTRAINT bookrent_instancebookid_fkey,
    ADD CONSTRAINT bookrent_instancebookid2_fkey FOREIGN KEY (instanceBookId) REFERENCES bookinstance (id);
commit;