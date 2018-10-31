创建自增表
CREATE TABLE PUBLIC.student2 ( ID INT NOT NULL, NAME CHARACTER ( 100 ), subjects CHARACTER ( 1 ), CONSTRAINT student2_pkey PRIMARY KEY ( ID ) ) WITH ( OIDS = FALSE );
ALTER TABLE PUBLIC.student2 OWNER TO postgres;
COMMENT ON TABLE PUBLIC.student2 IS '这是一个学生信息表2';
CREATE SEQUENCE student2_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;
ALTER TABLE student2 ALTER COLUMN ID 
SET DEFAULT nextval( 'student2_id_seq' );