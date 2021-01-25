import csv
import sys
import urllib.request
from datetime import datetime

from firebase_admin import firestore, credentials, initialize_app


def get_date(date_str) -> datetime:
    return datetime.strptime(date_str, "%Y-%m-%d")


def import_csv(url, collection_name, start_date, cert_path):
    cred = credentials.Certificate(cert_path)
    initialize_app(cred, {
        'projectId': 'another-covid-tracker-de304',
    })
    store = firestore.client()

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
            store.collection(collection_name).document(country).set({'entries': countries[country]})
        except Exception as e:
            print(e)


if __name__ == "__main__":
    url = sys.argv[1]
    collection_name = sys.argv[2]
    start_date = get_date(sys.argv[3])
    cert_path = sys.argv[4]
    import_csv(url, collection_name, start_date, cert_path)
