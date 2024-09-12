package ast

import (
	"config-lsp/utils"
	"testing"
)

func TestSimpleParserExample(
	t *testing.T,
) {
	input := utils.Dedent(`
PermitRootLogin no
PasswordAuthentication yes
`)
	p := NewSSHConfig()
	errors := p.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	if !(p.Options.Size() == 2 &&
		len(utils.KeysOfMap(p.CommentLines)) == 0) {
		t.Errorf("Expected 2 options and no comment lines, but got: %v, %v", p.Options, p.CommentLines)
	}

	rawFirstEntry, _ := p.Options.Get(uint32(0))
	firstEntry := rawFirstEntry.(*SSHOption)

	if !(firstEntry.Value == "PermitRootLogin no" &&
		firstEntry.LocationRange.Start.Line == 0 &&
		firstEntry.LocationRange.End.Line == 0 &&
		firstEntry.LocationRange.Start.Character == 0 &&
		firstEntry.LocationRange.End.Character == 17 &&
		firstEntry.Key.Value == "PermitRootLogin" &&
		firstEntry.Key.LocationRange.Start.Character == 0 &&
		firstEntry.Key.LocationRange.End.Character == 14 &&
		firstEntry.OptionValue.Value == "no" &&
		firstEntry.OptionValue.LocationRange.Start.Character == 16 &&
		firstEntry.OptionValue.LocationRange.End.Character == 17) {
		t.Errorf("Expected first entry to be PermitRootLogin no, but got: %v", firstEntry)
	}

	rawSecondEntry, _ := p.Options.Get(uint32(1))
	secondEntry := rawSecondEntry.(*SSHOption)

	if !(secondEntry.Value == "PasswordAuthentication yes" &&
		secondEntry.LocationRange.Start.Line == 1 &&
		secondEntry.LocationRange.End.Line == 1 &&
		secondEntry.LocationRange.Start.Character == 0 &&
		secondEntry.LocationRange.End.Character == 25 &&
		secondEntry.Key.Value == "PasswordAuthentication" &&
		secondEntry.Key.LocationRange.Start.Character == 0 &&
		secondEntry.Key.LocationRange.End.Character == 21 &&
		secondEntry.OptionValue.Value == "yes" &&
		secondEntry.OptionValue.LocationRange.Start.Character == 23 &&
		secondEntry.OptionValue.LocationRange.End.Character == 25) {
		t.Errorf("Expected second entry to be PasswordAuthentication yes, but got: %v", secondEntry)
	}
}

func TestMatchSimpleBlock(
	t *testing.T,
) {
	input := utils.Dedent(`
PermitRootLogin no

Match 192.168.0.1
	PasswordAuthentication yes
`)
	p := NewSSHConfig()
	errors := p.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	if !(p.Options.Size() == 2 &&
		len(utils.KeysOfMap(p.CommentLines)) == 0) {
		t.Errorf("Expected 1 option and no comment lines, but got: %v, %v", p.Options, p.CommentLines)
	}

	rawFirstEntry, _ := p.Options.Get(uint32(0))
	firstEntry := rawFirstEntry.(*SSHOption)
	if !(firstEntry.Value == "PermitRootLogin no") {
		t.Errorf("Expected first entry to be 'PermitRootLogin no', but got: %v", firstEntry.Value)
	}

	rawSecondEntry, _ := p.Options.Get(uint32(2))
	secondEntry := rawSecondEntry.(*SSHMatchBlock)
	if !(secondEntry.MatchEntry.Value == "Match 192.168.0.1") {
		t.Errorf("Expected second entry to be 'Match 192.168.0.1', but got: %v", secondEntry.MatchEntry.Value)
	}

	if !(secondEntry.Options.Size() == 1) {
		t.Errorf("Expected 1 option in match block, but got: %v", secondEntry.Options)
	}

	rawThirdEntry, _ := secondEntry.Options.Get(uint32(3))
	thirdEntry := rawThirdEntry.(*SSHOption)
	if !(thirdEntry.Key.Value == "PasswordAuthentication" && thirdEntry.OptionValue.Value == "yes") {
		t.Errorf("Expected third entry to be 'PasswordAuthentication yes', but got: %v", thirdEntry.Value)
	}
}

