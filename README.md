# NPS API to Postgres Database

I created this program to help me query the National Park Service API for all of it's parks and then input this information into my postgres database.

To use this as your own make sure to create a .env file in the project's root and inclde a DB_URL and API_KEY [get that here](https://www.nps.gov/subjects/developer/get-started.htm)

This program will print out each park name it successfully saves to the database and will stop when it is no longer receiving entries from the NPS API. My total count at the end was 474 different rows.
