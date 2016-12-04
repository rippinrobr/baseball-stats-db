from peewee import *
from playhouse.postgres_ext import *
import logging
import sys

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

class TeamsFranchises(BaseModel):
    franchID = CharField(null='false', primary_key=True)
    franchName = CharField(null='false')
    active = CharField(null='true')
    NAassoc = CharField(null='true')
    
class Teams(BaseModel):
    yearID = IntegerField(null='false')
    lgID = CharField(null='false')
    teamID = CharField(null='true')
    franchID = ForeignKeyField(TeamsFranchises, related_name='franchID')
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
    2B = IntegerField(null='false', default=0)
    3B = IntegerField(null='false', default=0)
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
    park = CharField(null='true')
    attendance = IntegerField(null='false', default=0)
    BPF = IntegerField(null='false', default=0)
    PPF = IntegerField(null='false', default=0)
    teamIDBR = CharField(null='false')
    teamIDlahman45 = CharField(null='false')
    teamIDretro = CharField(null='false')

    class Meta:
        indexes = (
            (('yearID', 'lgID', 'teamID'))
        )

class Managers(BaseModel):
    playerID = ForeignKeyField(Master, related_name='playerID')
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
        indexes = (
                # create a unique index
                (('playerID', 'teamID', 'lgID', 'yearID'), True),
        )

class ManagersHalf(BaseModel):
    playerID = ForeignKeyField(Master, related_name='playerID')
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
        indexes = (
                # create a unique index
                (('playerID', 'teamID', 'lgID', 'yearID', 'half'), True),
        )

class Schools(BaseModel):
    schoolID = CharField(null='false', primary_key=True)
    name_full = CharField(null='false', unique=true)
    city = CharField(null='false')
    state = CharField(null='false')
    country = CharField(null='false')
    
   
class BattingPost(BaseModel):
    yearID = IntegerField(null='false')
    round = CharField(null='true')
    playerID = ForeignKeyField(Master, related_name='playerID')
    teamID = CharField(null=False)
    lgID = CharField(null='false')
    G = IntegerField(null='false', default=0)
    AB = IntegerField(null='false', default=0)
    R = IntegerField(null='false', default=0)
    H = IntegerField(null='false', default=0)
    2B = IntegerField(null='false', default=0)
    3B = IntegerField(null='false', default=0)
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
        indexes = (
                # create a unique index
                (('yearID', 'round', 'playerID', 'teamID'), True),
        )
    
class CollegePlaying(BaseModel):
    playerID = ForeignKeyField(Master, related_name='playerID')
    schoolID = ForeignKeyField(Schools, related_name='schoolID')
    yearID = IntegerField(null='false')

    class Meta:
        index = (
            (('playerID', 'schoolID', 'yearID'), True),
        )
    
class Parks(BaseModel):
    parkID = CharField(null='false', primary_key=True)
    name = CharField(null='false')
    alias = CharField(null='true')
    city = CharField(null='false')
    state = CharField(null='false')
    country = CharField(null='false')
    
class AwardsPlayers(BaseModel):
    playerID = ForeignKeyField(Master, related_name='playerID')
    awardID = CharField(null='false')
    yearID = IntegerField(null='false')
    lgID = CharField(null='false')
    tie = CharField(null='true')
    notes = CharField(null='true')

    class Meta:
        index = (
            (('playerID', 'awardID', 'yearID'), True)
        )
    
class FieldingOF(BaseModel):
    playerID = ForeignKeyField(Master, related_name='playerID')
    yearID = IntegerField(null='false', default=0)
    stint = IntegerField(null='false', default=0)
    Glf = IntegerField(null='false', default=0)
    Gcf = IntegerField(null='false', default=0)
    Grf = IntegerField(null='false', default=0)

    class Meta:
        index = (
            (('playerID', 'yearID', 'stint'), True)
        )
    
class Pitching(BaseModel):
    playerID = ForeignKeyField(Master, related_name='playerID')
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
        index = (
            (('playerID', 'yearID', 'stint', 'teamID'), True)
        )

class AllstarFull(BaseModel):
    playerID = CharField(null='true')
    yearID = IntegerField(null='true')
    gameNum = IntegerField(null='true')
    gameID = CharField(null='true')
    teamID = CharField(null='true')
    lgID = CharField(null='true')
    GP = IntegerField(null='true')
    startingPos = IntegerField(null='true')

    class Meta:
        index = (
            (('playerID', 'yearID', 'gameNum'), True)
        )
    
class HomeGames(BaseModel):
    yearID = IntegerField(null='false')
    parkID = ForeignKeyField(Parks, related_name="parkID")
    teamID = charField(null=False)
    lgID = CharField(null=False)
    attendance = IntegerField(null='true')
    games = IntegerField(null='true')
    first_game = CharField(null='true')
    last_game = CharField(null='true')
    openings = IntegerField(null='true')

    class Meta:
        index = (
            (('yearID', 'parkID', 'teamID'), True)
        )
    
class Appearances(BaseModel):
    yearID = IntegerField(null='true')
    teamID = ForeignKeyField(Teams)
    lgID = CharField(null='true')
    playerID = CharField(null='true')
    G_all = IntegerField(null='true')
    GS = IntegerField(null='true')
    G_batting = IntegerField(null='true')
    G_defense = IntegerField(null='true')
    G_p = IntegerField(null='true')
    G_c = IntegerField(null='true')
    G_1b = IntegerField(null='true')
    G_2b = IntegerField(null='true')
    G_3b = IntegerField(null='true')
    G_ss = IntegerField(null='true')
    G_lf = IntegerField(null='true')
    G_cf = IntegerField(null='true')
    G_rf = IntegerField(null='true')
    G_of = IntegerField(null='true')
    G_dh = IntegerField(null='true')
    G_ph = IntegerField(null='true')
    G_pr = IntegerField(null='true')

    class Meta:
        index = (
            (('yearID', 'teamID', 'playerID'), True)
        )

