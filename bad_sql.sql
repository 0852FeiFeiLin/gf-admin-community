-- ----------------------------
-- Sequence structure for company_agent_fd_recharge_records_record_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."company_agent_fd_recharge_records_record_id_seq";
CREATE SEQUENCE "public"."company_agent_fd_recharge_records_record_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;
ALTER SEQUENCE "public"."company_agent_fd_recharge_records_record_id_seq" OWNER TO "ltdb";

-- ----------------------------
-- Sequence structure for sys_area_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."sys_area_id_seq";
CREATE SEQUENCE "public"."sys_area_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."sys_area_id_seq" OWNER TO "ltdb";

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."company_agent_fd_recharge_records_record_id_seq"
    OWNED BY "public"."company_agent_fd_recharge"."id";
SELECT setval('"public"."company_agent_fd_recharge_records_record_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."sys_area_id_seq"
    OWNED BY "public"."sys_area"."id";
SELECT setval('"public"."sys_area_id_seq"', 1, false);
