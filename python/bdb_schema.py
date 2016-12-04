from peewee import *
from playhouse.postgres_ext import *
import logging
import sys

BDB_DB = SqliteDatabase('db/bdb_v0.0.1.sqlite')

class BaseModel(Model):
    """BaseModel is used to connect to the database"""
    class Meta(object):
        """Where the connection is set"""
        database = BDB_DB

# Print the connection info here for sqlite
class Master(BaseModel):
    playerID = CharField(null='false',primary_key=True)
    birthYear = IntegerField(null='true')
    birthMonth = IntegerField(null='true')
    birthDay = IntegerField(null='true')
    birthCountry = CharField(null='true')
    birthState = CharField(null='true')
    birthCity = CharField(null='true')
    deathYear = IntegerField(null='true')
    deathMonth = IntegerField(null='true')
    deathDay = IntegerField(null='true')
    deathCountry = CharField(null='true')
    deathState = CharField(null='true')
    deathCity = CharField(null='true')
    nameFirst = CharField(null='false')
    nameLast = CharField(null='false')
    nameGiven = CharField(null='true')
    weight = IntegerField(null='true')
    height = IntegerField(null='true')
    bats = CharField(null='false')
    throws = CharField(null='false')
    debut = CharField(null='true')
    finalGame = CharField(null='true')
    retroID = CharField(null='false', unique=True)
    bbrefID = CharField(null='false', unique=True) 

class Teams_Franchises(BaseModel):
    franchID = CharField(null='false', primary_key=True)
    franchName = CharField(null='false')
    active = CharField(null='true')
    NAassoc = CharField(null='true')
    
class Parks(BaseModel):
    parkID = CharField(null='false', primary_key=True)
    name = CharField(null='false', unique=True)
    alias = CharField(null='true')
    city = CharField(null='false')
    state = CharField(null='false')
    country = CharField(null='false')
    

class Teams(BaseModel):
    yearID = IntegerField(null='false')
    lgID = CharField(null='false')
    teamID = CharField(null='true')
    franchID = ForeignKeyField(Teams_Franchises)
    divID = CharField(null='true')
    Rank = IntegerField(null='true')
    G = IntegerField(null='false', default=0)
    Ghome = IntegerField(null='false', default=0)
    W = IntegerField(null='false', default=0)
    L = IntegerField(null='false', default=0)
    DivWin = CharField(null='true')
    WCWin = CharField(null='true')
    LgWin = CharField(null='true')
    WSWin = CharField(null='true')
    R = IntegerField(null='false', default=0)
    AB = IntegerField(null='false', default=0)
    H = IntegerField(null='false', default=0)
    doubles = IntegerField(null='false', db_column="2B", default=0)
    triples = IntegerField(null='false', db_column="3B", default=0)
    HR = IntegerField(null='false', default=0)
    BB = IntegerField(null='false', default=0)
    SO = IntegerField(null='false', default=0)
    SB = IntegerField(null='false', default=0)
    CS = IntegerField(null='false', default=0)
    HBP = IntegerField(null='false', default=0)
    SF = IntegerField(null='false', default=0)
    RA = IntegerField(null='false', default=0)
    ER = IntegerField(null='false', default=0)
    ERA = FloatField(null='false', default=0.00)
    CG = IntegerField(null='false', default=0)
    SHO = IntegerField(null='false', default=0)
    SV = IntegerField(null='false', default=0)
    IPouts = IntegerField(null='false', default=0)
    HA = IntegerField(null='false', default=0)
    HRA = IntegerField(null='false', default=0)
    BBA = IntegerField(null='false', default=0)
    SOA = IntegerField(null='false', default=0)
    E = IntegerField(null='false', default=0)
    DP = IntegerField(null='false', default=0)
    FP = FloatField(null='true')
    name = CharField(null='false')
    park = ForeignKeyField(Parks, to_field='name')
    attendance = IntegerField(null='false', default=0)
    BPF = IntegerField(null='false', default=0)
    PPF = IntegerField(null='false', default=0)
    teamIDBR = CharField(null='false')
    teamIDlahman45 = CharField(null='false')
    teamIDretro = CharField(null='false')

    class Meta:
        primary_key = CompositeKey('yearID', 'lgID', 'teamID')

