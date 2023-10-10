#!/bin/sh

cd ../wavs

# upload correct track
TID="Everybody+(Backstreet's+Back)+(Radio+Edit)"
TAUDIO=`base64 -w 0 -i "$TID".wav`
ESCAPED=`perl -e "use URI::Escape; print uri_escape(\"$TID\")"`
TURL=http://localhost:3000/tracks/$ESCAPED
echo "{\"Id\":\"$TID\",\"Audio\":\"$TAUDIO\"}" > $TID.input
curl -X PUT -d @$TID.input $TURL
retcurl=$?
if [ $retcurl -ne 0 ]; then
	echo failed upload $SID
fi

# fetch sample from cooldown
SID="~Everybody+(Backstreet's+Back)+(Radio+Edit)"
SAUDIO=`base64 -i "$SID".wav`
SURL=localhost:3002/cooltown
echo "{\"Audio\":\"$SAUDIO\"}" > $SID.input
curl -s -X POST -d @$SID.input -o $SID.output $SURL
retcurl=$?
if [ $retcurl -ne 0 ]; then
	echo failed upload $SID
fi

# verify
cmp --ignore-initial=51:1 $TID.input $SID.output
retcmp=$?
if [ $retcmp -eq 0 ]; then
	echo success $SID
else
	echo failed $SID
fi

# clean up
rm $TID.input $SID.input $SID.output
