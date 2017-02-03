from __future__ import print_function
from factbeat_test import docker_stack  # noqa
from factbeat_test import query_all


def test_dots_in_fact_names_become_underscores(docker_stack):  # noqa
    hits = query_all()
    assert "dotted_fact_name" in hits[0]['_source']
