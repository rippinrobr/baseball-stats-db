from peewee import *
from playhouse.postgres_ext import *
import logging
import sys
import argparse

#  THIS IS HOW TO MAKE THIS PART DYNAMIC http://docs.peewee-orm.com/en/latest/peewee/database.html#dynamic-db
database_proxy =  Proxy() #SqliteDatabase('db/bdb_v0.0.1.sqlite')

class BaseModel(Model):
    """BaseModel is used to connect to the database"""
    class Meta(object):
        """Where the connection is set"""
        database = database_proxy

# Print the connection info here for sqlite
class People(BaseModel):
    playerID = CharField(null='false',primary_key=True, max_length=9)
    birthYear = IntegerField(null='true')
    birthMonth = IntegerField(null='true')
    birthDay = IntegerField(null='true')
    birthCountry = CharField(null='true', max_length=14)
    birthState = CharField(null='true', max_length=22)
    birthCity = CharField(null='true', max_length=26)
    deathYear = IntegerField(null='true')
    deathMonth = IntegerField(null='true')
    deathDay = IntegerField(null='true')
    deathCountry = CharField(null='true', max_length=14)
    deathState = CharField(null='true', max_length=20)
    deathCity = CharField(null='true', max_length=28)
    nameFirst = CharField(null='false', max_length=12)
    nameLast = CharField(null='false', max_length=14)
    nameGiven = CharField(null='true', max_length=43)
    weight = IntegerField(null='true')
    height = IntegerField(null='true')
    bats = CharField(null='false')
    throws = CharField(null='false')
    debut = CharField(null='true')
    finalGame = CharField(null='true')
    retroID = CharField(null='false')
    bbrefID = CharField(null='false', max_length=9) 

    def Meta(object):
        db_table = 'people'

def load_people_table(records):
    with BDB_DB.atomic() as txn:
        for rec in records:
            if any(rec):
                if not create_people(rec):
                    print "Didnt save %s" % rec

def create_people(rec):
    try:
        try:
            birthYear = None
            if 'birthYear' in rec and rec['birthYear']:
                birthYear = int(rec['birthYear'])
            
            birthMonth = None
            if 'birthMonth' in rec and rec['birthMonth'] and rec['birthMonth'] != '':
                birthMonth = int(rec['birthMonth'])
            
            birthDay = None
            if 'birthDay' in rec and rec['birthDay'] and rec['birthDay'] != '':
                birthDay = int(rec['birthDay'])

            deathYear = None
            if 'deathYear' in rec and rec['deathYear'] and rec['deathYear'] != '':
                deathYear = int(rec['deathYear'])
                
            deathMonth = None
            if 'deathMonth' in rec and  rec['deathMonth'] and rec['deathMonth'] != '':
                deathMonth = int(rec['deathMonth'])
            
            deathDay = None
            if 'deathDay' in rec and rec['deathDay'] and rec['deathDay'] != '':
                deathDay = int(rec['deathDay'])

            weight = None
            if 'weight' in rec and rec['weight'] and rec['weight'] != '':
                weight = int(rec['weight'])

            height = None
            if 'height' in rec and rec['height'] and rec['height'] != '':
                height = int(rec['height'])

        except:
            for e in sys.exc_info():
                print "converstions: %s" % e
                return False

        m = People.create(
            playerID=rec['playerID'],
            birthYear=birthYear,
            birthMonth=birthMonth,
            birthDay=birthDay,
            birthCountry=rec['birthCountry'],
            birthState=rec['birthState'],
            birthCity=rec['birthCity'],
            deathYear=deathYear,
            deathMonth=deathMonth,
            deathDay=deathDay,
            deathCountry=rec['deathCountry'],
            deathState=rec['deathState'],
            deathCity=rec['deathCity'],
            nameFirst=rec['nameFirst'],
            nameLast=rec['nameLast'],
            nameGiven=rec['nameGiven'],
            weight=weight,
            height=height,
            bats=rec['bats'],
            throws=rec['throws'],
        #    debut=rec['debut'],
        #    finalGame=rec['finalGame'],
        #    retroID=rec['retroID'],
            bbrefID=rec['bbrefID']
        )
        m.save()
        BDB_DB.commit()
        return True
    except:
        for e in sys.exc_info():
            print "%s" % e
        return False



class TeamsFranchises(BaseModel):
    franchID = CharField(null='false', primary_key=True)
    franchName = CharField(null='false')
    active = CharField(null='true')
    NAassoc = CharField(null='true')

    def Meta(object):
        db_table = 'teamsfranchises'

        
    
class Parks(BaseModel):
    parkkey = CharField(null='false', primary_key=True)
    parkname = CharField(null='false')
    parkalias = CharField(null='true')
    alias = CharField(null='true')
    city = CharField(null='false')
    state = CharField(null='false')
    country = CharField(null='false')

    def Meta(object):
        db_table = 'parks'
    