class Managers(BaseModel):
    playerID = ForeignKeyField(Master)
    yearID = IntegerField(null='false')
    teamID = CharField(null=False)
    lgID = CharField(null=False)
    inseason = IntegerField(null='false', default=0)
    G = IntegerField(null='false', default=0)
    W = IntegerField(null='false', default=0)
    L = IntegerField(null='false', default=0)
    rank = IntegerField(null='false', default=0)
    plyrMgr = CharField(null='true')

    class Meta:
        primary_key = CompositeKey('playerID', 'lgID', 'lgID', 'yearID')

class ManagersHalf(BaseModel):
    playerID = ForeignKeyField(Master)
    yearID = IntegerField(null='false')
    teamID = CharField(null=False)
    lgID = CharField(null='false')
    inseason = IntegerField(null='true')
    half = IntegerField(null='false')
    G = IntegerField(null='false', default=0)
    W = IntegerField(null='false', default=0)
    L = IntegerField(null='false', default=0)
    rank = IntegerField(null='true')
    
    class Meta:
        primary_key = CompositeKey('playerID', 'teamID', 'lgID', 'yearID', 'half')

class Schools(BaseModel):
    schoolID = CharField(null='false', primary_key=True)
    name_full = CharField(null='false', unique=True)
    city = CharField(null='false')
    state = CharField(null='false')
    country = CharField(null='false')
    
   
class Batting_Post(BaseModel):
    yearID = IntegerField(null='false')
    round = CharField(null='true')
    playerID = ForeignKeyField(Master)
    teamID = CharField(null=False)
    lgID = CharField(null='false')
    G = IntegerField(null='false', default=0)
    AB = IntegerField(null='false', default=0)
    R = IntegerField(null='false', default=0)
    H = IntegerField(null='false', default=0)
    doubles = IntegerField(null='false', db_column="2b", default=0)
    triples = IntegerField(null='false', db_column="3b", default=0)
    HR = IntegerField(null='false', default=0)
    RBI = IntegerField(null='false', default=0)
    SB = IntegerField(null='false', default=0)
    CS = IntegerField(null='false', default=0)
    BB = IntegerField(null='false', default=0)
    SO = IntegerField(null='false', default=0)
    IBB = IntegerField(null='false', default=0)
    HBP = IntegerField(null='false', default=0)
    SH = IntegerField(null='false', default=0)
    SF = IntegerField(null='false', default=0)
    GIDP = IntegerField(null='false', default=0)

    class Meta:
        primary_key = CompositeKey('yearID', 'round', 'playerID', 'teamID')
    
class College_Playing(BaseModel):
    playerID = ForeignKeyField(Master)
    schoolID = ForeignKeyField(Schools)
    yearID = IntegerField(null='false')

    class Meta:
        primary_key = CompositeKey('playerID', 'schoolID', 'yearID')

class Awards_Players(BaseModel):
    playerID = ForeignKeyField(Master)
    awardID = CharField(null='false')
    yearID = IntegerField(null='false')
    lgID = CharField(null='false')
    tie = CharField(null='true')
    notes = CharField(null='true')

    class Meta:
        primary_key = CompositeKey('playerID', 'awardID', 'yearID')
    
class Fielding_OF(BaseModel):
    playerID = ForeignKeyField(Master)
    yearID = IntegerField(null='false', default=0)
    stint = IntegerField(null='false', default=0)
    Glf = IntegerField(null='false', default=0)
    Gcf = IntegerField(null='false', default=0)
    Grf = IntegerField(null='false', default=0)

    class Meta:
        primary_key = CompositeKey('playerID', 'yearID', 'stint')
    
