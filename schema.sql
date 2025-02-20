CREATE TABLE "rooms" (
    "key" VARCHAR PRIMARY KEY,
    "is_active" BOOLEAN DEFAULT true NOT NULL
);

CREATE TABLE "users" (
    "id" serial PRIMARY KEY,
    "username" varchar NOT NULL,
    "room_key" varchar REFERENCES rooms(key) ON DELETE CASCADE,
    "hp" integer NOT NULL
);

CREATE TABLE "characters" (
    "id" serial PRIMARY KEY,
    "name" varchar NOT NULL,
    "description" varchar NOT NULL
);

CREATE TABLE "cards" (
    "id" serial PRIMARY KEY,
    "description" varchar NOT NULL,
    "damage" integer NOT NULL,
    "power" integer NOT NULL
);
