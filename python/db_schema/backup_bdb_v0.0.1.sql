PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE "people" 
( 
            "playerid"     VARCHAR(255) PRIMARY KEY, 
            "birthyear"    INTEGER, 
            "birthmonth"   INTEGER, 
            "birthday"     INTEGER, 
            "birthcountry" VARCHAR(255), 
            "birthstate"   VARCHAR(255), 
            "birthcity"    VARCHAR(255), 
            "deathyear"    INTEGER, 
            "deathmonth"   INTEGER, 
            "deathday"     INTEGER, 
            "deathcountry" VARCHAR(255), 
            "deathstate"   VARCHAR(255), 
            "eathcity"     VARCHAR(255), 
            "namefirst"    VARCHAR(255), 
            "namelast"     VARCHAR(255), 
            "namegiven"    VARCHAR(255), 
            "weight"       INTEGER, 
            "height"       INTEGER, 
            "bats"         VARCHAR(255), 
            "throws"       VARCHAR(255), 
            "debut"        VARCHAR(255), 
            "finalgame"    VARCHAR(255), 
            "retroid"      VARCHAR(255), 
            "bbrefid"      VARCHAR(255) 
);

CREATE TABLE "awards_managers" 
( 
            "playerid_id" VARCHAR(255) NOT NULL, 
            "awardid"     VARCHAR(255), 
            "yearid"      INTEGER, 
            "lgid"        VARCHAR(255), 
            "tie"         VARCHAR(255), 
            "notes"       VARCHAR(255), 
            PRIMARY KEY ("yearid", "playerid_id", "lgid", "tie"), 
            FOREIGN KEY ("playerid_id") REFERENCES "master" ("playerid") 
);

CREATE TABLE "fielding_of" 
( 
            "playerid_id" VARCHAR(255) NOT NULL, 
            "yearid"      INTEGER, 
            "stint"       INTEGER, 
            "glf"         INTEGER, 
            "gcf"         INTEGER, 
            "grf"         INTEGER, 
            PRIMARY KEY ("playerid_id", "yearid", "stint"), 
            FOREIGN KEY ("playerid_id") REFERENCES "master" ("playerid") 
);

CREATE TABLE "hall_of_fame" 
( 
            "playerid_id" VARCHAR(255) NOT NULL, 
            "yearid"      INTEGER, 
            "votedby"     VARCHAR(255), 
            "ballots"     INTEGER, 
            "needed"      INTEGER, 
            "votes"       INTEGER, 
            "inducted"    VARCHAR(255), 
            "category"    VARCHAR(255), 
            "needed_note" VARCHAR(255), 
            PRIMARY KEY ("playerid_id", "category"), 
            FOREIGN KEY ("playerid_id") REFERENCES "master" ("playerid") 
);

CREATE TABLE "managershalf" 
( 
            "playerid_id" VARCHAR(255) NOT NULL, 
            "yearid"      INTEGER, 
            "teamid"      VARCHAR(255) NOT NULL, 
            "lgid"        VARCHAR(255), 
            "inseason"    INTEGER, 
            "half"        INTEGER, 
            "g"           INTEGER, 
            "w"           INTEGER, 
            "l"           INTEGER, 
            "rank"        INTEGER, 
            PRIMARY KEY ("playerid_id", "teamid", "lgid", "yearid", "half"), 
            FOREIGN KEY ("playerid_id") REFERENCES "master" ("playerid") 
);

CREATE TABLE "fielding" 
( 
            "playerid_id" VARCHAR(255) NOT NULL, 
            "yearid"      INTEGER, 
            "stint"       INTEGER, 
            "teamid"      VARCHAR(255), 
            "lgid"        VARCHAR(255), 
            "pos"         VARCHAR(255), 
            "a"           INTEGER, 
            "gs"          INTEGER, 
            "g"           INTEGER, 
            "pb"          INTEGER, 
            "innouts"     INTEGER, 
            "zr"          INTEGER, 
            "po"          INTEGER, 
            "wp"          INTEGER, 
            "cs"          INTEGER, 
            "e"           INTEGER, 
            "dp"          INTEGER, 
            "sb"          INTEGER, 
            PRIMARY KEY ("playerid_id", "yearid", "stint"), 
            FOREIGN KEY ("playerid_id") REFERENCES "master" ("playerid") 
);