class Pitching(BaseModel):
    playerID = ForeignKeyField(Master)
    yearID = IntegerField(null='false', default=0)
    stint = IntegerField(null='false', default=0)
    teamID = CharField(null=False)
    lgID = CharField(null='false', default=0)
    W = IntegerField(null='false', default=0)
    L = IntegerField(null='false', default=0)
    G = IntegerField(null='false', default=0)
    GS = IntegerField(null='false', default=0)
    CG = IntegerField(null='false', default=0)
    SHO = IntegerField(null='false', default=0)
    SV = IntegerField(null='false', default=0)
    IPouts = IntegerField(null='false', default=0)
    H = IntegerField(null='false', default=0)
    ER = IntegerField(null='false', default=0)
    HR = IntegerField(null='false', default=0)
    BB = IntegerField(null='false', default=0)
    SO = IntegerField(null='false', default=0)
    BAOpp = FloatField(null='false', default=0)
    ERA = FloatField(null='false', default=0.00)
    IBB = IntegerField(null='false', default=0)
    WP = IntegerField(null='false', default=0)
    HBP = IntegerField(null='false', default=0)
    BK = IntegerField(null='false', default=0)
    BFP = IntegerField(null='false', default=0)
    GF = IntegerField(null='false', default=0)
    R = IntegerField(null='false', default=0)
    SH = IntegerField(null='false', default=0)
    SF = IntegerField(null='false', default=0)
    GIDP = IntegerField(null='false', default=0)
    
    class Meta:
        primary_key = CompositeKey('playerID', 'yearID', 'stint', 'teamID')

class Allstar_Full(BaseModel):
    playerID = ForeignKeyField(Master)
    yearID = IntegerField(null='false')
    gameNum = IntegerField(null='false')
    gameID = CharField(null='false')
    teamID = CharField(null='false')
    lgID = CharField(null='false')
    GP = IntegerField(null='false', default=0)
    startingPos = IntegerField(null='false', default=0)

    class Meta:
        primary_key = CompositeKey('playerID', 'yearID', 'gameNum')
    
class Home_Games(BaseModel):
    yearID = IntegerField(null='false')
    parkID = ForeignKeyField(Parks)
    teamID = CharField(null=False)
    lgID = CharField(null=False)
    attendance = IntegerField(null='true')
    games = IntegerField(null='true')
    first_game = CharField(null='true')
    last_game = CharField(null='true')
    openings = IntegerField(null='true')

    class Meta:
        primary_key = CompositeKey('yearID', 'parkID', 'teamID')
    
class Appearances(BaseModel):
    yearID = IntegerField(null='false')
    teamID = CharField(null='false')
    lgID = CharField(null='false')
    playerID = ForeignKeyField(Master)
    G_all = IntegerField(null='false', default=0)
    GS = IntegerField(null='false', default=0)
    G_batting = IntegerField(null='false', default=0)
    G_defense = IntegerField(null='false', default=0)
    G_p = IntegerField(null='false', default=0)
    G_c = IntegerField(null='false', default=0)
    G_1b = IntegerField(null='false', default=0)
    G_2b = IntegerField(null='false', default=0)
    G_3b = IntegerField(null='false', default=0)
    G_ss = IntegerField(null='false', default=0)
    G_lf = IntegerField(null='false', default=0)
    G_cf = IntegerField(null='false', default=0)
    G_rf = IntegerField(null='false', default=0)
    G_of = IntegerField(null='false', default=0)
    G_dh = IntegerField(null='false', default=0)
    G_ph = IntegerField(null='false', default=0)
    G_pr = IntegerField(null='false', default=0)

    class Meta:
        primary_key = CompositeKey('yearID', 'teamID', 'playerID')

class Hall_Of_Fame(BaseModel):
    playerID = ForeignKeyField(Master)
    yearID = IntegerField(null='true')
    votedBy = CharField(null='true')
    ballots = IntegerField(null='true')
    needed = IntegerField(null='true')
    votes = IntegerField(null='true')
    inducted = CharField(null='true')
    category = CharField(null='true')
    needed_note = CharField(null='true')

    class Meta:
        primary_key = CompositeKey('playerID', 'category')
    
