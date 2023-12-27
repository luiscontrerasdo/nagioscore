# Playing with Nagios Core
Install-nagioscore.sh -> This is a script for installing Nagios Core 4.4.14, Nagios Plugins 2.4.6 and NRPE 4.1.0

create-hosts.go -> This script in golang will perform the following tasks

Create a file named all_hosts.cfg with 100 hosts with the ping service check, the secuence of the names will be host###

It will create an entry to the nagios.cfg pointing the file cfg_file=/usr/local/nagios/etc/objects/all_hosts.cfg

Assign user and group permission to the file all_hosts.cfg

Nagios service will be restarted
