# Duet AI Bike Rental Demo

## Preparaton from within the cloud workstation

```sh
export PROJECT_ID=my-project-id
```

Clone the git repo from source code repo

```sh
gcloud source repos clone demo-repo --project=$PROJECT_ID
git checkout main
```

Get an API Key valid for Maps Embed API.

```sh
gcloud alpha services api-keys create \
    --project $PROJECT_ID --api-target="service=maps-embed-backend.googleapis.com"
KEY_NAME="$(gcloud alpha services api-keys list --format='value(name)' --limit 1 --project=$PROJECT_ID)"
export MAPS_KEY="$(gcloud alpha services api-keys get-key-string $KEY_NAME --format='value(keyString)' --project=$PROJECT_ID)"
echo "Maps API Key: $(echo $MAPS_KEY | cut -c1-7)..."
```

Prepare hot reload for Go

```sh
go install github.com/cosmtrek/air@latest
```

## Big Query - Data Engineer

### DEMO Steps

1. Open Bq editor in console and explore public dataset
2. Open [London Bike hire](https://console.cloud.google.com/bigquery?p=bigquery-public-data&d=london_bicycles&page=dataset)
3. Use autocompletion to come up with a query like this

    ```sql
    SELECT name, longitude, latitude, bikes_count as bikescount FROM `bigquery-public-data.london_bicycles.cycle_stations` ORDER BY bikes_count DESC LIMIT 50
    ```

## Local Development - Google SDK

Start `air` for hot reload of the go application

```sh
/home/user/go/bin/air 
```

1. fill in the function called `loadBikeStationData`

    ```go
    log.Println("Load london bike station data from big query with context PROJECT_ID...")
	stations := []BikeStation{}
	query := "SELECT name, longitude, latitude, bikes_count as bikescount FROM `bigquery-public-data.london_bicycles.cycle_stations` ORDER BY bikes_count DESC LIMIT 50"

	project := os.Getenv("PROJECT_ID")
    ctx := context.Background()
    client, err := bigquery.NewClient(ctx, project)
    ```

## Local Development - Generic Coding

```go
func repeatBikeEmoji(repeat int) string {
}
```

## Local Development - Test

1. Create a file called `main_test.go`
2. Ask chat to create a unit test for the emoji function
3. Run the test using `go test`

## Deploy to Cloud Run -- Operations Persona

```sh
gcloud iam service-accounts create bike-rental-service --project $PROJECT_ID
```

use that Service account to deploy to cloud run

```sh
gcloud run deploy bike-rental-demo-london --source . --project $PROJECT_ID --set-env-vars=MAPS_KEY=$MAPS_KEY,PROJECT_ID=$PROJECT_ID --service-account=bike-rental-service@$PROJECT_ID.iam.gserviceaccount.com --region=europe-west1
```

or quicker:

```sh
gcloud builds submit --substitutions=_MAPS_KEY=$MAPS_KEY --project $PROJECT_ID
```



### DEMO Hints

1. Check the endpoint and see the error
1. explore log output and use `explain this log entry` in logs explorer
1. Use Duet AI chat to figure out what permission I need to query the Big Query Dataset

    ```txt
    Which role do I need to create a bigquery job?
    ```

1. Fix the iam permissions

    ```sh
    gcloud projects add-iam-policy-binding $PROJECT_ID --member=serviceAccount:bike-rental-service@$PROJECT_ID.iam.gserviceaccount.com --role=roles/bigquery.user
    ```

1. Show that it worked

## Reset the demo

```sh
gcloud projects remove-iam-policy-binding $PROJECT_ID --member=serviceAccount:bike-rental-service@$PROJECT_ID.iam.gserviceaccount.com --role=roles/bigquery.user
```