CREATE TABLE "batting" 
( 
            "playerid_id" VARCHAR(255) NOT NULL, 
            "yearid"      INTEGER, 
            "stint"       INTEGER, 
            "teamid"      VARCHAR(255), 
            "lgid"        VARCHAR(255), 
            "g"           INTEGER, 
            "ab"          INTEGER, 
            "r"           INTEGER, 
            "h"           INTEGER, 
            "2b"          INTEGER, 
            "3b"          INTEGER, 
            "hr"          INTEGER, 
            "rbi"         INTEGER, 
            "sb"          INTEGER, 
            "cs"          INTEGER, 
            "bb"          INTEGER, 
            "so"          INTEGER, 
            "ibb"         INTEGER, 
            "hbp"         INTEGER, 
            "sh"          INTEGER, 
            "sf"          INTEGER, 
            "gidp"        INTEGER, 
            PRIMARY KEY ("playerid_id", "yearid", "stint"), 
            FOREIGN KEY ("playerid_id") REFERENCES "master" ("playerid") 
);

CREATE TABLE "appearances" 
( 
            "yearid"      INTEGER, 
            "teamid"      VARCHAR(255), 
            "lgid"        VARCHAR(255), 
            "playerid_id" VARCHAR(255) NOT NULL, 
            "g_all"       INTEGER, 
            "gs"          INTEGER, 
            "g_batting"   INTEGER, 
            "g_defense"   INTEGER, 
            "g_p"         INTEGER, 
            "g_c"         INTEGER, 
            "g_1b"        INTEGER, 
            "g_2b"        INTEGER, 
            "g_3b"        INTEGER, 
            "g_ss"        INTEGER, 
            "g_lf"        INTEGER, 
            "g_cf"        INTEGER, 
            "g_rf"        INTEGER, 
            "g_of"        INTEGER, 
            "g_dh"        INTEGER, 
            "g_ph"        INTEGER, 
            "g_pr"        INTEGER, 
            PRIMARY KEY ("yearid", "teamid", "playerid_id"), 
            FOREIGN KEY ("playerid_id") REFERENCES "master" ("playerid") 
);

CREATE TABLE "managers" 
( 
            "playerid_id" VARCHAR(255) NOT NULL, 
            "yearid"      INTEGER, 
            "teamid"      VARCHAR(255) NOT NULL, 
            "lgid"        VARCHAR(255) NOT NULL, 
            "inseason"    INTEGER, 
            "g"           INTEGER, 
            "w"           INTEGER, 
            "l"           INTEGER, 
            "rank"        INTEGER, 
            "plyrmgr"     VARCHAR(255), 
            PRIMARY KEY ("playerid_id", "lgid", "lgid", "yearid"), 
            FOREIGN KEY ("playerid_id") REFERENCES "master" ("playerid") 
);

CREATE TABLE "awards_players" 
( 
            "playerid_id" VARCHAR(255) NOT NULL, 
            "awardid"     VARCHAR(255), 
            "yearid"      INTEGER, 
            "lgid"        VARCHAR(255), 
            "tie"         VARCHAR(255), 
            "notes"       VARCHAR(255), 
            PRIMARY KEY ("playerid_id", "awardid", "yearid"), 
            FOREIGN KEY ("playerid_id") REFERENCES "master" ("playerid") 
);

CREATE TABLE "fielding_post" 
( 
            "playerid_id" VARCHAR(255) NOT NULL, 
            "yearid"      INTEGER, 
            "teamid"      VARCHAR(255), 
            "lgid"        VARCHAR(255), 
            "round"       VARCHAR(255), 
            "pos"         VARCHAR(255), 
            "g"           INTEGER, 
            "gs"          INTEGER, 
            "innouts"     INTEGER, 
            "po"          INTEGER, 
            "a"           INTEGER, 
            "e"           INTEGER, 
            "dp"          INTEGER, 
            "tp"          INTEGER, 
            "pb"          INTEGER, 
            "sb"          INTEGER, 
            "cs"          INTEGER, 
            PRIMARY KEY ("playerid_id", "yearid", "round", "pos"), 
            FOREIGN KEY ("playerid_id") REFERENCES "master" ("playerid") 
);

CREATE TABLE "awards_share_players" 
( 
            "awardid"     VARCHAR(255), 
            "yearid"      INTEGER, 
            "lgid"        VARCHAR(255), 
            "playerid_id" VARCHAR(255) NOT NULL, 
            "pointswon"   INTEGER, 
            "pointsmax"   INTEGER, 
            "votesfirst"  INTEGER, 
            PRIMARY KEY ("awardid", "yearid", "lgid", "playerid_id"), 
            FOREIGN KEY ("playerid_id") REFERENCES "master" ("playerid") 
);

