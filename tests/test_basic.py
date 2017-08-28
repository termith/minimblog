import subprocess

import pytest

mark = pytest.mark.usefixtures('run_engine')

def test_engine_running():
    """
        Check that engine is running
    """
    assert 'minimblog' in subprocess.check_output('ps ax')

 