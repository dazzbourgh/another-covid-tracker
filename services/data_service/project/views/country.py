from typing import Tuple

from google.cloud import firestore

db = firestore.Client()


def get(iso_code, from_date, to_date) -> Tuple[list, int]:
    documents: list = [e.to_dict() for e in db.collection('countries').where('iso_code', '==', iso_code)
        .where('date', '>=', from_date)
        .where('date', '<=', to_date)
        .stream()]

    def to_entry(document) -> dict:
        return {
            "iso_code": document["iso_code"],
            "cases": document["cases"],
            "vaccinations": document["vaccinations"]
        }

    return [map(to_entry, documents)], 200