CREATE TABLE "awards_share_managers" 
( 
            "awardid"     VARCHAR(255), 
            "yearid"      INTEGER, 
            "lgid"        VARCHAR(255), 
            "playerid_id" VARCHAR(255) NOT NULL, 
            "pointswon"   INTEGER, 
            "pointsmax"   INTEGER, 
            "votesfirst"  INTEGER, 
            PRIMARY KEY ("awardid", "yearid", "lgid", "playerid_id"), 
            FOREIGN KEY ("playerid_id") REFERENCES "master" ("playerid") 
);

CREATE TABLE "batting_post" 
( 
            "yearid"      INTEGER, 
            "round"       VARCHAR(255), 
            "playerid_id" VARCHAR(255) NOT NULL, 
            "teamid"      VARCHAR(255) NOT NULL, 
            "lgid"        VARCHAR(255), 
            "g"           INTEGER, 
            "ab"          INTEGER, 
            "r"           INTEGER, 
            "h"           INTEGER, 
            "2b"          INTEGER, 
            "3b"          INTEGER, 
            "hr"          INTEGER, 
            "rbi"         INTEGER, 
            "sb"          INTEGER, 
            "cs"          INTEGER, 
            "bb"          INTEGER, 
            "so"          INTEGER, 
            "ibb"         INTEGER, 
            "hbp"         INTEGER, 
            "sh"          INTEGER, 
            "sf"          INTEGER, 
            "gidp"        INTEGER, 
            PRIMARY KEY ("yearid", "round", "playerid_id", "teamid"), 
            FOREIGN KEY ("playerid_id") REFERENCES "master" ("playerid") 
);

CREATE TABLE "allstar_full" 
( 
            "playerid_id" VARCHAR(255) NOT NULL, 
            "yearid"      INTEGER, 
            "gamenum"     INTEGER, 
            "gameid"      VARCHAR(255), 
            "teamid"      VARCHAR(255), 
            "lgid"        VARCHAR(255), 
            "gp"          INTEGER, 
            "startingpos" INTEGER, 
            PRIMARY KEY ("playerid_id", "yearid", "gamenum"), 
            FOREIGN KEY ("playerid_id") REFERENCES "master" ("playerid") 
);

CREATE TABLE "parks" 
( 
            "parkid"  VARCHAR(255) PRIMARY KEY, 
            "name"    VARCHAR(255), 
            "alias"   VARCHAR(255), 
            "city"    VARCHAR(255), 
            "state"   VARCHAR(255), 
            "country" VARCHAR(255) 
);

CREATE TABLE "home_games" 
( 
            "yearid"     INTEGER, 
            "parkid_id"  VARCHAR(255) NOT NULL, 
            "teamid"     VARCHAR(255) NOT NULL, 
            "lgid"       VARCHAR(255) NOT NULL, 
            "attendance" INTEGER, 
            "games"      INTEGER, 
            "first_game" VARCHAR(255), 
            "last_game"  VARCHAR(255), 
            "openings"   INTEGER, 
            PRIMARY KEY ("yearid", "parkid_id", "teamid"), 
            FOREIGN KEY ("parkid_id") REFERENCES "parks" ("parkid") 
);

CREATE TABLE "pitching" 
( 
            "playerid_id" VARCHAR(255) NOT NULL, 
            "yearid"      INTEGER, 
            "stint"       INTEGER, 
            "teamid"      VARCHAR(255) NOT NULL, 
            "lgid"        VARCHAR(255), 
            "w"           INTEGER, 
            "l"           INTEGER, 
            "g"           INTEGER, 
            "gs"          INTEGER, 
            "cg"          INTEGER, 
            "sho"         INTEGER, 
            "sv"          INTEGER, 
            "ipouts"      INTEGER, 
            "h"           INTEGER, 
            "er"          INTEGER, 
            "hr"          INTEGER, 
            "bb"          INTEGER, 
            "so"          INTEGER, 
            "baopp"       REAL, 
            "era"         REAL, 
            "ibb"         INTEGER, 
            "wp"          INTEGER, 
            "hbp"         INTEGER, 
            "bk"          INTEGER, 
            "bfp"         INTEGER, 
            "gf"          INTEGER, 
            "r"           INTEGER, 
            "sh"          INTEGER, 
            "sf"          INTEGER, 
            "gidp"        INTEGER, 
            PRIMARY KEY ("playerid_id", "yearid", "stint", "teamid"), 
            FOREIGN KEY ("playerid_id") REFERENCES "master" ("playerid") 
);CREATE TABLE "pitching_post" 
( 
            "playerid_id" VARCHAR(255) NOT NULL, 
            "yearid"      INTEGER, 
            "round"       VARCHAR(255), 
            "teamid"      VARCHAR(255), 
            "lgid"        VARCHAR(255), 
            "w"           INTEGER, 
            "l"           INTEGER, 
            "g"           INTEGER, 
            "gs"          INTEGER, 
            "cg"          INTEGER, 
            "sho"         INTEGER, 
            "sv"          INTEGER, 
            "ipouts"      INTEGER, 
            "h"           INTEGER, 
            "er"          INTEGER, 
            "hr"          INTEGER, 
            "bb"          INTEGER, 
            "so"          INTEGER, 
            "baopp"       REAL, 
            "era"         REAL, 
            "ibb"         INTEGER, 
            "wp"          INTEGER, 
            "hbp"         INTEGER, 
            "bk"          INTEGER, 
            "bfp"         INTEGER, 
            "gf"          INTEGER, 
            "r"           INTEGER, 
            "sh"          INTEGER, 
            "sf"          INTEGER, 
            "gidp"        INTEGER, 
            PRIMARY KEY ("playerid_id", "yearid", "round"), 
            FOREIGN KEY ("playerid_id") REFERENCES "master" ("playerid") 
);

