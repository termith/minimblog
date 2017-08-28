import os
import pytest


@pytest.yield_fixture(scope="module")
def run_engine():
    os.system('../minimblog')
    print('Engine has been started')
    yield
    print('Stoping engine...')
    os.system('killall minimblog')