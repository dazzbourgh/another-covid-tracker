import sys

from firebase_admin import credentials, initialize_app, firestore


def delete_collection(coll_ref, batch_size):
    docs = coll_ref.limit(batch_size).stream()
    deleted = 0

    for doc in docs:
        print(f'Deleting doc {doc.id} => {doc.to_dict()}')
        doc.reference.delete()
        deleted = deleted + 1

    if deleted >= batch_size:
        return delete_collection(coll_ref, batch_size)


if __name__ == '__main__':
    cred = credentials.Certificate(sys.argv[1])
    initialize_app(cred, {
        'projectId': 'another-covid-tracker-de304',
    })
    store = firestore.client()
    for c in store.collection('countries').document('AFG').collection('items').stream():
        print(c)
    delete_collection(store.collection('countries'), 10)
