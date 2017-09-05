import os
import threading
import subprocess


def test_engine_running():
    """
        Check that engine is running
    """
    minimblog = threading.Thread(target=lambda: os.system('minimblog'))
    minimblog.start()
    print('Engine started...')
    assert 'minimblog' in str(subprocess.check_output(['ps', 'ax']))
    os.system('killall minimblog')
    minimblog.join()
