from datetime import datetime

from google.cloud import firestore

db = firestore.Client()


def get_countries(iso_code: str, from_date: datetime, to_date: datetime) -> list:
    return [e.to_dict() for e in db.collection('countries')
        .document(iso_code)
        .collection('items')
        .where('date', '>=', from_date)
        .where('date', '<=', to_date)
        .stream()]
