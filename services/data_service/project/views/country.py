from datetime import datetime
from typing import Tuple

from project.lookup.lookup import get_countries


def get(iso_code: str, from_date: str, to_date: str) -> Tuple[list, int]:
    documents: list = get_countries(iso_code.upper(),
                                    datetime.strptime(from_date, '%Y-%m-%d'),
                                    datetime.strptime(to_date, '%Y-%m-%d'))

    def to_entry(document) -> dict:
        return {
            'iso_code': document['iso_code'],
            'cases': float(document.get('new_cases', 0.0)),
            'vaccinations': float(document.get('people_fully_vaccinated', 0.0)),
            'date': document['date']
        }

    return [to_entry(doc) for doc in documents], 200
