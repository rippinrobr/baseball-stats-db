--' sqlite_schema_2017_retro_schedule.sql

DROP TABLE IF EXISTS "schedules";

CREATE TABLE "schedules" (
  "season" INTEGER NOT NULL,
  "game_date_str" VARCHAR(8) NOT NULL,
  "game_number" INTEGER NOT NULL,
  "visitors_team" VARCHAR(3) NOT NULL,
  "home_team" VARCHAR(3) NOT NULL,
  "game_month" INTEGER NOT NULL,
  "game_month_day" INTEGER NOT NULL,
  "week_day" INTEGER NOT NULL,
  "visitors_league" VARCHAR(3) NOT NULL,
  "visitors_season_game_number" INTEGER NOT NULL,
  "home_league" VARCHAR(3) NOT NULL,
  "home_season_game_number" INTEGER NOT NULL,
  "game_type" VARCHAR(1) NOT NULL,
  "postponed_canceled" VARCHAR(128) NOT NULL DEFAULT '',
  "make_up_date" VARCHAR(128) NOT NULL DEFAULT '',

  PRIMARY KEY(season, game_date_str, game_number, visitors_team, home_team, visitors_season_game_number)
);