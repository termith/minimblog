import os
import pytest


@pytest.yield_fixture(scope="module")
def run_engine():
    os.system('../minimblog')
    yield
    os.system('killall minimblog')