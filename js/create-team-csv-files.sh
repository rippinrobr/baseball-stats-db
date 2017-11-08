#!/bin/bash
# Future command line options
TEAMIDS='SFN|NY1'
OUTDIR=../giants
FRANCHISEIDS='SFG|NYI'
BDBDIR=../baseballdatabank/official

# internal variables
MGRIDSFILE='managerids'
PLAYERIDSFILE='playerids'
TEAMFILES=`ls $BDBDIR/*.csv | grep -v School | grep -v TeamsFranc | grep -v Awards | grep -v HallOfFame | grep -v Master | grep -v FieldingOF`

rm $OUTDIR/*.csv
rm $PLAYERIDSFILE
rm $MGRIDSFILE

for f in $TEAMFILES
do
  echo "Processing $f file..."
  BASEFILE=$(basename $f)
  head -1 $f > $OUTDIR/$BASEFILE
  cat $f | grep -E $TEAMIDS >>$OUTDIR/$BASEFILE
done

echo -n "Grabbing playerIds..."
cat $OUTDIR/Batting.csv $OUTDIR/Fielding.csv $OUTDIR/Pitching.csv | cut -d, -f1 | grep -v yearID |  sort | uniq >$PLAYERIDSFILE
echo "Done!"

echo -n "Grabbing the managerIds..."
cat $OUTDIR/Managers*.csv | cut -d, -f1 | grep -v yearID | sort | uniq >$MGRIDSFILE
echo "Done!"

echo -n "Creating the $OUTDIR/Players.csv file..."
cat $BDBDIR/Master.csv | grep -f $PLAYERIDSFILE | cut -d, -f 2-33 | uniq >$OUTDIR/Players.csv
echo "Done!"

echo -n "Adding Managers to the $OUTDIR/Master.csv file..."
cat $BDBDIR/Master.csv | grep -f $MGRIDSFILE | cut -d, -f 2-33 | uniq >$OUTDIR/MgrDemographics.csv
echo "Done!"

echo -n "Creating the $OUTDIR/FieldingOF.csv file..."
head -1 $BDBDIR/FieldingOF.csv >$OUTDIR/FieldingOF.csv
cat $BDBDIR/FieldingOF.csv | grep -f $PLAYERIDSFILE >>$OUTDIR/FieldingOF.csv
echo "Done!"

echo -n "Creating the $OUTDIR/Pos.csv file..."
cat $BDBDIR/Fielding.csv  | cut -d, -f1,6 | uniq >$OUTDIR/Pos.csv
cat $BDBDIR/FieldingPost.csv | grep -v playerID | cut -d, -f1,6 | uniq >>$OUTDIR/Pos.csv
cat $OUTDIR/Pos.csv | uniq >$OUTDIR/Pos.csv

echo -n "Creating the $OUTDIR/AwardsPlayers.csv file..."
head -1 $BDBDIR/AwardsPlayers.csv >$OUTDIR/AwardsPlayers.csv
cat $BDBDIR/AwardsPlayers.csv | grep -f $PLAYERIDSFILE >>$OUTDIR/AwardsPlayers.csv
echo "Done!"

echo -n "Creating the $OUTDIR/AwardsManagers.csv file..."
head -1 $BDBDIR/AwardsManagers.csv >$OUTDIR/AwardsManagers.csv
cat $BDBDIR/AwardsManagers.csv | grep -f $MGRIDSFILE >>$OUTDIR/AwardsManagers.csv
echo "Done!"

echo -n "Creating the $OUTDIR/TeamsFranchise.csv file..."
head -1 $BDBDIR/TeamsFranchises.csv >$OUTDIR/TeamsFranchises.csv
cat $BDBDIR/TeamsFranchises.csv | egrep "$FRANCHISEIDS" >>$OUTDIR/TeamsFranchises.csv
echo "Done!"
