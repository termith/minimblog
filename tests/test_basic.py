import os
import subprocess


def test_engine_running():
    """
        Check that engine is running
    """
    os.system('minimblog')
    assert 'minimblog' in str(subprocess.check_output(['ps', 'ax']))
    os.system('killall minimblog')