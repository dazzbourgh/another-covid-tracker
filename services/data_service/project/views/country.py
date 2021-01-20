from google.cloud import firestore

db = firestore.Client()


def get(iso_code, from_date, to_date):
    return db.collection('countries').where('iso_code', '==', iso_code)\
               .where('date', '>=', from_date)\
               .where('date', '<=', to_date)\
               .stream(), 200


def search(iso_code, from_date, to_date):
    return get(iso_code, from_date, to_date)