import os
import threading
import subprocess
import requests

import pytest


@pytest.yield_fixture(scope="function")
def run_engine():
    minimblog = threading.Thread(target=lambda: os.system('minimblog'))
    minimblog.start()
    print('Engine started...')
    yield
    os.system('killall minimblog')
    minimblog.join()

def test_engine_running(run_engine):
    """
        Check that engine is running
    """
    assert 'minimblog' in str(subprocess.check_output(['ps', 'ax']))
    

def test_ping_engine(run_engine):
    response = requests.get("http://localhost:8080/ping")
    assert response.status_code == 200
    assert response.text == "ok"