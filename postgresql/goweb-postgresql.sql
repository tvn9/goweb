CREATE TABLE "users" (
  "User_ID" bigserial PRIMARY KEY NOT NULL,
  "User_Name" text NOT NULL,
  "Pass_Word_Hash" text NOT NULL,
  "Name" text NOT NULL,
  "Config" jsonb NOT NULL,
  "Created_At" timestamp NOT NULL DEFAULT (now()),
  "Is_Enabled" boolean NOT NULL DEFAULT true
);

CREATE TABLE "images" (
  "Image_ID" bigserial NOT NULL,
  "Content_Type" text NOT NULL,
  "Image_Data" bytea NOT NULL
);

CREATE TABLE "workouts" (
  "Workout_ID" bigserial PRIMARY KEY NOT NULL,
  "Set_ID_1" bigserial NOT NULL,
  "User_ID_1" bigserial NOT NULL,
  "Exercise_ID" bigserial NOT NULL,
  "Start_Date" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "sets" (
  "Set_ID" bigserial PRIMARY KEY NOT NULL,
  "Exercise_ID" bigserial NOT NULL,
  "Weight" int NOT NULL
);

CREATE TABLE "exercises" (
  "Exercise_ID" bigserial PRIMARY KEY NOT NULL,
  "Exercise_Name" text NOT NULL
);

CREATE INDEX ON "users" ("User_Name");

CREATE INDEX ON "images" ("Image_ID");

CREATE INDEX ON "workouts" ("Workout_ID");

CREATE INDEX ON "workouts" ("User_ID_1");

CREATE INDEX ON "workouts" ("Set_ID_1");

CREATE INDEX ON "workouts" ("Exercise_ID");

CREATE INDEX ON "workouts" ("Workout_ID", "User_ID_1", "Set_ID_1", "Exercise_ID");

CREATE INDEX ON "sets" ("Set_ID");

CREATE INDEX ON "exercises" ("Exercise_Name");

ALTER TABLE "images" ADD FOREIGN KEY ("Image_ID") REFERENCES "users" ("User_ID");

ALTER TABLE "workouts" ADD FOREIGN KEY ("Set_ID_1") REFERENCES "sets" ("Set_ID");

ALTER TABLE "workouts" ADD FOREIGN KEY ("User_ID_1") REFERENCES "users" ("User_ID");

ALTER TABLE "workouts" ADD FOREIGN KEY ("Exercise_ID") REFERENCES "exercises" ("Exercise_ID");

ALTER TABLE "sets" ADD FOREIGN KEY ("Exercise_ID") REFERENCES "exercises" ("Exercise_ID");