CREATE TABLE "salaries" 
( 
            "yearid"      INTEGER, 
            "teamid"      VARCHAR(255), 
            "lgid"        VARCHAR(255), 
            "playerid_id" VARCHAR(255) NOT NULL, 
            "salary"      INTEGER, 
            PRIMARY KEY ("yearid", "teamid", "playerid_id"), 
            FOREIGN KEY ("playerid_id") REFERENCES "master" ("playerid") 
);

CREATE TABLE "schools" 
( 
            "schoolid"  VARCHAR(255) PRIMARY KEY, 
            "name_full" VARCHAR(255), 
            "city"      VARCHAR(255), 
            "state"     VARCHAR(255), 
            "country"   VARCHAR(255) 
);

CREATE TABLE "college_playing" 
( 
            "playerid_id" VARCHAR(255) NOT NULL, 
            "schoolid_id" VARCHAR(255) NOT NULL, 
            "yearid"      INTEGER, 
            PRIMARY KEY ("playerid_id", "schoolid_id", "yearid"), 
            FOREIGN KEY ("playerid_id") REFERENCES "master" ("playerid"), 
            FOREIGN KEY ("schoolid_id") REFERENCES "schools" ("schoolid") 
);

CREATE TABLE "series_post" 
( 
            "yearid"       INTEGER, 
            "round"        VARCHAR(255), 
            "teamidwinner" VARCHAR(255), 
            "lgidwinner"   VARCHAR(255), 
            "teamidloser"  VARCHAR(255), 
            "lgidloser"    VARCHAR(255), 
            "wins"         INTEGER, 
            "losses"       INTEGER, 
            "ties"         INTEGER, 
            PRIMARY KEY ("yearid", "round") 
);

CREATE TABLE "teams_franchises" 
( 
            "franchid"   VARCHAR(255) PRIMARY KEY, 
            "franchname" VARCHAR(255), 
            "active"     VARCHAR(255), 
            "naassoc"    VARCHAR(255) 
);

CREATE TABLE "teams" 
( 
            "yearid"         INTEGER, 
            "lgid"           VARCHAR(255), 
            "teamid"         VARCHAR(255), 
            "franchid_id"    VARCHAR(255) NOT NULL, 
            "divid"          VARCHAR(255), 
            "rank"           INTEGER, 
            "g"              INTEGER, 
            "ghome"          INTEGER, 
            "w"              INTEGER, 
            "l"              INTEGER, 
            "divwin"         VARCHAR(255), 
            "wcwin"          VARCHAR(255), 
            "lgwin"          VARCHAR(255), 
            "wswin"          VARCHAR(255), 
            "r"              INTEGER, 
            "ab"             INTEGER, 
            "h"              INTEGER, 
            "2b"             INTEGER, 
            "3b"             INTEGER, 
            "hr"             INTEGER, 
            "bb"             INTEGER, 
            "so"             INTEGER, 
            "sb"             INTEGER, 
            "cs"             INTEGER, 
            "hbp"            INTEGER, 
            "sf"             INTEGER, 
            "ra"             INTEGER, 
            "er"             INTEGER, 
            "era"            REAL, 
            "cg"             INTEGER, 
            "sho"            INTEGER, 
            "sv"             INTEGER, 
            "ipouts"         INTEGER, 
            "ha"             INTEGER, 
            "hra"            INTEGER, 
            "bba"            INTEGER, 
            "soa"            INTEGER, 
            "e"              INTEGER, 
            "dp"             INTEGER, 
            "fp"             REAL, 
            "name"           VARCHAR(255), 
            "park_id"        VARCHAR(255) NOT NULL, 
            "attendance"     INTEGER, 
            "bpf"            INTEGER, 
            "ppf"            INTEGER, 
            "teamidbr"       VARCHAR(255), 
            "teamidlahman45" VARCHAR(255), 
            "teamidretro"    VARCHAR(255), 
            PRIMARY KEY ("yearid", "lgid", "teamid"), 
            FOREIGN KEY ("franchid_id") REFERENCES "teams_franchises" ("franchid"),
            FOREIGN KEY ("park_id") REFERENCES "parks" ("name") 
);

