import requests
from retrying import retry
from pytest import fixture


class DockerStackError(Exception):
    pass


class FactbeatError(Exception):
    pass


def query_all():
    response = requests.get(
        "http://elasticsearch:9200/factbeat-*/_search").json()
    return response['hits']['hits']


@retry(stop_max_attempt_number=8, wait_exponential_multiplier=100)
def wait_for_elasticsearch():
    try:
        reply = requests.get('http://elasticsearch:9200')
    except requests.ConnectionError:
        raise DockerStackError("Elasticsearch is not answering.")

    status = reply.status_code
    if status != 200:
        raise DockerStackError("Elasticsearch returned HTTP %d." % status)

    cluster_name = reply.json()['cluster_name']
    if cluster_name != 'docker-factbeat':
        raise DockerStackError(
            "Elasticsearch cluster has the wrong name: '%s'." % cluster_name)


@retry(stop_max_attempt_number=8, wait_exponential_multiplier=100)
def wait_for_factbeat_data():
    wait_for_elasticsearch()
    try:
        assert len(query_all()) > 0
    except AssertionError:
        raise FactbeatError("Factbeat is not populating Elasticsearch.")


def setup_elasticsearch():
    wait_for_elasticsearch()
    requests.delete('http://elasticsearch:9200/factbeat-*')


@fixture
def docker_stack():
    """Ensure the Docker stack is up and available for testing."""
    # subprocess.call(['docker-compose', 'up', '-d'])
    setup_elasticsearch()
    wait_for_factbeat_data()