class Fielding_Post(BaseModel):
    playerID = ForeignKeyField(Master)
    yearID = IntegerField(null='false')
    teamID = CharField(null='false')
    lgID = CharField(null='false')
    round = CharField(null='false')
    pos = CharField(null='false')
    G = IntegerField(null='false', default=0)
    GS = IntegerField(null='false', default=0)
    InnOuts = IntegerField(null='false', default=0)
    PO = IntegerField(null='false', default=0)
    A = IntegerField(null='false', default=0)
    E = IntegerField(null='false', default=0)
    DP = IntegerField(null='false', default=0)
    TP = IntegerField(null='false', default=0)
    PB = IntegerField(null='false', default=0)
    SB = IntegerField(null='false', default=0)
    CS = IntegerField(null='false', default=0)

    class Meta:
        primary_key = CompositeKey('playerID', 'yearID', 'round', 'pos')
        
    
class Series_Post(BaseModel):
    yearID = IntegerField(null='false')
    round = CharField(null='false')
    teamIDwinner = CharField(null='false')
    lgIDwinner = CharField(null='false')
    teamIDloser = CharField(null='false')
    lgIDloser = CharField(null='false')
    wins = IntegerField(null='false', default=0)
    losses = IntegerField(null='false', default=0)
    ties = IntegerField(null='false', default=0)
    
    class Meta:
        primary_key = CompositeKey('yearID', 'round')

class Awards_Managers(BaseModel):
    playerID = ForeignKeyField(Master)
    awardID = CharField(null='false')
    yearID = IntegerField(null='false')
    lgID = CharField(null='false')
    tie = CharField(null='false', default='N')
    notes = CharField(null='true')
    
    class Meta:
        primary_key = CompositeKey('yearID', 'playerID', 'lgID', 'tie')

class Awards_Share_Managers(BaseModel):
    awardID = CharField(null='true')
    yearID = IntegerField(null='true')
    lgID = CharField(null='true')
    playerID = ForeignKeyField(Master)
    pointsWon = IntegerField(null='true')
    pointsMax = IntegerField(null='true')
    votesFirst = IntegerField(null='true')
    
    class Meta:
        primary_key = CompositeKey('awardID', 'yearID', 'lgID', 'playerID')

class Salaries(BaseModel):
    yearID = IntegerField(null='true')
    teamID = CharField(null='true')
    lgID = CharField(null='true')
    playerID = ForeignKeyField(Master)
    salary = IntegerField(null='true')

    class Meta:
        primary_key = CompositeKey('yearID', 'teamID', 'playerID')
    
class Pitching_Post(BaseModel):
    playerID = ForeignKeyField(Master)
    yearID = IntegerField(null='true')
    round = CharField(null='true')
    teamID = CharField(null='true')
    lgID = CharField(null='true')
    W = IntegerField(null='false', default=0)
    L = IntegerField(null='false', default=0)
    G = IntegerField(null='false', default=0)
    GS = IntegerField(null='false', default=0)
    CG = IntegerField(null='false', default=0)
    SHO = IntegerField(null='false', default=0)
    SV = IntegerField(null='false', default=0)
    IPouts = IntegerField(null='false', default=0)
    H = IntegerField(null='false', default=0)
    ER = IntegerField(null='false', default=0)
    HR = IntegerField(null='false', default=0)
    BB = IntegerField(null='false', default=0)
    SO = IntegerField(null='false', default=0)
    BAOpp = FloatField(null='false', default=0.000)
    ERA = FloatField(null='false', default=0.00)
    IBB = IntegerField(null='false', default=0)
    WP = IntegerField(null='false', default=0)
    HBP = IntegerField(null='false', default=0)
    BK = IntegerField(null='false', default=0)
    BFP = IntegerField(null='false', default=0)
    GF = IntegerField(null='false', default=0)
    R = IntegerField(null='false', default=0)
    SH = IntegerField(null='false', default=0)
    SF = IntegerField(null='false', default=0)
    GIDP = IntegerField(null='false', default=0)

    class Meta:
        primary_key = CompositeKey('playerID', 'yearID', 'round')
    
