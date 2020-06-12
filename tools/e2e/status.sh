#!/bin/bash

if type "grpcurl" > /dev/null 2>&1; then
  echo "grpcurl is exist"
else
  echo "install grpcurl"
  brew install grpcurl
fi


url=localhost:18080
data=(
#    "GetIconImageList"
    "GetSubscriptions"
#    "GetPopularSubscriptions"
#    "GetRecommendSubscriptions"
#    "GetMySubscription"
#    "CreateSubscription"
#    "UpdateSubscription"
#    "RegisterSubscription"
#    "UnregisterSubscription"
)


RESULT=$(cat <<- EndOfMessage
EndOfMessage)


for name in "${data[@]}"
do
    out_file=./out/$name.txt
    if [ ! -e "$out_file" ] ; then
      touch "$out_file"
    fi

    res=$(grpcurl -plaintext $url subscription.SubscriptionService/$name)
    if [[ $res != $(cat $out_file) ]] ; then
      echo "$name: Test Failed!"
    else
      echo "$name: Test Ok!"
    fi
done

echo "All Test Finished!"
