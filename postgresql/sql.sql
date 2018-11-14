创建自增表
CREATE TABLE PUBLIC.student2 ( ID INT NOT NULL, NAME CHARACTER ( 100 ), subjects CHARACTER ( 1 ), CONSTRAINT student2_pkey PRIMARY KEY ( ID ) ) WITH ( OIDS = FALSE );
ALTER TABLE PUBLIC.student2 OWNER TO postgres;
COMMENT ON TABLE PUBLIC.student2 IS '这是一个学生信息表2';
CREATE SEQUENCE student2_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;
ALTER TABLE student2 ALTER COLUMN ID 
SET DEFAULT nextval( 'student2_id_seq' );

-- 这里的"test"专指postgre中的表空间(模式)，默认的表空间是"public"  
DROP SEQUENCE if EXISTS "test"."testseq_id_seq";  
CREATE SEQUENCE "test"."testseq_id_seq"  
 INCREMENT 1  
 MINVALUE 1  
 MAXVALUE 9223372036854775807  
 START 1  
 CACHE 1;  

DROP TABLE if EXISTS "test"."testtable";  
CREATE TABLE "test"."testtable" (  
"id" int4 DEFAULT nextval('testseq_id_seq'::regclass) NOT NULL, -- 表数据关联SEQUENCE，每次插入取nextval('testseq_id_seq')<pre name="code" class="sql"><pre name="code" class="sql">nextval('testseq_id_seq'  
"create_date" timestamp(6),  
"age" int4,  
"name" varchar(100),  
"grade" float4  
)  
WITH (OIDS=FALSE)  
; 
