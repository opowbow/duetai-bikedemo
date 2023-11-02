#! /bin/bash

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
echo $SCRIPT_DIR
source $SCRIPT_DIR/.secret-env
gcloud projects remove-iam-policy-binding $PROJECT_ID --member=serviceAccount:bike-rental-service@$PROJECT_ID.iam.gserviceaccount.com --role=roles/bigquery.user
rm $SCRIPT_DIR/main_test.go
git checkout $SCRIPT_DIR/main.go