#!/bin/sh
cd ../wavs
#ID="Everybody+(Backstreet's+Back)+(Radio+Edit)"

echo UPLOADING
for f in [A-Za-z]*.wav; do
	ID=`basename $f .wav`
	#echo $ID
	ESCAPED=`perl -e "use URI::Escape; print uri_escape(\"$ID\")"`
	#echo $ESCAPED
	AUDIO=`base64 -w 0 -i "$ID".wav`
	RESOURCE=localhost:3000/tracks/$ESCAPED
	#echo $RESOURCE
	echo "{\"Id\":\"$ID\",\"Audio\":\"$AUDIO\"}" > $ID.input
	curl -X PUT -d @$ID.input $RESOURCE
	retcurl=$?
	if [ $retcurl -eq 0 ]; then
		echo success $ID
	else
		echo failed $ID $retcurl
	fi
done
echo

echo LISTING -- expecting 4 items
curl -X GET localhost:3000/tracks
echo

echo DOWNLOADING
for f in [A-Za-z]*.wav; do
	ID=`basename $f .wav`
	#echo $ID
	ESCAPED=`perl -e "use URI::Escape; print uri_escape(\"$ID\")"`
	#echo $ESCAPED
	RESOURCE=localhost:3000/tracks/$ESCAPED
	#echo $RESOURCE
	#curl -s -X GET $RESOURCE | base64 -d > $ID.output
	curl -s -X GET $RESOURCE -o $ID.output
	retcurl=$?
	if [ $retcurl -eq 0 ]; then
		cmp $ID.input $ID.output
		retcmp=$?
		if [ $retcmp -eq 0 ]; then
			echo success $ID
		else
			echo mismatch $ID
		fi
	else
		echo failed $ID $retcurl
	fi
done
echo

echo DELETING
for f in [A-Za-z]*.wav; do
	ID=`basename $f .wav`
	ESCAPED=`perl -e "use URI::Escape; print uri_escape(\"$ID\")"`
	RESOURCE=localhost:3000/tracks/$ESCAPED
	curl -s -X DELETE $RESOURCE
	retcurl=$?
	if [ $retcurl -eq 0 ]; then
		echo success $ID
	else
		echo failed $ID $retcurl
	fi
done
echo

echo LISTING -- expecting empty list
curl -X GET localhost:3000/tracks
echo

# clean up
for f in [A-Za-z]*.wav; do
	ID=`basename $f .wav`
	rm $ID.input $ID.output
done
