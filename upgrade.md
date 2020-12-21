#### How to upgrade your node to the latest version
* Login into the server hosting client software.
* Download the upgrade script by:
    - `curl https://raw.githubusercontent.com/DOSNetwork/core/master/upgrade.sh --output upgrade.sh`
    - Or `wget https://raw.githubusercontent.com/DOSNetwork/core/master/upgrade.sh -O upgrade.sh`
* Make the script runnable: `chmod a+x upgrade.sh`
* Export password of the keystore file the node operates with: `export PASSWORD="your-own-password"`
* Start the updrade process and run the node with: `./upgrade.sh`
* Congratulations, now you should've been done! Check node status by log files `tail -n 100 -f $(ls -l ~/vault/doslog.txt | cut -d '>' -f2)` or simply by querying rest endpoint `curl localhost:8080`.
* Ask for help or additional questions in our node runner channel [here](https://t.me/dos_node).
