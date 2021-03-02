import csv
import sys
import urllib.request
from datetime import datetime

from google.cloud import firestore
from google.oauth2.credentials import UserAccessTokenCredentials


def get_date(date_str) -> datetime:
    return datetime.strptime(date_str, "%Y-%m-%d")


def import_csv(url: str, collection_name: str, start_date: datetime):
    db = firestore.Client(credentials=UserAccessTokenCredentials())

    lines = map(lambda bs: bs.decode('utf-8'), urllib.request.urlopen(url))

    countries = {}

    csv_reader = csv.reader(lines, delimiter=',')
    headers = list(next(csv_reader, None))
    for row in csv_reader:
        obj = {}
        lst = list(row)
        for idx, value in enumerate(lst):
            if value != '':
                obj[headers[idx]] = value
        if 'iso_code' not in obj or get_date(obj['date']) < start_date:
            continue
        if obj['iso_code'] not in countries:
            countries[obj['iso_code']] = [obj]
        else:
            countries[obj['iso_code']] += [obj]

    for country in filter(lambda c: c != '', countries):
        try:
            batch = db.batch()
            for entry in countries[country]:
                date_string = entry['date']
                copy = dict(entry)
                copy['date'] = get_date(date_string)
                db.collection(collection_name) \
                    .document(country) \
                    .collection('items') \
                    .document(entry['date']) \
                    .set(copy)
            batch.commit()
        except Exception as e:
            print(e)


if __name__ == "__main__":
    url = sys.argv[1]
    collection_name = sys.argv[2]
    start_date = get_date(sys.argv[3])
    import_csv(url, collection_name, start_date)