class Teams(BaseModel):
    yearID = IntegerField(null='false')
    lgID = CharField(null='false')
    teamID = CharField(null='true')
    franchID = CharField(null='false')
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
    park = CharField(null='true')
    attendance = IntegerField(null='false', default=0)
    BPF = IntegerField(null='false', default=0)
    PPF = IntegerField(null='false', default=0)
    teamIDBR = CharField(null='false')
    teamIDlahman45 = CharField(null='false')
    teamIDretro = CharField(null='false')

    class Meta(object):
        primary_key = CompositeKey('yearID', 'lgID', 'teamID')
        db_table = 'teams'


class Managers(BaseModel):
    playerID = CharField(null='false')
    yearID = IntegerField(null='false')
    teamID = CharField(null=False)
    lgID = CharField(null=False)
    inseason = IntegerField(null='false', default=0)
    G = IntegerField(null='false', default=0)
    W = IntegerField(null='false', default=0)
    L = IntegerField(null='false', default=0)
    rank = IntegerField(null='false', default=0)
    plyrMgr = CharField(null='true')

    class Meta(object):
        #primary_key = CompositeKey('playerID', 'yearID', 'lgID', 'teamID', 'rank')
        db_table = 'managers'

class ManagersHalf(BaseModel):
    playerID = CharField(null='false')
    yearID = IntegerField(null='false')
    teamID = CharField(null=False)
    lgID = CharField(null='false')
    inseason = IntegerField(null='true')
    half = IntegerField(null='false')
    G = IntegerField(null='false', default=0)
    W = IntegerField(null='false', default=0)
    L = IntegerField(null='false', default=0)
    rank = IntegerField(null='true')
    
    class Meta(object):
        primary_key = CompositeKey('playerID', 'teamID', 'lgID', 'yearID', 'half')
        db_table = 'managershalf'

class Schools(BaseModel):
    schoolID = CharField(null='false', primary_key=True)
    name_full = CharField(null='false')
    city = CharField(null='false')
    state = CharField(null='false')
    country = CharField(null='false')

    def Meta(object):
        db_table = 'schools'

   
class Batting_Post(BaseModel):
    yearID = IntegerField(null='false')
    round = CharField(null='true')
    playerID = CharField(null='false')
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

    class Meta(object):
        primary_key = CompositeKey('yearID', 'round', 'playerID', 'teamID')
        db_table = 'battingpost'
    
class College_Playing(BaseModel):
    playerID = CharField(null='false')
    schoolID = CharField(null='false')
    yearID = IntegerField(null='false')

    class Meta(object):
        primary_key = CompositeKey('playerID', 'schoolID', 'yearID')
        db_table = 'collegeplaying'

class Awards_Players(BaseModel):
    playerID = CharField(null='false')
    awardID = CharField(null='false')
    yearID = IntegerField(null='false')
    lgID = CharField(null='false')
    tie = CharField(null='true')
    notes = CharField(null='true')

    class Meta(object):
        #primary_key = CompositeKey('playerID', 'awardID', 'yearID')
        db_table = 'awardsplayers'
    
class Fielding_OF(BaseModel):
    playerID = CharField(null='false')
    yearID = IntegerField(null='false', default=0)
    stint = IntegerField(null='false', default=0)
    Glf = IntegerField(null='false', default=0)
    Gcf = IntegerField(null='false', default=0)
    Grf = IntegerField(null='false', default=0)
    Grf = IntegerField(null='false', default=0)

    class Meta(object):
        primary_key = CompositeKey('playerID', 'yearID', 'stint')
        db_table = 'fieldingof'
    
    
class Fielding_OF_Split(BaseModel):
    playerID = CharField(null='false')
    yearID = IntegerField(null='false', default=0)
    stint = IntegerField(null='false', default=0)
    teamid  = CharField(null='false')
    lgid = CharField(null='false')
    pos = CharField(null='false')
    g = IntegerField(null='false',default=0)
    gs = IntegerField(null='false',default=0)
    innouts = IntegerField(null='false',default=0)
    po = IntegerField(null='false',default=0)
    A = IntegerField(null='false',default=0)
    E = IntegerField(null='false',default=0)
    DP = IntegerField(null='false',default=0)
    PB = IntegerField(null='false',default=0)
    WP = IntegerField(null='false',default=0)
    SB = IntegerField(null='false',default=0)
    CS = IntegerField(null='false',default=0)
    ZR = IntegerField(null='false',default=0)
   
    class Meta(object):
        primary_key = CompositeKey('playerID', 'yearID', 'teamid', 'pos', 'stint')
        db_table = 'fieldingofsplit'
    
class Pitching(BaseModel):
    playerID = CharField(null='false')
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
    
    class Meta(object):
        primary_key = CompositeKey('playerID', 'yearID', 'stint', 'teamID')
        db_table = 'pitching'

class Allstar_Full(BaseModel):
    playerID = CharField(null='false')
    yearID = IntegerField(null='false')
    gameNum = IntegerField(null='false')
    gameID = CharField(null='false')
    teamID = CharField(null='false')
    lgID = CharField(null='false')
    GP = IntegerField(null='false', default=0)
    startingPos = IntegerField(null='false', default=0)

    class Meta(object):
        primary_key = CompositeKey('playerID', 'yearID', 'gameNum')
        db_table = 'allstarfull'
    