class Awards_Share_Players(BaseModel):
    awardID = CharField(null='false')
    yearID = IntegerField(null='false')
    lgID = CharField(null='false')
    playerID = ForeignKeyField(Master)
    pointsWon = IntegerField(null='fails', default=0)
    pointsMax = IntegerField(null='fails', default=0)
    votesFirst = IntegerField(null='fails', default=0)

    class Meta:
        primary_key = CompositeKey('awardID', 'yearID', 'lgID', 'playerID')
    
class Fielding(BaseModel):
    playerID = ForeignKeyField(Master)
    yearID = IntegerField(null='false', default=0)
    stint = IntegerField(null='false', default=0)
    teamID = CharField(null='true')
    lgID = CharField(null='true')
    Pos = CharField(null='true')
    A = IntegerField(null='false', default=0)
    GS = IntegerField(null='false', default=0)
    G = IntegerField(null='false', default=0)
    PB = IntegerField(null='false', default=0)
    InnOuts = IntegerField(null='false', default=0)
    ZR = IntegerField(null='false', default=0)
    PO = IntegerField(null='false', default=0)
    WP = IntegerField(null='false', default=0)
    CS = IntegerField(null='false', default=0)
    E = IntegerField(null='false', default=0)
    DP = IntegerField(null='false', default=0)
    SB = IntegerField(null='false', default=0)

    class Meta:
        primary_key = CompositeKey('playerID', 'yearID', 'stint')


class Teams_Half(BaseModel):
    yearID = IntegerField(null='true')
    lgID = CharField(null='true')
    teamID = CharField(null='true')
    half = IntegerField(null='false', default=1)
    divID = CharField(null='true')
    divWin = CharField(null='true')
    Rank = IntegerField(null='false', default=0)
    G = IntegerField(null='false', default=0)
    W = IntegerField(null='false', default=0)
    L = IntegerField(null='false', default=0)

    class Meta:
        primary_key = CompositeKey('yearID', 'teamID', 'half')
    
class Batting(BaseModel):
    playerID = ForeignKeyField(Master)
    yearID = IntegerField(null='false')
    stint = IntegerField(null='false')
    teamID = CharField(null='true')
    lgID = CharField(null='true')
    G = IntegerField(null='false', default=0)
    AB = IntegerField(null='false', default=0)
    R = IntegerField(null='false', default=0)
    H = IntegerField(null='false', default=0)
    doubles = IntegerField(null='false', db_column="2B", default=0)
    triples = IntegerField(null='false', db_column="3B", default=0)
    HR = IntegerField(null='false', default=0)
    RBI = IntegerField(null='false', default=0)
    SB = IntegerField(null='false', default=0)
    CS = IntegerField(null='false', default=0)
    BB = IntegerField(null='false', default=0)
    SO = IntegerField(null='false', default=0)
    IBB = IntegerField(null='false', default=0)
    HBP = IntegerField(null='false', default=0)
    SH = IntegerField(null='false', default=0)
    SF = IntegerField(null='false', default=0)
    GIDP = IntegerField(null='false', default=0)
    
    class Meta:
        primary_key = CompositeKey('playerID', 'yearID', 'stint')


TABLE_NAMES = [ 
    Master, Teams_Franchises, Parks, Managers, Teams , ManagersHalf, Schools,
    Batting_Post, College_Playing, Awards_Players, Fielding_OF, 
    Pitching, Allstar_Full, Home_Games, Appearances , Hall_Of_Fame, Fielding_Post,
    Series_Post, Awards_Managers, Awards_Share_Managers, Salaries, Pitching_Post,
    Awards_Share_Players, Fielding, Teams_Half, Batting
]

def Create_Tables():
    BDB_DB.connect()
    BDB_DB.create_tables(TABLE_NAMES)

