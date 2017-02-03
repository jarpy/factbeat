from factbeat import BaseTest

import os


class Test(BaseTest):

    def test_base(self):
        """
        Basic test with exiting Factbeat normally
        """
        self.render_config_template(
                path=os.path.abspath(self.working_dir) + "/log/*"
        )

        factbeat_proc = self.start_beat()
        self.wait_until( lambda: self.log_contains("factbeat is running"))
        exit_code = factbeat_proc.kill_and_wait()
        assert exit_code == 0
