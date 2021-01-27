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
            'cases': document.get('total_cases', None),
            'vaccinations': document.get('people_vaccinated', None),
            'date': document['date']
        }

    return list(map(to_entry, documents)), 200
