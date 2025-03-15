-- ===============================================
-- Configuration Schema: reference data
-- ===============================================
CREATE SCHEMA IF NOT EXISTS conf;

CREATE TABLE conf.language (
    language_id INT PRIMARY KEY,
    code VARCHAR(10),
    name VARCHAR(50)
);

-- Comments for conf.language
COMMENT ON TABLE conf.language IS 'Table of available languages for the application';
COMMENT ON COLUMN conf.language.language_id IS 'Unique identifier for the language';
COMMENT ON COLUMN conf.language.code IS 'Language code (e.g., en, es)';
COMMENT ON COLUMN conf.language.name IS 'Full name of the language';

-- ===============================================
-- Main Schema: core application data
-- ===============================================
CREATE SCHEMA IF NOT EXISTS public;

CREATE TABLE public."group" (
    group_id INT PRIMARY KEY,
    default_weighing DECIMAL(10, 2),
    name VARCHAR(50),
    description TEXT,
    photo_url TEXT CHECK (photo_url ~* '^https?://.+'),
    created_at TIMESTAMP,
    created_by INT,
    updated_at TIMESTAMP,
    updated_by INT
);

-- Comments for public.group
COMMENT ON TABLE public."group" IS 'Table of groups in the application';
COMMENT ON COLUMN public."group".group_id IS 'Unique identifier for the group';
COMMENT ON COLUMN public."group".default_weighing IS 'Default weighing value for the group';
COMMENT ON COLUMN public."group".name IS 'Name of the group';
COMMENT ON COLUMN public."group".description IS 'Description of the group';
COMMENT ON COLUMN public."group".photo_url IS 'URL of the representative photo for the group';

CREATE TABLE public.movement (
    movement_id INT PRIMARY KEY,
    group_id INT,
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
    group_id INT,
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
-- Authorization Schema: users, roles, permissions
-- ===============================================
CREATE SCHEMA IF NOT EXISTS auth;

CREATE TABLE auth."user" (
    user_id UUID DEFAULT gen_random_uuid(),
    user_name VARCHAR(50) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL, -- Campo de email con restricción de unicidad
    password TEXT NOT NULL, -- Almacena la contraseña (se recomienda almacenar un hash)
    photo_url TEXT CHECK (photo_url ~* '^https?://.+'),
    language_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by INT,
    updated_at TIMESTAMP,
    updated_by INT,
    CONSTRAINT fk_auth_user_language FOREIGN KEY (language_id) REFERENCES conf.language(language_id)
);

-- Comments for auth.user
COMMENT ON TABLE auth."user" IS 'Table of application users';
COMMENT ON COLUMN auth."user".user_id IS 'Unique identifier for the user';
COMMENT ON COLUMN auth."user".user_name IS 'Name of the user';
COMMENT ON COLUMN auth."user".photo_url IS 'URL of the user''s photo';
COMMENT ON COLUMN auth."user".language_id IS 'Identifier of the user''s preferred language';

CREATE TABLE auth.role (
    role_id INT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description TEXT,
    created_by INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    updated_by INT
);

-- Comments for auth.role
COMMENT ON TABLE auth.role IS 'Table of defined roles in the application';
COMMENT ON COLUMN auth.role.role_id IS 'Unique identifier for the role';
COMMENT ON COLUMN auth.role.name IS 'Name of the role';
COMMENT ON COLUMN auth.role.description IS 'Description of the role';

CREATE TABLE auth.permission (
    permission_id INT PRIMARY KEY,
    description TEXT,
    status VARCHAR(20)
);

-- Comments for auth.permission
COMMENT ON TABLE auth.permission IS 'Table of available permissions in the application';
COMMENT ON COLUMN auth.permission.permission_id IS 'Unique identifier for the permission';
COMMENT ON COLUMN auth.permission.description IS 'Description of the permission';
COMMENT ON COLUMN auth.permission.status IS 'Status of the permission (e.g., active, inactive)';

CREATE TABLE auth.user_role (
    user_id UUID,
    role_id INT,
    group_id INT,
    created_at TIMESTAMP,
    created_by INT,
    updated_at TIMESTAMP,
    updated_by INT,
    valid_until TIMESTAMP,
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

CREATE TABLE auth.role_permission (
    role_id INT,
    permission_id INT,
    created_at TIMESTAMP,
    created_by INT,
    updated_at TIMESTAMP,
    updated_by INT,
    PRIMARY KEY (role_id, permission_id),
    CONSTRAINT fk_role_permission_role FOREIGN KEY (role_id) REFERENCES auth.role(role_id),
    CONSTRAINT fk_role_permission_permission FOREIGN KEY (permission_id) REFERENCES auth.permission(permission_id)
);

-- Comments for auth.role_permission
COMMENT ON TABLE auth.role_permission IS 'Table that assigns permissions to roles';
COMMENT ON COLUMN auth.role_permission.role_id IS 'Identifier of the role';
COMMENT ON COLUMN auth.role_permission.permission_id IS 'Identifier of the assigned permission';