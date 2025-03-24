-- ===============================================
-- Configuration Schema: reference data
-- ===============================================
CREATE SCHEMA IF NOT EXISTS conf;

CREATE TABLE conf.language (
    code VARCHAR(10) PRIMARY KEY,
    name VARCHAR(50)
);

-- Comments for conf.language
COMMENT ON TABLE conf.language IS 'Table of available languages for the application';
COMMENT ON COLUMN conf.language.code IS 'Language code (e.g., en, es)';
COMMENT ON COLUMN conf.language.name IS 'Full name of the language';

INSERT INTO conf.language (code, name) VALUES
    ('en', 'English'),
    ('es', 'Español'),
    ('zh', '中文 (Chinese)');
    
-- ===============================================
-- Main Schema: core application data
-- ===============================================
CREATE SCHEMA IF NOT EXISTS public;

CREATE TABLE public."group" (
    group_id UUID PRIMARY KEY DEFAULT gen_random_uuid() UNIQUE NOT NULL,
    group_name VARCHAR(50),
    description TEXT,
    photo_url TEXT CHECK (photo_url ~* '^https?://.+') DEFAULT 'https://groupphoto.png',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
    created_by UUID,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by UUID
);

-- Comments for public.group
COMMENT ON TABLE public."group" IS 'Table of groups in the application';
COMMENT ON COLUMN public."group".group_id IS 'Unique identifier for the group';
COMMENT ON COLUMN public."group".group_name IS 'Name of the group';
COMMENT ON COLUMN public."group".description IS 'Description of the group';
COMMENT ON COLUMN public."group".photo_url IS 'URL of the representative photo for the group';

CREATE TABLE public.movement (
    movement_id INT PRIMARY KEY,
    group_id UUID,
    amount DECIMAL(10, 2),
    created_at TIMESTAMP,
    created_by INT,
    CONSTRAINT fk_movement_group FOREIGN KEY (group_id) REFERENCES public."group"(group_id)
);

-- Comments for public.movement
COMMENT ON TABLE public.movement IS 'Table of movements associated with a group';
COMMENT ON COLUMN public.movement.movement_id IS 'Unique identifier for the movement';
COMMENT ON COLUMN public.movement.group_id IS 'Identifier of the group to which the movement belongs';
COMMENT ON COLUMN public.movement.amount IS 'Amount of the movement';

CREATE TABLE public.movement_field (
    movement_field_id INT PRIMARY KEY,
    group_id UUID,
    movement_id INT,
    name VARCHAR(50),
    type VARCHAR(50),
    required BOOLEAN,
    value VARCHAR(255),
    CONSTRAINT fk_movement_field_group FOREIGN KEY (group_id) REFERENCES public."group"(group_id),
    CONSTRAINT fk_movement_field_movement FOREIGN KEY (movement_id) REFERENCES public.movement(movement_id)
);

-- Comments for public.movement_field
COMMENT ON TABLE public.movement_field IS 'Table of fields associated with a movement';
COMMENT ON COLUMN public.movement_field.movement_field_id IS 'Unique identifier for the movement field';
COMMENT ON COLUMN public.movement_field.group_id IS 'Identifier of the group associated with the field';
COMMENT ON COLUMN public.movement_field.movement_id IS 'Identifier of the movement associated with the field';
COMMENT ON COLUMN public.movement_field.name IS 'Name of the field';
COMMENT ON COLUMN public.movement_field.type IS 'Data type or input type of the field';
COMMENT ON COLUMN public.movement_field.required IS 'Indicates if the field is mandatory';
COMMENT ON COLUMN public.movement_field.value IS 'Value assigned to the field';

CREATE TABLE public.movement_field_options (
    movement_field_options_id INT PRIMARY KEY,
    movement_field_id INT,
    value VARCHAR(255),
    CONSTRAINT fk_movement_field_options FOREIGN KEY (movement_field_id) REFERENCES public.movement_field(movement_field_id)
);

-- Comments for public.movement_field_options
COMMENT ON TABLE public.movement_field_options IS 'Table of options available for a movement field';
COMMENT ON COLUMN public.movement_field_options.movement_field_options_id IS 'Unique identifier for the option';
COMMENT ON COLUMN public.movement_field_options.movement_field_id IS 'Identifier of the associated movement field';
COMMENT ON COLUMN public.movement_field_options.value IS 'Value of the option';

-- ===============================================
-- Authorization Schema: users, roles
-- ===============================================
CREATE SCHEMA IF NOT EXISTS auth;

CREATE TABLE auth."user" (
    user_id UUID DEFAULT gen_random_uuid() UNIQUE NOT NULL,
    user_name VARCHAR(50) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    photo_url TEXT CHECK (photo_url ~* '^https?://.+') DEFAULT 'https://userphoto.png',
    language_code VARCHAR(10) DEFAULT 'es',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_auth_user_language FOREIGN KEY (language_code) REFERENCES conf.language(code)
);

-- Comments for auth.user
COMMENT ON TABLE auth."user" IS 'Table of application users';
COMMENT ON COLUMN auth."user".user_id IS 'Unique identifier for the user';
COMMENT ON COLUMN auth."user".user_name IS 'Name of the user';
COMMENT ON COLUMN auth."user".photo_url IS 'URL of the user''s photo';
COMMENT ON COLUMN auth."user".language_code IS 'Identifier of the user''s preferred language';

CREATE TABLE auth.role (
    role_id VARCHAR(1) PRIMARY KEY,
    role_name VARCHAR(50) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Comments for auth.role
COMMENT ON TABLE auth.role IS 'Table of defined roles in the application';
COMMENT ON COLUMN auth.role.role_id IS 'Unique identifier for the role';
COMMENT ON COLUMN auth.role.role_name IS 'Name of the role';
COMMENT ON COLUMN auth.role.description IS 'Description of the role';

INSERT INTO auth.role (role_id, role_name, description)
VALUES
    ('V', 'VIEWER', 'User with read-only access'),
    ('E', 'EDITOR', 'User with editing capabilities'),
    ('A', 'ADMIN', 'User with full administrative privileges');


CREATE TABLE auth.user_role (
    user_id UUID,
    role_id VARCHAR(1),
    group_id UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by UUID,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by UUID,
    valid_until TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, role_id),
    CONSTRAINT fk_user_role_user FOREIGN KEY (user_id) REFERENCES auth."user"(user_id),
    CONSTRAINT fk_user_role_role FOREIGN KEY (role_id) REFERENCES auth.role(role_id),
    CONSTRAINT fk_user_role_group FOREIGN KEY (group_id) REFERENCES public."group"(group_id)
);

-- Comments for auth.user_role
COMMENT ON TABLE auth.user_role IS 'Table for assigning roles to users for specific groups';
COMMENT ON COLUMN auth.user_role.user_id IS 'Identifier of the user';
COMMENT ON COLUMN auth.user_role.role_id IS 'Identifier of the assigned role';
COMMENT ON COLUMN auth.user_role.group_id IS 'Identifier of the group in which the role is assigned';
COMMENT ON COLUMN auth.user_role.valid_until IS 'Date until the assignment is valid';