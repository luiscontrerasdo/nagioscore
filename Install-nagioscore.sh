#!/bin/bash -x

sed -i 's/SELINUX=.*/SELINUX=disabled/g' /etc/selinux/config
setenforce 0

dnf install -y gcc glibc glibc-common perl httpd php wget gd gd-devel
dnf install openssl-devel
dnf update -y

cd /tmp
wget -O nagioscore.tar.gz https://github.com/NagiosEnterprises/nagioscore/archive/nagios-4.4.14.tar.gz
tar xzf nagioscore.tar.gz

cd /tmp/nagioscore-nagios-4.4.14/
./configure
make all

make install-groups-users
usermod -a -G nagios apache

make install

make install-daemoninit
systemctl enable httpd.service

make install-commandmode
make install-config
make install-webconf

firewall-cmd --zone=public --add-port=80/tcp
firewall-cmd --zone=public --add-port=80/tcp --permanent


htpasswd -c /usr/local/nagios/etc/htpasswd.users nagiosadmin

systemctl start httpd.service
systemctl start nagios.service
systemctl enable httpd.service
systemctl enable nagios.service

yum install -y gcc glibc glibc-common make gettext automake autoconf wget openssl-devel net-snmp net-snmp-utils epel-release
yum --enablerepo=powertools,epel install perl-Net-SNMP
yum --enablerepo=PowerTools,epel install perl-Net-SNMP

cd /tmp
wget --no-check-certificate -O nagios-plugins.tar.gz https://github.com/nagios-plugins/nagios-plugins/archive/release-2.4.6.tar.gz
tar zxf nagios-plugins.tar.gz


cd /tmp/nagios-plugins-release-2.4.6/
./tools/setup
./configure
make
make install

systemctl start nagios.service
tar zxf nagios-plugins.tar.gz


cd /tmp/nagios-plugins-release-2.4.6/
./tools/setup
./configure
make
make install

systemctl start nagios.service
systemctl stop nagios.service
systemctl restart nagios.service
systemctl status nagios.service

yum install -y gcc glibc glibc-common openssl openssl-devel perl wget

cd /tmp
wget --no-check-certificate -O nrpe.tar.gz https://github.com/NagiosEnterprises/nrpe/archive/nrpe-4.1.0.tar.gz
tar xzf nrpe.tar.gz

cd /tmp/nrpe-nrpe-4.1.0/
./configure --enable-command-args

make all
make install-groups-users
make install
make install-config
#make install-plugins
make install-plugin