class Home_Games(BaseModel):
    yearkey = IntegerField(null='false')
    parkkey = CharField(null='false')
    teamkey = CharField(null='false')
    leaguekey = CharField(null='false')
    attendance = IntegerField(null='true')
    games = IntegerField(null='true')
    spanfirst = CharField(null='true')
    spanlast = CharField(null='true')
    openings = IntegerField(null='true')

    class Meta(object):
        primary_key = CompositeKey('yearkey', 'parkkey', 'teamkey')
        db_table = 'homegames'
    
class Appearances(BaseModel):
    yearID = IntegerField(null='false')
    teamID = CharField(null='false')
    lgID = CharField(null='false')
    playerID = CharField(null='false')
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

    class Meta(object):
        primary_key = CompositeKey('yearID', 'teamID', 'playerID', 'G_all')
        db_table = 'appearances'

class Hall_Of_Fame(BaseModel):
    playerID = CharField(null='false')
    yearID = IntegerField(null='true')
    votedBy = CharField(null='true')
    ballots = IntegerField(null='true')
    needed = IntegerField(null='true')
    votes = IntegerField(null='true')
    inducted = CharField(null='true')
    category = CharField(null='true')
    needed_note = CharField(null='true')

    class Meta(object):
        primary_key = CompositeKey('playerID', 'category', 'yearID', 'votes', 'ballots')
        db_table = 'halloffame'
    
class Fielding_Post(BaseModel):
    playerID = CharField(null='false')
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

    class Meta(object):
        primary_key = CompositeKey('playerID', 'yearID', 'round', 'pos')
        db_table = 'fieldingpost'
        
    
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
    
    class Meta(object):
        primary_key = CompositeKey('yearID', 'round')
        db_table = 'seriespost'

class Awards_Managers(BaseModel):
    playerID = CharField(null='false')
    awardID = CharField(null='false')
    yearID = IntegerField(null='false')
    lgID = CharField(null='false')
    tie = CharField(null='false', default='N')
    notes = CharField(null='true')
    
    class Meta(object):
        #primary_key = CompositeKey('yearID', 'playerID', 'lgID', 'tie')
        db_table = 'awardsmanagers'

class Awards_Share_Managers(BaseModel):
    awardID = CharField(null='true')
    yearID = IntegerField(null='true')
    lgID = CharField(null='true')
    playerID = CharField(null='false')
    pointsWon = IntegerField(null='true')
    pointsMax = IntegerField(null='true')
    votesFirst = IntegerField(null='true')
    
    class Meta(object):
        primary_key = CompositeKey('awardID', 'yearID', 'lgID', 'playerID')
        db_table = 'awardssharemanagers'

class Salaries(BaseModel):
    yearID = IntegerField(null='true')
    teamID = CharField(null='true')
    lgID = CharField(null='true')
    playerID = CharField(null='false')
    salary = IntegerField(null='true')

    class Meta(object):
        primary_key = CompositeKey('yearID', 'teamID', 'playerID')
        db_table = 'salaries'
    
class Pitching_Post(BaseModel):
    playerID = CharField(null='false')
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

    class Meta(object):
        primary_key = CompositeKey('playerID', 'yearID', 'round')
        db_table = 'pitchingpost'

class Awards_Share_Players(BaseModel):
    awardID = CharField(null='false')
    yearID = IntegerField(null='false')
    lgID = CharField(null='false')
    playerID = CharField(null='false')
    pointsWon = DecimalField(null='fails', default=0)
    pointsMax = IntegerField(null='fails', default=0)
    votesFirst = DecimalField(null='fails', default=0)

    class Meta(object):
        primary_key = CompositeKey('awardID', 'yearID', 'lgID', 'playerID')
        db_table = 'awardsshareplayers'
    
class Fielding(BaseModel):
    playerID = CharField(null='false')
    yearID = IntegerField(null='false', default=0)
    stint = IntegerField(null='false', default=0)
    teamID = CharField(null='false')
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

    class Meta(object):
        primary_key = CompositeKey('playerID', 'yearID', 'teamID', 'Pos', 'stint')
        db_table = 'fielding'

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

    class Meta(object):
        primary_key = CompositeKey('yearID', 'teamID', 'half')
        db_table = 'teamshalf'

class Batting(BaseModel):
    playerID = CharField(null='false')
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
    
    class Meta(object):
        primary_key = CompositeKey('playerID', 'yearID', 'stint')
        db_table = 'batting'

def initialize_db_and_connect(bdb_db):
    database_proxy.initialize(bdb_db)
    database_proxy.connect()
    database_proxy.create_tables(TABLE_NAMES)
    

TABLE_NAMES = [ 
    People, TeamsFranchises, Parks, Managers, Teams , ManagersHalf, Schools,
    Batting_Post, College_Playing, Awards_Players, Fielding_OF, Fielding_OF_Split,
    Pitching, Allstar_Full, Home_Games, Appearances , Hall_Of_Fame, Fielding_Post,
    Series_Post, Awards_Managers, Awards_Share_Managers, Salaries, Pitching_Post,
    Awards_Share_Players, Fielding, Teams_Half, Batting
]