class HallOfFame(BaseModel):
    playerID = ForeignKeyField(Master, related_name='playerID')
    yearID = IntegerField(null='true')
    votedBy = CharField(null='true')
    ballots = IntegerField(null='true')
    needed = IntegerField(null='true')
    votes = IntegerField(null='true')
    inducted = CharField(null='true')
    category = CharField(null='true')
    needed_note = CharField(null='true')

    class Meta:
        index = (
            (('playerID', 'category'), True)
        )
    
class FieldingPost(BaseModel):
    playerID = ForeignKeyField(Master, related_name='playerID')
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

    def Meta:
        index = (
            (('playerID', 'yearID', 'round', 'pos'), True)
        )
        
    
class SeriesPost(BaseModel):
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
        index = (
            (('yearID', 'round'), True)
        )

class AwardsManagers(BaseModel):
    playerID = ForeignKeyField(Master, relalted_name='playerID')
    awardID = CharField(null='false')
    yearID = IntegerField(null='false')
    lgID = CharField(null='false')
    tie = CharField(null='false', default='N')
    notes = CharField(null='true')
    
    class Meta:
        index (
            (('yearID', 'playerID', 'lgID', 'tie'), True)
        )

class AwardsShareManagers(BaseModel):
    awardID = CharField(null='true')
    yearID = IntegerField(null='true')
    lgID = CharField(null='true')
    playerID = ForeignFieldKey(Master, related_name='playerID')
    pointsWon = IntegerField(null='true')
    pointsMax = IntegerField(null='true')
    votesFirst = IntegerField(null='true')
    
    class Meta:
        index (
            (('awardID', 'yearID', 'lgID', 'playerID'), True)
        )

class Salaries(BaseModel):
    yearID = IntegerField(null='true')
    teamID = CharField(null='true')
    lgID = CharField(null='true')
    playerID = ForeignKeyField(Master, related_name='playerID')
    salary = IntegerField(null='true')

    def Meta:
        index = (
            (('yearID', 'teamID', 'playerID'), True)
        )
    
class PitchingPost(BaseModel):
    playerID = ForeignKeyField(Master, related_name='playerID')
    yearID = IntegerField(null='true')
    round = CharField(null='true')
    teamID = CharField(null='true')
    lgID = CharField(null='true')
    W = IntegerField(null='false', default=0))
    L = IntegerField(null='false', default=0))
    G = IntegerField(null='false', default=0))
    GS = IntegerField(null='false', default=0))
    CG = IntegerField(null='false', default=0))
    SHO = IntegerField(null='false', default=0))
    SV = IntegerField(null='false', default=0))
    IPouts = IntegerField(null='false', default=0))
    H = IntegerField(null='false', default=0))
    ER = IntegerField(null='false', default=0))
    HR = IntegerField(null='false', default=0))
    BB = IntegerField(null='false', default=0))
    SO = IntegerField(null='false', default=0))
    BAOpp = FloatField(null='false', default=0.000)
    ERA = FloatField(null='false', default=0.00)
    IBB = IntegerField(null='false', default=0))
    WP = IntegerField(null='false', default=0))
    HBP = IntegerField(null='false', default=0))
    BK = IntegerField(null='false', default=0))
    BFP = IntegerField(null='false', default=0))
    GF = IntegerField(null='false', default=0))
    R = IntegerField(null='false', default=0))
    SH = IntegerField(null='false', default=0))
    SF = IntegerField(null='false', default=0))
    GIDP = IntegerField(null='false', default=0))

    def Meta:
        index = (
            (('playerID', 'yearID', 'round'), True)
        )
    
class AwardsSharePlayers(BaseModel):
    awardID = CharField(null='false')
    yearID = IntegerField(null='false')
    lgID = CharField(null='false')
    playerID = ForeignKeyField(Master, related_name='playerID')
    pointsWon = IntegerField(null='fails', default=0)
    pointsMax = IntegerField(null='fails', default=0)
    votesFirst = IntegerField(null='fails', default=0)

    class Meta:
        index = (
            (('awardID', 'yearID', 'lgID', 'playerID'), True)
        )
    
class Fielding(BaseModel):
    playerID = ForeignKeyField(Master, related_name='playerID')
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
        index = (
            (('playerID', 'yearID', 'stint'), True)
        )


class TeamsHalf(BaseModel):
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
        index = (
            (('yearID', 'teamID', 'half'),  True)
        )
    
class Batting(BaseModel):
    playerID = ForeignKeyField(Master, related_name='playerID')
    yearID = IntegerField(null='true')
    stint = IntegerField(null='true')
    teamID = CharField(null='true')
    lgID = CharField(null='true')
    G = IntegerField(null='true')
    AB = IntegerField(null='true')
    R = IntegerField(null='true')
    H = IntegerField(null='true')
    2B = IntegerField(null='true')
    3B = IntegerField(null='true')
    HR = IntegerField(null='true')
    RBI = IntegerField(null='true')
    SB = IntegerField(null='true')
    CS = IntegerField(null='true')
    BB = IntegerField(null='true')
    SO = IntegerField(null='true')
    IBB = IntegerField(null='true')
    HBP = IntegerField(null='true')
    SH = IntegerField(null='true')
    SF = IntegerField(null='true')
    GIDP = IntegerField(null='true')
    
    class Meta:
        index = (
            (('playerID', 'yearID', 'stint'), True)
        )