CREATE TABLE "teams_half" 
( 
            "yearid" INTEGER, 
            "lgid"   VARCHAR(255), 
            "teamid" VARCHAR(255), 
            "half"   INTEGER, 
            "divid"  VARCHAR(255), 
            "divwin" VARCHAR(255), 
            "rank"   INTEGER, 
            "g"      INTEGER, 
            "w"      INTEGER, 
            "l"      INTEGER, 
            PRIMARY KEY ("yearid", "teamid", "half") 
);

CREATE UNIQUE INDEX "master_retroID" 
ON "master" 
( 
        "retroid" 
);

CREATE UNIQUE INDEX "master_bbrefID" 
ON "master" 
( 
        "bbrefid" 
);

CREATE INDEX "awards_managers_playerID_id" 
ON "awards_managers" 
( 
            "playerid_id" 
);

CREATE INDEX "fielding_of_playerID_id" 
ON "fielding_of" 
( 
        "playerid_id" 
);

CREATE INDEX "hall_of_fame_playerID_id" 
ON "hall_of_fame" 
( 
        "playerid_id" 
);
             
CREATE INDEX "managershalf_playerID_id" 
ON "managershalf" 
( 
        "playerid_id" 
);

CREATE INDEX "fielding_playerID_id" 
ON "fielding" 
( 
        "playerid_id" 
);

CREATE INDEX "batting_playerID_id" 
ON "batting" 
( 
        "playerid_id" 
);

CREATE INDEX "appearances_playerID_id" 
ON "appearances" 
( 
        "playerid_id" 
);

CREATE INDEX "managers_playerID_id" 
ON "managers" 
( 
        "playerid_id" 
);

CREATE INDEX "awards_players_playerID_id" 
ON "awards_players" 
( 
        "playerid_id" 
);

CREATE INDEX "fielding_post_playerID_id" 
ON "fielding_post" 
( 
        "playerid_id" 
);

CREATE INDEX "awards_share_players_playerID_id" 
ON "awards_share_players" 
( 
        "playerid_id" 
);

CREATE INDEX "awards_share_managers_playerID_id" 
ON "awards_share_managers" 
( 
        "playerid_id" 
);

CREATE INDEX "batting_post_playerID_id" 
ON "batting_post" 
( 
        "playerid_id" 
);

CREATE INDEX "allstar_full_playerID_id" 
ON "allstar_full" 
( 
        "playerid_id" 
);

CREATE UNIQUE INDEX "parks_name" 
ON "parks" 
( 
        "name" 
);

CREATE INDEX "home_games_parkID_id" 
ON "home_games" 
( 
        "parkid_id" 
);

CREATE INDEX "pitching_playerID_id" 
ON "pitching" 
( 
        "playerid_id" 
);

CREATE INDEX "pitching_post_playerID_id" 
ON "pitching_post" 
( 
        "playerid_id" 
);

CREATE INDEX "salaries_playerID_id" 
ON "salaries" 
( 
        "playerid_id" 
);

CREATE UNIQUE INDEX "schools_name_full" 
ON "schools" 
( 
        "name_full" 
);

CREATE INDEX "college_playing_playerID_id" 
ON "college_playing" 
( 
        "playerid_id" 
);

CREATE INDEX "college_playing_schoolID_id" 
ON "college_playing" 
( 
            "schoolid_id" 
);

CREATE INDEX "teams_franchID_id" 
ON "teams" 
             ( 
                          "franchid_id" 
             );CREATE INDEX "teams_park_id" 
ON "teams" 
             ( 
                          "park_id" 
             );COMMIT;