func TestMultipleMatchBlocks(
	t *testing.T,
) {
	input := utils.Dedent(`
PermitRootLogin no

Match 192.168.0.1
	PasswordAuthentication yes
	AllowUsers root user

Match 192.168.0.2
	MaxAuthTries 3
`)
	p := NewSSHConfig()
	errors := p.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	if !(p.Options.Size() == 3 &&
		len(utils.KeysOfMap(p.CommentLines)) == 0) {
		t.Errorf("Expected 3 options and no comment lines, but got: %v, %v", p.Options, p.CommentLines)
	}

	rawSecondEntry, _ := p.Options.Get(uint32(2))
	secondEntry := rawSecondEntry.(*SSHMatchBlock)
	if !(secondEntry.Options.Size() == 2) {
		t.Errorf("Expected 2 options in second match block, but got: %v", secondEntry.Options)
	}

	rawThirdEntry, _ := secondEntry.Options.Get(uint32(3))
	thirdEntry := rawThirdEntry.(*SSHOption)
	if !(thirdEntry.Key.Value == "PasswordAuthentication" && thirdEntry.OptionValue.Value == "yes") {
		t.Errorf("Expected third entry to be 'PasswordAuthentication yes', but got: %v", thirdEntry.Value)
	}

	rawFourthEntry, _ := secondEntry.Options.Get(uint32(4))
	fourthEntry := rawFourthEntry.(*SSHOption)
	if !(fourthEntry.Key.Value == "AllowUsers" && fourthEntry.OptionValue.Value == "root user") {
		t.Errorf("Expected fourth entry to be 'AllowUsers root user', but got: %v", fourthEntry.Value)
	}

	rawFifthEntry, _ := p.Options.Get(uint32(6))
	fifthEntry := rawFifthEntry.(*SSHMatchBlock)
	if !(fifthEntry.Options.Size() == 1) {
		t.Errorf("Expected 1 option in fifth match block, but got: %v", fifthEntry.Options)
	}

	rawSixthEntry, _ := fifthEntry.Options.Get(uint32(7))
	sixthEntry := rawSixthEntry.(*SSHOption)
	if !(sixthEntry.Key.Value == "MaxAuthTries" && sixthEntry.OptionValue.Value == "3") {
		t.Errorf("Expected sixth entry to be 'MaxAuthTries 3', but got: %v", sixthEntry.Value)
	}

	firstOption, firstMatchBlock := p.FindOption(uint32(3))

	if !(firstOption.Key.Value == "PasswordAuthentication" && firstOption.OptionValue.Value == "yes" && firstMatchBlock.MatchEntry.Value == "Match 192.168.0.1") {
		t.Errorf("Expected first option to be 'PasswordAuthentication yes' and first match block to be 'Match 192.168.0.1', but got: %v, %v", firstOption, firstMatchBlock)
	}
}

func TestSimpleExampleWithComments(
	t *testing.T,
) {
	input := utils.Dedent(`
# Test
PermitRootLogin no
Port 22
# Second test
AddressFamily any
Sample
`)
	p := NewSSHConfig()
	errors := p.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	if !(p.Options.Size() == 4 &&
		len(utils.KeysOfMap(p.CommentLines)) == 2) {
		t.Errorf("Expected 3 options and 2 comment lines, but got: %v, %v", p.Options, p.CommentLines)
	}

	rawFirstEntry, _ := p.Options.Get(uint32(1))
	firstEntry := rawFirstEntry.(*SSHOption)
	if !(firstEntry.Value == "PermitRootLogin no") {
		t.Errorf("Expected first entry to be 'PermitRootLogin no', but got: %v", firstEntry.Value)
	}

	if len(p.CommentLines) != 2 {
		t.Errorf("Expected 2 comment lines, but got: %v", p.CommentLines)
	}

	if !utils.KeyExists(p.CommentLines, uint32(0)) {
		t.Errorf("Expected comment line 0 to not exist, but it does")
	}

	if !(utils.KeyExists(p.CommentLines, uint32(3))) {
		t.Errorf("Expected comment line 2 to exist, but it does not")
	}

	rawSecondEntry, _ := p.Options.Get(uint32(5))
	secondEntry := rawSecondEntry.(*SSHOption)

	if !(secondEntry.Value == "Sample") {
		t.Errorf("Expected second entry to be 'Sample', but got: %v", secondEntry.Value)
	}

}

