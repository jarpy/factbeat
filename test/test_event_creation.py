from factbeat_test import docker_stack  # noqa
from factbeat_test import query_all


def test_events_are_in_elasticserch(docker_stack):  # noqa
    assert len(query_all()) > 0
