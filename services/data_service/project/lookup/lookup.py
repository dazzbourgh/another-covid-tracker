from datetime import datetime

from google.cloud import firestore
from google.oauth2.credentials import UserAccessTokenCredentials

db = firestore.Client(credentials=UserAccessTokenCredentials())


def get_countries(iso_code: str, from_date: datetime, to_date: datetime) -> list:
    return [e.to_dict() for e in db.collection('countries')
        .document(iso_code)
        .collection('items')
        .where('date', '>=', from_date)
        .where('date', '<=', to_date)
        .stream()]