func TestComplexExample(
	t *testing.T,
) {
	// From https://gist.github.com/kjellski/5940875
	input := utils.Dedent(`
# This is the sshd server system-wide configuration file.  See
# sshd_config(5) for more information.

# This sshd was compiled with PATH=/usr/bin:/bin:/usr/sbin:/sbin

# The strategy used for options in the default sshd_config shipped with
# OpenSSH is to specify options with their default value where
# possible, but leave them commented.  Uncommented options change a
# default value.

#Port 22
#AddressFamily any
#ListenAddress 0.0.0.0
#ListenAddress ::

# The default requires explicit activation of protocol 1
#Protocol 2

# HostKey for protocol version 1
#HostKey /etc/ssh/ssh_host_key
# HostKeys for protocol version 2
#HostKey /etc/ssh/ssh_host_rsa_key
#HostKey /etc/ssh/ssh_host_dsa_key
#HostKey /etc/ssh/ssh_host_ecdsa_key

# Lifetime and size of ephemeral version 1 server key
#KeyRegenerationInterval 1h
#ServerKeyBits 1024

# Logging
# obsoletes QuietMode and FascistLogging
#SyslogFacility AUTH
#LogLevel INFO

# Authentication:

#LoginGraceTime 2m
#BC# Root only allowed to login from LAN IP ranges listed at end
PermitRootLogin no
#PermitRootLogin yes
#StrictModes yes
#MaxAuthTries 6
#MaxSessions 10

#RSAAuthentication yes
#PubkeyAuthentication yes
#AuthorizedKeysFile  .ssh/authorized_keys

# For this to work you will also need host keys in /etc/ssh/ssh_known_hosts
#RhostsRSAAuthentication no
# similar for protocol version 2
#HostbasedAuthentication no
# Change to yes if you don't trust ~/.ssh/known_hosts for
# RhostsRSAAuthentication and HostbasedAuthentication
#IgnoreUserKnownHosts no
# Don't read the user's ~/.rhosts and ~/.shosts files
#IgnoreRhosts yes

# To disable tunneled clear text passwords, change to no here!
#BC# Disable password authentication by default (except for LAN IP ranges listed later)
PasswordAuthentication no
PermitEmptyPasswords no
#BC# Have to allow root here because AllowUsers not allowed in Match block.  It will not work though because of PermitRootLogin.
#BC# This is no longer true as of 6.1.  AllowUsers is now allowed in a Match block.
AllowUsers kmk root

# Change to no to disable s/key passwords
#BC# I occasionally use s/key one time passwords generated by a phone app
ChallengeResponseAuthentication yes

# Kerberos options
#KerberosAuthentication no
#KerberosOrLocalPasswd yes
#KerberosTicketCleanup yes
#KerberosGetAFSToken no

# GSSAPI options
#GSSAPIAuthentication no
#GSSAPICleanupCredentials yes

# Set this to 'yes' to enable PAM authentication, account processing, 
# and session processing. If this is enabled, PAM authentication will 
# be allowed through the ChallengeResponseAuthentication and
# PasswordAuthentication.  Depending on your PAM configuration,
# PAM authentication via ChallengeResponseAuthentication may bypass
# the setting of "PermitRootLogin without-password".
# If you just want the PAM account and session checks to run without
# PAM authentication, then enable this but set PasswordAuthentication
# and ChallengeResponseAuthentication to 'no'.
#BC# I would turn this off but I compiled ssh without PAM support so it errors if I set this.
#UsePAM no

#AllowAgentForwarding yes
#AllowTcpForwarding yes
#GatewayPorts no
X11Forwarding yes
#X11DisplayOffset 10
#X11UseLocalhost yes
#PrintMotd yes
#PrintLastLog yes
#TCPKeepAlive yes
#UseLogin no
#UsePrivilegeSeparation yes
#PermitUserEnvironment no
#Compression delayed
#ClientAliveInterval 0
#ClientAliveCountMax 3
#UseDNS yes
#PidFile /var/run/sshd.pid
#MaxStartups 10
#PermitTunnel no
#ChrootDirectory none

# no default banner path
#Banner none

# override default of no subsystems
#Subsystem	sftp	/usr/lib/misc/sftp-server
Subsystem	sftp	internal-sftp

# the following are HPN related configuration options
# tcp receive buffer polling. disable in non autotuning kernels
#TcpRcvBufPoll yes
 
# allow the use of the none cipher
#NoneEnabled no

# disable hpn performance boosts. 
#HPNDisabled no

# buffer size for hpn to non-hpn connections
#HPNBufferSize 2048


# Example of overriding settings on a per-user basis
Match User anoncvs
	X11Forwarding no
	AllowTcpForwarding no
	ForceCommand cvs server

#BC# My internal networks
#BC# Root can log in from here but only with a key and kmk can log in here with a password.
Match Address 172.22.100.0/24,172.22.5.0/24,127.0.0.1
  PermitRootLogin without-password
  PasswordAuthentication yes
`)
	p := NewSSHConfig()
	errors := p.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	if !(p.Options.Size() == 9 &&
		len(utils.KeysOfMap(p.CommentLines)) == 105) {
		t.Errorf("Expected 9 options and 105 comment lines, but got: %v, %v", p.Options, p.CommentLines)
	}

	rawFirstEntry, _ := p.Options.Get(uint32(38))
	firstEntry := rawFirstEntry.(*SSHOption)
	if !(firstEntry.Key.Value == "PermitRootLogin" && firstEntry.OptionValue.Value == "no") {
		t.Errorf("Expected first entry to be 'PermitRootLogin no', but got: %v", firstEntry.Value)
	}

	rawSecondEntry, _ := p.Options.Get(uint32(60))
	secondEntry := rawSecondEntry.(*SSHOption)
	if !(secondEntry.Key.Value == "PasswordAuthentication" && secondEntry.OptionValue.Value == "no") {
		t.Errorf("Expected second entry to be 'PasswordAuthentication no', but got: %v", secondEntry.Value)
	}

	rawThirdEntry, _ := p.Options.Get(uint32(118))
	thirdEntry := rawThirdEntry.(*SSHOption)
	if !(thirdEntry.Key.Value == "Subsystem" && thirdEntry.OptionValue.Value == "sftp\tinternal-sftp") {
		t.Errorf("Expected third entry to be 'Subsystem sftp internal-sftp', but got: %v", thirdEntry.Value)
	}

	rawFourthEntry, _ := p.Options.Get(uint32(135))
	fourthEntry := rawFourthEntry.(*SSHMatchBlock)
	if !(fourthEntry.MatchEntry.Value == "Match User anoncvs") {
		t.Errorf("Expected fourth entry to be 'Match User anoncvs', but got: %v", fourthEntry.MatchEntry.Value)
	}

	if !(fourthEntry.Options.Size() == 3) {
		t.Errorf("Expected 3 options in fourth match block, but got: %v", fourthEntry.Options)
	}

	rawFifthEntry, _ := fourthEntry.Options.Get(uint32(136))
	fifthEntry := rawFifthEntry.(*SSHOption)
	if !(fifthEntry.Key.Value == "X11Forwarding" && fifthEntry.OptionValue.Value == "no") {
		t.Errorf("Expected fifth entry to be 'X11Forwarding no', but got: %v", fifthEntry.Value)
	}

	rawSixthEntry, _ := p.Options.Get(uint32(142))
	sixthEntry := rawSixthEntry.(*SSHMatchBlock)
	if !(sixthEntry.MatchEntry.Value == "Match Address 172.22.100.0/24,172.22.5.0/24,127.0.0.1") {
		t.Errorf("Expected sixth entry to be 'Match Address 172.22.100.0/24,172.22.5.0/24,127.0.0.1', but got: %v", sixthEntry.MatchEntry.Value)
	}

	if !(sixthEntry.MatchEntry.Key.Value == "Match" && sixthEntry.MatchEntry.OptionValue.Value == "Address 172.22.100.0/24,172.22.5.0/24,127.0.0.1") {
		t.Errorf("Expected sixth entry to be 'Match Address 172.22.100.0/24,172.22.5.0/24,127.0.0.1', but got: %v", sixthEntry.MatchEntry.Value)
	}

	if !(sixthEntry.Options.Size() == 2) {
		t.Errorf("Expected 2 options in sixth match block, but got: %v", sixthEntry.Options)
	}

	rawSeventhEntry, _ := sixthEntry.Options.Get(uint32(143))
	seventhEntry := rawSeventhEntry.(*SSHOption)
	if !(seventhEntry.Key.Value == "PermitRootLogin" && seventhEntry.OptionValue.Value == "without-password") {
		t.Errorf("Expected seventh entry to be 'PermitRootLogin without-password', but got: %v", seventhEntry.Value)
	}
}
